package request

type OnebotApiRequestClient interface {
	// set_friend_add_request
	SetFriendAddRequest(flag string, approve bool, remark string) error
	// set_group_add_request
	SetGroupAddRequest(flag string, subType string, approve bool, reason string) error
}
