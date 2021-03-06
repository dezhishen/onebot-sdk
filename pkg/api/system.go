package api

import (
	"github.com/dezhishen/onebot-sdk/pkg/cli"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

func CleanCache() error {
	return cli.Post("/clean_cache")
}

func SetRestart(delay int64) error {
	reqMap := make(map[string]interface{})
	reqMap["delay"] = delay

	return cli.PostWithRequest("/set_restart", reqMap)
}

func GetVersionInfo() (*model.VersionInfoResult, error) {
	var result model.VersionInfoResult
	if err := cli.PostForResult("/get_version_info", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func GetStatus() (*model.StatusInfoResult, error) {
	var result model.StatusInfoResult
	if err := cli.PostForResult("/get_status", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func CanSendImage() (*model.BoolYesOfResult, error) {
	var result model.BoolYesOfResult
	if err := cli.PostForResult("/can_send_image", &result); err != nil {
		return &result, err
	}
	return &result, nil
}

func CanSendRecord() (*model.BoolYesOfResult, error) {
	var result model.BoolYesOfResult
	if err := cli.PostForResult("/can_send_record", &result); err != nil {
		return &result, err
	}
	return &result, nil
}
