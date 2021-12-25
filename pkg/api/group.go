package api

import (
	"github.com/dezhishen/onebot-sdk/pkg/cli"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

// 群组踢人
func SetGroupKick(groupID, userID int, rejectAddRequest bool) error {
	req := make(map[string]interface{})
	req["user_id"] = userID
	req["group_id"] = groupID
	req["reject_add_request"] = rejectAddRequest

	url := "/set_group_kick"
	if err := cli.PostWithRequest(url, req); err != nil {
		return err
	}

	return nil
}

// 群组禁言
func SetGroupBan(groupID, userID, duration int) error {
	req := make(map[string]int)
	req["user_id"] = userID
	req["group_id"] = groupID
	req["duration"] = duration

	url := "/set_group_ban"
	if err := cli.PostWithRequest(url, req); err != nil {
		return err
	}
	return nil
}

// 群组匿名用户禁言
func SetGroupAnonymousBan(groupID, duration int, anonymousFlag string) error {
	req := make(map[string]interface{})
	req["group_id"] = groupID
	req["anonymous_flag"] = anonymousFlag
	req["duration"] = duration
	url := "/set_group_anonymous_ban"
	if err := cli.PostWithRequest(url, req); err != nil {
		return err
	}
	return nil
}

//群组全员禁言
func SetGroupWholeBan(groupID int, enable bool) error {
	req := make(map[string]interface{})
	req["group_id"] = groupID
	req["enable"] = enable

	url := "/set_group_whole_ban"
	if err := cli.PostWithRequest(url, req); err != nil {
		return err
	}
	return nil

}

//群组设置管理员
func SetGroupAdmin(groupID, userID int, enable bool) error {
	req := make(map[string]interface{})
	req["group_id"] = groupID
	req["user_id"] = userID
	req["enable"] = enable

	url := "/set_group_admin"
	if err := cli.PostWithRequest(url, req); err != nil {
		return err
	}
	return nil

}

//群组匿名
func SetGroupAnonymous(groupID int, enable bool) error {
	req := make(map[string]interface{})
	req["group_id"] = groupID
	req["enable"] = enable

	url := "/set_group_anonymous"
	if err := cli.PostWithRequest(url, req); err != nil {
		return err
	}
	return nil
}

//设置群名片（群备注）
func SetGroupCard(groupID, userID int, card string) error {
	req := make(map[string]interface{})
	req["group_id"] = groupID
	req["user_id"] = userID
	req["card"] = card

	url := "/set_group_card"
	if err := cli.PostWithRequest(url, req); err != nil {
		return err
	}
	return nil
}

//设置群名
func SetGroupName(groupID int, groupName string) error {
	req := make(map[string]interface{})
	req["group_id"] = groupID
	req["group_name"] = groupName

	url := "/set_group_name"
	if err := cli.PostWithRequest(url, req); err != nil {
		return err
	}
	return nil
}

//退出群组
func SetGroupLeave(groupID int, isDismiss bool) error {
	req := make(map[string]interface{})
	req["group_id"] = groupID
	req["is_dismiss"] = isDismiss

	url := "/set_group_leave"
	if err := cli.PostWithRequest(url, req); err != nil {
		return err
	}
	return nil
}

//设置群组专属头衔
func SetGroupSpecialTitle(groupID, userID, duration int, specialTitle string) error {
	req := make(map[string]interface{})
	req["group_id"] = groupID
	req["user_id"] = userID
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
func GetGroupInfo(groupID int, noCache bool) (*model.Group, error) {
	req := make(map[string]interface{})
	req["group_id"] = groupID
	req["no_cache"] = noCache

	var result model.GroupResult
	url := "/get_group_info"
	if err := cli.PostWithRequsetForResult(url, req, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

//获取群列表
func GetGroupList() ([]*model.Group, error) {
	var result model.GroupListResult
	url := "/get_group_list"

	if err := cli.PostForResult(url, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

//获取群成员信息
func GetGroupMemberInfo(groupID, userID int, noCache bool) (*model.GroupMember, error) {
	req := make(map[string]interface{})
	req["group_id"] = groupID
	req["user_id"] = userID
	req["no_cache"] = noCache

	var result model.GroupMemberResult
	url := "/get_group_member_info"
	if err := cli.PostWithRequsetForResult(url, req, &result); err != nil {
		return nil, err
	}
	return result.Data, nil
}

//获取群成员列表
func GetGroupMemberListInfo() ([]*model.GroupMember, error) {
	var result model.GroupMemberListResult
	url := "/get_group_member_list"

	if err := cli.PostForResult(url, &result); err != nil {
		return nil, err
	}
	return result.Data, nil
}

//获取群荣誉信息
func GetGroupHonorInfo(groupID int, honorType string) (*model.GroupHonorInfo, error) {
	req := make(map[string]interface{})
	req["group_id"] = groupID
	req["type"] = honorType

	var result model.GroupHonorInfoResult
	url := "/get_group_honor_info"
	if err := cli.PostWithRequsetForResult(url, req, &result); err != nil {
		return nil, err
	}
	return result.Data, nil
}
