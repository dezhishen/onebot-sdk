package friendopt

type OnebotApiFriendOptClient interface {
	// 删除好友
	// delete_friend
	// user_id: 对方 QQ 号
	DeleteFriend(userId int64) error
	// 删除单向好友
	// delete_unidirectional_friend
	// user_id: 对方 QQ 号
	DeleteUnidirectionalFriend(userId int64) error
}
