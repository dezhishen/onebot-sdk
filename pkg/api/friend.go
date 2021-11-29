package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/dezhishen/onebot-sdk/pkg/config"
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
