package model

type EventRequestBase struct {
	EventBase
	//请求类型
	RequestType string `json:"request_type"`
}

func (e *EventRequestBase) ToGRPC() *EventRequestBaseGRPC {
	return &EventRequestBaseGRPC{
		EventBase:   e.EventBase.ToGRPC(),
		RequestType: e.RequestType,
	}
}

func (e *EventRequestBaseGRPC) ToStruct() *EventRequestBase {
	return &EventRequestBase{
		EventBase:   *e.EventBase.ToStruct(),
		RequestType: e.RequestType,
	}
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

func (e *EventRequestFriend) ToGRPC() *EventRequestFriendGRPC {
	return &EventRequestFriendGRPC{
		EventRequestBase: e.EventRequestBase.ToGRPC(),
		UserId:           e.UserId,
		Comment:          e.Comment,
		Flag:             e.Flag,
	}
}

func (e *EventRequestFriendGRPC) ToStruct() *EventRequestFriend {
	return &EventRequestFriend{
		EventRequestBase: *e.EventRequestBase.ToStruct(),
		UserId:           e.UserId,
		Comment:          e.Comment,
		Flag:             e.Flag,
	}
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

func (e *EventRequestGroup) ToGRPC() *EventRequestGroupGRPC {
	return &EventRequestGroupGRPC{
		EventRequestBase: e.EventRequestBase.ToGRPC(),
		SubType:          e.SubType,
		GroupId:          e.GroupId,
		UserId:           e.UserId,
		Comment:          e.Comment,
		Flag:             e.Flag,
	}
}

func (e *EventRequestGroupGRPC) ToStruct() *EventRequestGroup {
	return &EventRequestGroup{
		EventRequestBase: *e.EventRequestBase.ToStruct(),
		SubType:          e.SubType,
		GroupId:          e.GroupId,
		UserId:           e.UserId,
		Comment:          e.Comment,
		Flag:             e.Flag,
	}
}
