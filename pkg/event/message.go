package event

import (
	"fmt"
	"net/url"

	"github.com/dezhishen/onebot-sdk/pkg/config"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

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
		log.Printf("%s", message)
	}
}

func ListenReciverPrivateMsg(func(data []byte) error) {

}
