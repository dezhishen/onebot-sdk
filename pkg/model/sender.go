package model

type Sender struct {
	//	发送者 QQ 号
	UserID int64 `json:"user_id"`
	//昵称
	Nickname string `json:"nickname"`
	//群名片／备注
	Card string `json:"card"`
	//性别，male 或 female 或 unknown
	Sex string `json:"sex"`
	//	年龄
	Age int32 `json:"age"`
	//群名片／备注
	Area string `json:"area"`
	//成员等级
	Level string `json:"level"`
	//角色 owner/admin/member
	Role string `json:"role"`
	//头衔
	Title string `json:"title"`
}
