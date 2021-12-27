package api

import (
	"testing"

	"github.com/dezhishen/onebot-sdk/pkg/model"
)

var userId int64 = 123456
var groupId int64 = 123456

func TestSendMsg(t *testing.T) {
	msg := model.MsgForSend{
		MessageType: "private",
		UserId:      userId,
		Message: []*model.MessageSegment{
			{
				Type: "text",
				Data: &model.MessageElementText{
					Text: "测试消息",
				},
			},
		},
	}
	got, err := SendMsg(&msg)
	if err != nil {
		t.Errorf("SendMsg() error = %v", err)
		return
	}
	print(got)
}

func TestSendGroupMsg(t *testing.T) {
	msg := model.GroupMsg{
		GroupId: groupId,
		Message: []*model.MessageSegment{
			{
				Type: "text",
				Data: &model.MessageElementText{
					Text: "测试消息",
				},
			},
		},
	}
	got, err := SendGroupMsg(&msg)
	if err != nil {
		t.Errorf("SendGroupMsg() error = %v", err)
		return
	}
	print(got)
}

func TestSendPrivateMsg(t *testing.T) {
	msg := model.PrivateMsg{
		UserId: userId,
		Message: []*model.MessageSegment{
			{
				Type: "text",
				Data: &model.MessageElementText{
					Text: "测试消息",
				},
			},
		},
	}
	got, err := SendPrivateMsg(&msg)
	if err != nil {
		t.Errorf("SendPrivate() error = %v", err)
		return
	}
	print(got)
}
