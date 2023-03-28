package message

import (
	"fmt"

	"github.com/dezhishen/onebot-sdk/pkg/channel"
	"github.com/dezhishen/onebot-sdk/pkg/config"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type ChannelApiMessageClient struct {
	channel.ApiChannel
	conf *config.OnebotApiConfig
}

func NewChannelApiMessageClient(channel channel.ApiChannel, conf *config.OnebotApiConfig) (OnebotApiMessageClient, error) {
	return &ChannelApiMessageClient{
		ApiChannel: channel,
		conf:       conf,
	}, nil
}

// 发送私信
// send_private_msg
// msg 消息
func (cli *ChannelApiMessageClient) SendPrivateMsg(msg *model.PrivateMsg) (*model.SendMessageResult, error) {
	var result model.SendMessageResult
	if err := cli.PostByRequestForResult(API_SEND_PRIVATE_MSG, msg, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 发送群消息
// send_group_msg
// msg 消息
func (cli *ChannelApiMessageClient) SendGroupMsg(msg *model.GroupMsg) (*model.SendMessageResult, error) {
	var result model.SendMessageResult
	if err := cli.PostByRequestForResult(API_SEND_GROUP_MSG, msg, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 发送消息
func (cli *ChannelApiMessageClient) SendMsg(msg *model.MsgForSend) (*model.SendMessageResult, error) {
	if msg.MessageType == model.PrivateMessageType {
		return cli.SendPrivateMsg(&model.PrivateMsg{
			UserId:     msg.UserId,
			Message:    msg.Message,
			AutoEscape: msg.AutoEscape,
		})
	}
	if msg.MessageType == model.GroupMessageType {
		return cli.SendGroupMsg(&model.GroupMsg{
			GroupId:    msg.GroupId,
			Message:    msg.Message,
			AutoEscape: msg.AutoEscape,
		})
	}
	return nil, fmt.Errorf("message type error")
}

// 获取消息
// get_msg
func (cli *ChannelApiMessageClient) GetMsg(id int64) (*model.MessageDataResult, error) {
	var result model.MessageDataResult
	if err := cli.PostByRequestForResult(API_GET_MSG, map[string]interface{}{
		"message_id": id,
	}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 撤回消息
// delete_msg
func (cli *ChannelApiMessageClient) DeleteMsg(id int64) error {
	return cli.PostByRequest(API_DELETE_MSG, map[string]interface{}{
		"message_id": id,
	})
}

// 标记消息已读
// mark_msg_as_read
func (cli *ChannelApiMessageClient) MarkMsgAsRead(id int64) error {
	return cli.PostByRequest(API_MARK_MSG_AS_READ, map[string]interface{}{
		"message_id": id,
	})
}

// 获取合并转发消息
// get_forward_msg
func (cli *ChannelApiMessageClient) GetForwardMsg(id int64) (*model.ForwardMessageDataResult, error) {
	var result model.ForwardMessageDataResult
	if err := cli.PostByRequestForResult(API_GET_FORWARD_MSG, map[string]interface{}{
		"id": id,
	}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 发送合并转发 ( 群聊 )
// send_group_forward_msg
func (cli *ChannelApiMessageClient) SendGroupForwardMsg(groupId int64, messages []*model.MessageSegment) (*model.SendForwardMessageDataResult, error) {
	var result model.SendForwardMessageDataResult
	if err := cli.PostByRequestForResult(API_SEND_GROUP_FORWARD_MSG, map[string]interface{}{
		"group_id": groupId,
		"messages": messages,
	}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 发送合并转发 ( 好友 )
// send_private_forward_msg
func (cli *ChannelApiMessageClient) SendPrivateForwardMsg(userId int64, messages []*model.MessageSegment) (*model.SendForwardMessageDataResult, error) {
	var result model.SendForwardMessageDataResult
	if err := cli.PostByRequestForResult(API_SEND_PRIVATE_FORWARD_MSG, map[string]interface{}{
		"user_id":  userId,
		"messages": messages,
	}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 获取群消息历史记录
// get_group_msg_history
func (cli *ChannelApiMessageClient) GetGroupMsgHistory(message_seq int64, groupId int64) (*model.MessagesResult, error) {
	var result model.MessagesResult
	if err := cli.PostByRequestForResult(API_GET_GROUP_MSG_HISTORY, map[string]interface{}{
		"message_seq": message_seq,
		"group_id":    groupId,
	}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
