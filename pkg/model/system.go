package model

type MapInfoResult struct {
	Data    map[string]interface{} `json:"data"`
	RetCode int                    `json:"retcode"`
	Status  string                 `json:"status"`
}
