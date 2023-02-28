package record

import (
	"github.com/dezhishen/onebot-sdk/pkg/channel"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type ChannelApiRecordClient struct {
	channel.ApiChannel
}

func NewChannelApiRecordClient(channel channel.ApiChannel) (OnebotApiRecordClient, error) {
	return &ChannelApiRecordClient{
		channel,
	}, nil
}

// 获取语音
// get_record
func (cli *ChannelApiRecordClient) GetRecord(file string, outFormat string) (*model.RecordResult, error) {
	var result model.RecordResult
	if err := cli.PostByRequestForResult(API_GET_RECORD, map[string]interface{}{
		"file":       file,
		"out_format": outFormat,
	}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 检查是否可以发送语音
// can_send_record
func (cli *ChannelApiRecordClient) CanSendRecord() (*model.BoolYesOfResult, error) {
	var result model.BoolYesOfResult
	if err := cli.PostForResult(API_CAN_SEND_RECORD, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
