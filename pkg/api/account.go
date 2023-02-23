package api

import (
	"fmt"

	"github.com/dezhishen/onebot-sdk/pkg/client"
	"github.com/dezhishen/onebot-sdk/pkg/config"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type onebotApiAccountClient interface {
	GetLoginInfo() (*model.AccountResult, error)
	GetStrangerInfo(userId int64, noCache bool) (*model.AccountResult, error)
	GetCookies(domin string) (*model.CokiesResult, error)
	GetCSRFToken() (*model.CSRFTokenResult, error)
	GetCredentials(domin string) (*model.CredentialsResult, error)
	GetRecord(file string, outFormat string) (*model.FileResult, error)
	GetImage(file string) (*model.FileResult, error)
}

type httpOnebotApiAccountClient struct {
	*client.HttpCli
}

func NewOnebotApiAccountClient(conf *config.OnebotConfig) (onebotApiAccountClient, error) {
	if conf.Type == "http" {
		return &httpOnebotApiAccountClient{
			client.NewHttpCli(conf),
		}, nil
	}
	return nil, fmt.Errorf("not support type %s", conf.Type)
}

func (cli *httpOnebotApiAccountClient) GetLoginInfo() (*model.AccountResult, error) {
	url := "/get_login_info"
	var result model.AccountResult
	if err := cli.PostForResult(url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *httpOnebotApiAccountClient) GetStrangerInfo(userId int64, noCache bool) (*model.AccountResult, error) {
	req := make(map[string]interface{})
	req["user_id"] = userId
	req["no_cache"] = noCache
	url := "/get_stranger_info"
	var result model.AccountResult
	if err := cli.PostWithRequsetForResult(url, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *httpOnebotApiAccountClient) GetCookies(domin string) (*model.CokiesResult, error) {
	req := make(map[string]interface{})
	req["domain"] = domin
	url := "/get_cookies"
	var result model.CokiesResult
	if err := cli.PostWithRequsetForResult(url, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *httpOnebotApiAccountClient) GetCSRFToken() (*model.CSRFTokenResult, error) {
	url := "/get_csrf_token"
	var result model.CSRFTokenResult
	if err := cli.PostForResult(url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *httpOnebotApiAccountClient) GetCredentials(domin string) (*model.CredentialsResult, error) {
	req := make(map[string]interface{})
	req["domain"] = domin
	url := "/get_credentials"
	var result model.CredentialsResult
	if err := cli.PostWithRequsetForResult(url, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *httpOnebotApiAccountClient) GetRecord(file string, out_format string) (*model.FileResult, error) {
	req := make(map[string]interface{})
	req["file"] = file
	req["out_format"] = out_format
	url := "/get_record"
	var result model.FileResult
	if err := cli.PostWithRequsetForResult(url, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *httpOnebotApiAccountClient) GetImage(file string) (*model.FileResult, error) {
	req := make(map[string]interface{})
	req["file"] = file
	url := "/get_image"
	var result model.FileResult
	if err := cli.PostWithRequsetForResult(url, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
