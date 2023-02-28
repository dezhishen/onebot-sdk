package image

import (
	"github.com/dezhishen/onebot-sdk/pkg/channel"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type ChannelApiImageClient struct {
	channel.ApiChannel
}

func NewChannelApiImageClient(channel channel.ApiChannel) (OnebotApiImageClient, error) {
	return &ChannelApiImageClient{
		channel,
	}, nil
}

// 获取图片
// get_image
func (cli *ChannelApiImageClient) GetImage(file string) (*model.ImageResult, error) {
	var result model.ImageResult
	if err := cli.PostByRequestForResult(API_GET_IMAGE, map[string]interface{}{
		"file": file,
	}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 检查是否可以发送图片
// can_send_image
func (cli *ChannelApiImageClient) CanSendImage() (*model.BoolYesOfResult, error) {
	var result model.BoolYesOfResult
	if err := cli.PostByRequestForResult(API_CAN_SEND_IMAGE, map[string]interface{}{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 图片ORC识别
// ocr_image
func (cli *ChannelApiImageClient) OcrImage(image string) (*model.OcrImageResult, error) {
	var result model.OcrImageResult
	if err := cli.PostByRequestForResult(API_OCR_IMAGE, map[string]interface{}{
		"image": image,
	}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
