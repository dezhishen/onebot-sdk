package cqhttp

/*
	type OnebotApiCqhttp interface {
		// 获取 CSRF Token
		// get_csrf_token
		GetCsrfToken() (*model.CsrfTokenResult, error)
		// 获取 QQ 相关接口凭证
		// get_credentials
		GetCredentials() (*model.CredentialsResult, error)
		// 获取版本信息
		// get_version_info
		GetVersionInfo() (*model.VersionInfoResult, error)
		// 获取状态
		// get_status
		GetStatus() (*model.StatusResult, error)
		// 重启 Go-CqHttp
		// set_restart
		// delay 要延迟的毫秒数, 如果默认情况下无法重启, 可以尝试设置延迟为 2000 左右
		SetRestart(delay int32) error
		// 清理缓存
		// clean_cache
		CleanCache() error
		// 重载事件过滤器
		// reload_event_filter
		// file 事件过滤器文件路径
		ReloadEventFilter(file string) error
		// 下载文件到缓存目录
		// download_file
		// url 文件链接
		// thread_count 线程数
		// headers 请求头
		DownloadFile(url string, threadCount int32, headers []string) (*model.DownloadFileResult, error)
		// 检查链接安全性
		// check_url_safely
		// url 链接
		CheckUrlSafely(url string) (*model.CheckUrlSafelyResult, error)
		// 获取中文分词 ( 隐藏 API )
		// .get_word_slices
		// content 内容
		HiddenGetWordSlices(content string) (*model.WordSlicesResult, error)
		// 对事件执行快速操作 ( 隐藏 API )
		// .handle_quick_operation
		// context 事件上下文
		// operation 操作
		HiddenHandleQuickOperation(context interface{}, operation interface{}) error
	}
*/
const (
	API_GET_COOKIES                   = "get_cookies"
	API_GET_CSRF_TOKEN                = "get_csrf_token"
	API_GET_CREDENTIALS               = "get_credentials"
	API_GET_VERSION_INFO              = "get_version_info"
	API_GET_STATUS                    = "get_status"
	API_SET_RESTART                   = "set_restart"
	API_CLEAN_CACHE                   = "clean_cache"
	API_RELOAD_EVENT_FILTER           = "reload_event_filter"
	API_DOWNLOAD_FILE                 = "download_file"
	API_CHECK_URL_SAFELY              = "check_url_safely"
	API_HIDDEN_GET_WORD_SLICES        = ".get_word_slices"
	API_HIDDEN_HANDLE_QUICK_OPERATION = ".handle_quick_operation"
)
