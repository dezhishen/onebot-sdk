package event

import (
	"context"
	"encoding/json"
	"time"

	"github.com/dezhishen/onebot-sdk/pkg/model"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/websocket"
)

const (
	EventTypeMessage   string = "message"
	EventTypeNotice    string = "notice"
	EventTypeRequest   string = "request"
	EventTypeMetaEvent string = "meta_event"
)

type eventsFactory struct {
	allhandler map[string]func(data []byte) error
}

func NewEventsFactory() *eventsFactory {
	return &eventsFactory{
		allhandler: make(map[string]func(data []byte) error),
	}
}

var allhandler = make(map[string]func(data []byte) error)

// // StartWs 开始监听
// // 监听前请注册相关事件
// func StartWs() error {
// 	log.Warn("使用该方法将无法控制WS的关闭，推荐使用StartWsWithContext()")
// 	url := fmt.Sprintf("ws://%v:%v", config.GetWsHost(), config.GetWsPort())
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()
// 	return listen(url, ctx)

// }

// // StartWs 开始监听
// // 监听前请注册相关事件
// func StartWsWithContext(ctx context.Context) error {
// 	url := fmt.Sprintf("ws://%v:%v", config.GetWsHost(), config.GetWsPort())
// 	return listen(url, ctx)
// }

func runHandler(callback func(data []byte) error, message []byte) {
	err := callback(message)
	if err != nil {
		log.Errorf("handler event error:%v", err)
	}
}

func setHandler(eventType string, handler func(data []byte) error) {
	allhandler[eventType] = handler
}

func listen(url string, ctx context.Context) error {
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
			log.Printf("handle listen decoder err :%v,raw:%v", err, message)
		}
		log.Tracef("received event,post_type:%v", eventBaseInfo.PostType)
		handler, ok := allhandler[eventBaseInfo.PostType]
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
