package account

import (
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type OnebotApiAccountClient interface {
	GetLoginInfo() (*model.AccountResult, error)
	SetQQProfile(profile *model.QQProfile) error
	GetModelShow() (*model.ModelShowResult, error)
	// _set_model_show
	SetModelShow(model, modelshow string) error
	// get_online_clients //todo
	// GetOnlineClients(noCache bool) (*model.OnlineClientsResult, error)
}
