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
func SendMsg(msg *model.MsgForSend) (string, error) {
	if msg.MessageType == "private" {
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
func SendPrivateMsg(msg *model.PrivateMsg) (string, error) {
	requestBody, _ := json.Marshal(msg)
	resp, err := http.Post(
		config.GetUrl()+"/send_private_msg",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	return string(respBodyContent), nil

}

// 发送群消息
func SendGroupMsg(msg *model.GroupMsg) (string, error) {
	requestBody, _ := json.Marshal(msg)
	resp, err := http.Post(
		config.GetUrl()+"/send_group_msg",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	return string(respBodyContent), nil
}
