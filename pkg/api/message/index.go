package message

import "github.com/dezhishen/onebot-sdk/pkg/model"

type OnebotApiMessageClient interface {
	// 发送私信
	SendPrivateMsg(msg *model.PrivateMsg) (*model.SendMessageResult, error)
	// 发送群消息
	SendGroupMsg(msg *model.GroupMsg) (*model.SendMessageResult, error)
	// 发送消息
	SendMsg(msg *model.MsgForSend) (*model.SendMessageResult, error)
	// 获取消息
	GetMsg(id int64) (*model.MessageDataResult, error)
	// 撤回消息
	DelMsg(id int64) error
	// 标记消息已读
	MarkMsgAsRead(id int64) error
	// 获取合并转发消息
	GetForwardMsg(id int64) (*model.ForwardMessageDataResult, error)
	// 发送合并转发 ( 群聊 )
	SendGroupForwardMsg(groupId int64, messages []*model.MessageSegment) (*model.SendForwardMessageDataResult, error)
	// 发送合并转发 ( 好友 )
	// send_private_forward_msg
	SendPrivateForwardMsg(userId int64, messages []*model.MessageSegment) (*model.SendForwardMessageDataResult, error)
	// 获取群消息历史记录
	// get_group_msg_history
	GetGroupMsgHistory(message_seq, groupId int64) (*model.MessagesResult, error)
}
