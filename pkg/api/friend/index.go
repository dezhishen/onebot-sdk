package friend

import (
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type OnebotApiFriendClient interface {
	// 获取陌生人信息
	// get_stranger_info
	// user_id: 对方 QQ 号
	// no_cache: 是否不使用缓存（使用缓存可能更新不及时，但响应更快）
	GetStrangerInfo(userId int64, noCache bool) (*model.StrangerInfoResult, error)
	// 获取好友列表
	// get_friend_list
	GetFriendList() (*model.FriendListResult, error)
	// 获取单向好友列表
	// get_unidirectional_friend_list
	GetUnidirectionalFriendList() (*model.FriendListResult, error)
}
