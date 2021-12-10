package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/dezhishen/onebot-sdk/pkg/config"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

// 群组踢人
func SetGroupKick(groupID, userID int, rejectAddRequest bool) error {
	req := make(map[string]interface{})
	req["user_id"] = userID
	req["group_id"] = groupID
	req["reject_add_request"] = rejectAddRequest
	requestBody, _ := json.Marshal(req)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_group_kick",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

// 群组禁言
func SetGroupBan(groupID, userID, duration int) error {
	req := make(map[string]int)
	req["user_id"] = userID
	req["group_id"] = groupID
	req["duration"] = duration
	requestBody, _ := json.Marshal(req)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_group_ban",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

// 群组匿名用户禁言
func SetGroupAnonymousBan(groupID, duration int, anonymousFlag string) error {
	req := make(map[string]interface{})
	req["group_id"] = groupID
	req["anonymous_flag"] = anonymousFlag
	req["duration"] = duration
	requestBody, _ := json.Marshal(req)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_group_anonymous_ban",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

//群组全员禁言
func SetGroupWholeBan(groupID int, enable bool) error {
	req := make(map[string]interface{})
	req["group_id"] = groupID
	req["enable"] = enable
	requestBody, _ := json.Marshal(req)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_group_whole_ban",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

//群组设置管理员
func SetGroupAdmin(groupID, userID int, enable bool) error {
	req := make(map[string]interface{})
	req["group_id"] = groupID
	req["user_id"] = userID
	req["enable"] = enable
	requestBody, _ := json.Marshal(req)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_group_admin",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

//群组匿名
func SetGroupAnonymous(groupID int, enable bool) error {
	req := make(map[string]interface{})
	req["group_id"] = groupID
	req["enable"] = enable
	requestBody, _ := json.Marshal(req)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_group_anonymous",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

//设置群名片（群备注）
func SetGroupCard(groupID, userID int, card string) error {
	req := make(map[string]interface{})
	req["group_id"] = groupID
	req["user_id"] = userID
	req["card"] = card
	requestBody, _ := json.Marshal(req)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_group_card",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

//设置群名
func SetGroupName(groupID int, groupName string) error {
	req := make(map[string]interface{})
	req["group_id"] = groupID
	req["group_name"] = groupName
	requestBody, _ := json.Marshal(req)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_group_name",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

//退出群组
func SetGroupLeave(groupID int, isDismiss bool) error {
	req := make(map[string]interface{})
	req["group_id"] = groupID
	req["is_dismiss"] = isDismiss
	requestBody, _ := json.Marshal(req)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_group_leave",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

//设置群组专属头衔
func SetGroupSpecialTitle(groupID, userID, duration int, specialTitle string) error {
	req := make(map[string]interface{})
	req["group_id"] = groupID
	req["user_id"] = userID
	req["special_title"] = specialTitle
	req["duration"] = duration
	requestBody, _ := json.Marshal(req)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_group_special_title",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

//处理加群请求／邀请
func SetGroupAddRequest(flag, subType, reason string, approve bool) error {
	req := make(map[string]interface{})
	req["flag"] = flag
	req["sub_type"] = subType
	req["approve"] = approve
	req["reason"] = reason
	requestBody, _ := json.Marshal(req)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_group_add_request",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

//获取群信息
func GetGroupInfo(groupID int, noCache bool) (*model.Group, error) {
	req := make(map[string]interface{})
	req["group_id"] = groupID
	req["no_cache"] = noCache
	requestBody, _ := json.Marshal(req)
	resp, err := http.Post(
		config.GetHttpUrl()+"/get_group_info",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	var result model.GroupResult
	json.Unmarshal(respBodyContent, &result)
	return result.Data, nil
}

//获取群列表
func GetGroupList() ([]model.Group, error) {
	resp, err := http.Post(
		config.GetHttpUrl()+"/get_group_list",
		"application/json",
		nil,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	var result model.GroupListResult
	json.Unmarshal(respBodyContent, &result)
	return result.Data, nil
}

//获取群成员信息
func GetGroupMemberInfo(groupID, userID int, noCache bool) (*model.GroupMember, error) {
	req := make(map[string]interface{})
	req["group_id"] = groupID
	req["user_id"] = userID
	req["no_cache"] = noCache
	requestBody, _ := json.Marshal(req)
	resp, err := http.Post(
		config.GetHttpUrl()+"/get_group_member_info",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	var result model.GroupMemberResult
	json.Unmarshal(respBodyContent, &result)
	return result.Data, nil
}

//获取群成员列表
func GetGroupMemberListInfo() ([]model.GroupMember, error) {
	resp, err := http.Post(
		config.GetHttpUrl()+"/get_group_member_list",
		"application/json",
		nil,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	var result model.GroupMemberListResult
	json.Unmarshal(respBodyContent, &result)
	return result.Data, nil
}

//获取群荣誉信息
func GetGroupHonorInfo(groupID int, honorType string) (*model.GroupHonorInfo, error) {
	req := make(map[string]interface{})
	req["group_id"] = groupID
	req["type"] = honorType
	requestBody, _ := json.Marshal(req)
	resp, err := http.Post(
		config.GetHttpUrl()+"/get_group_honor_info",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	var result model.GroupHonorInfoResult
	json.Unmarshal(respBodyContent, &result)
	return result.Data, nil

}
