package api

import (
	"github.com/dezhishen/onebot-sdk/pkg/cli"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

// 群组踢人
func SetGroupKick(groupId, userId int64, rejectAddRequest bool) error {
	req := make(map[string]interface{})
	req["user_id"] = userId
	req["group_id"] = groupId
	req["reject_add_request"] = rejectAddRequest

	url := "/set_group_kick"
	if err := cli.PostWithRequest(url, req); err != nil {
		return err
	}

	return nil
}

// 群组禁言
func SetGroupBan(groupId, userId, duration int64) error {
	req := make(map[string]int64)
	req["user_id"] = userId
	req["group_id"] = groupId
	req["duration"] = duration

	url := "/set_group_ban"
	if err := cli.PostWithRequest(url, req); err != nil {
		return err
	}
	return nil
}

// 群组匿名用户禁言
func SetGroupAnonymousBan(groupId, duration int64, anonymousFlag string) error {
	req := make(map[string]interface{})
	req["group_id"] = groupId
	req["anonymous_flag"] = anonymousFlag
	req["duration"] = duration
	url := "/set_group_anonymous_ban"
	if err := cli.PostWithRequest(url, req); err != nil {
		return err
	}
	return nil
}

//群组全员禁言
func SetGroupWholeBan(groupId int64, enable bool) error {
	req := make(map[string]interface{})
	req["group_id"] = groupId
	req["enable"] = enable

	url := "/set_group_whole_ban"
	if err := cli.PostWithRequest(url, req); err != nil {
		return err
	}
	return nil

}

//群组设置管理员
func SetGroupAdmin(groupId, userId int64, enable bool) error {
	req := make(map[string]interface{})
	req["group_id"] = groupId
	req["user_id"] = userId
	req["enable"] = enable

	url := "/set_group_admin"
	if err := cli.PostWithRequest(url, req); err != nil {
		return err
	}
	return nil

}

//群组匿名
func SetGroupAnonymous(groupId int64, enable bool) error {
	req := make(map[string]interface{})
	req["group_id"] = groupId
	req["enable"] = enable

	url := "/set_group_anonymous"
	if err := cli.PostWithRequest(url, req); err != nil {
		return err
	}
	return nil
}

//设置群名片（群备注）
func SetGroupCard(groupId, userId int64, card string) error {
	req := make(map[string]interface{})
	req["group_id"] = groupId
	req["user_id"] = userId
	req["card"] = card

	url := "/set_group_card"
	if err := cli.PostWithRequest(url, req); err != nil {
		return err
	}
	return nil
}

//设置群名
func SetGroupName(groupId int64, groupName string) error {
	req := make(map[string]interface{})
	req["group_id"] = groupId
	req["group_name"] = groupName

	url := "/set_group_name"
	if err := cli.PostWithRequest(url, req); err != nil {
		return err
	}
	return nil
}

//退出群组
func SetGroupLeave(groupId int64, isDismiss bool) error {
	req := make(map[string]interface{})
	req["group_id"] = groupId
	req["is_dismiss"] = isDismiss

	url := "/set_group_leave"
	if err := cli.PostWithRequest(url, req); err != nil {
		return err
	}
	return nil
}

//设置群组专属头衔
func SetGroupSpecialTitle(groupId, userId, duration int64, specialTitle string) error {
	req := make(map[string]interface{})
	req["group_id"] = groupId
	req["user_id"] = userId
	req["special_title"] = specialTitle
	req["duration"] = duration

	url := "/set_group_special_title"
	if err := cli.PostWithRequest(url, req); err != nil {
		return err
	}
	return nil
}

//处理加群请求／邀请
func SetGroupAddRequest(flag, subType, reason string, approve bool) error {
	req := make(map[string]interface{})
	req["flag"] = flag
	req["sub_type"] = subType
	req["approve"] = approve
	req["reason"] = reason

	url := "/set_group_add_request"
	if err := cli.PostWithRequest(url, req); err != nil {
		return err
	}
	return nil
}

//获取群信息
func GetGroupInfo(groupId int64, noCache bool) (*model.GroupResult, error) {
	req := make(map[string]interface{})
	req["group_id"] = groupId
	req["no_cache"] = noCache

	var result model.GroupResult
	url := "/get_group_info"
	if err := cli.PostWithRequsetForResult(url, req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

//获取群列表
func GetGroupList() (*model.GroupListResult, error) {
	var result model.GroupListResult
	url := "/get_group_list"

	if err := cli.PostForResult(url, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

//获取群成员信息
func GetGroupMemberInfo(groupId, userId int64, noCache bool) (*model.GroupMemberResult, error) {
	req := make(map[string]interface{})
	req["group_id"] = groupId
	req["user_id"] = userId
	req["no_cache"] = noCache

	var result model.GroupMemberResult
	url := "/get_group_member_info"
	if err := cli.PostWithRequsetForResult(url, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

//获取群成员列表
func GetGroupMemberListInfo() (*model.GroupMemberListResult, error) {
	var result model.GroupMemberListResult
	url := "/get_group_member_list"

	if err := cli.PostForResult(url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

//获取群荣誉信息
func GetGroupHonorInfo(groupId int64, honorType string) (*model.GroupHonorInfoResult, error) {
	req := make(map[string]interface{})
	req["group_id"] = groupId
	req["type"] = honorType

	var result model.GroupHonorInfoResult
	url := "/get_group_honor_info"
	if err := cli.PostWithRequsetForResult(url, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
