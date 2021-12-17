package api

import (
	"github.com/dezhishen/onebot-sdk/pkg/cli"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

func CleanCache() error {
	return cli.Post("/clean_cache")
}

func SetRestart(delay int) error {
	reqMap := make(map[string]interface{})
	reqMap["delay"] = delay

	return cli.PostWithRequest("/set_restart", reqMap)
}

func GetVersionInfo() (map[string]interface{}, error) {
	var result model.MapInfoResult
	if err := cli.PostForResult("/get_version_info", &result); err != nil {
		return nil, err
	}
	return result.Data, nil
}

func GetStatus() (map[string]interface{}, error) {
	var result model.MapInfoResult
	if err := cli.PostForResult("/get_status", &result); err != nil {
		return nil, err
	}
	return result.Data, nil
}

func CanSendImage() (bool, error) {
	var result model.BoolYesOfResult
	if err := cli.PostForResult("/can_send_image", &result); err != nil {
		return false, err
	}
	return result.Data.Yes, nil
}

func CanSendRecord() (bool, error) {
	var result model.BoolYesOfResult
	if err := cli.PostForResult("/can_send_record", &result); err != nil {
		return false, err
	}
	return result.Data.Yes, nil
}
