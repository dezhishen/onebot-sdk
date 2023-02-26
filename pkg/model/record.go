package model

type Record struct {
	File      string `json:"file"`
	OutFormat string `json:"out_format"`
}

func (a *Record) ToGRPC() *RecordGRPC {
	return &RecordGRPC{
		File:      a.File,
		OutFormat: a.OutFormat,
	}
}

func (a *RecordGRPC) ToStruct() *Record {
	return &Record{
		File:      a.File,
		OutFormat: a.OutFormat,
	}
}

type RecordResult struct {
	Data    *Record `json:"data"`
	Retcode int64   `json:"retcode"`
	Status  string  `json:"status"`
	Msg     string  `json:"msg"`
	Wording string  `json:"wording"`
}

func (a *RecordResult) ToGRPC() *RecordResultGRPC {
	result := &RecordResultGRPC{
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

func (a *RecordResultGRPC) ToStruct() *RecordResult {
	result := &RecordResult{
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
