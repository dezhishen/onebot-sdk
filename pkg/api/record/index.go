package record

import (
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type OnebotApiRecordClient interface {
	// 获取语音
	// get_record
	GetRecord(file string, outFormat string) (*model.RecordResult, error)
	// 检查是否可以发送语音
	// can_send_record
	CanSendRecord() (*model.BoolYesOfResult, error)
}
