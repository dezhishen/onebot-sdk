package model

type Cookies struct {
	// cookies
	Cookies string `json:"cookies"`
}

func (c *Cookies) ToGRPC() *CookiesGRPC {
	return &CookiesGRPC{
		Cookies: c.Cookies,
	}
}

func (C *CookiesGRPC) ToStruct() *Cookies {
	return &Cookies{
		Cookies: C.Cookies,
	}
}

type CookiesResult struct {
	Data    *Cookies `json:"data"`
	Retcode int64    `json:"retcode"`
	Status  string   `json:"status"`
	Msg     string   `json:"msg"`
	Wording string   `json:"wording"`
}

func (c *CookiesResult) ToGRPC() *CookiesResultGRPC {
	result := &CookiesResultGRPC{
		Retcode: c.Retcode,
		Status:  c.Status,
		Msg:     c.Msg,
		Wording: c.Wording,
	}
	if c.Data != nil {
		result.Data = c.Data.ToGRPC()
	}
	return result
}

func (C *CookiesResultGRPC) ToStruct() *CookiesResult {
	result := &CookiesResult{
		Retcode: C.Retcode,
		Status:  C.Status,
		Msg:     C.Msg,
		Wording: C.Wording,
	}
	if C.Data != nil {
		result.Data = C.Data.ToStruct()
	}
	return result
}

type CsrfToken struct {
	Token int32 `json:"token"`
}

func (c *CsrfToken) ToGRPC() *CsrfTokenGRPC {
	return &CsrfTokenGRPC{
		Token: c.Token,
	}
}

func (C *CsrfTokenGRPC) ToStruct() *CsrfToken {
	return &CsrfToken{
		Token: C.Token,
	}
}

type CsrfTokenResult struct {
	Data    *CsrfToken `json:"data"`
	Retcode int64      `json:"retcode"`
	Status  string     `json:"status"`
	Msg     string     `json:"msg"`
	Wording string     `json:"wording"`
}

func (c *CsrfTokenResult) ToGRPC() *CsrfTokenResultGRPC {
	result := &CsrfTokenResultGRPC{
		Retcode: c.Retcode,
		Status:  c.Status,
		Msg:     c.Msg,
		Wording: c.Wording,
	}
	if c.Data != nil {
		result.Data = c.Data.ToGRPC()
	}
	return result
}

func (c *CsrfTokenResultGRPC) ToStruct() *CsrfTokenResult {
	result := &CsrfTokenResult{
		Retcode: c.Retcode,
		Status:  c.Status,
		Msg:     c.Msg,
		Wording: c.Wording,
	}
	if c.Data != nil {
		result.Data = c.Data.ToStruct()
	}
	return result
}

type Credentials struct {
	Cookies string `json:"cookies"`
	Token   int32  `json:"token"`
}

func (c *Credentials) ToGRPC() *CredentialsGRPC {
	return &CredentialsGRPC{
		Cookies: c.Cookies,
		Token:   c.Token,
	}
}

func (C *CredentialsGRPC) ToStruct() *Credentials {
	return &Credentials{
		Cookies: C.Cookies,
		Token:   C.Token,
	}
}

type CredentialsResult struct {
	Data    *Credentials `json:"data"`
	Retcode int64        `json:"retcode"`
	Status  string       `json:"status"`
	Msg     string       `json:"msg"`
	Wording string       `json:"wording"`
}

func (c *CredentialsResult) ToGRPC() *CredentialsResultGRPC {
	result := &CredentialsResultGRPC{
		Retcode: c.Retcode,
		Status:  c.Status,
		Msg:     c.Msg,
		Wording: c.Wording,
	}
	if c.Data != nil {
		result.Data = c.Data.ToGRPC()
	}
	return result
}

func (c *CredentialsResultGRPC) ToStruct() *CredentialsResult {
	result := &CredentialsResult{
		Retcode: c.Retcode,
		Status:  c.Status,
		Msg:     c.Msg,
		Wording: c.Wording,
	}
	if c.Data != nil {
		result.Data = c.Data.ToStruct()
	}
	return result
}

type VersionInfo struct {
	// app_name	string	go-cqhttp	应用标识, 如 go-cqhttp 固定值
	AppName string `json:"app_name"`
	// app_version	string		应用版本, 如 v0.9.40-fix4
	AppVersion string `json:"app_version"`
	// app_full_name	string		应用完整名称
	AppFullName string `json:"app_full_name"`
	// protocol_version	string	v11	OneBot 标准版本 固定值
	ProtocolVersion string `json:"protocol_version"`
	// coolq_edition	string	pro	原Coolq版本 固定值
	CoolqEdition string `json:"coolq_edition"`
	// coolq_directory	string
	CoolqDirectory string `json:"coolq_directory"`
	// go-cqhttp	bool	true	是否为go-cqhttp 固定值
	GoCqhttp bool `json:"go-cqhttp"`
	// plugin_version	string	4.15.0	固定值
	PluginVersion string `json:"plugin_version"`
	// plugin_build_number	int	99	固定值
	PluginBuildNumber int32 `json:"plugin_build_number"`
	// plugin_build_configuration	string	release	固定值
	PluginBuildConfiguration string `json:"plugin_build_configuration"`
	// runtime_version	string
	RuntimeVersion string `json:"runtime_version"`
	// runtime_os	string
	RuntimeOs string `json:"runtime_os"`
	// version	string		应用版本, 如 v0.9.40-fix4
	Version string `json:"version"`
	// protocol	int	0/1/2/3/-1	当前登陆使用协议类型
	Protocol int32 `json:"protocol"`
}

func (c *VersionInfo) ToGRPC() *VersionInfoGRPC {
	return &VersionInfoGRPC{
		AppName:                  c.AppName,
		AppVersion:               c.AppVersion,
		AppFullName:              c.AppFullName,
		ProtocolVersion:          c.ProtocolVersion,
		CoolqEdition:             c.CoolqEdition,
		CoolqDirectory:           c.CoolqDirectory,
		GoCqhttp:                 c.GoCqhttp,
		PluginVersion:            c.PluginVersion,
		PluginBuildNumber:        c.PluginBuildNumber,
		PluginBuildConfiguration: c.PluginBuildConfiguration,
		RuntimeVersion:           c.RuntimeVersion,
		RuntimeOs:                c.RuntimeOs,
		Version:                  c.Version,
		Protocol:                 c.Protocol,
	}
}

func (c *VersionInfoGRPC) ToStruct() *VersionInfo {
	return &VersionInfo{
		AppName:                  c.AppName,
		AppVersion:               c.AppVersion,
		AppFullName:              c.AppFullName,
		ProtocolVersion:          c.ProtocolVersion,
		CoolqEdition:             c.CoolqEdition,
		CoolqDirectory:           c.CoolqDirectory,
		GoCqhttp:                 c.GoCqhttp,
		PluginVersion:            c.PluginVersion,
		PluginBuildNumber:        c.PluginBuildNumber,
		PluginBuildConfiguration: c.PluginBuildConfiguration,
		RuntimeVersion:           c.RuntimeVersion,
		RuntimeOs:                c.RuntimeOs,
		Version:                  c.Version,
		Protocol:                 c.Protocol,
	}
}

type VersionInfoResult struct {
	Data    *VersionInfo `json:"data"`
	Retcode int64        `json:"ret_code"`
	Status  string       `json:"status"` //
	Msg     string       `json:"msg"`
	Wording string       `json:"wording"`
}

func (c *VersionInfoResult) ToGRPC() *VersionInfoResultGRPC {
	result := &VersionInfoResultGRPC{
		Retcode: c.Retcode,
		Status:  c.Status,
		Msg:     c.Msg,
		Wording: c.Wording,
	}
	if c.Data != nil {
		result.Data = c.Data.ToGRPC()
	}
	return result
}

func (c *VersionInfoResultGRPC) ToStruct() *VersionInfoResult {
	result := &VersionInfoResult{
		Retcode: c.Retcode,
		Status:  c.Status,
		Msg:     c.Msg,
		Wording: c.Wording,
	}
	if c.Data != nil {
		result.Data = c.Data.ToStruct()
	}
	return result
}

type Statistics struct {
	// PacketReceived	uint64	收到的数据包总数
	PacketReceived uint64 `json:"packet_received"`
	// PacketSent	uint64	发送的数据包总数
	PacketSent uint64 `json:"packet_sent"`
	// PacketLost	uint32	数据包丢失总数
	PacketLost uint32 `json:"packet_lost"`
	// MessageReceived	uint64	接受信息总数
	MessageReceived uint64 `json:"message_received"`
	// MessageSent	uint64	发送信息总数
	MessageSent uint64 `json:"message_sent"`
	// DisconnectTimes	uint32	TCP 链接断开次数
	DisconnectTimes uint32 `json:"disconnect_times"`
	// LostTimes	uint32	账号掉线次数
	LostTimes uint32 `json:"lost_times"`
	// LastMessageTime	int64	最后一条消息时间
	LastMessageTime int64 `json:"last_message_time"`
}

func (c *Statistics) ToGRPC() *StatisticsGRPC {
	return &StatisticsGRPC{
		PacketReceived:  c.PacketReceived,
		PacketSent:      c.PacketSent,
		PacketLost:      c.PacketLost,
		MessageReceived: c.MessageReceived,
		MessageSent:     c.MessageSent,
		DisconnectTimes: c.DisconnectTimes,
		LostTimes:       c.LostTimes,
		LastMessageTime: c.LastMessageTime,
	}
}

func (c *StatisticsGRPC) ToStruct() *Statistics {
	return &Statistics{
		PacketReceived:  c.PacketReceived,
		PacketSent:      c.PacketSent,
		PacketLost:      c.PacketLost,
		MessageReceived: c.MessageReceived,
		MessageSent:     c.MessageSent,
		DisconnectTimes: c.DisconnectTimes,
		LostTimes:       c.LostTimes,
		LastMessageTime: c.LastMessageTime,
	}
}

type Status struct {
	// app_initialized	bool	原 CQHTTP 字段, 恒定为 true
	AppInitialized bool `json:"app_initialized"`
	// app_enabled	bool	原 CQHTTP 字段, 恒定为 true
	AppEnabled bool `json:"app_enabled"`
	// plugins_good	bool	原 CQHTTP 字段, 恒定为 true
	PluginsGood bool `json:"plugins_good"`
	// app_good	bool	原 CQHTTP 字段, 恒定为 true
	AppGood bool `json:"app_good"`
	// online	bool	表示BOT是否在线
	Online bool `json:"online"`
	// good	bool	同 online
	Good bool `json:"good"`
	// stat	Statistics	运行统计
	Stat *Statistics `json:"stat"`
}

func (c *Status) ToGRPC() *StatusGRPC {
	result := &StatusGRPC{
		AppInitialized: c.AppInitialized,
		AppEnabled:     c.AppEnabled,
		PluginsGood:    c.PluginsGood,
		AppGood:        c.AppGood,
		Online:         c.Online,
		Good:           c.Good,
	}
	if c.Stat != nil {
		result.Stat = c.Stat.ToGRPC()
	}
	return result
}

func (c *StatusGRPC) ToStruct() *Status {
	result := &Status{
		AppInitialized: c.AppInitialized,
		AppEnabled:     c.AppEnabled,
		PluginsGood:    c.PluginsGood,
		AppGood:        c.AppGood,
		Online:         c.Online,
		Good:           c.Good,
	}
	if c.Stat != nil {
		result.Stat = c.Stat.ToStruct()
	}
	return result
}

type StatusResult struct {
	Data    *Status `json:"data"`
	Retcode int64   `json:"ret_code"`
	Status  string  `json:"status"`
	Msg     string  `json:"msg"`
	Wording string  `json:"wording"`
}

func (c *StatusResult) ToGRPC() *StatusResultGRPC {
	result := &StatusResultGRPC{
		Retcode: c.Retcode,
		Status:  c.Status,
		Msg:     c.Msg,
		Wording: c.Wording,
	}
	if c.Data != nil {
		result.Data = c.Data.ToGRPC()
	}
	return result
}

func (c *StatusResultGRPC) ToStruct() *StatusResult {
	result := &StatusResult{
		Retcode: c.Retcode,
		Status:  c.Status,
		Msg:     c.Msg,
		Wording: c.Wording,
	}
	if c.Data != nil {
		result.Data = c.Data.ToStruct()
	}
	return result
}

type DownloadFile struct {
	// file	string	文件路径
	File string `json:"file"`
}

func (c *DownloadFile) ToGRPC() *DownloadFileGRPC {
	return &DownloadFileGRPC{
		File: c.File,
	}
}

func (c *DownloadFileGRPC) ToStruct() *DownloadFile {
	return &DownloadFile{
		File: c.File,
	}
}

type DownloadFileResult struct {
	Data    *DownloadFile `json:"data"`
	Retcode int64         `json:"ret_code"`
	Status  string        `json:"status"`
	Msg     string        `json:"msg"`
	Wording string        `json:"wording"`
}

func (c *DownloadFileResult) ToGRPC() *DownloadFileResultGRPC {
	result := &DownloadFileResultGRPC{
		Retcode: c.Retcode,
		Status:  c.Status,
		Msg:     c.Msg,
		Wording: c.Wording,
	}
	if c.Data != nil {
		result.Data = c.Data.ToGRPC()
	}
	return result
}

func (c *DownloadFileResultGRPC) ToStruct() *DownloadFileResult {
	result := &DownloadFileResult{
		Retcode: c.Retcode,
		Status:  c.Status,
		Msg:     c.Msg,
		Wording: c.Wording,
	}
	if c.Data != nil {
		result.Data = c.Data.ToStruct()
	}
	return result
}

type CheckUrlSafely struct {
	// level	int	安全等级, 1: 安全 2: 未知 3: 危险
	Level int32 `json:"level"`
}

func (c *CheckUrlSafely) ToGRPC() *CheckUrlSafelyGRPC {
	return &CheckUrlSafelyGRPC{
		Level: c.Level,
	}
}

func (c *CheckUrlSafelyGRPC) ToStruct() *CheckUrlSafely {
	return &CheckUrlSafely{
		Level: c.Level,
	}
}

type CheckUrlSafelyResult struct {
	Data    *CheckUrlSafely `json:"data"`
	Retcode int64           `json:"ret_code"`
	Status  string          `json:"status"`
	Msg     string          `json:"msg"`
	Wording string          `json:"wording"`
}

func (c *CheckUrlSafelyResult) ToGRPC() *CheckUrlSafelyResultGRPC {
	result := &CheckUrlSafelyResultGRPC{
		Retcode: c.Retcode,
		Status:  c.Status,
		Msg:     c.Msg,
		Wording: c.Wording,
	}
	if c.Data != nil {
		result.Data = c.Data.ToGRPC()
	}
	return result
}

func (c *CheckUrlSafelyResultGRPC) ToStruct() *CheckUrlSafelyResult {
	result := &CheckUrlSafelyResult{
		Retcode: c.Retcode,
		Status:  c.Status,
		Msg:     c.Msg,
		Wording: c.Wording,
	}
	if c.Data != nil {
		result.Data = c.Data.ToStruct()
	}
	return result
}

type WordSlices struct {
	// slices string[]	词组
	Slices []string `json:"slices"`
}

func (c *WordSlices) ToGRPC() *WordSlicesGRPC {
	return &WordSlicesGRPC{
		Slices: c.Slices,
	}
}

func (c *WordSlicesGRPC) ToStruct() *WordSlices {
	return &WordSlices{
		Slices: c.Slices,
	}
}

type WordSlicesResult struct {
	Data    *WordSlices `json:"data"`
	Retcode int64       `json:"ret_code"`
	Status  string      `json:"status"`
	Msg     string      `json:"msg"`
	Wording string      `json:"wording"`
}

func (c *WordSlicesResult) ToGRPC() *WordSlicesResultGRPC {
	result := &WordSlicesResultGRPC{
		Retcode: c.Retcode,
		Status:  c.Status,
		Msg:     c.Msg,
		Wording: c.Wording,
	}
	if c.Data != nil {
		result.Data = c.Data.ToGRPC()
	}
	return result
}

func (c *WordSlicesResultGRPC) ToStruct() *WordSlicesResult {
	result := &WordSlicesResult{
		Retcode: c.Retcode,
		Status:  c.Status,
		Msg:     c.Msg,
		Wording: c.Wording,
	}
	if c.Data != nil {
		result.Data = c.Data.ToStruct()
	}
	return result
}
