package group

import (
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type OnebotApiGroupClient interface {
	// get_group_info
	GetGroupInfo(groupId int64, noCache bool) (*model.GroupInfoResult, error)
	// get_group_list
	GetGroupList(noCache bool) (*model.GroupListResult, error)
	// get_group_member_info
	GetGroupMemberInfo(groupId int64, userId int64, noCache bool) (*model.GroupMemberInfoResult, error)
	// get_group_member_list
	GetGroupMemberList(groupId int64, noCache bool) (*model.GroupMemberListResult, error)
	// get_group_honor_info
	GetGroupHonorInfo(groupId int64, honorType string) (*model.GroupHonorInfoResult, error)
	// get_group_system_msg
	GetGroupSystemMsg() (*model.GroupSystemMsgResult, error)
	// get_essence_msg_list
	// 获取精华消息列表
	GetEssenceMsgList(groupId int64) (*model.EssenceMsgListResult, error)
	// get_group_at_all_remain
	// 获取群 @全体成员 剩余次数
	GetGroupAtAllRemain(groupId int64) (*model.GroupAtAllRemainResult, error)
}
