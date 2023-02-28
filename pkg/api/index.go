package api

import (
	"fmt"

	"github.com/dezhishen/onebot-sdk/pkg/api/account"
	"github.com/dezhishen/onebot-sdk/pkg/api/cqhttp"
	"github.com/dezhishen/onebot-sdk/pkg/api/file"
	"github.com/dezhishen/onebot-sdk/pkg/api/friend"
	"github.com/dezhishen/onebot-sdk/pkg/api/friendopt"
	"github.com/dezhishen/onebot-sdk/pkg/api/group"
	"github.com/dezhishen/onebot-sdk/pkg/api/groupopt"
	"github.com/dezhishen/onebot-sdk/pkg/api/message"
	"github.com/dezhishen/onebot-sdk/pkg/api/record"
	"github.com/dezhishen/onebot-sdk/pkg/api/request"
	"github.com/dezhishen/onebot-sdk/pkg/channel"
	"github.com/dezhishen/onebot-sdk/pkg/config"
)

type OnebotApiClient struct {
	account.OnebotApiAccountClient
	cqhttp.OnebotApiCqhttpClient
	file.OnebotApiFileClient
	friend.OnebotApiFriendClient
	friendopt.OnebotApiFriendOptClient
	group.OnebotApiGroupClient
	groupopt.OnebotApiGroupOptClient
	message.OnebotApiMessageClient
	record.OnebotApiRecordClient
	request.OnebotApiRequestClient
}

func NewOnebotApiClientByConfigPath(path string) (*OnebotApiClient, error) {
	conf, err := config.LoadConfig(path)
	if err != nil {
		return nil, err
	}
	return NewOnebotApiClientByConfig(conf)
}

func NewOnebotApiClientByConfig(config *config.OnebotConfig) (*OnebotApiClient, error) {
	if config == nil {
		return nil, fmt.Errorf("config is nil")
	}
	_channel, err := channel.NewApiChannel(config)
	if err != nil {
		return nil, err
	}
	accountCli, err := account.NewChannelApiAccountClient(_channel)
	if err != nil {
		return nil, err
	}
	cqhttpCli, err := cqhttp.NewChannelApiCqhttpClient(_channel)
	if err != nil {
		return nil, err
	}
	fileCli, err := file.NewChannelApiFileClient(_channel)
	if err != nil {
		return nil, err
	}
	friendCli, err := friend.NewChannelApiFriendClient(_channel)
	if err != nil {
		return nil, err
	}
	friendoptCli, err := friendopt.NewChannelApiFriendOptClient(_channel)
	if err != nil {
		return nil, err
	}
	groupCli, err := group.NewChannelApiGroupClient(_channel)
	if err != nil {
		return nil, err
	}
	groupoptCli, err := groupopt.NewChannelApiGroupOptClient(_channel)
	if err != nil {
		return nil, err
	}
	messageCli, err := message.NewChannelApiMessageClient(_channel)
	if err != nil {
		return nil, err
	}
	recordCli, err := record.NewChannelApiRecordClient(_channel)
	if err != nil {
		return nil, err
	}
	requestCli, err := request.NewChannelApiRequestClient(_channel)
	if err != nil {
		return nil, err
	}
	return &OnebotApiClient{
		accountCli,
		cqhttpCli,
		fileCli,
		friendCli,
		friendoptCli,
		groupCli,
		groupoptCli,
		messageCli,
		recordCli,
		requestCli,
	}, nil
}
