package model

type BoolYes struct {
	Yes bool `json:"yes"`
}

func (a *BoolYes) ToGRPC() *BoolYesGRPC {
	return &BoolYesGRPC{
		Yes: a.Yes,
	}
}

func (a *BoolYesGRPC) ToStruct() *BoolYes {
	return &BoolYes{
		Yes: a.Yes,
	}
}

type BoolYesOfResult struct {
	Data    *BoolYes `json:"data"`
	Retcode int64    `json:"retcode"`
	Status  string   `json:"status"`
	Msg     string   `json:"msg"`
	Wording string   `json:"wording"`
}

func (a *BoolYesOfResult) ToGRPC() *BoolYesOfResultGRPC {
	result := &BoolYesOfResultGRPC{
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
	if a.Data != nil {
		result.Data = a.Data.ToGRPC()
	}
	return result
}

func (a *BoolYesOfResultGRPC) ToStruct() *BoolYesOfResult {
	result := &BoolYesOfResult{
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
	if a.Data != nil {
		result.Data = a.Data.ToStruct()
	}
	return result
}

type VersionInfo struct {
	AppName         string `json:"app_name"`        //app_name  = 1;
	AppVersion      string `json:"app_version"`     //app_version = 2;
	ProtocolVersion string `json:"protocl_version"` //protocol_version  = 3;
}

func (a *VersionInfo) ToGRPC() *VersionInfoGRPC {
	return &VersionInfoGRPC{
		AppName:         a.AppName,
		AppVersion:      a.AppVersion,
		ProtocolVersion: a.ProtocolVersion,
	}
}

func (a *VersionInfoGRPC) ToStruct() *VersionInfo {
	return &VersionInfo{
		AppName:         a.AppName,
		AppVersion:      a.AppVersion,
		ProtocolVersion: a.ProtocolVersion,
	}
}

type VersionInfoResult struct {
	Data    *VersionInfo `json:"data"`
	Retcode int64        `json:"ret_code"`
	Status  string       `json:"status"` //
	Msg     string       `json:"msg"`
	Wording string       `json:"wording"`
}

func (a *VersionInfoResult) ToGRPC() *VersionInfoResultGRPC {
	result := &VersionInfoResultGRPC{
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
	if a.Data != nil {
		result.Data = a.Data.ToGRPC()
	}
	return result
}

func (a *VersionInfoResultGRPC) ToStruct() *VersionInfoResult {
	result := &VersionInfoResult{
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
	if a.Data != nil {
		result.Data = a.Data.ToStruct()
	}
	return result
}

type StatusInfo struct {
	Online bool `json:"online"`
	Good   bool `json:"good"`
}

func (a *StatusInfo) ToGRPC() *StatusInfoGRPC {
	return &StatusInfoGRPC{
		Online: a.Online,
		Good:   a.Good,
	}
}

func (a *StatusInfoGRPC) ToStruct() *StatusInfo {
	return &StatusInfo{
		Online: a.Online,
		Good:   a.Good,
	}
}

type StatusInfoResult struct {
	Data    *StatusInfo `json:"data"`
	Retcode int64       `json:"ret_code"`
	Status  string      `json:"status"`
	Msg     string      `json:"msg"`
	Wording string      `json:"wording"`
}

func (a *StatusInfoResult) ToGRPC() *StatusInfoResultGRPC {
	result := &StatusInfoResultGRPC{
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
	if a.Data != nil {
		result.Data = a.Data.ToGRPC()
	}
	return result
}

func (a *StatusInfoResultGRPC) ToStruct() *StatusInfoResult {
	result := &StatusInfoResult{
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
	if a.Data != nil {
		result.Data = a.Data.ToStruct()
	}
	return result
}
