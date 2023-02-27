package group

import (
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type OnebotApiGroupClient interface {
	// 获取群信息
	// get_group_info
	// group_id: 群号
	// no_cache: 是否不使用缓存（使用缓存可能更新不及时，但响应更快）
	GetGroupInfo(groupId int64, noCache bool) (*model.GroupInfoResult, error)
	// 获取群列表
	// get_group_list
	// no_cache: 是否不使用缓存（使用缓存可能更新不及时，但响应更快）
	GetGroupList(noCache bool) (*model.GroupListResult, error)
	// 获取群成员信息
	// get_group_member_info
	// group_id: 群号
	// user_id: QQ 号
	// no_cache: 是否不使用缓存（使用缓存可能更新不及时，但响应更快）
	GetGroupMemberInfo(groupId int64, userId int64, noCache bool) (*model.GroupMemberInfoResult, error)
	// 获取群成员列表
	// get_group_member_list
	// group_id: 群号
	// no_cache: 是否不使用缓存（使用缓存可能更新不及时，但响应更快）
	GetGroupMemberList(groupId int64, noCache bool) (*model.GroupMemberListResult, error)
	// 获取群荣誉信息
	// get_group_honor_info
	// group_id: 群号
	// type: 群荣誉类型，目前支持 talkative（群聊之火）、performer（群聊炽焰）、legend（群聊传说）、strong_newbie（群聊新星）、emotion（群表情之火）
	GetGroupHonorInfo(groupId int64, honorType string) (*model.GroupHonorInfoResult, error)
	// 获取群系统消息
	// get_group_system_msg
	GetGroupSystemMsg() (*model.GroupSystemMsgResult, error)
	// 获取精华消息列表
	// get_essence_msg_list
	// group_id: 群号
	GetEssenceMsgList(groupId int64) (*model.EssenceMsgListResult, error)
	// 获取群 @全体成员 剩余次数
	// get_group_at_all_remain
	// group_id: 群号
	GetGroupAtAllRemain(groupId int64) (*model.GroupAtAllRemainResult, error)
}
