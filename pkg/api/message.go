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
func SendMsg(msg *model.MsgForSend) (int, error) {
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
func SendPrivateMsg(msg *model.PrivateMsg) (int, error) {
	requestBody, _ := json.Marshal(msg)
	resp, err := http.Post(
		config.GetHttpUrl()+"/send_private_msg",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	var result model.SendMessageResult
	err = json.Unmarshal(respBodyContent, &result)
	return result.Data.MessageId, err
}

// 发送群消息
func SendGroupMsg(msg *model.GroupMsg) (int, error) {
	requestBody, _ := json.Marshal(msg)
	resp, err := http.Post(
		config.GetHttpUrl()+"/send_group_msg",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	var result model.SendMessageResult
	json.Unmarshal(respBodyContent, &result)
	return result.Data.MessageId, nil
}

func DelMsg(id int) error {
	reqMap := make(map[string]int)
	reqMap["message_id"] = id
	requestBody, _ := json.Marshal(reqMap)
	_, err := http.Post(
		config.GetHttpUrl()+"/get_msg",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

func GetMsg(id int) ([]*model.MessageResultData, error) {
	reqMap := make(map[string]int)
	reqMap["message_id"] = id
	requestBody, _ := json.Marshal(reqMap)
	resp, err := http.Post(
		config.GetHttpUrl()+"/get_msg",
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
	return result.Data, nil
}

func GetForwardMsg(id int) ([]*model.MessageSegment, error) {
	reqMap := make(map[string]int)
	reqMap["message_id"] = id
	requestBody, _ := json.Marshal(reqMap)
	resp, err := http.Post(
		config.GetHttpUrl()+"/get_msg",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	var result model.ForwardMessageResult
	json.Unmarshal(respBodyContent, &result)
	return result.Data, nil
}
