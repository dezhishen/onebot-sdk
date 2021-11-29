package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/dezhishen/onebot-sdk/pkg/config"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

func GetLoginInfo() (*model.Account, error) {
	resp, err := http.Post(
		config.GetHttpUrl()+"/get_login_info",
		"application/json",
		nil,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	var result model.AccountResult
	json.Unmarshal(respBodyContent, &result)
	return result.Data, nil
}

func GetStrangerInfo(userId int, noCache bool) (*model.Account, error) {
	reqMap := make(map[string]interface{})
	reqMap["user_id"] = userId
	reqMap["no_cache"] = noCache
	requestBody, _ := json.Marshal(reqMap)
	resp, err := http.Post(
		config.GetHttpUrl()+"/get_stranger_info",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	var result model.AccountResult
	json.Unmarshal(respBodyContent, &result)
	return result.Data, nil
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
