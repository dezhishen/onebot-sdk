package account

// import (
// 	"fmt"

// 	"github.com/dezhishen/onebot-sdk/pkg/client"
// 	"github.com/dezhishen/onebot-sdk/pkg/config"
// 	"github.com/dezhishen/onebot-sdk/pkg/model"
// )

// type httpOnebotApiAccountClient struct {
// 	*client.HttpCli
// }

// func NewOnebotApiAccountClient(conf *config.OnebotConfig) (OnebotApiAccountClient, error) {
// 	if conf.Type == "http" {
// 		return &httpOnebotApiAccountClient{
// 			client.NewHttpCli(conf),
// 		}, nil
// 	}
// 	return nil, fmt.Errorf("not support type %s", conf.Type)
// }

// func (cli *httpOnebotApiAccountClient) GetLoginInfo() (*model.AccountResult, error) {
// 	url := API_GET_LOGIN_INFO
// 	var result model.AccountResult
// 	if err := cli.PostForResult(url, &result); err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }

// func (cli *httpOnebotApiAccountClient) GetStrangerInfo(userId int64, noCache bool) (*model.AccountResult, error) {
// 	req := make(map[string]interface{})
// 	req["user_id"] = userId
// 	req["no_cache"] = noCache
// 	url := API_GET_STRANGER_INFO // "/get_stranger_info"
// 	var result model.AccountResult
// 	if err := cli.PostWithRequsetForResult(url, req, &result); err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }

// func (cli *httpOnebotApiAccountClient) GetCookies(domin string) (*model.CokiesResult, error) {
// 	req := make(map[string]interface{})
// 	req["domain"] = domin
// 	url := API_GET_COOKIES //  "/get_cookies"
// 	var result model.CokiesResult
// 	if err := cli.PostWithRequsetForResult(url, req, &result); err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }

// func (cli *httpOnebotApiAccountClient) GetCSRFToken() (*model.CSRFTokenResult, error) {
// 	url := API_GET_CSRF_TOKEN //"/get_csrf_token"
// 	var result model.CSRFTokenResult
// 	if err := cli.PostForResult(url, &result); err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }

// func (cli *httpOnebotApiAccountClient) GetCredentials(domin string) (*model.CredentialsResult, error) {
// 	req := make(map[string]interface{})
// 	req["domain"] = domin
// 	url := API_GET_CREDENTIALS // "/get_credentials"
// 	var result model.CredentialsResult
// 	if err := cli.PostWithRequsetForResult(url, req, &result); err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }

// func (cli *httpOnebotApiAccountClient) GetRecord(file string, out_format string) (*model.FileResult, error) {
// 	req := make(map[string]interface{})
// 	req["file"] = file
// 	req["out_format"] = out_format
// 	url := API_GET_RECORD //"/get_record"
// 	var result model.FileResult
// 	if err := cli.PostWithRequsetForResult(url, req, &result); err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }

// func (cli *httpOnebotApiAccountClient) GetImage(file string) (*model.FileResult, error) {
// 	req := make(map[string]interface{})
// 	req["file"] = file
// 	url := API_GET_IMAGE //"get_image"
// 	var result model.FileResult
// 	if err := cli.PostWithRequsetForResult(url, req, &result); err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }
