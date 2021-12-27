package model

type Group struct {
	GroupId        int64  `json:"group_id"`
	GroupName      string `json:"group_name"`
	MemberCount    int32  `json:"member_count"`
	MaxMemberCount int32  `json:"max_member_count"`
}

func (g *Group) ToGRPC() *GroupGRPC {
	return &GroupGRPC{
		GroupId:        g.GroupId,
		GroupName:      g.GroupName,
		MemberCount:    g.MemberCount,
		MaxMemberCount: g.MaxMemberCount,
	}
}
func (g *GroupGRPC) ToStruct() *Group {
	return &Group{
		GroupId:        g.GroupId,
		GroupName:      g.GroupName,
		MemberCount:    g.MemberCount,
		MaxMemberCount: g.MaxMemberCount,
	}
}

type GroupResult struct {
	Data    *Group `json:"data"`
	Retcode int64  `json:"retcode"`
	Status  string `json:"status"`
}

func GroupArray2GroupGRPCArray(source []*Group) []*GroupGRPC {
	var result []*GroupGRPC
	for _, e := range source {
		result = append(result, e.ToGRPC())
	}
	return result
}

func GroupGRPCArray2GroupArray(source []*GroupGRPC) []*Group {
	var result []*Group
	for _, e := range source {
		result = append(result, e.ToStruct())
	}
	return result
}

func (a *GroupResult) ToGRPC() *GroupResultGRPC {
	return &GroupResultGRPC{
		Data:    a.Data.ToGRPC(),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

func (a *GroupResultGRPC) ToStruct() *GroupResult {
	return &GroupResult{
		Data:    a.Data.ToStruct(),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

type GroupListResult struct {
	Data    []*Group `json:"data"`
	Retcode int64    `json:"retcode"`
	Status  string   `json:"status"`
}

func (a *GroupListResult) ToGRPC() *GroupListResultGRPC {
	return &GroupListResultGRPC{
		Data:    GroupArray2GroupGRPCArray(a.Data),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

func (a *GroupListResultGRPC) ToStruct() *GroupListResult {
	return &GroupListResult{
		Data:    GroupGRPCArray2GroupArray(a.Data),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

type GroupMember struct {
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

func (a *GroupMember) ToGRPC() *GroupMemberGRPC {
	return &GroupMemberGRPC{
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

func (a *GroupMemberGRPC) ToStruct() *GroupMember {
	return &GroupMember{
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

func GroupMemberArray2GroupMemberGRPCArray(source []*GroupMember) []*GroupMemberGRPC {
	var result []*GroupMemberGRPC
	for _, e := range source {
		result = append(result, e.ToGRPC())
	}
	return result
}

func GroupMemberGRPCArray2GroupMemberArray(source []*GroupMemberGRPC) []*GroupMember {
	var result []*GroupMember
	for _, e := range source {
		result = append(result, e.ToStruct())
	}
	return result
}

type GroupMemberResult struct {
	Data    *GroupMember `json:"data"`
	Retcode int64        `json:"retcode"`
	Status  string       `json:"status"`
}

func (a *GroupMemberResult) ToGRPC() *GroupMemberResultGRPC {
	return &GroupMemberResultGRPC{
		Data:    a.Data.ToGRPC(),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

func (a *GroupMemberResultGRPC) ToStruct() *GroupMemberResult {
	return &GroupMemberResult{
		Data:    a.Data.ToStruct(),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

type GroupMemberListResult struct {
	Data    []*GroupMember `json:"data"`
	Retcode int64          `json:"retcode"`
	Status  string         `json:"status"`
}

func (a *GroupMemberListResult) ToGRPC() *GroupMemberListResultGRPC {
	return &GroupMemberListResultGRPC{
		Data:    GroupMemberArray2GroupMemberGRPCArray(a.Data),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

func (a *GroupMemberListResultGRPC) ToStruct() *GroupMemberListResult {
	return &GroupMemberListResult{
		Data:    GroupMemberGRPCArray2GroupMemberArray(a.Data),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

type GroupHonorInfoResult struct {
	Data    *GroupHonorInfo `json:"data"`
	Retcode int64           `json:"retcode"`
	Status  string          `json:"status"`
}

func (a *GroupHonorInfoResult) ToGRPC() *GroupHonorInfoResultGRPC {
	return &GroupHonorInfoResultGRPC{
		Data:    a.Data.ToGRPC(),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

func (a *GroupHonorInfoResultGRPC) ToStruct() *GroupHonorInfoResult {
	return &GroupHonorInfoResult{
		Data:    a.Data.ToStruct(),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
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
	return &GroupHonorInfoGRPC{
		GroupId:          a.GroupId,
		CurrentTalkative: a.CurrentTalkative.ToGRPC(),
		TalkativeList:    HonorOwnerInfoArray2HonorOwnerInfoGRPCArray(a.TalkativeList),
		PerformerList:    HonorOwnerInfoArray2HonorOwnerInfoGRPCArray(a.PerformerList),
		LegendList:       HonorOwnerInfoArray2HonorOwnerInfoGRPCArray(a.LegendList),
		StrongNewbieList: HonorOwnerInfoArray2HonorOwnerInfoGRPCArray(a.StrongNewbieList),
		EmotionList:      HonorOwnerInfoArray2HonorOwnerInfoGRPCArray(a.EmotionList),
	}
}

func (a *GroupHonorInfoGRPC) ToStruct() *GroupHonorInfo {
	return &GroupHonorInfo{
		GroupId:          a.GroupId,
		CurrentTalkative: a.CurrentTalkative.ToStruct(),
		TalkativeList:    HonorOwnerInfoGRPCArray2HonorOwnerInfoArray(a.TalkativeList),
		PerformerList:    HonorOwnerInfoGRPCArray2HonorOwnerInfoArray(a.PerformerList),
		LegendList:       HonorOwnerInfoGRPCArray2HonorOwnerInfoArray(a.LegendList),
		StrongNewbieList: HonorOwnerInfoGRPCArray2HonorOwnerInfoArray(a.StrongNewbieList),
		EmotionList:      HonorOwnerInfoGRPCArray2HonorOwnerInfoArray(a.EmotionList),
	}
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
