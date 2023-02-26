package image

import "github.com/dezhishen/onebot-sdk/pkg/model"

type OnebotApiImageClient interface {
	// 获取图片
	// get_image
	GetImage(file string) (*model.ImageResult, error)
	// 检查是否可以发送图片
	// can_send_image
	CanSendImage() (*model.BoolYesOfResult, error)
	// 图片ORC识别
	// ocr_image
	OcrImage(image string) (*model.OcrImageResult, error)
}
