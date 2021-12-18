package event

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/dezhishen/onebot-sdk/pkg/config"
	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
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
func StartWs() {
	host := fmt.Sprintf("%v:%v", config.GetWsHost(), config.GetWsPort())
	u := url.URL{Scheme: "ws", Host: host, Path: ""}
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
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
