package model

type StrangerInfo struct {
	// user_id	int64	QQ 号
	UserId int64 `json:"user_id"`
	// nickname	string	昵称
	Nickname string `json:"nickname"`
	// sex	string	性别, male 或 female 或 unknown
	Sex string `json:"sex"`
	// age	int32	年龄
	Age int32 `json:"age"`
	// qid	string	qid ID身份卡
	Qid string `json:"qid"`
	// level	int32	等级
	Level int32 `json:"level"`
	// login_days	int32	等级
	LoginDays int32 `json:"login_days"`
}

// ToGRPC
func (s *StrangerInfo) ToGRPC() *StrangerInfoGRPC {
	return &StrangerInfoGRPC{
		UserId:    s.UserId,
		Nickname:  s.Nickname,
		Sex:       s.Sex,
		Age:       s.Age,
		Qid:       s.Qid,
		Level:     s.Level,
		LoginDays: s.LoginDays,
	}
}

// ToStruct
func (s *StrangerInfoGRPC) ToStruct() *StrangerInfo {
	return &StrangerInfo{
		UserId:    s.UserId,
		Nickname:  s.Nickname,
		Sex:       s.Sex,
		Age:       s.Age,
		Qid:       s.Qid,
		Level:     s.Level,
		LoginDays: s.LoginDays,
	}
}

type StrangerInfoResult struct {
	Data    *StrangerInfo `json:"data"`
	Retcode int64         `json:"retcode"`
	Status  string        `json:"status"`
	Msg     string        `json:"msg"`
	Wording string        `json:"wording"`
}

// ToGRPC
func (s *StrangerInfoResult) ToGRPC() *StrangerInfoResultGRPC {
	res := &StrangerInfoResultGRPC{
		Retcode: s.Retcode,
		Status:  s.Status,
		Msg:     s.Msg,
		Wording: s.Wording,
	}
	if s.Data != nil {
		res.Data = s.Data.ToGRPC()
	}
	return res
}

// ToStruct
func (s *StrangerInfoResultGRPC) ToStruct() *StrangerInfoResult {
	res := &StrangerInfoResult{
		Retcode: s.Retcode,
		Status:  s.Status,
		Msg:     s.Msg,
		Wording: s.Wording,
	}
	if s.Data != nil {
		res.Data = s.Data.ToStruct()
	}
	return res
}

type Friend struct {
	// user_id	int64	QQ 号
	UserId int64 `json:"user_id"`
	// nickname	string	昵称
	Nickname string `json:"nickname"`
	// remark	string	好友备注
	Remark string `json:"remark"`
}

// ToGRPC
func (f *Friend) ToGRPC() *FriendGRPC {
	return &FriendGRPC{
		UserId:   f.UserId,
		Nickname: f.Nickname,
		Remark:   f.Remark,
	}
}

// ToStruct
func (f *FriendGRPC) ToStruct() *Friend {
	return &Friend{
		UserId:   f.UserId,
		Nickname: f.Nickname,
		Remark:   f.Remark,
	}
}

type FriendListResult struct {
	Data    []*Friend `json:"data"`
	Retcode int64     `json:"retcode"`
	Status  string    `json:"status"`
	Msg     string    `json:"msg"`
	Wording string    `json:"wording"`
}

// ToGRPC
func (f *FriendListResult) ToGRPC() *FriendListResultGRPC {
	res := &FriendListResultGRPC{
		Retcode: f.Retcode,
		Status:  f.Status,
		Msg:     f.Msg,
		Wording: f.Wording,
	}
	if f.Data != nil {
		for _, v := range f.Data {
			res.Data = append(res.Data, v.ToGRPC())
		}
	}
	return res
}

// ToStruct
func (f *FriendListResultGRPC) ToStruct() *FriendListResult {
	res := &FriendListResult{
		Retcode: f.Retcode,
		Status:  f.Status,
		Msg:     f.Msg,
		Wording: f.Wording,
	}
	if f.Data != nil {
		for _, v := range f.Data {
			res.Data = append(res.Data, v.ToStruct())
		}
	}
	return res
}
