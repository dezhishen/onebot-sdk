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

type CredentialsResult struct {
	Data    *Credentials `json:"data"`
	RetCode int          `json:"retcode"`
	Status  string       `json:"status"`
}

type File struct {
	File string `json:"file"`
}

type FileResult struct {
	Data    *File  `json:"data"`
	RetCode int    `json:"retcode"`
	Status  string `json:"status"`
}
