package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/dezhishen/onebot-sdk/pkg/config"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type WebsocketCli struct {
	actionCli   *websocket.Conn
	endpoint    string
	accessToken string
}
type actionReq struct {
	Action string      `json:"action"`
	Params interface{} `json:"params"`
	Echo   string      `json:"echo"`
}

func NewWebsocketCli(conf *config.OnebotConfig) *WebsocketCli {
	if strings.HasSuffix("/", conf.Endpoint) {
		conf.Endpoint = strings.TrimSuffix(conf.Endpoint, "/")
	}
	return &WebsocketCli{
		endpoint:    conf.Endpoint,
		accessToken: conf.AccessToken,
	}
}

func (cli *WebsocketCli) SendWithOutParam(action string) error {
	url := fmt.Sprintf("%s/api", cli.endpoint)
	if cli.accessToken != "" {
		url = fmt.Sprintf("%s?access_token=%s", url, cli.accessToken)
	}
	if cli.actionCli == nil {
		var err error
		cli.actionCli, _, err = websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			return err
		}
	}
	reqData := actionReq{
		Action: action,
	}
	err := cli.actionCli.WriteJSON(reqData)
	if err != nil {
		cli.actionCli.Close()
		cli.actionCli = nil
	}
	return err
}

func (cli *WebsocketCli) SendWithOutParamForResult(action string, result interface{}) error {
	url := fmt.Sprintf("%s/api", cli.endpoint)
	if cli.accessToken != "" {
		url = fmt.Sprintf("%s?access_token=%s", url, cli.accessToken)
	}
	if cli.actionCli == nil {
		var err error
		cli.actionCli, _, err = websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			return err
		}
	}
	reqData := actionReq{
		Action: action,
	}
	err := cli.actionCli.WriteJSON(reqData)
	if err != nil {
		cli.actionCli.Close()
		cli.actionCli = nil
	}
	return err
}

func (cli *WebsocketCli) Send(action string, params interface{}) error {
	url := fmt.Sprintf("%s/api", cli.endpoint)
	if cli.accessToken != "" {
		url = fmt.Sprintf("%s?access_token=%s", url, cli.accessToken)
	}
	if cli.actionCli == nil {
		var err error
		cli.actionCli, _, err = websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			return err
		}
	}
	reqData := actionReq{
		Action: action,
		Params: params,
	}
	err := cli.actionCli.WriteJSON(reqData)
	if err != nil {
		cli.actionCli.Close()
		cli.actionCli = nil
	}
	return err
}

func (cli *WebsocketCli) SendForResult(action string, params interface{}, result interface{}) error {
	url := fmt.Sprintf("%s/api", cli.endpoint)
	if cli.accessToken != "" {
		url = fmt.Sprintf("%s?access_token=%s", url, cli.accessToken)
	}
	if cli.actionCli == nil {
		var err error
		cli.actionCli, _, err = websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			return err
		}
	}
	reqData := actionReq{
		Action: action,
		Params: params,
	}
	err := cli.actionCli.WriteJSON(reqData)
	if err != nil {
		cli.actionCli.Close()
		cli.actionCli = nil
	}
	return err
}

func (cli *WebsocketCli) Listen(ctx context.Context, callback func(message []byte) error) error {
	url := fmt.Sprintf("%s/event", cli.endpoint)
	if cli.accessToken != "" {
		url = fmt.Sprintf("%s?access_token=%s", url, cli.accessToken)
	}
	// url := cli.endpoint + "?access_token=" + cli.accessToken
	log.Infof("开始建立连接...地址:%v", url)
	c, _, err := websocket.DefaultDialer.DialContext(ctx, url, nil)
	if err != nil {
		log.Infof("建立连接出现错误...,错误信息%v", err)
		return err
	}
	log.Info("建立连接成功！")
	defer c.Close()
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Errorf("websocket发生错误:%v,将在3s后重试...", err)
				time.Sleep(3 * time.Second)
				for i := 1; i < 5; i++ {
					log.Infof("尝试重连....,第%v次", i)
					c, _, err = websocket.DefaultDialer.DialContext(ctx, url, nil)
					if err != nil {
						internal := 3 * i
						log.Errorf("第%v次重连失败，将在%vs后重试", i, internal)
						// for j := 0; j < i; j++ {
						time.Sleep(time.Duration(internal) * time.Second)
						// }
					}
				}
				if err != nil {
					log.Errorf("重连失败...,%v", err)
					return err
				}
				continue
			}
			// recover from panic
			defer func() {
				if err := recover(); err != nil {
					log.Errorf("处理回调函数发生错误...%v", err)
				}
			}()
			callback(message)
		}
	}
}
