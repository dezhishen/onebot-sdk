package event

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/dezhishen/onebot-sdk/pkg/config"
	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/gorilla/websocket"
)

type onebotBaseEventCli interface {
	EventType() OnebotEventType
	Handler(data []byte) error
}

type onebotEventcli struct {
	onebotMessageEventCli // 消息事件
	onebotNoticeEventCli  // 通知事件
	onebotRequestEventCli // 请求事件
	onebotMetaEventCli    // 元事件
	conf                  config.OnebotConfig
	allHandler            map[OnebotEventType]onebotBaseEventCli
	// StartWsWithContext(ctx context.Context) error
}

func NewOnebotEventCli(conf *config.OnebotConfig) (*onebotEventcli, error) {
	if conf == nil {
		return nil, fmt.Errorf("config is nil")
	}
	if conf.Type == "" {
		return nil, fmt.Errorf("type is empty")
	}
	if conf.Type != "websocket" {
		return nil, fmt.Errorf("type is not websocket")
	}
	if conf.Endpoint == "" {
		return nil, fmt.Errorf("endpoint is empty")
	}
	var result = &onebotEventcli{
		conf:       *conf,
		allHandler: make(map[OnebotEventType]onebotBaseEventCli),
	}
	// 消息事件
	messageEventHandler, err := NewOnebotMessageEventCli()
	if err != nil {
		return nil, err
	}
	result.allHandler[OnebotEventTypeMessage] = messageEventHandler
	result.onebotMessageEventCli = messageEventHandler
	// 元事件
	metaEventHandler, err := NewOnebotMetaEventCli()
	if err != nil {
		return nil, err
	}
	result.allHandler[OnebotEventTypeMetaEvent] = metaEventHandler
	result.onebotMetaEventCli = metaEventHandler
	// 通知事件
	noticeEventHandler, err := NewOnebotNoticeEventCli()
	if err != nil {
		return nil, err
	}
	result.allHandler[OnebotEventTypeNotice] = noticeEventHandler
	result.onebotNoticeEventCli = noticeEventHandler
	// 请求事件
	requestEventHandler, err := NewOnebotRequestEventCli()
	if err != nil {
		return nil, err
	}
	result.allHandler[OnebotEventTypeRequest] = requestEventHandler
	result.onebotRequestEventCli = requestEventHandler
	return result, nil
}

func (cli *onebotEventcli) StartWithContext(ctx context.Context) error {
	var url = cli.conf.Endpoint
	// url := cli.conf.Endpoint
	if cli.conf.AccessToken != "" {
		url += "?access_token=" + cli.conf.AccessToken
	}
	return cli.listen(url, ctx)
}

func (cli *onebotEventcli) listen(url string, ctx context.Context) error {
	log.Infof("开始建立连接...地址:%v", url)
	c, _, err := websocket.DefaultDialer.DialContext(ctx, url, nil)
	if err != nil {
		log.Infof("建立连接出现错误...,错误信息%v", err)
		return err
	}
	log.Info("建立连接成功！")
	defer c.Close()
	for {
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
		var eventBaseInfo model.EventBase
		err = json.Unmarshal(message, &eventBaseInfo)
		if err != nil {
			log.Errorf("handle listen decoder err :%v,raw:%v", err, message)
		}
		log.Infof("received event,post_type:%v", eventBaseInfo.PostType)
		var eventType = OnebotEventType(eventBaseInfo.PostType)
		handler, ok := cli.allHandler[eventType]
		if ok {
			defer func() {
				if err := recover(); err != nil {
					// 打印异常，关闭资源，退出此函数
					log.Errorf("处理回调函数发生错误...%v", err)
				}
			}()
			go runHandler(handler, message)
		} else {
			log.Errorf("undefine event post_type:%v", eventBaseInfo.PostType)
		}
	}
}

func runHandler(handler onebotBaseEventCli, message []byte) {
	defer func() {
		if err := recover(); err != nil {
			// 打印异常，关闭资源，退出此函数
			log.Errorf("处理回调函数发生错误...%v", err)
		}
	}()
	err := handler.Handler(message)
	if err != nil {
		log.Errorf("handle listen decoder err :%v,raw:%v", err, message)
	}
}
