package event

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dezhishen/onebot-sdk/pkg/config"
	"github.com/dezhishen/onebot-sdk/pkg/model"
	log "github.com/sirupsen/logrus"
	"nhooyr.io/websocket"
)

const (
	EventTypeMessage   string = "message"
	EventTypeNotice    string = "notice"
	EventTypeRequest   string = "request"
	EventTypeMetaEvent string = "meta_event"
)

var allhandler = make(map[string]func(data []byte) error)

//StartWs 开始监听
//监听前请注册相关事件
func StartWs() error {
	log.Warn("使用该方法将无法控制WS的关闭，推荐使用StartWsWithContext()")
	url := fmt.Sprintf("ws://%v:%v", config.GetWsHost(), config.GetWsPort())
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	log.Infof("开始建立连接...地址:%v", url)
	c, _, err := websocket.Dial(ctx, url, nil)
	if err != nil {
		log.Infof("建立连接出现错误...,错误信息%v", err)
		return err
	}
	log.Info("建立连接成功！")
	defer c.Close(websocket.StatusNormalClosure, "disconnected")
	for {
		_, message, err := c.Read(ctx)
		if err != nil {
			log.Println("read:", err)
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
			err = handler(message)
			if err != nil {
				log.Errorf("handler event error:%v", err)

			}
		} else {
			log.Errorf("undefine event post_type:%v", eventBaseInfo.PostType)
		}
	}
}

// StartWs 开始监听
// 监听前请注册相关事件
func StartWsWithContext(ctx context.Context) error {
	url := fmt.Sprintf("ws://%v:%v", config.GetWsHost(), config.GetWsPort())
	log.Infof("开始建立连接...地址:%v", url)
	c, _, err := websocket.Dial(ctx, url, nil)
	if err != nil {
		log.Errorf("建立连接出现错误...,错误信息%v", err)
		return err
	}
	log.Info("建立连接成功！")
	defer c.Close(websocket.StatusNormalClosure, "disconnected")
	for {
		select {
		case <-ctx.Done():
			log.Info("关闭连接...")
			return nil
		default:
			_, message, err := c.Read(ctx)
			if err != nil {
				log.Println("read:", err)
				continue
			}
			var eventBaseInfo model.EventBase
			err = json.Unmarshal(message, &eventBaseInfo)
			if err != nil {
				log.Printf("handle listen decoder err :%v,raw:%v", err, message)
				continue
			}
			// log.Infof("received event,post_type:%v", eventBaseInfo.PostType)
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
}

func runHandler(callback func(data []byte) error, message []byte) {
	err := callback(message)
	if err != nil {
		log.Errorf("handler event error:%v", err)
	}
}

func setHandler(eventType string, handler func(data []byte) error) {
	allhandler[eventType] = handler
}
