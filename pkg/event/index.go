package event

import (
	"context"
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/dezhishen/onebot-sdk/pkg/client"
	"github.com/dezhishen/onebot-sdk/pkg/config"
	"github.com/dezhishen/onebot-sdk/pkg/model"
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
	ws                    *client.WebsocketCli
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
	var url = cli.conf.Endpoint + "/event"
	// url := cli.conf.Endpoint
	if cli.conf.AccessToken != "" {
		url += "?access_token=" + cli.conf.AccessToken
	}
	return cli.listen(url, ctx)
}

func (cli *onebotEventcli) listen(url string, ctx context.Context) error {
	err := cli.ws.Listen(ctx, func(message []byte) error {
		var eventBaseInfo model.EventBase
		err := json.Unmarshal(message, &eventBaseInfo)
		if err != nil {
			return fmt.Errorf("handle listen decoder err :%v,raw:%v", err, message)
		}
		log.Debugf("received event,post_type:%v", eventBaseInfo.PostType)
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
		}
		return nil
	})
	return err
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
