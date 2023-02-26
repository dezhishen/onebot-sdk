package api

// var userIdForTest int64 = 1179551960
// var groupIdForTest int64 = 727670105

// func TestSendMsg(t *testing.T) {
// 	msg := model.MsgForSend{
// 		MessageType: "private",
// 		UserId:      userIdForTest,
// 		Message: []*model.MessageSegment{
// 			{
// 				Type: "text",
// 				Data: &model.MessageElementText{
// 					Text: "测试消息",
// 				},
// 			},
// 		},
// 	}
// 	cli, err := NewOnebotApiClientByConfigPath("")
// 	if err != nil {
// 		t.Errorf("NewOnebotApiClientByConfigPath() error = %v", err)
// 		return
// 	}
// 	got, err := cli.SendMsg(&msg)
// 	if err != nil {
// 		t.Errorf("SendMsg() error = %v", err)
// 		return
// 	}
// 	print(got)
// }

// func TestSendGroupMsg(t *testing.T) {
// 	msg := model.GroupMsg{
// 		GroupId: groupIdForTest,
// 		Message: []*model.MessageSegment{
// 			{
// 				Type: "text",
// 				Data: &model.MessageElementText{
// 					Text: "测试消息",
// 				},
// 			},
// 		},
// 	}
// 	cli, err := NewOnebotApiClientByConfigPath("")
// 	if err != nil {
// 		t.Errorf("NewOnebotApiClientByConfigPath() error = %v", err)
// 		return
// 	}
// 	got, err := cli.SendGroupMsg(&msg)
// 	if err != nil {
// 		t.Errorf("SendGroupMsg() error = %v", err)
// 		return
// 	}
// 	print(got)
// }

// func TestSendPrivateMsg(t *testing.T) {
// 	msg := model.PrivateMsg{
// 		UserId: userIdForTest,
// 		Message: []*model.MessageSegment{
// 			{
// 				Type: "text",
// 				Data: &model.MessageElementText{
// 					Text: "测试消息",
// 				},
// 			},
// 		},
// 	}
// 	cli, err := NewOnebotApiClientByConfigPath("")
// 	if err != nil {
// 		t.Errorf("NewOnebotApiClientByConfigPath() error = %v", err)
// 		return
// 	}
// 	got, err := cli.SendPrivateMsg(&msg)
// 	if err != nil {
// 		t.Errorf("SendPrivate() error = %v", err)
// 		return
// 	}
// 	print(got)
// }

// func TestSendGroupForwardMessage(t *testing.T) {
// 	msg := []*model.MessageSegment{
// 		{
// 			Type: "text",
// 			Data: &model.MessageElementText{
// 				Text: "测试转发消息",
// 			},
// 		},
// 	}
// 	cli, err := NewOnebotApiClientByConfigPath("")
// 	if err != nil {
// 		t.Errorf("NewOnebotApiClientByConfigPath() error = %v", err)
// 		return
// 	}
// 	loginInfo, err := cli.GetLoginInfo()
// 	if err != nil {
// 		t.Errorf("get login info () error = %v", err)
// 		return
// 	}
// 	var forwardMsg []*model.MessageSegment
// 	forwardMsg = append(forwardMsg, &model.MessageSegment{
// 		Type: "node",
// 		Data: &model.MessageElementNode{
// 			Uin:     loginInfo.Data.UserId,
// 			Name:    loginInfo.Data.Nickname,
// 			Content: msg,
// 		},
// 	})
// 	cli.SendGroupForwardMsg(groupIdForTest, forwardMsg)
// 	cli.SendGroupForwardMsgByRawMsg(groupIdForTest, loginInfo.Data.UserId, loginInfo.Data.Nickname, msg)
// }
