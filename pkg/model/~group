package model

type Group struct {
	GroupId        int    `json:"group_id"`
	GroupName      int    `json:"group_name"`
	MemberCount    string `json:"member_count"`
	MaxMemberCount string `json:"max_member_count"`
}

type GroupResult struct {
	Data    *Group `json:"data"`
	RetCode int    `json:"retcode"`
	Status  string `json:"status"`
}

type GroupListResult struct {
	Data    []Group `json:"data"`
	RetCode int     `json:"retcode"`
	Status  string  `json:"status"`
}

type GroupMember struct {
	GroupId         int    `json:"group_id"`
	UserID          int    `json:"user_id"`
	NickName        string `json:"nickname"`
	Card            string `json:"card"`
	Sex             int    `json:"sex"`
	Age             int    `json:"age"`
	Area            string `json:"area"`
	JoinTime        string `json:"join_time"`
	LastSentTime    int    `json:"last_sent_time"`
	Level           int    `json:"level"`
	Unfriendly      string `json:"unfriendly"`
	Titel           string `json:"title"`
	TitleExpireTime string `json:"title_expire_time"`
	CardChangeable  string `json:"card_changeable"`
}

type GroupMemberResult struct {
	Data    *GroupMember `json:"data"`
	RetCode int          `json:"retcode"`
	Status  string       `json:"status"`
}

type GroupMemberListResult struct {
	Data    []GroupMember `json:"data"`
	RetCode int           `json:"retcode"`
	Status  string        `json:"status"`
}

type GroupHonorInfoResult struct {
	Data    *GroupHonorInfo `json:"data"`
	RetCode int             `json:"retcode"`
	Status  string          `json:"status"`
}

type GroupHonorInfo struct {
	GroupId          int               `json:"group_id"`
	CurrentTalkative *CurrentTalkative `json:"current_talkative"`  //龙王
	TalkativeList    []HonorOwnerInfo  `json:"talkative_list"`     //历史龙王
	PerformerList    []HonorOwnerInfo  `json:"performer_list"`     //群聊之火
	LegendList       []HonorOwnerInfo  `json:"legend_list"`        //群聊炽焰
	StrongNewbieList []HonorOwnerInfo  `json:"strong_newbie_list"` //冒尖小春笋
	EmotionList      []HonorOwnerInfo  `json:"emotion_list"`       //快乐之源
}

type CurrentTalkative struct {
	UserID   int    `json:"user_id"`
	NickName string `json:"nickname"`
	Avatar   string `json:"avatar"`    //头像url
	DayCount int    `json:"day_count"` //持续天数
}

type HonorOwnerInfo struct {
	UserID      int    `json:"user_id"`
	NickName    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Description int    `json:"description"`
}
