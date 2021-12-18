package event

import (
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/dezhishen/onebot-sdk/pkg/model"
)

func TestListenPrivateMessage(t *testing.T) {
	ListenReciverPrivateMsg(func(data model.EventPrivateMsg) error {
		log.Println(data.Message)
		return nil
	})
	StartWs()
}

func TestListenGroupMsg(t *testing.T) {
	ListenReciverGroupMsg(func(data model.EventGroupMsg) error {
		log.Println(data.Message)
		return nil
	})
	StartWs()
}
