package channel

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"net/http"

	"github.com/dezhishen/onebot-sdk/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebsocketReverseEventChannel struct {
	svc     *http.Server
	handler func(message []byte) error
	Conn    *websocket.Conn
}

func NewWebsocketReverseEventChannel(conf *config.OnebotEventConfig) (EventChannel, error) {
	router := gin.Default()
	/*
		X-Signature的内容中，sha1的值的计算方式为：
		获取请求体的内容，即 requestBody
		使用密钥为 secret的HMAC算法加密 requestBody
		得到的HMAC加密值以Hex格式输出为字符串
	*/
	ws := &WebsocketReverseEventChannel{}
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
	router.Any("/event", func(c *gin.Context) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		var err error
		ws.Conn, err = upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		for {
			_, message, err := ws.Conn.ReadMessage()
			if err != nil {
				break
			}
			if ws.handler != nil {
				if err := ws.handler(message); err != nil {
					break
				}
			}
		}
	})
	ws.svc = &http.Server{
		Addr:    conf.Addr,
		Handler: router,
	}
	return ws, nil
}

func (ws *WebsocketReverseEventChannel) StartListen(ctx context.Context, handler func(message []byte) error) error {
	ws.handler = handler
	// 异常通过chan 传递
	errChan := make(chan error)
	go func() {
		if err := ws.svc.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// panic(err)
			errChan <- err
		}
	}()
	select {
	case <-ctx.Done():
		return ws.svc.Shutdown(ctx)
	case err := <-errChan:
		return err
	}
}
