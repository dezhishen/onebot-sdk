package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/dezhishen/onebot-sdk/pkg/config"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

// 发送消息
func SendMsg(msg *model.MsgForSend) (*model.MessageResult, error) {
	if msg.MessageType == model.PrivateMessageType {
		return SendPrivateMsg(&model.PrivateMsg{
			UserID:     msg.UserID,
			GroupID:    msg.GroupID,
			Message:    msg.Message,
			AutoEscape: msg.AutoEscape,
		})
	}
	return SendGroupMsg(&model.GroupMsg{
		GroupID:    msg.GroupID,
		Message:    msg.Message,
		AutoEscape: msg.AutoEscape})
}

// 发送私信
func SendPrivateMsg(msg *model.PrivateMsg) (*model.MessageResult, error) {
	requestBody, _ := json.Marshal(msg)
	resp, err := http.Post(
		config.GetHttpUrl()+"/send_private_msg",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	var result model.MessageResult
	err = json.Unmarshal(respBodyContent, &result)
	return &result, err
}

// 发送群消息
func SendGroupMsg(msg *model.GroupMsg) (*model.MessageResult, error) {
	requestBody, _ := json.Marshal(msg)
	resp, err := http.Post(
		config.GetHttpUrl()+"/send_group_msg",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	var result model.MessageResult
	json.Unmarshal(respBodyContent, &result)
	return &result, nil
}
