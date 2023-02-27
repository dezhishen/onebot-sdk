package account

import (
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type OnebotApiAccountClient interface {
	// 获取登录号信息
	// get_login_info
	GetLoginInfo() (*model.AccountResult, error)
	// 设置登录号资料
	// set_login_info
	SetQQProfile(profile *model.QQProfile) error
	// 获取在线机型
	// _get_model_show
	GetModelShow() (*model.ModelShowResult, error)
	// 设置在线机型
	// _set_model_show
	// model: 机型
	// modelshow: 机型展示
	SetModelShow(model, modelshow string) error
	// 获取在线客户端列表
	// get_online_clients
	GetOnlineClients() (*model.OnlineClientsResult, error)
}
