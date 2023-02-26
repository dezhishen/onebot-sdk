package account

// import (
// 	"github.com/dezhishen/onebot-sdk/pkg/client"
// 	"github.com/dezhishen/onebot-sdk/pkg/config"
// 	"github.com/dezhishen/onebot-sdk/pkg/model"
// )

// type websocketOnebotApiAccountClient struct {
// 	*client.WebsocketCli
// }

// func NewWebsocketOnebotApiAccountClient(conf *config.OnebotConfig) (OnebotApiAccountClient, error) {
// 	return &websocketOnebotApiAccountClient{
// 		client.NewWebsocketCli(conf),
// 	}, nil
// }

// func (cli *websocketOnebotApiAccountClient) GetLoginInfo() (*model.AccountResult, error) {
// 	url := API_GET_LOGIN_INFO
// 	var result model.AccountResult
// 	if err := cli.SendWithOutParam(url); err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }

// func (cli *websocketOnebotApiAccountClient) GetStrangerInfo(userId int64, noCache bool) (*model.AccountResult, error) {
// 	req := make(map[string]interface{})
// 	req["user_id"] = userId
// 	req["no_cache"] = noCache
// 	url := API_GET_STRANGER_INFO
// 	var result model.AccountResult
// 	if err := cli.Send(url, req); err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }

// func (cli *websocketOnebotApiAccountClient) GetCookies(domin string) (*model.CokiesResult, error) {
// 	req := make(map[string]interface{})
// 	req["domain"] = domin
// 	url := API_GET_COOKIES
// 	var result model.CokiesResult
// 	if err := cli.Send(url, req); err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }

// func (cli *websocketOnebotApiAccountClient) GetCSRFToken() (*model.CSRFTokenResult, error) {
// 	url := API_GET_CSRF_TOKEN
// 	var result model.CSRFTokenResult
// 	if err := cli.SendWithOutParam(url); err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }

// func (cli *websocketOnebotApiAccountClient) GetCredentials(domin string) (*model.CredentialsResult, error) {
// 	req := make(map[string]interface{})
// 	req["domain"] = domin
// 	url := API_GET_CREDENTIALS
// 	var result model.CredentialsResult
// 	if err := cli.Send(url, req); err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }

// func (cli *websocketOnebotApiAccountClient) GetRecord(file string, outFormat string) (*model.FileResult, error) {
// 	req := make(map[string]interface{})
// 	req["file"] = file
// 	req["out_format"] = outFormat
// 	url := API_GET_RECORD
// 	var result model.FileResult
// 	if err := cli.Send(url, req); err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }

// func (cli *websocketOnebotApiAccountClient) GetImage(file string) (*model.FileResult, error) {
// 	req := make(map[string]interface{})
// 	req["file"] = file
// 	url := API_GET_IMAGE
// 	var result model.FileResult
// 	if err := cli.Send(url, req); err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }
