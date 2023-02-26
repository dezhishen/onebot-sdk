package model

type GroupInfo struct {
	GroupId        int64  `json:"group_id"`
	GroupName      string `json:"group_name"`
	MemberCount    int32  `json:"member_count"`
	MaxMemberCount int32  `json:"max_member_count"`
}

func (g *GroupInfo) ToGRPC() *GroupInfoGRPC {
	return &GroupInfoGRPC{
		GroupId:        g.GroupId,
		GroupName:      g.GroupName,
		MemberCount:    g.MemberCount,
		MaxMemberCount: g.MaxMemberCount,
	}
}
func (g *GroupInfoGRPC) ToStruct() *GroupInfo {
	return &GroupInfo{
		GroupId:        g.GroupId,
		GroupName:      g.GroupName,
		MemberCount:    g.MemberCount,
		MaxMemberCount: g.MaxMemberCount,
	}
}

type GroupInfoResult struct {
	Data    *GroupInfo `json:"data"`
	Retcode int64      `json:"retcode"`
	Status  string     `json:"status"`
	Msg     string     `json:"msg"`
	Wording string     `json:"wording"`
}

func GroupInfoArray2GroupInfoGRPCArray(source []*GroupInfo) []*GroupInfoGRPC {
	var result []*GroupInfoGRPC
	for _, e := range source {
		result = append(result, e.ToGRPC())
	}
	return result
}

func GroupInfoGRPCArray2GroupInfoArray(source []*GroupInfoGRPC) []*GroupInfo {
	var result []*GroupInfo
	for _, e := range source {
		result = append(result, e.ToStruct())
	}
	return result
}

func (a *GroupInfoResult) ToGRPC() *GroupInfoResultGRPC {
	result := &GroupInfoResultGRPC{
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

func (a *GroupInfoResultGRPC) ToStruct() *GroupInfoResult {
	result := &GroupInfoResult{
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

type GroupListResult struct {
	Data    []*GroupInfo `json:"data"`
	Retcode int64        `json:"retcode"`
	Status  string       `json:"status"`
	Msg     string       `json:"msg"`
	Wording string       `json:"wording"`
}

func (a *GroupListResult) ToGRPC() *GroupListResultGRPC {
	return &GroupListResultGRPC{
		Data:    GroupInfoArray2GroupInfoGRPCArray(a.Data),
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
}

func (a *GroupListResultGRPC) ToStruct() *GroupListResult {
	return &GroupListResult{
		Data:    GroupInfoGRPCArray2GroupInfoArray(a.Data),
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
}

type GroupMemberInfo struct {
	GroupId         int64  `json:"group_id"`
	UserId          int64  `json:"user_id"`
	Nickname        string `json:"nickname"`
	Card            string `json:"card"`
	Sex             string `json:"sex"`
	Age             int32  `json:"age"`
	Area            string `json:"area"`
	JoinTime        int32  `json:"join_time"`
	LastSentTime    int32  `json:"last_sent_time"`
	Level           string `json:"level"`
	Unfriendly      bool   `json:"unfriendly"`
	Title           string `json:"title"`
	TitleExpireTime int32  `json:"title_expire_time"`
	CardChangeable  bool   `json:"card_changeable"`
}

func (a *GroupMemberInfo) ToGRPC() *GroupMemberInfoGRPC {
	return &GroupMemberInfoGRPC{
		UserId:          a.UserId,
		Sex:             a.Sex,
		Nickname:        a.Nickname,
		Card:            a.Card,
		Age:             a.Age,
		Area:            a.Area,
		JoinTime:        a.JoinTime,
		LastSentTime:    a.LastSentTime,
		Level:           a.Level,
		Unfriendly:      a.Unfriendly,
		Title:           a.Title,
		TitleExpireTime: a.TitleExpireTime,
		CardChangeable:  a.CardChangeable,
	}
}

func (a *GroupMemberInfoGRPC) ToStruct() *GroupMemberInfo {
	return &GroupMemberInfo{
		UserId:          a.UserId,
		Sex:             a.Sex,
		Nickname:        a.Nickname,
		Card:            a.Card,
		Age:             a.Age,
		Area:            a.Area,
		JoinTime:        a.JoinTime,
		LastSentTime:    a.LastSentTime,
		Level:           a.Level,
		Unfriendly:      a.Unfriendly,
		Title:           a.Title,
		TitleExpireTime: a.TitleExpireTime,
		CardChangeable:  a.CardChangeable,
	}
}

func GroupMemberInfoArray2GroupMemberInfoGRPCArray(source []*GroupMemberInfo) []*GroupMemberInfoGRPC {
	var result []*GroupMemberInfoGRPC
	for _, e := range source {
		result = append(result, e.ToGRPC())
	}
	return result
}

func GroupMemberInfoGRPCArray2GroupMemberInfoArray(source []*GroupMemberInfoGRPC) []*GroupMemberInfo {
	var result []*GroupMemberInfo
	for _, e := range source {
		result = append(result, e.ToStruct())
	}
	return result
}

type GroupMemberInfoResult struct {
	Data    *GroupMemberInfo `json:"data"`
	Retcode int64            `json:"retcode"`
	Status  string           `json:"status"`
	Msg     string           `json:"msg"`
	Wording string           `json:"wording"`
}

func (a *GroupMemberInfoResult) ToGRPC() *GroupMemberInfoResultGRPC {
	result := &GroupMemberInfoResultGRPC{
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

func (a *GroupMemberInfoResultGRPC) ToStruct() *GroupMemberInfoResult {
	result := &GroupMemberInfoResult{
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

type GroupMemberListResult struct {
	Data    []*GroupMemberInfo `json:"data"`
	Retcode int64              `json:"retcode"`
	Status  string             `json:"status"`
	Msg     string             `json:"msg"`
	Wording string             `json:"wording"`
}

func (a *GroupMemberListResult) ToGRPC() *GroupMemberListResultGRPC {
	return &GroupMemberListResultGRPC{
		Data:    GroupMemberInfoArray2GroupMemberInfoGRPCArray(a.Data),
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
}

func (a *GroupMemberListResultGRPC) ToStruct() *GroupMemberListResult {
	return &GroupMemberListResult{
		Data:    GroupMemberInfoGRPCArray2GroupMemberInfoArray(a.Data),
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
}

type GroupHonorInfoResult struct {
	Data    *GroupHonorInfo `json:"data"`
	Retcode int64           `json:"retcode"`
	Status  string          `json:"status"`
	Msg     string          `json:"msg"`
	Wording string          `json:"wording"`
}

func (a *GroupHonorInfoResult) ToGRPC() *GroupHonorInfoResultGRPC {
	result := &GroupHonorInfoResultGRPC{
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

func (a *GroupHonorInfoResultGRPC) ToStruct() *GroupHonorInfoResult {
	result := &GroupHonorInfoResult{
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

type GroupHonorInfo struct {
	GroupId          int64             `json:"group_id"`
	CurrentTalkative *CurrentTalkative `json:"current_talkative"`  //龙王
	TalkativeList    []*HonorOwnerInfo `json:"talkative_list"`     //历史龙王
	PerformerList    []*HonorOwnerInfo `json:"performer_list"`     //群聊之火
	LegendList       []*HonorOwnerInfo `json:"legend_list"`        //群聊炽焰
	StrongNewbieList []*HonorOwnerInfo `json:"strong_newbie_list"` //冒尖小春笋
	EmotionList      []*HonorOwnerInfo `json:"emotion_list"`       //快乐之源
}

func (a *GroupHonorInfo) ToGRPC() *GroupHonorInfoGRPC {
	result := &GroupHonorInfoGRPC{
		GroupId:          a.GroupId,
		TalkativeList:    HonorOwnerInfoArray2HonorOwnerInfoGRPCArray(a.TalkativeList),
		PerformerList:    HonorOwnerInfoArray2HonorOwnerInfoGRPCArray(a.PerformerList),
		LegendList:       HonorOwnerInfoArray2HonorOwnerInfoGRPCArray(a.LegendList),
		StrongNewbieList: HonorOwnerInfoArray2HonorOwnerInfoGRPCArray(a.StrongNewbieList),
		EmotionList:      HonorOwnerInfoArray2HonorOwnerInfoGRPCArray(a.EmotionList),
	}
	if a.CurrentTalkative != nil {
		result.CurrentTalkative = a.CurrentTalkative.ToGRPC()
	}
	return result
}

func (a *GroupHonorInfoGRPC) ToStruct() *GroupHonorInfo {
	result := &GroupHonorInfo{
		GroupId:          a.GroupId,
		TalkativeList:    HonorOwnerInfoGRPCArray2HonorOwnerInfoArray(a.TalkativeList),
		PerformerList:    HonorOwnerInfoGRPCArray2HonorOwnerInfoArray(a.PerformerList),
		LegendList:       HonorOwnerInfoGRPCArray2HonorOwnerInfoArray(a.LegendList),
		StrongNewbieList: HonorOwnerInfoGRPCArray2HonorOwnerInfoArray(a.StrongNewbieList),
		EmotionList:      HonorOwnerInfoGRPCArray2HonorOwnerInfoArray(a.EmotionList),
	}
	if a.CurrentTalkative != nil {
		result.CurrentTalkative = a.CurrentTalkative.ToStruct()
	}
	return result
}

type CurrentTalkative struct {
	UserId   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`    //头像url
	DayCount int32  `json:"day_count"` //持续天数
}

func (a *CurrentTalkative) ToGRPC() *CurrentTalkativeGRPC {
	return &CurrentTalkativeGRPC{
		UserId:   a.UserId,
		Avatar:   a.Avatar,
		Nickname: a.Nickname,
		DayCount: a.DayCount,
	}
}

func (a *CurrentTalkativeGRPC) ToStruct() *CurrentTalkative {
	return &CurrentTalkative{
		UserId:   a.UserId,
		Avatar:   a.Avatar,
		Nickname: a.Nickname,
		DayCount: a.DayCount,
	}
}

type HonorOwnerInfo struct {
	UserId      int64  `json:"user_id"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
}

func (a *HonorOwnerInfo) ToGRPC() *HonorOwnerInfoGRPC {
	return &HonorOwnerInfoGRPC{
		UserId:      a.UserId,
		Avatar:      a.Avatar,
		Nickname:    a.Nickname,
		Description: a.Description,
	}
}

func (a *HonorOwnerInfoGRPC) ToStruct() *HonorOwnerInfo {
	return &HonorOwnerInfo{
		UserId:      a.UserId,
		Avatar:      a.Avatar,
		Nickname:    a.Nickname,
		Description: a.Description,
	}
}

func HonorOwnerInfoArray2HonorOwnerInfoGRPCArray(source []*HonorOwnerInfo) []*HonorOwnerInfoGRPC {
	var result []*HonorOwnerInfoGRPC
	for _, e := range source {
		result = append(result, e.ToGRPC())
	}
	return result
}

func HonorOwnerInfoGRPCArray2HonorOwnerInfoArray(source []*HonorOwnerInfoGRPC) []*HonorOwnerInfo {
	var result []*HonorOwnerInfo
	for _, e := range source {
		result = append(result, e.ToStruct())
	}
	return result
}

type InvitedRequest struct {
	// request_id	int64	请求ID
	RequestId int64 `json:"request_id"`
	// invitor_uin	int64	邀请者
	InvitorUin int64 `json:"invitor_uin"`
	// invitor_nick	string	邀请者昵称
	InvitorNick string `json:"invitor_nick"`
	// group_id	int64	群号
	GroupId int64 `json:"group_id"`
	// group_name	string	群名
	GroupName string `json:"group_name"`
	// checked	bool	是否已被处理
	Checked bool `json:"checked"`
	// actor	int64	处理者, 未处理为0
	Actor int64 `json:"actor"`
}

func (a *InvitedRequest) ToGRPC() *InvitedRequestGRPC {
	return &InvitedRequestGRPC{
		RequestId:   a.RequestId,
		InvitorUin:  a.InvitorUin,
		InvitorNick: a.InvitorNick,
		GroupId:     a.GroupId,
		GroupName:   a.GroupName,
		Checked:     a.Checked,
		Actor:       a.Actor,
	}
}

func (a *InvitedRequestGRPC) ToStruct() *InvitedRequest {
	return &InvitedRequest{
		RequestId:   a.RequestId,
		InvitorUin:  a.InvitorUin,
		InvitorNick: a.InvitorNick,
		GroupId:     a.GroupId,
		GroupName:   a.GroupName,
		Checked:     a.Checked,
		Actor:       a.Actor,
	}
}

func InvitedRequestArray2InvitedRequestGRPCArray(source []*InvitedRequest) []*InvitedRequestGRPC {
	var result []*InvitedRequestGRPC
	for _, e := range source {
		result = append(result, e.ToGRPC())
	}
	return result
}

func InvitedRequestGRPCArray2InvitedRequestArray(source []*InvitedRequestGRPC) []*InvitedRequest {
	var result []*InvitedRequest
	for _, e := range source {
		result = append(result, e.ToStruct())
	}
	return result
}

type JoinRequest struct {
	// request_id	int64	请求ID
	RequestId int64 `json:"request_id"`
	// requester_uin	int64	请求者ID
	RequesterUin int64 `json:"requester_uin"`
	// requester_nick	string	请求者昵称
	RequesterNick string `json:"requester_nick"`
	// message	string	验证消息
	Message string `json:"message"`
	// group_id	int64	群号
	GroupId int64 `json:"group_id"`
	// group_name	string	群名
	GroupName string `json:"group_name"`
	// checked	bool	是否已被处理
	Checked bool `json:"checked"`
	// actor	int64	处理者, 未处理为0
	Actor int64 `json:"actor"`
}

func (a *JoinRequest) ToGRPC() *JoinRequestGRPC {
	return &JoinRequestGRPC{
		RequestId:     a.RequestId,
		RequesterUin:  a.RequesterUin,
		RequesterNick: a.RequesterNick,
		Message:       a.Message,
		GroupId:       a.GroupId,
		GroupName:     a.GroupName,
		Checked:       a.Checked,
		Actor:         a.Actor,
	}
}

func (a *JoinRequestGRPC) ToStruct() *JoinRequest {
	return &JoinRequest{
		RequestId:     a.RequestId,
		RequesterUin:  a.RequesterUin,
		RequesterNick: a.RequesterNick,
		Message:       a.Message,
		GroupId:       a.GroupId,
		GroupName:     a.GroupName,
		Checked:       a.Checked,
		Actor:         a.Actor,
	}
}

func JoinRequestArray2JoinRequestGRPCArray(source []*JoinRequest) []*JoinRequestGRPC {
	var result []*JoinRequestGRPC
	for _, e := range source {
		result = append(result, e.ToGRPC())
	}
	return result
}

func JoinRequestGRPCArray2JoinRequestArray(source []*JoinRequestGRPC) []*JoinRequest {
	var result []*JoinRequest
	for _, e := range source {
		result = append(result, e.ToStruct())
	}
	return result
}

type GroupSystemMsg struct {
	InvitedRequests []*InvitedRequest `json:"invited_requests"`
	JoinRequests    []*JoinRequest    `json:"join_requests"`
}

func (a *GroupSystemMsg) ToGRPC() *GroupSystemMsgGRPC {
	return &GroupSystemMsgGRPC{
		InvitedRequests: InvitedRequestArray2InvitedRequestGRPCArray(a.InvitedRequests),
		JoinRequests:    JoinRequestArray2JoinRequestGRPCArray(a.JoinRequests),
	}
}

func (a *GroupSystemMsgGRPC) ToStruct() *GroupSystemMsg {
	return &GroupSystemMsg{
		InvitedRequests: InvitedRequestGRPCArray2InvitedRequestArray(a.InvitedRequests),
		JoinRequests:    JoinRequestGRPCArray2JoinRequestArray(a.JoinRequests),
	}
}

type GroupSystemMsgResult struct {
	Data    *GroupSystemMsg `json:"data"`
	Retcode int64           `json:"retcode"`
	Status  string          `json:"status"`
	Msg     string          `json:"msg"`
	Wording string          `json:"wording"`
}

func (a *GroupSystemMsgResult) ToGRPC() *GroupSystemMsgResultGRPC {
	result := &GroupSystemMsgResultGRPC{
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

func (a *GroupSystemMsgResultGRPC) ToStruct() *GroupSystemMsgResult {
	result := &GroupSystemMsgResult{
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

type EssenceMsg struct {
	// sender_id	int64	发送者QQ 号
	SenderId int64 `json:"sender_id"`
	// sender_nick	string	发送者昵称
	SenderNick string `json:"sender_nick"`
	// sender_time	int64	消息发送时间
	SenderTime int64 `json:"sender_time"`
	// operator_id	int64	操作者QQ 号
	OperatorId int64 `json:"operator_id"`
	// operator_nick	string	操作者昵称
	OperatorNick string `json:"operator_nick"`
	// operator_time	int64	精华设置时间
	OperatorTime int64 `json:"operator_time"`
	// message_id	int32	消息ID
	MessageId int32 `json:"message_id"`
}

func (a *EssenceMsg) ToGRPC() *EssenceMsgGRPC {
	return &EssenceMsgGRPC{
		SenderId:     a.SenderId,
		SenderNick:   a.SenderNick,
		SenderTime:   a.SenderTime,
		OperatorId:   a.OperatorId,
		OperatorNick: a.OperatorNick,
		OperatorTime: a.OperatorTime,
		MessageId:    a.MessageId,
	}
}

func (a *EssenceMsgGRPC) ToStruct() *EssenceMsg {
	return &EssenceMsg{
		SenderId:     a.SenderId,
		SenderNick:   a.SenderNick,
		SenderTime:   a.SenderTime,
		OperatorId:   a.OperatorId,
		OperatorNick: a.OperatorNick,
		OperatorTime: a.OperatorTime,
		MessageId:    a.MessageId,
	}
}

func EssenceMsgArray2EssenceMsgGRPCArray(source []*EssenceMsg) []*EssenceMsgGRPC {
	var result []*EssenceMsgGRPC
	for _, e := range source {
		result = append(result, e.ToGRPC())
	}
	return result
}

func EssenceMsgGRPCArray2EssenceMsgArray(source []*EssenceMsgGRPC) []*EssenceMsg {
	var result []*EssenceMsg
	for _, e := range source {
		result = append(result, e.ToStruct())
	}
	return result
}

type EssenceMsgListResult struct {
	Data    []*EssenceMsg `json:"data"`
	Retcode int64         `json:"retcode"`
	Status  string        `json:"status"`
	Msg     string        `json:"msg"`
	Wording string        `json:"wording"`
}

func (a *EssenceMsgListResult) ToGRPC() *EssenceMsgListResultGRPC {
	return &EssenceMsgListResultGRPC{
		Data:    EssenceMsgArray2EssenceMsgGRPCArray(a.Data),
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
}

func (a *EssenceMsgListResultGRPC) ToStruct() *EssenceMsgListResult {
	return &EssenceMsgListResult{
		Data:    EssenceMsgGRPCArray2EssenceMsgArray(a.Data),
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
}

type GroupAtAllRemain struct {
	// can_at_all	bool	是否可以 @全体成员
	CanAtAll bool `json:"can_at_all"`
	// remain_at_all_count_for_group	int16	群内所有管理当天剩余 @全体成员 次数
	RemainAtAllCountForGroup int32 `json:"remain_at_all_count_for_group"`
	// remain_at_all_count_for_uin	int16	Bot 当天剩余 @全体成员 次数
	RemainAtAllCountForUin int32 `json:"remain_at_all_count_for_uin"`
}

func (a *GroupAtAllRemain) ToGRPC() *GroupAtAllRemainGRPC {
	return &GroupAtAllRemainGRPC{
		CanAtAll:                 a.CanAtAll,
		RemainAtAllCountForGroup: a.RemainAtAllCountForGroup,
		RemainAtAllCountForUin:   a.RemainAtAllCountForUin,
	}
}

func (a *GroupAtAllRemainGRPC) ToStruct() *GroupAtAllRemain {
	return &GroupAtAllRemain{
		CanAtAll:                 a.CanAtAll,
		RemainAtAllCountForGroup: a.RemainAtAllCountForGroup,
		RemainAtAllCountForUin:   a.RemainAtAllCountForUin,
	}
}

type GroupAtAllRemainResult struct {
	Data    *GroupAtAllRemain `json:"data"`
	Retcode int64             `json:"retcode"`
	Status  string            `json:"status"`
	Msg     string            `json:"msg"`
	Wording string            `json:"wording"`
}

func (a *GroupAtAllRemainResult) ToGRPC() *GroupAtAllRemainResultGRPC {
	result := &GroupAtAllRemainResultGRPC{
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

func (a *GroupAtAllRemainResultGRPC) ToStruct() *GroupAtAllRemainResult {
	result := &GroupAtAllRemainResult{
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

type GroupNoticeMessage struct {
	// text	string	公告内容
	Text string `json:"text"`
	// images	array	公告图片
	Images []string `json:"images"`
}

func (a *GroupNoticeMessage) ToGRPC() *GroupNoticeMessageGRPC {
	return &GroupNoticeMessageGRPC{
		Text:   a.Text,
		Images: a.Images,
	}
}

func (a *GroupNoticeMessageGRPC) ToStruct() *GroupNoticeMessage {
	return &GroupNoticeMessage{
		Text:   a.Text,
		Images: a.Images,
	}
}

type GroupNotice struct {
	// sender_id	int64	公告发表者
	SenderId int64 `json:"sender_id"`
	// publish_time	int64	公告发表时间
	PublishTime int64 `json:"publish_time"`
	// message	object	公告内容
	Message *GroupNoticeMessage `json:"message"`
}

func (a *GroupNotice) ToGRPC() *GroupNoticeGRPC {
	return &GroupNoticeGRPC{
		SenderId:    a.SenderId,
		PublishTime: a.PublishTime,
		Message:     a.Message.ToGRPC(),
	}
}

func (a *GroupNoticeGRPC) ToStruct() *GroupNotice {
	return &GroupNotice{
		SenderId:    a.SenderId,
		PublishTime: a.PublishTime,
		Message:     a.Message.ToStruct(),
	}
}

type GroupNoticeResult struct {
	Data    *GroupNotice `json:"data"`
	Retcode int64        `json:"retcode"`
	Status  string       `json:"status"`
	Msg     string       `json:"msg"`
	Wording string       `json:"wording"`
}

func (a *GroupNoticeResult) ToGRPC() *GroupNoticeResultGRPC {
	return &GroupNoticeResultGRPC{
		Data:    a.Data.ToGRPC(),
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
}

func (a *GroupNoticeResultGRPC) ToStruct() *GroupNoticeResult {
	return &GroupNoticeResult{
		Data:    a.Data.ToStruct(),
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
}
