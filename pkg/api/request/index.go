package request

type OnebotApiRequestClient interface {
	// 处理加好友请求
	// set_friend_add_request
	// flag: 加好友请求的 flag（需从上报的数据中获得）
	// approve: 是否同意请求
	// remark: 添加后的好友备注（仅在同意时有效）
	SetFriendAddRequest(flag string, approve bool, remark string) error
	// 处理加群请求／邀请
	// set_group_add_request
	// flag: 加群请求的 flag（需从上报的数据中获得）
	// sub_type: add 或 invite，请求类型（需要和上报消息中的 sub_type 字段相符）
	// approve: 是否同意请求／邀请
	// reason: 拒绝理由（仅在拒绝时有效）
	SetGroupAddRequest(flag string, subType string, approve bool, reason string) error
}
