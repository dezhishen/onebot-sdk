package model

type EventNoticeBase struct {
	EventBase
	NoticeType string `json:"notice_type"`
}

func (a *EventNoticeBase) ToGRPC() *EventNoticeBaseGRPC {
	return &EventNoticeBaseGRPC{
		EventBase:  a.EventBase.ToGRPC(),
		NoticeType: a.NoticeType,
	}
}

func (a *EventNoticeBaseGRPC) ToStruct() *EventNoticeBase {
	return &EventNoticeBase{
		EventBase:  *a.EventBase.ToStruct(),
		NoticeType: a.NoticeType,
	}
}

type EventNoticeGroupBase struct {
	EventNoticeBase
	//群号
	GroupId int64 `json:"group_id"`
}

func (a *EventNoticeGroupBase) ToGRPC() *EventNoticeGroupBaseGRPC {
	return &EventNoticeGroupBaseGRPC{
		EventNoticeBase: a.EventNoticeBase.ToGRPC(),
		GroupId:         a.GroupId,
	}
}

func (a *EventNoticeGroupBaseGRPC) ToStruct() *EventNoticeGroupBase {
	return &EventNoticeGroupBase{
		EventNoticeBase: *a.EventNoticeBase.ToStruct(),
		GroupId:         a.GroupId,
	}
}

type EventNoticeNotifyBase struct {
	EventNoticeBase
	// honor,lucky_king,poke 群荣誉,红包运气王，戳一戳
	SubType string `json:"sub_type"`
}

func (a *EventNoticeNotifyBase) ToGRPC() *EventNoticeNotifyBaseGRPC {
	return &EventNoticeNotifyBaseGRPC{
		EventNoticeBase: a.EventNoticeBase.ToGRPC(),
		SubType:         a.SubType,
	}
}

func (a *EventNoticeNotifyBaseGRPC) ToStruct() *EventNoticeNotifyBase {
	return &EventNoticeNotifyBase{
		EventNoticeBase: *a.EventNoticeBase.ToStruct(),
		SubType:         a.SubType,
	}
}

type EventNoticeNotifyGroupBase struct {
	EventNoticeGroupBase
	// honor,lucky_king,poke 群荣誉,红包运气王，戳一戳
	SubType string `json:"sub_type"`
}

func (a *EventNoticeNotifyGroupBase) ToGRPC() *EventNoticeNotifyGroupBaseGRPC {
	return &EventNoticeNotifyGroupBaseGRPC{
		EventNoticeGroupBase: a.EventNoticeGroupBase.ToGRPC(),
		SubType:              a.SubType,
	}
}

func (a *EventNoticeNotifyGroupBaseGRPC) ToStruct() *EventNoticeNotifyGroupBase {
	return &EventNoticeNotifyGroupBase{
		EventNoticeGroupBase: *a.EventNoticeGroupBase.ToStruct(),
		SubType:              a.SubType,
	}
}

type QQFile struct {
	//文件 Id
	Id string `json:"id"`
	//文件名
	Name string `json:"name"`
	//文件大小（字节数）
	Size int64 `json:"size"`
	//busid（目前不清楚有什么作用）
	Busid int64 `json:"busid"`
}

func (a *QQFile) ToGRPC() *QQFileGRPC {
	return &QQFileGRPC{
		Id:    a.Id,
		Name:  a.Name,
		Size:  a.Size,
		Busid: a.Busid,
	}
}

func (a *QQFileGRPC) ToStruct() *QQFile {
	return &QQFile{
		Id:    a.Id,
		Name:  a.Name,
		Size:  a.Size,
		Busid: a.Busid,
	}
}

type EventNoticeGroupUpload struct {
	EventNoticeGroupBase
	//发送者 QQ 号
	UserId int64   `json:"user_id"`
	File   *QQFile `json:"file"`
}

func (a *EventNoticeGroupUpload) ToGRPC() *EventNoticeGroupUploadGRPC {
	return &EventNoticeGroupUploadGRPC{
		EventNoticeGroupBase: a.EventNoticeGroupBase.ToGRPC(),
		UserId:               a.UserId,
		File:                 a.File.ToGRPC(),
	}
}

func (a *EventNoticeGroupUploadGRPC) ToStruct() *EventNoticeGroupUpload {
	return &EventNoticeGroupUpload{
		EventNoticeGroupBase: *a.EventNoticeGroupBase.ToStruct(),
		UserId:               a.UserId,
		File:                 a.File.ToStruct(),
	}
}

type EventNoticeGroupAdmin struct {
	EventNoticeGroupBase
	//set、unset	事件子类型，分别表示设置和取消管理员
	SubType string `json:"sub_type"`
	//管理员 QQ 号
	UserId int64 `json:"user_id"`
}

//群成员减少
type EventNoticeGroupDecrease struct {
	EventNoticeGroupBase
	//leave、kick、kick_me	事件子类型，分别表示主动退群、成员被踢、登录号被踢
	SubType string `json:"sub_type"`
	//离开者 QQ 号
	UserId int64 `json:"user_id"`
	//操作者 QQ 号（如果是主动退群，则和 user_id 相同）
	OperatorId int64 `json:"operator_id"`
}

//群成员增加
type EventNoticeGroupIncrease struct {
	EventNoticeGroupBase
	//approve、invite	事件子类型，分别表示管理员已同意入群、管理员邀请入群
	SubType string `json:"sub_type"`
	//加入者 QQ 号
	UserId int64 `json:"user_id"`
	//操作者 QQ 号
	OperatorId int64 `json:"operator_id"`
}

//群禁言
type EventNoticeGroupBan struct {
	EventNoticeGroupBase
	//ban、lift_ban	事件子类型，分别表示禁言、解除禁言
	SubType string `json:"sub_type"`
	//加入者 QQ 号
	UserId int64 `json:"user_id"`
	//操作者 QQ 号
	OperatorId int64 `json:"operator_id"`
	//禁言时长，单位秒
	Duration int64 `json:"duration"`
}

//好友添加
type EventNoticeFriendAdd struct {
	EventNoticeBase
	//新添加好友 QQ 号
	UserId int64 `json:"user_id"`
}

//群消息撤回
type EventNoticeGroupRecall struct {
	EventNoticeGroupBase
	//消息发送者 QQ 号
	UserId int64 `json:"user_id"`
	//操作者 QQ 号
	OperatorId int64 `json:"operator_id"`
	//被撤回的消息 Id
	MessageId int64 `json:"message_id"`
}

//好友消息撤回
type EventNoticeFriendRecall struct {
	EventNoticeBase
	//好友 QQ 号
	UserId int64 `json:"user_id"`
	//被撤回的消息 Id
	MessageId int64 `json:"message_id"`
}

//群内戳一戳
type EventNoticeGroupNotifyPoke struct {
	EventNoticeNotifyGroupBase
	//发送者 QQ 号
	UserId int64 `json:"user_id"`
	//被戳者 QQ 号
	TargetId int64 `json:"target_id"`
}

//群红包运气王
type EventNoticeGroupNotifyLuckyKing struct {
	EventNoticeNotifyGroupBase
	//红包发送者 QQ 号
	UserId int64 `json:"user_id"`
	//运气王 QQ 号
	TargetId int64 `json:"target_id"`
}

//群成员荣誉变更
type EventNoticeGroupNotifyHonor struct {
	EventNoticeNotifyGroupBase
	//talkative、performer、emotion	荣誉类型，分别表示龙王、群聊之火、快乐源泉
	HonorType string `json:"honor_type"`
	//成员 QQ 号
	UserId int64 `json:"user_id"`
}
