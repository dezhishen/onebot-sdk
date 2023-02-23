package api

import (
	"github.com/dezhishen/onebot-sdk/pkg/client"
	"github.com/dezhishen/onebot-sdk/pkg/config"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type onebotApiSystemClient interface {
	CleanCache() error
	SetRestart(delay int64) error
	GetVersionInfo() (*model.VersionInfoResult, error)

	GetStatus() (*model.StatusInfoResult, error)

	CanSendImage() (*model.BoolYesOfResult, error)
	CanSendRecord() (*model.BoolYesOfResult, error)
}
type httpOnebotApiSystemClient struct {
	*client.HttpCli
}

func NewOnebotApiSystemClient(conf *config.OnebotConfig) (onebotApiSystemClient, error) {
	return &httpOnebotApiSystemClient{
		client.NewHttpCli(conf),
	}, nil
}

func (cli *httpOnebotApiSystemClient) CleanCache() error {
	return cli.Post("/clean_cache")
}

func (cli *httpOnebotApiSystemClient) SetRestart(delay int64) error {
	reqMap := make(map[string]interface{})
	reqMap["delay"] = delay

	return cli.PostWithRequest("/set_restart", reqMap)
}

func (cli *httpOnebotApiSystemClient) GetVersionInfo() (*model.VersionInfoResult, error) {
	var result model.VersionInfoResult
	if err := cli.PostForResult("/get_version_info", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *httpOnebotApiSystemClient) GetStatus() (*model.StatusInfoResult, error) {
	var result model.StatusInfoResult
	if err := cli.PostForResult("/get_status", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *httpOnebotApiSystemClient) CanSendImage() (*model.BoolYesOfResult, error) {
	var result model.BoolYesOfResult
	if err := cli.PostForResult("/can_send_image", &result); err != nil {
		return &result, err
	}
	return &result, nil
}

func (cli *httpOnebotApiSystemClient) CanSendRecord() (*model.BoolYesOfResult, error) {
	var result model.BoolYesOfResult
	if err := cli.PostForResult("/can_send_record", &result); err != nil {
		return &result, err
	}
	return &result, nil
}
