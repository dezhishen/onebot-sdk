package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/dezhishen/onebot-sdk/pkg/config"
)

// 群组踢人
func SetGroupKick(groupID, userID int, rejectAddRequest bool) error {
	reqMap := make(map[string]interface{})
	reqMap["user_id"] = userID
	reqMap["group_id"] = groupID
	reqMap["reject_add_request"] = rejectAddRequest
	requestBody, _ := json.Marshal(reqMap)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_group_kick",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

// 群组禁言
func SetGroupBan(groupID, userID, duration int) error {
	reqMap := make(map[string]int)
	reqMap["user_id"] = userID
	reqMap["group_id"] = groupID
	reqMap["duration"] = duration
	requestBody, _ := json.Marshal(reqMap)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_group_ban",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

// // 群组匿名用户禁言
// func SetGroupAnonymousBan(groupID, userID, duration int) error {
// 	reqMap := make(map[string]int)
// 	reqMap["user_id"] = userID
// 	reqMap["group_id"] = groupID
// 	reqMap["duration"] = duration
// 	requestBody, _ := json.Marshal(reqMap)
// 	_, err := http.Post(
// 		config.GetHttpUrl()+"/set_group_anonymous_ban",
// 		"application/json",
// 		bytes.NewBuffer(requestBody),
// 	)
// 	return err
// }
