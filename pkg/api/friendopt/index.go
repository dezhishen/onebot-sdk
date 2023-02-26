package friendopt

type OnebotApiFriendOptClient interface {
	DeleteFriend(userId int64) error
	DeleteUnidirectionalFriend(userId int64) error
}
