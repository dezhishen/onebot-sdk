package account

import (
	"github.com/dezhishen/onebot-sdk/pkg/client"
	"github.com/dezhishen/onebot-sdk/pkg/config"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type websocketOnebotApiAccountClient struct {
	*client.WebsocketCli
}

func NewWebsocketOnebotApiAccountClient(conf *config.OnebotConfig) (OnebotApiAccountClient, error) {
	return &websocketOnebotApiAccountClient{
		client.NewWebsocketCli(conf),
	}, nil
}

func (cli *websocketOnebotApiAccountClient) GetLoginInfo() (*model.AccountResult, error) {
	var result model.AccountResult
	if err := cli.SendWithOutParamForResult(API_GET_LOGIN_INFO, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *websocketOnebotApiAccountClient) SetQQProfile(profile *model.QQProfile) error {
	return cli.Send(API_SET_QQ_PROFILE, profile)
}

func (cli *websocketOnebotApiAccountClient) GetModelShow() (*model.ModelShowResult, error) {
	var result model.ModelShowResult
	if err := cli.SendWithOutParamForResult(API__GET_MODEL_SHOW, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *websocketOnebotApiAccountClient) SetModelShow(model, modelshow string) error {
	req := make(map[string]interface{})
	req["model"] = model
	req["modelshow"] = modelshow
	return cli.Send(API__SET_MODEL_SHOW, req)
}

func (cli *websocketOnebotApiAccountClient) GetOnlineClients() (*model.OnlineClientsResult, error) {
	var result model.OnlineClientsResult
	if err := cli.SendWithOutParamForResult(API_GET_ONLINE_CLIENTS, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
