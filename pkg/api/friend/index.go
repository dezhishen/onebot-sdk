package friend

import (
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type OnebotApiFriendClient interface {
	GetStrangerInfo(userId int64, noCache bool) (*model.AccountResult, error)
	GetFriendList() (*model.FriendListResult, error)
	GetUnidirectionalFriendList() (*model.FriendListResult, error)
	DeleteFriend(userId int64) error
	DeleteUnidirectionalFriend(userId int64) error
}
