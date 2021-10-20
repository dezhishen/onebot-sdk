package event

import (
	"fmt"
	"net/url"
	"os"
	"time"

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
	done := make(chan struct{})
	interrupt := make(chan os.Signal, 1)
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("%s", message)
		}
	}()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func ListenReciverPrivateMsg() {

}
