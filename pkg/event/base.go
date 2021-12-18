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
	// "github.com/gorilla/websocket"
)

const (
	EventTypeMessage   string = "message"
	EventTypeNotice    string = "notice"
	EventTypeRequest   string = "request"
	EventTypeMetaEvent string = "meta_event"
)

var allhandler = make(map[string]func(data []byte) error)

// StartWs 开始监听
// 监听前请注册相关事件
func StartWs() error {
	url := fmt.Sprintf("ws://%v:%v", config.GetWsHost(), config.GetWsPort())
	// url.URL{Scheme: "ws", Host: host, Path: ""}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	c, _, err := websocket.Dial(ctx, url, nil)
	if err != nil {
		return err
	}
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
		log.Infof("received event,post_type:%v", eventBaseInfo.PostType)
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

func setHandler(eventType string, handler func(data []byte) error) {
	allhandler[eventType] = handler
}
