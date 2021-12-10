package api

import (
	"github.com/dezhishen/onebot-sdk/pkg/cli"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

func GetLoginInfo() (*model.Account, error) {
	url := "/get_login_info"
	var result model.AccountResult
	if err := cli.PostForResult(url, &result); err != nil {
		return nil, err
	}
	return result.Data, nil
}

func GetStrangerInfo(userId int, noCache bool) (*model.Account, error) {
	req := make(map[string]interface{})
	req["user_id"] = userId
	req["no_cache"] = noCache
	url := "/get_stranger_info"
	var result model.AccountResult
	if err := cli.PostWithRequsetForResult(url, req, &result); err != nil {
		return nil, err
	}
	return result.Data, nil
}

func GetCookies(domin string) (*model.Cokies, error) {
	req := make(map[string]interface{})
	req["domain"] = domin
	url := "/get_cookies"
	var result model.CokiesResult
	if err := cli.PostWithRequsetForResult(url, req, &result); err != nil {
		return nil, err
	}
	return result.Data, nil
}

func GetCSRFToken() (*model.CSRFToken, error) {
	url := "/get_csrf_token"
	var result model.CSRFTokenResult
	if err := cli.PostForResult(url, &result); err != nil {
		return nil, err
	}
	return result.Data, nil
}

func GetCredentials(domin string) (*model.Credentials, error) {
	req := make(map[string]interface{})
	req["domain"] = domin
	url := "/get_credentials"
	var result model.CredentialsResult
	if err := cli.PostWithRequsetForResult(url, req, &result); err != nil {
		return nil, err
	}
	return result.Data, nil
}

func GetRecord(file string, out_format string) (*model.File, error) {
	req := make(map[string]interface{})
	req["file"] = file
	req["out_format"] = out_format
	url := "/get_record"
	var result model.FileResult
	if err := cli.PostWithRequsetForResult(url, req, &result); err != nil {
		return nil, err
	}
	return result.Data, nil
}

func GetImage(file string) (*model.File, error) {
	req := make(map[string]interface{})
	req["file"] = file
	url := "/get_image"
	var result model.FileResult
	if err := cli.PostWithRequsetForResult(url, req, &result); err != nil {
		return nil, err
	}
	return result.Data, nil
}
