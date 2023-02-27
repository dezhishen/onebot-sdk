package event

import (
	"context"
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/dezhishen/onebot-sdk/pkg/client"
	"github.com/dezhishen/onebot-sdk/pkg/config"
	"github.com/dezhishen/onebot-sdk/pkg/event/base"
	"github.com/dezhishen/onebot-sdk/pkg/event/message"
	"github.com/dezhishen/onebot-sdk/pkg/event/meta"
	"github.com/dezhishen/onebot-sdk/pkg/event/notice"
	"github.com/dezhishen/onebot-sdk/pkg/event/request"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type onebotEventcli struct {
	message.OnebotMessageEventCli // 消息事件
	notice.OnebotNoticeEventCli   // 通知事件
	request.OnebotRequestEventCli // 请求事件
	meta.OnebotMetaEventCli       // 元事件
	conf                          config.OnebotConfig
	ws                            *client.WebsocketCli
	allHandler                    map[base.OnebotEventType]base.OnebotBaseEventCli
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
		allHandler: make(map[base.OnebotEventType]base.OnebotBaseEventCli),
	}
	// 消息事件
	messageEventHandler, err := message.NewOnebotMessageEventCli()
	if err != nil {
		return nil, err
	}
	result.allHandler[base.OnebotEventTypeMessage] = messageEventHandler
	result.OnebotMessageEventCli = messageEventHandler
	// 元事件
	metaEventHandler, err := meta.NewOnebotMetaEventCli()
	if err != nil {
		return nil, err
	}
	result.allHandler[base.OnebotEventTypeMetaEvent] = metaEventHandler
	result.OnebotMetaEventCli = metaEventHandler
	// 通知事件
	noticeEventHandler, err := notice.NewOnebotNoticeEventCli()
	if err != nil {
		return nil, err
	}
	result.allHandler[base.OnebotEventTypeNotice] = noticeEventHandler
	result.OnebotNoticeEventCli = noticeEventHandler
	// 请求事件
	requestEventHandler, err := request.NewOnebotRequestEventCli()
	if err != nil {
		return nil, err
	}
	result.allHandler[base.OnebotEventTypeRequest] = requestEventHandler
	result.OnebotRequestEventCli = requestEventHandler
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
		var eventType = base.OnebotEventType(eventBaseInfo.PostType)
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

func runHandler(handler base.OnebotBaseEventCli, message []byte) {
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
