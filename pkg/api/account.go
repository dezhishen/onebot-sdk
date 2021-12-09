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

func GetCookies(domin string) (string, error) {
	reqMap := make(map[string]interface{})
	reqMap["domain"] = domin
	requestBody, _ := json.Marshal(reqMap)
	resp, err := http.Post(
		config.GetHttpUrl()+"/get_cookies",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	var result model.Credentials
	json.Unmarshal(respBodyContent, &result)
	return result.Cookies, nil
}

func GetCSRFToken() (int32, error) {
	resp, err := http.Post(
		config.GetHttpUrl()+"/get_csrf_token ",
		"application/json",
		nil,
	)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	var result model.Credentials
	json.Unmarshal(respBodyContent, &result)
	return result.CSRFToken, nil
}

func GetCredentials(domin string) (*model.Credentials, error) {
	reqMap := make(map[string]interface{})
	reqMap["domain"] = domin
	requestBody, _ := json.Marshal(reqMap)
	resp, err := http.Post(
		config.GetHttpUrl()+"/get_credentials",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	var result model.Credentials
	json.Unmarshal(respBodyContent, &result)
	return &result, nil
}

func GetRecord(file string, out_format string) (map[string]interface{}, error) {
	reqMap := make(map[string]interface{})
	reqMap["file"] = file
	reqMap["out_format"] = out_format
	requestBody, _ := json.Marshal(reqMap)
	resp, err := http.Post(
		config.GetHttpUrl()+"/get_record",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	result := make(map[string]interface{})
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)

	json.Unmarshal(respBodyContent, &result)
	return result, nil
}

func GetImage(file string) (map[string]interface{}, error) {
	reqMap := make(map[string]interface{})
	reqMap["file"] = file
	requestBody, _ := json.Marshal(reqMap)
	resp, err := http.Post(
		config.GetHttpUrl()+"/get_image",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	result := make(map[string]interface{})
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)

	json.Unmarshal(respBodyContent, &result)
	return result, nil
}
