package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/dezhishen/onebot-sdk/pkg/config"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

func CleanCache() error {
	_, err := http.Post(
		config.GetHttpUrl()+"/clean_cache",
		"application/json",
		nil,
	)
	return err
}

func SetRestart(delay int) error {
	reqMap := make(map[string]interface{})
	reqMap["delay"] = delay
	requestBody, _ := json.Marshal(reqMap)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_restart",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

func GetVersionInfo() (map[string]interface{}, error) {
	resp, err := http.Post(
		config.GetHttpUrl()+"/get_version_info",
		"application/json",
		nil,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	var result model.MapInfoResult
	json.Unmarshal(respBodyContent, &result)
	return result.Data, nil
}

func GetStatus() (map[string]interface{}, error) {
	resp, err := http.Post(
		config.GetHttpUrl()+"/get_status",
		"application/json",
		nil,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	var result model.MapInfoResult
	json.Unmarshal(respBodyContent, &result)
	return result.Data, nil
}

func CanSendImage() (map[string]interface{}, error) {
	resp, err := http.Post(
		config.GetHttpUrl()+"/can_send_image",
		"application/json",
		nil,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	var result model.MapInfoResult
	json.Unmarshal(respBodyContent, &result)
	return result.Data, nil
}

func CanSendRecord() (map[string]interface{}, error) {
	resp, err := http.Post(
		config.GetHttpUrl()+"/can_send_record",
		"application/json",
		nil,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	var result model.MapInfoResult
	json.Unmarshal(respBodyContent, &result)
	return result.Data, nil
}
