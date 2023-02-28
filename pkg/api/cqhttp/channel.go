package cqhttp

import (
	"github.com/dezhishen/onebot-sdk/pkg/channel"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type ChannelApiCqhttpClient struct {
	channel.ApiChannel
}

func NewChannelApiCqhttpClient(channel channel.ApiChannel) (OnebotApiCqhttpClient, error) {
	return &ChannelApiCqhttpClient{
		channel,
	}, nil
}

// 获取 Cookies
// get_cookies
// domain 指定域名
func (cli *ChannelApiCqhttpClient) GetCookies(domain string) (*model.CookiesResult, error) {
	var result model.CookiesResult
	if err := cli.PostByRequestForResult(API_GET_COOKIES, map[string]interface{}{
		"domain": domain,
	}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 获取 CSRF Token
// get_csrf_token
func (cli *ChannelApiCqhttpClient) GetCsrfToken() (*model.CsrfTokenResult, error) {
	var result model.CsrfTokenResult
	if err := cli.PostForResult(API_GET_CSRF_TOKEN, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 获取 QQ 相关接口凭证
// get_credentials
func (cli *ChannelApiCqhttpClient) GetCredentials() (*model.CredentialsResult, error) {
	var result model.CredentialsResult
	if err := cli.PostForResult(API_GET_CREDENTIALS, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 获取版本信息
// get_version_info
func (cli *ChannelApiCqhttpClient) GetVersionInfo() (*model.VersionInfoResult, error) {
	var result model.VersionInfoResult
	if err := cli.PostForResult(API_GET_VERSION_INFO, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 获取状态
// get_status
func (cli *ChannelApiCqhttpClient) GetStatus() (*model.StatusResult, error) {
	var result model.StatusResult
	if err := cli.PostForResult(API_GET_STATUS, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 重启 Go-CqHttp
// set_restart
// delay 要延迟的毫秒数, 如果默认情况下无法重启, 可以尝试设置延迟为 2000 左右
func (cli *ChannelApiCqhttpClient) SetRestart(delay int32) error {
	return cli.PostByRequest(API_SET_RESTART, map[string]interface{}{
		"delay": delay,
	})
}

// 清理缓存
// clean_cache
func (cli *ChannelApiCqhttpClient) CleanCache() error {
	return cli.Post(API_CLEAN_CACHE)
}

// 重载事件过滤器
// reload_event_filter
// file 事件过滤器文件路径
func (cli *ChannelApiCqhttpClient) ReloadEventFilter(file string) error {
	return cli.PostByRequest(API_RELOAD_EVENT_FILTER, map[string]interface{}{
		"file": file,
	})
}

// 下载文件到缓存目录
// download_file
// url 文件链接
// thread_count 线程数
// headers 请求头
func (cli *ChannelApiCqhttpClient) DownloadFile(url string, threadCount int32, headers []string) (*model.DownloadFileResult, error) {
	var result model.DownloadFileResult
	if err := cli.PostByRequestForResult(API_DOWNLOAD_FILE, map[string]interface{}{
		"url":          url,
		"thread_count": threadCount,
		"headers":      headers,
	}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 检查链接安全性
// check_url_safely
// url 链接
func (cli *ChannelApiCqhttpClient) CheckUrlSafely(url string) (*model.CheckUrlSafelyResult, error) {
	var result model.CheckUrlSafelyResult
	if err := cli.PostByRequestForResult(API_CHECK_URL_SAFELY, map[string]interface{}{
		"url": url,
	}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 获取中文分词 ( 隐藏 API )
// .get_word_slices
// content 内容
func (cli *ChannelApiCqhttpClient) HiddenGetWordSlices(content string) (*model.WordSlicesResult, error) {
	var result model.WordSlicesResult
	if err := cli.PostByRequestForResult(API_HIDDEN_GET_WORD_SLICES, map[string]interface{}{
		"content": content,
	}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 对事件执行快速操作 ( 隐藏 API )
// .handle_quick_operation
// context 事件上下文
// operation 操作
func (cli *ChannelApiCqhttpClient) HiddenHandleQuickOperation(context interface{}, operation interface{}) error {
	return cli.PostByRequest(
		API_HIDDEN_HANDLE_QUICK_OPERATION,
		map[string]interface{}{
			"context":   context,
			"operation": operation,
		},
	)
}
