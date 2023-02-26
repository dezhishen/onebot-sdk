package api

// import (
// 	"fmt"

// 	"github.com/dezhishen/onebot-sdk/pkg/api/account"

// 	"github.com/dezhishen/onebot-sdk/pkg/config"
// )

// type OnebotApiClient struct {
// 	account.OnebotApiAccountClient
// 	onebotApiGroupClient
// 	onebotApiFriendClient
// 	onebotApiMessageClient
// 	onebotApiSystemClient
// }

// func NewOnebotApiClientByConfigPath(path string) (*OnebotApiClient, error) {
// 	conf, err := config.LoadConfig(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return NewOnebotApiClientByConfig(conf)
// }

// func NewOnebotApiClientByConfig(config *config.OnebotConfig) (*OnebotApiClient, error) {
// 	if config == nil {
// 		return nil, fmt.Errorf("config is nil")
// 	}
// 	if config.Endpoint == "" {
// 		return nil, fmt.Errorf("endpoint is empty")
// 	}
// 	if config.Type == "" {
// 		return nil, fmt.Errorf("type is empty")
// 	}
// 	accountCli, err := NewOnebotApiAccountClient(config)
// 	if err != nil {
// 		return nil, err
// 	}
// 	groupCli, err := NewOnebotApiGroupClient(config)
// 	if err != nil {
// 		return nil, err
// 	}
// 	friendCli, err := NewOnebotApiFriendClient(config)
// 	if err != nil {
// 		return nil, err
// 	}
// 	messageCli, err := NewOnebotApiMessageClient(config)
// 	if err != nil {
// 		return nil, err
// 	}
// 	systemCli, err := NewOnebotApiSystemClient(config)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &OnebotApiClient{
// 		onebotApiAccountClient: accountCli,
// 		onebotApiGroupClient:   groupCli,
// 		onebotApiFriendClient:  friendCli,
// 		onebotApiMessageClient: messageCli,
// 		onebotApiSystemClient:  systemCli,
// 	}, nil
// }
