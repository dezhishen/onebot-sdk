package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/dezhishen/onebot-sdk/pkg/config"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

func SendLike(userID int, times int) error {
	reqMap := make(map[string]int)
	reqMap["user_id"] = userID
	reqMap["times"] = times
	requestBody, _ := json.Marshal(reqMap)
	_, err := http.Post(
		config.GetHttpUrl()+"/send_like",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

//处理加好友请求
func SetFriendAddRequest(flag string, approve bool, remark string) error {
	reqMap := make(map[string]interface{})
	reqMap["flag"] = flag
	reqMap["approve"] = approve
	reqMap["remark"] = remark
	requestBody, _ := json.Marshal(reqMap)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_friend_add_request",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

func GetFriendList() ([]model.Account, error) {
	resp, err := http.Post(
		config.GetHttpUrl()+"/get_friend_list",
		"application/json",
		nil,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	var result model.FriendListResult
	json.Unmarshal(respBodyContent, &result)
	return result.Data, nil
}
