package model

type BoolYes struct {
	Yes bool `json:"yes"`
}

func (a *BoolYes) ToGRPC() *BoolYesGRPC {
	return &BoolYesGRPC{
		Yes: a.Yes,
	}
}

func (a *BoolYesGRPC) ToStruct() *BoolYes {
	return &BoolYes{
		Yes: a.Yes,
	}
}

type BoolYesOfResult struct {
	Data    *BoolYes `json:"data"`
	Retcode int64    `json:"retcode"`
	Status  string   `json:"status"`
	Msg     string   `json:"msg"`
	Wording string   `json:"wording"`
}

func (a *BoolYesOfResult) ToGRPC() *BoolYesOfResultGRPC {
	result := &BoolYesOfResultGRPC{
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

func (a *BoolYesOfResultGRPC) ToStruct() *BoolYesOfResult {
	result := &BoolYesOfResult{
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

type Anonymous struct {
	// id	int64	匿名用户 ID
	Id int64 `json:"id"`
	// name	string	匿名用户名称
	Name string `json:"name"`
	// flag	string	匿名用户 flag, 在调用禁言 API 时需要传入
	Flag string `json:"flag"`
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
