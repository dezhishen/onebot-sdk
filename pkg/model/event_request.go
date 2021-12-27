package model

type EventRequestBase struct {
	EventBase
	//请求类型
	RequestType string `json:"request_type"`
}

type EventRequestFriend struct {
	EventRequestBase
	//发送请求的QQ
	UserId int64 `json:"user_id"`
	//验证信息
	Comment string `json:"comment"`
	//请求 flag，在调用处理请求的 API 时需要传入
	Flag string `json:"flag"`
}

type EventRequestGroup struct {
	EventRequestBase
	//add、invite	请求子类型，分别表示加群请求、邀请登录号入群
	SubType string `json:"sub_type"`
	//群号
	GroupId int64 `json:"group_id"`
	//发送请求的QQ
	UserId int64 `json:"user_id"`
	//验证信息
	Comment string `json:"comment"`
	//请求 flag，在调用处理请求的 API 时需要传入
	Flag string `json:"flag"`
}
