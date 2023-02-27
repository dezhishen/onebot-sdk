package model

type Account struct {
	// QQ 号
	UserId int64 `json:"user_id"`
	// 昵称
	Nickname string `json:"nickname"`
}

func (a *Account) ToGRPC() *AccountGRPC {
	return &AccountGRPC{
		UserId:   a.UserId,
		Nickname: a.Nickname,
	}
}

func (a *AccountGRPC) ToStruct() *Account {
	return &Account{
		UserId:   a.UserId,
		Nickname: a.Nickname,
	}
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

type Device struct {
	// app_id	int64	客户端ID
	AppId int64 `json:"app_id"`
	// device_name	string	设备名称
	DeviceName string `json:"device_name"`
	// device_kind	string	设备类型
	DeviceKind string `json:"device_kind"`
}

func (a *Device) ToGRPC() *DeviceGRPC {
	return &DeviceGRPC{
		AppId:      a.AppId,
		DeviceName: a.DeviceName,
		DeviceKind: a.DeviceKind,
	}
}

func (a *DeviceGRPC) ToStruct() *Device {
	return &Device{
		AppId:      a.AppId,
		DeviceName: a.DeviceName,
		DeviceKind: a.DeviceKind,
	}
}

type OnlineClients struct {
	Clients []Device `json:"clients"`
}

func (a *OnlineClients) ToGRPC() *OnlineClientsGRPC {
	result := &OnlineClientsGRPC{}
	for _, v := range a.Clients {
		result.Clients = append(result.Clients, v.ToGRPC())
	}
	return result
}

func (a *OnlineClientsGRPC) ToStruct() *OnlineClients {
	result := &OnlineClients{}
	for _, v := range a.Clients {
		result.Clients = append(result.Clients, *v.ToStruct())
	}
	return result
}

type OnlineClientsResult struct {
	Data    *OnlineClients `json:"data"`
	Retcode int64          `json:"retcode"`
	Status  string         `json:"status"`
	Msg     string         `json:"msg"`
	Wording string         `json:"wording"`
}

func (a *OnlineClientsResult) ToGRPC() *OnlineClientsResultGRPC {
	result := &OnlineClientsResultGRPC{
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

func (a *OnlineClientsResultGRPC) ToStruct() *OnlineClientsResult {
	result := &OnlineClientsResult{
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
