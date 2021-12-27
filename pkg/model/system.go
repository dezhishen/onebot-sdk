package model

type BoolYes struct {
	Yes bool `json:"yes"`
}

type BoolYesOfResult struct {
	Data    *BoolYes `json:"data"`
	RetCode int      `json:"retcode"`
	Status  string   `json:"status"`
}

type VersionInfoData struct {
	AppName        string `json:"app_name"`        //app_name  = 1;
	AppVersion     string `json:"app_version"`     //app_version = 2;
	ProtoclVersion string `json:"protocl_version"` //protocol_version  = 3;
}

type VersionInfoResult struct {
	Data    *VersionInfoData `json:"data"`
	RetCode int64            `json:"ret_code"`
	Status  string           `json:"status"` //
}

type StatusInfoData struct {
	Online bool `json:"online"`
	Good   bool `json:"good"`
}

type StatusInfoResult struct {
	Data    *StatusInfoData `json:"data"`
	RetCode int64           `json:"ret_code"`
	Status  string          `json:"status"`
}
