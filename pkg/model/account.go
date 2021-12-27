package model

type Account struct {
	UserId int64 `json:"user_id"`
	//性别，male 或 female 或 unknown
	Sex string `json:"sex"`
	//昵称
	Nickname string `json:"nickname"`
	Age      uint32 `json:"age"`
	Remark   string `json:"remark"`
}

func (a *Account) ToGRPC() *AccountGRPC {
	return &AccountGRPC{
		UserId:   a.UserId,
		Sex:      a.Sex,
		Nickname: a.Nickname,
		Age:      a.Age,
		Remark:   a.Remark,
	}
}

func (a *AccountGRPC) ToStruct() *Account {
	return &Account{
		UserId:   a.UserId,
		Sex:      a.Sex,
		Nickname: a.Nickname,
		Age:      a.Age,
		Remark:   a.Remark,
	}
}

func AccountArray2AccountGRPCArray(source []*Account) []*AccountGRPC {
	var result []*AccountGRPC
	for _, e := range source {
		result = append(result, e.ToGRPC())
	}
	return result
}

func AccountGRPCArray2AccountArray(source []*AccountGRPC) []*Account {
	var result []*Account
	for _, e := range source {
		result = append(result, e.ToStruct())
	}
	return result
}

type AccountResult struct {
	Data    *Account `json:"data"`
	Retcode int64    `json:"retcode"`
	Status  string   `json:"status"`
}

func (a *AccountResult) ToGRPC() *AccountResultGRPC {
	return &AccountResultGRPC{
		Data:    a.Data.ToGRPC(),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

func (a *AccountResultGRPC) ToStruct() *AccountResult {
	return &AccountResult{
		Data:    a.Data.ToStruct(),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

type FriendListResult struct {
	Data    []*Account `json:"data"`
	Retcode int64      `json:"retcode"`
	Status  string     `json:"status"`
}

func (a *FriendListResult) ToGRPC() *FriendListResultGRPC {
	return &FriendListResultGRPC{
		Data:    AccountArray2AccountGRPCArray(a.Data),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

func (a *FriendListResultGRPC) ToStruct() *FriendListResult {
	return &FriendListResult{
		Data:    AccountGRPCArray2AccountArray(a.Data),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

type Anonymous struct {
	//匿名用户 Id
	Id int64 `json:"id"`
	//匿名用户 Id
	Name int64 `json:"name"`
	//匿名用户 flag，在调用禁言 API 时需要传入
	Flag int64 `json:"flag"`
}

func (a *Anonymous) ToGRPC() *AnonymousGRPC {
	return &AnonymousGRPC{
		Id:   a.Id,
		Name: a.Name,
		Flag: a.Flag,
	}
}

func (a *AnonymousGRPC) ToStruct() *Anonymous {
	return &Anonymous{
		Id:   a.Id,
		Name: a.Name,
		Flag: a.Flag,
	}
}

type Cokies struct {
	Cookies string `json:"cookies"`
}

func (a *Cokies) ToGRPC() *CokiesGRPC {
	return &CokiesGRPC{
		Cookies: a.Cookies,
	}
}

func (a *CokiesGRPC) ToStruct() *Cokies {
	return &Cokies{
		Cookies: a.Cookies,
	}
}

type CokiesResult struct {
	Data    *Cokies `json:"data"`
	Retcode int64   `json:"retcode"`
	Status  string  `json:"status"`
}

func (a *CokiesResult) ToGRPC() *CokiesResultGRPC {
	return &CokiesResultGRPC{
		Data:    a.Data.ToGRPC(),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

func (a *CokiesResultGRPC) ToStruct() *CokiesResult {
	return &CokiesResult{
		Data:    a.Data.ToStruct(),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

type CSRFToken struct {
	Token string `json:"token"`
}

func (a *CSRFToken) ToGRPC() *CSRFTokenGRPC {
	return &CSRFTokenGRPC{
		Token: a.Token,
	}
}

func (a *CSRFTokenGRPC) ToStruct() *CSRFToken {
	return &CSRFToken{
		Token: a.Token,
	}
}

type CSRFTokenResult struct {
	Data    *CSRFToken `json:"data"`
	Retcode int64      `json:"retcode"`
	Status  string     `json:"status"`
}

func (a *CSRFTokenResult) ToGRPC() *CSRFTokenResultGRPC {
	return &CSRFTokenResultGRPC{
		Data:    a.Data.ToGRPC(),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

func (a *CSRFTokenResultGRPC) ToStruct() *CSRFTokenResult {
	return &CSRFTokenResult{
		Data:    a.Data.ToStruct(),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

type Credentials struct {
	Cookies   string `json:"cookies"`
	CSRFToken int32  `json:"csrf_token"`
}

func (a *Credentials) ToGRPC() *CredentialsGRPC {
	return &CredentialsGRPC{
		Cookies:   a.Cookies,
		CSRFToken: a.CSRFToken,
	}
}

func (a *CredentialsGRPC) ToStruct() *Credentials {
	return &Credentials{
		Cookies:   a.Cookies,
		CSRFToken: a.CSRFToken,
	}
}

type CredentialsResult struct {
	Data    *Credentials `json:"data"`
	Retcode int64        `json:"retcode"`
	Status  string       `json:"status"`
}

func (a *CredentialsResult) ToGRPC() *CredentialsResultGRPC {
	return &CredentialsResultGRPC{
		Data:    a.Data.ToGRPC(),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

func (a *CredentialsResultGRPC) ToStruct() *CredentialsResult {
	return &CredentialsResult{
		Data:    a.Data.ToStruct(),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

type File struct {
	File string `json:"file"`
}

func (a *File) ToGRPC() *FileGRPC {
	return &FileGRPC{
		File: a.File,
	}
}

func (a *FileGRPC) ToStruct() *File {
	return &File{
		File: a.File,
	}
}

type FileResult struct {
	Data    *File  `json:"data"`
	Retcode int64  `json:"retcode"`
	Status  string `json:"status"`
}

func (a *FileResult) ToGRPC() *FileResultGRPC {
	return &FileResultGRPC{
		Data:    a.Data.ToGRPC(),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

func (a *FileResultGRPC) ToStruct() *FileResult {
	return &FileResult{
		Data:    a.Data.ToStruct(),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}
