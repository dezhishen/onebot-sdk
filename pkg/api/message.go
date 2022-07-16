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
func SendMsg(msg *model.MsgForSend) (*model.SendMessageResult, error) {
	if msg.MessageType == model.PrivateMessageType {
		return SendPrivateMsg(&model.PrivateMsg{
			UserId:     msg.UserId,
			GroupId:    msg.GroupId,
			Message:    msg.Message,
			AutoEscape: msg.AutoEscape,
		})
	}
	return SendGroupMsg(
		&model.GroupMsg{
			GroupId:    msg.GroupId,
			Message:    msg.Message,
			AutoEscape: msg.AutoEscape,
		},
	)
}

// 发送私信
func SendPrivateMsg(msg *model.PrivateMsg) (*model.SendMessageResult, error) {
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
	var result model.SendMessageResult
	err = json.Unmarshal(respBodyContent, &result)
	return &result, err
}

// 发送群消息
func SendGroupMsg(msg *model.GroupMsg) (*model.SendMessageResult, error) {
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
	var result model.SendMessageResult
	json.Unmarshal(respBodyContent, &result)
	return &result, nil
}

func DelMsg(id int64) error {
	reqMap := make(map[string]int64)
	reqMap["message_id"] = id
	requestBody, _ := json.Marshal(reqMap)
	_, err := http.Post(
		config.GetHttpUrl()+"/get_msg",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

func GetMsg(id int64) (*model.MessageDataResult, error) {
	reqMap := make(map[string]int64)
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
	var result model.MessageDataResult
	json.Unmarshal(respBodyContent, &result)
	return &result, nil
}

func GetForwardMsg(id int64) (*model.ForwardMessageDataResult, error) {
	reqMap := make(map[string]int64)
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
	var result model.ForwardMessageDataResult
	json.Unmarshal(respBodyContent, &result)
	return &result, nil
}

func SendGroupForwardMsg(GroupId int64, messages []*model.MessageSegment) (*model.SendGroupForwardMessageDataResult, error) {
	reqMap := make(map[string]interface{})
	reqMap["group_id"] = GroupId
	reqMap["messages"] = messages
	requestBody, _ := json.Marshal(reqMap)
	println(string(requestBody))
	resp, err := http.Post(
		config.GetHttpUrl()+"/send_group_forward_msg",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	var result model.SendGroupForwardMessageDataResult
	json.Unmarshal(respBodyContent, &result)
	return &result, nil
}
