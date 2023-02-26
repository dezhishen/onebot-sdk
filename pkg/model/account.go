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
	Msg     string   `json:"msg"`
	Wording string   `json:"wording"`
}

func (a *AccountResult) ToGRPC() *AccountResultGRPC {
	res := &AccountResultGRPC{
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
	if a.Data != nil {
		res.Data = a.Data.ToGRPC()
	}
	return res
}

func (a *AccountResultGRPC) ToStruct() *AccountResult {
	res := &AccountResult{
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
	if a.Data != nil {
		res.Data = a.Data.ToStruct()
	}
	return res
}

type FriendListResult struct {
	Data    []*Account `json:"data"`
	Retcode int64      `json:"retcode"`
	Status  string     `json:"status"`
	Msg     string     `json:"msg"`
	Wording string     `json:"wording"`
}

func (a *FriendListResult) ToGRPC() *FriendListResultGRPC {
	return &FriendListResultGRPC{
		Data:    AccountArray2AccountGRPCArray(a.Data),
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
}

func (a *FriendListResultGRPC) ToStruct() *FriendListResult {
	return &FriendListResult{
		Data:    AccountGRPCArray2AccountArray(a.Data),
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
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
	Msg     string  `json:"msg"`
	Wording string  `json:"wording"`
}

func (a *CokiesResult) ToGRPC() *CokiesResultGRPC {
	result := &CokiesResultGRPC{
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

func (a *CokiesResultGRPC) ToStruct() *CokiesResult {
	result := &CokiesResult{
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
	Msg     string     `json:"msg"`
	Wording string     `json:"wording"`
}

func (a *CSRFTokenResult) ToGRPC() *CSRFTokenResultGRPC {
	result := &CSRFTokenResultGRPC{
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

func (a *CSRFTokenResultGRPC) ToStruct() *CSRFTokenResult {
	result := &CSRFTokenResult{
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
	Msg     string       `json:"msg"`
	Wording string       `json:"wording"`
}

func (a *CredentialsResult) ToGRPC() *CredentialsResultGRPC {
	result := &CredentialsResultGRPC{
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

func (a *CredentialsResultGRPC) ToStruct() *CredentialsResult {
	result := &CredentialsResult{
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

type Variant struct {
	ModelShow string `json:"model_show"`
	NeedPay   bool   `json:"need_pay"`
}

func (a *Variant) ToGRPC() *VariantGRPC {
	return &VariantGRPC{
		ModelShow: a.ModelShow,
		NeedPay:   a.NeedPay,
	}
}

func (a *VariantGRPC) ToStruct() *Variant {
	return &Variant{
		ModelShow: a.ModelShow,
		NeedPay:   a.NeedPay,
	}
}

type ModelShow struct {
	Variants []Variant `json:"variants"`
}

func (a *ModelShow) ToGRPC() *ModelShowGRPC {
	result := &ModelShowGRPC{}
	for _, v := range a.Variants {
		result.Variants = append(result.Variants, v.ToGRPC())
	}
	return result
}

func (a *ModelShowGRPC) ToStruct() *ModelShow {
	result := &ModelShow{}
	for _, v := range a.Variants {
		result.Variants = append(result.Variants, *v.ToStruct())
	}
	return result
}

type ModelShowResult struct {
	Data    *ModelShow `json:"data"`
	Retcode int64      `json:"retcode"`
	Status  string     `json:"status"`
	Msg     string     `json:"msg"`
	Wording string     `json:"wording"`
}

func (a *ModelShowResult) ToGRPC() *ModelShowResultGRPC {
	result := &ModelShowResultGRPC{
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

func (a *ModelShowResultGRPC) ToStruct() *ModelShowResult {
	result := &ModelShowResult{
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

type QQProfileResult struct {
	Data    *QQProfile `json:"data"`
	Retcode int64      `json:"retcode"`
	Status  string     `json:"status"`
	Msg     string     `json:"msg"`
	Wording string     `json:"wording"`
}

type QQProfile struct {
	Nickname     string `json:"nickname"`
	Company      string `json:"company"`
	Email        string `json:"email"`
	Collage      string `json:"collage"`
	PersonalNote string `json:"personal_note"`
}

func (a *QQProfile) ToGRPC() *QQProfileGRPC {
	return &QQProfileGRPC{
		Nickname:     a.Nickname,
		Company:      a.Company,
		Email:        a.Email,
		Collage:      a.Collage,
		PersonalNote: a.PersonalNote,
	}
}

func (a *QQProfileGRPC) ToStruct() *QQProfile {
	return &QQProfile{
		Nickname:     a.Nickname,
		Company:      a.Company,
		Email:        a.Email,
		Collage:      a.Collage,
		PersonalNote: a.PersonalNote,
	}
}

func (a *QQProfileResult) ToGRPC() *QQProfileResultGRPC {
	result := &QQProfileResultGRPC{
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

func (a *QQProfileResultGRPC) ToStruct() *QQProfileResult {
	result := &QQProfileResult{
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
