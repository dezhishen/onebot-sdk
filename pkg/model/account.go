package model

type AccountResult struct {
	Data    *Account `json:"data"`
	RetCode int      `json:"retcode"`
	Status  string   `json:"status"`
}

type FriendListResult struct {
	Data    []Account `json:"data"`
	RetCode int       `json:"retcode"`
	Status  string    `json:"status"`
}

type Account struct {
	UserID   int    `json:"user_id"`
	Sex      string `json:"sex"`
	Nickname string `json:"nickname"`
	Age      int    `json:"age"`
	Remark   string `json:"remark"`
}

type Credentials struct {
	Cookies   string `json:"cookies"`
	CSRFToken int32  `json:"csrf_token"`
}
