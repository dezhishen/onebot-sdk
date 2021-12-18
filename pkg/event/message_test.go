package event

import (
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/dezhishen/onebot-sdk/pkg/model"
)

func TestListenMessagePrivate(t *testing.T) {
	ListenMessagePrivate(func(data model.EventMessagePrivate) error {
		log.Printf("1:%v", data.Message)
		return nil
	})
	ListenMessagePrivate(func(data model.EventMessagePrivate) error {
		log.Printf("2:%v", data.Message)
		return nil
	})
	StartWs()
}

func TestListenMessageGroup(t *testing.T) {
	ListenMessageGroup(func(data model.EventMessageGroup) error {
		log.Printf("1:%v", data.Message)
		return nil
	})
	ListenMessageGroup(func(data model.EventMessageGroup) error {
		log.Printf("2:%v", data.Message)
		return nil
	})
	StartWs()
}
