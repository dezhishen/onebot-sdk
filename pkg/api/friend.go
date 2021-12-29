package api

import (
	"github.com/dezhishen/onebot-sdk/pkg/cli"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

func SendLike(userId int64, times int64) error {
	req := make(map[string]int64)
	req["user_id"] = userId
	req["times"] = times
	url := "/send_like"
	return cli.PostWithRequest(url, req)
}

//处理加好友请求
func SetFriendAddRequest(flag string, approve bool, remark string) error {
	req := make(map[string]interface{})
	req["flag"] = flag
	req["approve"] = approve
	req["remark"] = remark
	url := "/set_friend_add_request"
	return cli.PostWithRequest(url, req)
}

func GetFriendList() ([]*model.Account, error) {
	var result model.FriendListResult
	url := "/get_friend_list"
	if err := cli.PostForResult(url, &result); err != nil {
		return nil, err
	}
	return result.Data, nil
}
