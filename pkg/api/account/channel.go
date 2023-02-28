package account

import (
	"github.com/dezhishen/onebot-sdk/pkg/channel"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type ChannelApiAccountClient struct {
	channel.ApiChannel
}

func NewChannelApiAccountClient(channel channel.ApiChannel) (OnebotApiAccountClient, error) {
	return &ChannelApiAccountClient{
		channel,
	}, nil
}

func (cli *ChannelApiAccountClient) GetLoginInfo() (*model.AccountResult, error) {
	var result model.AccountResult
	if err := cli.PostForResult(API_GET_LOGIN_INFO, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *ChannelApiAccountClient) SetQQProfile(profile *model.QQProfile) error {
	return cli.PostByRequest(API_SET_QQ_PROFILE, profile)
}

func (cli *ChannelApiAccountClient) GetModelShow() (*model.ModelShowResult, error) {
	var result model.ModelShowResult
	if err := cli.PostForResult(API__GET_MODEL_SHOW, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *ChannelApiAccountClient) SetModelShow(model, modelshow string) error {
	req := make(map[string]interface{})
	req["model"] = model
	req["modelshow"] = modelshow
	return cli.PostByRequest(API__SET_MODEL_SHOW, req)
}

func (cli *ChannelApiAccountClient) GetOnlineClients() (*model.OnlineClientsResult, error) {
	var result model.OnlineClientsResult
	if err := cli.PostForResult(API_GET_ONLINE_CLIENTS, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
