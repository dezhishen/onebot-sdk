package message

import "github.com/dezhishen/onebot-sdk/pkg/model"

type OnebotApiMessageClient interface {
	// 发送私信
	// send_private_msg
	// msg 消息
	SendPrivateMsg(msg *model.PrivateMsg) (*model.SendMessageResult, error)
	// 发送群消息
	// send_group_msg
	// msg 消息
	SendGroupMsg(msg *model.GroupMsg) (*model.SendMessageResult, error)
	// 发送消息
	SendMsg(msg *model.MsgForSend) (*model.SendMessageResult, error)
	// 获取消息
	// get_msg
	GetMsg(id int64) (*model.MessageDataResult, error)
	// 撤回消息
	// delete_msg
	DeleteMsg(id int64) error
	// 标记消息已读
	// mark_msg_as_read
	MarkMsgAsRead(id int64) error
	// 获取合并转发消息
	// get_forward_msg
	GetForwardMsg(id int64) (*model.ForwardMessageDataResult, error)
	// 发送合并转发 ( 群聊 )
	// send_group_forward_msg
	SendGroupForwardMsg(groupId int64, messages []*model.MessageSegment) (*model.SendForwardMessageDataResult, error)
	// 发送合并转发 ( 好友 )
	// send_private_forward_msg
	SendPrivateForwardMsg(userId int64, messages []*model.MessageSegment) (*model.SendForwardMessageDataResult, error)
	// 获取群消息历史记录
	// get_group_msg_history
	GetGroupMsgHistory(message_seq, groupId int64) (*model.MessagesResult, error)
}
