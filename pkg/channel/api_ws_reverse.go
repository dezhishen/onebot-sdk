package channel

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/dezhishen/onebot-sdk/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type WebsocketReverseApiChannel struct {
	Ctx             context.Context
	Conn            *websocket.Conn
	ready           chan bool
	svc             *http.Server
	resultCallbacks map[string]map[string]func(data []byte) error
	timeout         time.Duration
}

func NewWebsocketReverseApiChannel(conf *config.OnebotApiConfig) (ApiChannel, error) {
	router := gin.Default()
	/*
		X-Signature的内容中，sha1的值的计算方式为：
		获取请求体的内容，即 requestBody
		使用密钥为 secret的HMAC算法加密 requestBody
		得到的HMAC加密值以Hex格式输出为字符串
	*/
	if conf.Timeout == 0 {
		// 默认10s
		conf.Timeout = DEFAULT_TIMEOUT
	}
	ws := &WebsocketReverseApiChannel{
		ready:           make(chan bool),
		timeout:         time.Duration(conf.Timeout) * time.Millisecond,
		resultCallbacks: make(map[string]map[string]func(data []byte) error),
	}
	router.Use(func(c *gin.Context) {
		hSignature := c.GetHeader("X-Signature")
		if hSignature != "" {
			if conf.Secret == "" {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			// 读取 request body
			body, err := io.ReadAll(c.Request.Body)
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			// gin框架中request body只能读取一次，需要复写 request body
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			// 使用密钥计算request body的 hmac码
			mac := hmac.New(sha1.New, []byte(conf.Secret))
			if _, err := mac.Write(body); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			// 校验hmac签名
			if "sha1="+hex.EncodeToString(mac.Sum(nil)) != hSignature {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}
		c.Next()
	})
	uri, err := url.Parse(conf.Endpoint)
	if err != nil {
		return nil, err
	}
	addr := uri.Host
	path := uri.Path
	router.Any(path, func(c *gin.Context) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		var err error
		ws.Conn, err = upgrader.Upgrade(c.Writer, c.Request, nil)
		ws.ready <- true
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		for {
			_, message, err := ws.Conn.ReadMessage()
			if err != nil {
				break
			}
			// 处理消息
			var result actionResult
			if err := json.Unmarshal(message, &result); err != nil {
				log.Errorf("unmarshal result error: %s", err)
				continue
			}
			// 处理回调
			// recover from panic
			defer func() {
				if err := recover(); err != nil {
					log.Errorf("处理回调函数发生错误...%v", err)
				}
			}()
			if strings.Contains(result.Echo, "/") {
				action := strings.Split(result.Echo, "/")[0]
				reqId := strings.Split(result.Echo, "/")[1]
				if cli, ok := ws.resultCallbacks[action]; ok {
					if callback, ok := cli[reqId]; ok {
						if err := callback(message); err != nil {
							break
						}
					}
				}
				continue
			}
		}
	})
	ws.svc = &http.Server{
		Addr:    addr,
		Handler: router,
	}
	// 异常通过chan 传递
	go func() {
		if err := ws.svc.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	// 等待连接
	select {
	case <-ws.ready:
		return ws, nil
	case <-time.After(30 * time.Second):
		return ws, fmt.Errorf("连接超时")
	}
}

func (cli *WebsocketReverseApiChannel) Post(action string) error {
	echo := fmt.Sprintf("%s/%d", action, time.Now().UnixNano())
	reqData := actionReq{
		Action: action,
		Echo:   echo,
	}
	err := cli.Conn.WriteJSON(reqData)
	if err != nil {
		cli.Conn.Close()
		cli.Conn = nil
	}
	return err
}

func (cli *WebsocketReverseApiChannel) PostForResult(action string, result interface{}) error {
	reqId := fmt.Sprintf("%d", time.Now().UnixNano())
	echo := fmt.Sprintf("%s/%s", action, reqId)
	reqData := actionReq{
		Action: action,
		Echo:   echo,
	}
	// 注册回调
	channel := make(chan byte)
	cli.registerCallbackWithChannel(action, reqId, channel, result)
	defer func() {
		cli.unregisterCallback(action, reqId)
		close(channel)
	}()
	err := cli.Conn.WriteJSON(reqData)
	if err != nil {
		cli.Conn.Close()
		cli.Conn = nil
	}
	// 等待回调
	select {
	case <-channel:
		return nil
	case <-time.After(cli.timeout):
		return fmt.Errorf("request timeout")
	}
}

func (cli *WebsocketReverseApiChannel) PostByRequest(action string, params interface{}) error {
	reqData := actionReq{
		Action: action,
		Params: params,
	}
	err := cli.Conn.WriteJSON(reqData)
	if err != nil {
		cli.Conn.Close()
		cli.Conn = nil
	}
	return err
}

func (cli *WebsocketReverseApiChannel) PostByRequestForResult(action string, params interface{}, result interface{}) error {
	reqId := fmt.Sprintf("%d", time.Now().UnixNano())
	echo := fmt.Sprintf("%s/%s", action, reqId)
	reqData := actionReq{
		Action: action,
		Params: params,
		Echo:   echo,
	}
	// 注册回调
	channel := make(chan byte)
	cli.registerCallbackWithChannel(action, reqId, channel, result)
	defer func() {
		cli.unregisterCallback(action, reqId)
		close(channel)
	}()
	err := cli.Conn.WriteJSON(reqData)
	if err != nil {
		cli.Conn.Close()
		cli.Conn = nil
	}
	select {
	case <-channel:
		return nil
	case <-time.After(cli.timeout):
		return fmt.Errorf("request timeout")
	}
}

func (cli *WebsocketReverseApiChannel) registerCallbackWithChannel(
	action string,
	reqId string,
	channel chan byte,
	result interface{},
) {
	if _, ok := cli.resultCallbacks[action]; !ok {
		cli.resultCallbacks[action] = make(map[string]func(data []byte) error)
	}
	cli.resultCallbacks[action][reqId] =
		func(data []byte) error {
			json.Unmarshal(data, result)
			channel <- 1
			return nil
		}
}

func (cli *WebsocketReverseApiChannel) unregisterCallback(
	action string,
	reqId string,
) {
	if _, ok := cli.resultCallbacks[action]; ok {
		delete(cli.resultCallbacks[action], reqId)
	}
}
