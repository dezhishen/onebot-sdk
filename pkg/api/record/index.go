package record

import (
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type OnebotApiRecordClient interface {
	// get_record
	GetRecord(file string, outFormat string) (*model.RecordResult, error)
	// can_send_record
	CanSendRecord() (*model.BoolYesOfResult, error)
}
