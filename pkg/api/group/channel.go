package group

import (
	"github.com/dezhishen/onebot-sdk/pkg/channel"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type ChannelApiGroupClient struct {
	channel.ApiChannel
}

func NewChannelApiGroupClient(channel channel.ApiChannel) (OnebotApiGroupClient, error) {
	return &ChannelApiGroupClient{
		channel,
	}, nil
}

// 获取群信息
// get_group_info
// group_id: 群号
// no_cache: 是否不使用缓存（使用缓存可能更新不及时，但响应更快）
func (cli *ChannelApiGroupClient) GetGroupInfo(groupId int64, noCache bool) (*model.GroupInfoResult, error) {
	var result model.GroupInfoResult
	if err := cli.PostByRequestForResult(API_GET_GROUP_INFO, map[string]interface{}{
		"group_id": groupId,
		"no_cache": noCache,
	}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 获取群列表
// get_group_list
// no_cache: 是否不使用缓存（使用缓存可能更新不及时，但响应更快）
func (cli *ChannelApiGroupClient) GetGroupList(noCache bool) (*model.GroupListResult, error) {
	var result model.GroupListResult
	if err := cli.PostByRequestForResult(API_GET_GROUP_LIST, map[string]interface{}{
		"no_cache": noCache,
	}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 获取群成员信息
// get_group_member_info
// group_id: 群号
// user_id: QQ 号
// no_cache: 是否不使用缓存（使用缓存可能更新不及时，但响应更快）
func (cli *ChannelApiGroupClient) GetGroupMemberInfo(groupId int64, userId int64, noCache bool) (*model.GroupMemberInfoResult, error) {
	var result model.GroupMemberInfoResult
	if err := cli.PostByRequestForResult(API_GET_GROUP_MEMBER_INFO, map[string]interface{}{
		"group_id": groupId,
		"user_id":  userId,
		"no_cache": noCache,
	}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 获取群成员列表
// get_group_member_list
// group_id: 群号
// no_cache: 是否不使用缓存（使用缓存可能更新不及时，但响应更快）
func (cli *ChannelApiGroupClient) GetGroupMemberList(groupId int64, noCache bool) (*model.GroupMemberListResult, error) {
	var result model.GroupMemberListResult
	if err := cli.PostByRequestForResult(API_GET_GROUP_MEMBER_LIST, map[string]interface{}{
		"group_id": groupId,
		"no_cache": noCache,
	}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 获取群荣誉信息
// get_group_honor_info
// group_id: 群号
// type: 群荣誉类型，目前支持 talkative（群聊之火）、performer（群聊炽焰）、legend（群聊传说）、strong_newbie（群聊新星）、emotion（群表情之火）
func (cli *ChannelApiGroupClient) GetGroupHonorInfo(groupId int64, honorType string) (*model.GroupHonorInfoResult, error) {
	var result model.GroupHonorInfoResult
	if err := cli.PostByRequestForResult(API_GET_GROUP_HONOR_INFO, map[string]interface{}{
		"group_id": groupId,
		"type":     honorType,
	}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 获取群系统消息
// get_group_system_msg
func (cli *ChannelApiGroupClient) GetGroupSystemMsg() (*model.GroupSystemMsgResult, error) {
	var result model.GroupSystemMsgResult
	if err := cli.PostByRequestForResult(API_GET_GROUP_SYSTEM_MSG, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 获取精华消息列表
// get_essence_msg_list
// group_id: 群号
func (cli *ChannelApiGroupClient) GetEssenceMsgList(groupId int64) (*model.EssenceMsgListResult, error) {
	var result model.EssenceMsgListResult
	if err := cli.PostByRequestForResult(API_GET_ESSENCE_MSG_LIST, map[string]interface{}{
		"group_id": groupId,
	}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 获取群 @全体成员 剩余次数
// get_group_at_all_remain
// group_id: 群号
func (cli *ChannelApiGroupClient) GetGroupAtAllRemain(groupId int64) (*model.GroupAtAllRemainResult, error) {
	var result model.GroupAtAllRemainResult
	if err := cli.PostByRequestForResult(API_GET_GROUP_AT_ALL_REMAIN, map[string]interface{}{
		"group_id": groupId,
	}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
