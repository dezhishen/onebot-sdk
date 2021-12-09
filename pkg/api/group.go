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
	reqMap := make(map[string]interface{})
	reqMap["user_id"] = userID
	reqMap["group_id"] = groupID
	reqMap["reject_add_request"] = rejectAddRequest
	requestBody, _ := json.Marshal(reqMap)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_group_kick",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

// 群组禁言
func SetGroupBan(groupID, userID, duration int) error {
	reqMap := make(map[string]int)
	reqMap["user_id"] = userID
	reqMap["group_id"] = groupID
	reqMap["duration"] = duration
	requestBody, _ := json.Marshal(reqMap)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_group_ban",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

// 群组匿名用户禁言
func SetGroupAnonymousBan(groupID, duration int, anonymousFlag string) error {
	reqMap := make(map[string]interface{})
	reqMap["group_id"] = groupID
	reqMap["anonymous_flag"] = anonymousFlag
	reqMap["duration"] = duration
	requestBody, _ := json.Marshal(reqMap)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_group_anonymous_ban",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

func SetGroupWholeBan(groupID int, enable bool) error {
	reqMap := make(map[string]interface{})
	reqMap["group_id"] = groupID
	reqMap["enable"] = enable
	requestBody, _ := json.Marshal(reqMap)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_group_whole_ban",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

func SetGroupAdmin(groupID, userID int, enable bool) error {
	reqMap := make(map[string]interface{})
	reqMap["group_id"] = groupID
	reqMap["user_id"] = userID
	reqMap["enable"] = enable
	requestBody, _ := json.Marshal(reqMap)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_group_admin",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

func SetGroupAnonymous(groupID int, enable bool) error {
	reqMap := make(map[string]interface{})
	reqMap["group_id"] = groupID
	reqMap["enable"] = enable
	requestBody, _ := json.Marshal(reqMap)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_group_anonymous",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

func SetGroupCard(groupID, userID int, card string) error {
	reqMap := make(map[string]interface{})
	reqMap["group_id"] = groupID
	reqMap["user_id"] = userID
	reqMap["card"] = card
	requestBody, _ := json.Marshal(reqMap)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_group_card",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

func SetGroupName(groupID int, groupName string) error {
	reqMap := make(map[string]interface{})
	reqMap["group_id"] = groupID
	reqMap["group_name"] = groupName
	requestBody, _ := json.Marshal(reqMap)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_group_name",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

func SetGroupLeave(groupID int, isDismiss bool) error {
	reqMap := make(map[string]interface{})
	reqMap["group_id"] = groupID
	reqMap["is_dismiss"] = isDismiss
	requestBody, _ := json.Marshal(reqMap)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_group_leave",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

func SetGroupSpecialTitle(groupID, userID, duration int, specialTitle string) error {
	reqMap := make(map[string]interface{})
	reqMap["group_id"] = groupID
	reqMap["user_id"] = userID
	reqMap["special_title"] = specialTitle
	reqMap["duration"] = duration
	requestBody, _ := json.Marshal(reqMap)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_group_special_title",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

func SetGroupAddRequest(flag, subType, reason string, approve bool) error {
	reqMap := make(map[string]interface{})
	reqMap["flag"] = flag
	reqMap["sub_type"] = subType
	reqMap["approve"] = approve
	reqMap["reason"] = reason
	requestBody, _ := json.Marshal(reqMap)
	_, err := http.Post(
		config.GetHttpUrl()+"/set_group_add_request",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

func GetGroupInfo(groupID int, noCache bool) (*model.Group, error) {
	reqMap := make(map[string]interface{})
	reqMap["group_id"] = groupID
	reqMap["no_cache"] = noCache
	requestBody, _ := json.Marshal(reqMap)
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

func GetGroupMemberInfo(groupID, userID int, noCache bool) (*model.GroupMember, error) {
	reqMap := make(map[string]interface{})
	reqMap["group_id"] = groupID
	reqMap["user_id"] = userID
	reqMap["no_cache"] = noCache
	requestBody, _ := json.Marshal(reqMap)
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

func GetGroupHonorInfo(groupID int, honorType string) (*model.GroupHonorInfo, error) {
	reqMap := make(map[string]interface{})
	reqMap["group_id"] = groupID
	reqMap["type"] = honorType
	requestBody, _ := json.Marshal(reqMap)
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
