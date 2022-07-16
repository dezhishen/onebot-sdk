package api

import (
	"testing"

	"github.com/dezhishen/onebot-sdk/pkg/model"
)

var userId int64 = 1179551960
var groupId int64 = 727670105

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

func TestSendGroupForwardMessage(t *testing.T) {
	msg := []*model.MessageSegment{
		{
			Type: "text",
			Data: &model.MessageElementText{
				Text: "测试消息",
			},
		},
	}
	loginInfo, err := GetLoginInfo()
	if err != nil {
		t.Errorf("get login info () error = %v", err)
		return
	}
	var forwardMsg []*model.MessageSegment
	forwardMsg = append(forwardMsg, &model.MessageSegment{
		Type: "node",
		Data: &model.MessageElementNode{
			Uin:     loginInfo.Data.UserId,
			Name:    loginInfo.Data.Nickname,
			Content: msg,
		},
	})

	SendGroupForwardMsg(groupId, forwardMsg)
}
