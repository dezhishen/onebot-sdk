package account

import (
	"github.com/dezhishen/onebot-sdk/pkg/client"
	"github.com/dezhishen/onebot-sdk/pkg/config"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type httpOnebotApiAccountClient struct {
	*client.HttpCli
}

func NewHttpOnebotApiAccountClient(conf *config.OnebotConfig) (OnebotApiAccountClient, error) {
	return &httpOnebotApiAccountClient{
		client.NewHttpCli(conf),
	}, nil
}

func (cli *httpOnebotApiAccountClient) GetLoginInfo() (*model.AccountResult, error) {
	var result model.AccountResult
	if err := cli.PostForResult(API_GET_LOGIN_INFO, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *httpOnebotApiAccountClient) SetQQProfile(profile *model.QQProfile) error {
	return cli.PostWithRequest(API_SET_QQ_PROFILE, profile)
}

func (cli *httpOnebotApiAccountClient) GetModelShow() (*model.ModelShowResult, error) {
	var result model.ModelShowResult
	if err := cli.PostForResult(API__GET_MODEL_SHOW, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *httpOnebotApiAccountClient) SetModelShow(model, modelshow string) error {
	req := make(map[string]interface{})
	req["model"] = model
	req["modelshow"] = modelshow
	return cli.PostWithRequest(API__SET_MODEL_SHOW, req)
}

func (cli *httpOnebotApiAccountClient) GetOnlineClients() (*model.OnlineClientsResult, error) {
	var result model.OnlineClientsResult
	if err := cli.PostForResult(API_GET_ONLINE_CLIENTS, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
