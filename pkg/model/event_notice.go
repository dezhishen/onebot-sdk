package model

type EventNoticeBase struct {
	EventBase
	NoticeType string `json:"notice_type"`
}

func (e *EventNoticeBase) ToGRPC() *EventNoticeBaseGRPC {
	return &EventNoticeBaseGRPC{
		EventBase:  e.EventBase.ToGRPC(),
		NoticeType: e.NoticeType,
	}
}

func (e *EventNoticeBaseGRPC) ToStruct() *EventNoticeBase {
	return &EventNoticeBase{
		EventBase:  *e.EventBase.ToStruct(),
		NoticeType: e.NoticeType,
	}
}

type EventNoticeGroupBase struct {
	EventNoticeBase
	//群号
	GroupId int64 `json:"group_id"`
}

func (e *EventNoticeGroupBase) ToGRPC() *EventNoticeGroupBaseGRPC {
	return &EventNoticeGroupBaseGRPC{
		EventNoticeBase: e.EventNoticeBase.ToGRPC(),
		GroupId:         e.GroupId,
	}
}

func (e *EventNoticeGroupBaseGRPC) ToStruct() *EventNoticeGroupBase {
	return &EventNoticeGroupBase{
		EventNoticeBase: *e.EventNoticeBase.ToStruct(),
		GroupId:         e.GroupId,
	}
}

type EventNoticeNotifyBase struct {
	EventNoticeBase
	// honor,lucky_king,poke 群荣誉,红包运气王，戳一戳
	SubType string `json:"sub_type"`
}

func (e *EventNoticeNotifyBase) ToGRPC() *EventNoticeNotifyBaseGRPC {
	return &EventNoticeNotifyBaseGRPC{
		EventNoticeBase: e.EventNoticeBase.ToGRPC(),
		SubType:         e.SubType,
	}
}

func (e *EventNoticeNotifyBaseGRPC) ToStruct() *EventNoticeNotifyBase {
	return &EventNoticeNotifyBase{
		EventNoticeBase: *e.EventNoticeBase.ToStruct(),
		SubType:         e.SubType,
	}
}

type EventNoticeNotifyGroupBase struct {
	EventNoticeGroupBase
	// honor,lucky_king,poke 群荣誉,红包运气王，戳一戳
	SubType string `json:"sub_type"`
}

func (e *EventNoticeNotifyGroupBase) ToGRPC() *EventNoticeNotifyGroupBaseGRPC {
	return &EventNoticeNotifyGroupBaseGRPC{
		EventNoticeGroupBase: e.EventNoticeGroupBase.ToGRPC(),
		SubType:              e.SubType,
	}
}

func (e *EventNoticeNotifyGroupBaseGRPC) ToStruct() *EventNoticeNotifyGroupBase {
	return &EventNoticeNotifyGroupBase{
		EventNoticeGroupBase: *e.EventNoticeGroupBase.ToStruct(),
		SubType:              e.SubType,
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

func (e *QQFile) ToGRPC() *QQFileGRPC {
	return &QQFileGRPC{
		Id:    e.Id,
		Name:  e.Name,
		Size:  e.Size,
		Busid: e.Busid,
	}
}

func (e *QQFileGRPC) ToStruct() *QQFile {
	return &QQFile{
		Id:    e.Id,
		Name:  e.Name,
		Size:  e.Size,
		Busid: e.Busid,
	}
}

type EventNoticeGroupUpload struct {
	EventNoticeGroupBase
	//发送者 QQ 号
	UserId int64   `json:"user_id"`
	File   *QQFile `json:"file"`
}

func (e *EventNoticeGroupUpload) ToGRPC() *EventNoticeGroupUploadGRPC {
	result := &EventNoticeGroupUploadGRPC{
		EventNoticeGroupBase: e.EventNoticeGroupBase.ToGRPC(),
		UserId:               e.UserId,
	}
	if e.File != nil {
		result.File = e.File.ToGRPC()
	}
	return result
}

func (e *EventNoticeGroupUploadGRPC) ToStruct() *EventNoticeGroupUpload {
	result := &EventNoticeGroupUpload{
		EventNoticeGroupBase: *e.EventNoticeGroupBase.ToStruct(),
		UserId:               e.UserId,
	}
	if e.File != nil {
		result.File = e.File.ToStruct()
	}
	return result
}

type EventNoticeGroupAdmin struct {
	EventNoticeGroupBase
	//set、unset	事件子类型，分别表示设置和取消管理员
	SubType string `json:"sub_type"`
	//管理员 QQ 号
	UserId int64 `json:"user_id"`
}

func (e *EventNoticeGroupAdmin) ToGRPC() *EventNoticeGroupAdminGRPC {
	return &EventNoticeGroupAdminGRPC{
		EventNoticeGroupBase: e.EventNoticeGroupBase.ToGRPC(),
		UserId:               e.UserId,
		SubType:              e.SubType,
	}
}

func (e *EventNoticeGroupAdminGRPC) ToStruct() *EventNoticeGroupAdmin {
	return &EventNoticeGroupAdmin{
		EventNoticeGroupBase: *e.EventNoticeGroupBase.ToStruct(),
		UserId:               e.UserId,
		SubType:              e.SubType,
	}
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

func (e *EventNoticeGroupDecrease) ToGRPC() *EventNoticeGroupDecreaseGRPC {
	return &EventNoticeGroupDecreaseGRPC{
		EventNoticeGroupBase: e.EventNoticeGroupBase.ToGRPC(),
		UserId:               e.UserId,
		SubType:              e.SubType,
		OperatorId:           e.OperatorId,
	}
}

func (e *EventNoticeGroupDecreaseGRPC) ToStruct() *EventNoticeGroupDecrease {
	return &EventNoticeGroupDecrease{
		EventNoticeGroupBase: *e.EventNoticeGroupBase.ToStruct(),
		UserId:               e.UserId,
		SubType:              e.SubType,
		OperatorId:           e.OperatorId,
	}
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

func (e *EventNoticeGroupIncrease) ToGRPC() *EventNoticeGroupIncreaseGRPC {
	return &EventNoticeGroupIncreaseGRPC{
		EventNoticeGroupBase: e.EventNoticeGroupBase.ToGRPC(),
		UserId:               e.UserId,
		SubType:              e.SubType,
		OperatorId:           e.OperatorId,
	}
}

func (e *EventNoticeGroupIncreaseGRPC) ToStruct() *EventNoticeGroupIncrease {
	return &EventNoticeGroupIncrease{
		EventNoticeGroupBase: *e.EventNoticeGroupBase.ToStruct(),
		UserId:               e.UserId,
		SubType:              e.SubType,
		OperatorId:           e.OperatorId,
	}
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

func (e *EventNoticeGroupBan) ToGRPC() *EventNoticeGroupBanGRPC {
	return &EventNoticeGroupBanGRPC{
		EventNoticeGroupBase: e.EventNoticeGroupBase.ToGRPC(),
		UserId:               e.UserId,
		SubType:              e.SubType,
		OperatorId:           e.OperatorId,
		Duration:             e.Duration,
	}
}

func (e *EventNoticeGroupBanGRPC) ToStruct() *EventNoticeGroupBan {
	return &EventNoticeGroupBan{
		EventNoticeGroupBase: *e.EventNoticeGroupBase.ToStruct(),
		UserId:               e.UserId,
		SubType:              e.SubType,
		OperatorId:           e.OperatorId,
		Duration:             e.Duration,
	}
}

//好友添加
type EventNoticeFriendAdd struct {
	EventNoticeBase
	//新添加好友 QQ 号
	UserId int64 `json:"user_id"`
}

func (e *EventNoticeFriendAdd) ToGRPC() *EventNoticeFriendAddGRPC {
	return &EventNoticeFriendAddGRPC{
		EventNoticeBase: e.EventNoticeBase.ToGRPC(),
		UserId:          e.UserId,
	}
}

func (e *EventNoticeFriendAddGRPC) ToStruct() *EventNoticeFriendAdd {
	return &EventNoticeFriendAdd{
		EventNoticeBase: *e.EventNoticeBase.ToStruct(),
		UserId:          e.UserId,
	}
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

func (e *EventNoticeGroupRecall) ToGRPC() *EventNoticeGroupRecallGRPC {
	return &EventNoticeGroupRecallGRPC{
		EventNoticeGroupBase: e.EventNoticeGroupBase.ToGRPC(),
		UserId:               e.UserId,
		OperatorId:           e.OperatorId,
		MessageId:            e.MessageId,
	}
}

func (e *EventNoticeGroupRecallGRPC) ToStruct() *EventNoticeGroupRecall {
	return &EventNoticeGroupRecall{
		EventNoticeGroupBase: *e.EventNoticeGroupBase.ToStruct(),
		UserId:               e.UserId,
		OperatorId:           e.OperatorId,
		MessageId:            e.MessageId,
	}
}

//好友消息撤回
type EventNoticeFriendRecall struct {
	EventNoticeBase
	//好友 QQ 号
	UserId int64 `json:"user_id"`
	//被撤回的消息 Id
	MessageId int64 `json:"message_id"`
}

func (e *EventNoticeFriendRecall) ToGRPC() *EventNoticeFriendRecallGRPC {
	return &EventNoticeFriendRecallGRPC{
		EventNoticeBase: e.EventNoticeBase.ToGRPC(),
		UserId:          e.UserId,
		MessageId:       e.MessageId,
	}
}

func (e *EventNoticeFriendRecallGRPC) ToStruct() *EventNoticeFriendRecall {
	return &EventNoticeFriendRecall{
		EventNoticeBase: *e.EventNoticeBase.ToStruct(),
		UserId:          e.UserId,
		MessageId:       e.MessageId,
	}
}

//群内戳一戳
type EventNoticeGroupNotifyPoke struct {
	EventNoticeNotifyGroupBase
	//发送者 QQ 号
	UserId int64 `json:"user_id"`
	//被戳者 QQ 号
	TargetId int64 `json:"target_id"`
}

func (e *EventNoticeGroupNotifyPoke) ToGRPC() *EventNoticeGroupNotifyPokeGRPC {
	return &EventNoticeGroupNotifyPokeGRPC{
		EventNoticeNotifyGroupBase: e.EventNoticeNotifyGroupBase.ToGRPC(),
		UserId:                     e.UserId,
		TargetId:                   e.TargetId,
	}
}

func (e *EventNoticeGroupNotifyPokeGRPC) ToStruct() *EventNoticeGroupNotifyPoke {
	return &EventNoticeGroupNotifyPoke{
		EventNoticeNotifyGroupBase: *e.EventNoticeNotifyGroupBase.ToStruct(),
		UserId:                     e.UserId,
		TargetId:                   e.TargetId,
	}
}

//群红包运气王
type EventNoticeGroupNotifyLuckyKing struct {
	EventNoticeNotifyGroupBase
	//红包发送者 QQ 号
	UserId int64 `json:"user_id"`
	//运气王 QQ 号
	TargetId int64 `json:"target_id"`
}

func (e *EventNoticeGroupNotifyLuckyKing) ToGRPC() *EventNoticeGroupNotifyLuckyKingGRPC {
	return &EventNoticeGroupNotifyLuckyKingGRPC{
		EventNoticeNotifyGroupBase: e.EventNoticeNotifyGroupBase.ToGRPC(),
		UserId:                     e.UserId,
		TargetId:                   e.TargetId,
	}
}

func (e *EventNoticeGroupNotifyLuckyKingGRPC) ToStruct() *EventNoticeGroupNotifyLuckyKing {
	return &EventNoticeGroupNotifyLuckyKing{
		EventNoticeNotifyGroupBase: *e.EventNoticeNotifyGroupBase.ToStruct(),
		UserId:                     e.UserId,
		TargetId:                   e.TargetId,
	}
}

//群成员荣誉变更
type EventNoticeGroupNotifyHonor struct {
	EventNoticeNotifyGroupBase
	//talkative、performer、emotion	荣誉类型，分别表示龙王、群聊之火、快乐源泉
	HonorType string `json:"honor_type"`
	//成员 QQ 号
	UserId int64 `json:"user_id"`
}

func (e *EventNoticeGroupNotifyHonor) ToGRPC() *EventNoticeGroupNotifyHonorGRPC {
	return &EventNoticeGroupNotifyHonorGRPC{
		EventNoticeNotifyGroupBase: e.EventNoticeNotifyGroupBase.ToGRPC(),
		UserId:                     e.UserId,
		HonorType:                  e.HonorType,
	}
}

func (e *EventNoticeGroupNotifyHonorGRPC) ToStruct() *EventNoticeGroupNotifyHonor {
	return &EventNoticeGroupNotifyHonor{
		EventNoticeNotifyGroupBase: *e.EventNoticeNotifyGroupBase.ToStruct(),
		UserId:                     e.UserId,
		HonorType:                  e.HonorType,
	}
}
