package friend

import (
	"github.com/dezhishen/onebot-sdk/pkg/channel"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type ChannelApiFriendClient struct {
	channel.ApiChannel
}

func NewChannelApiFriendClient(channel channel.ApiChannel) (OnebotApiFriendClient, error) {
	return &ChannelApiFriendClient{
		channel,
	}, nil
}

// 获取陌生人信息
// get_stranger_info
// user_id: 对方 QQ 号
// no_cache: 是否不使用缓存（使用缓存可能更新不及时，但响应更快）
func (cli *ChannelApiFriendClient) GetStrangerInfo(userId int64, noCache bool) (*model.StrangerInfoResult, error) {
	var result model.StrangerInfoResult
	if err := cli.PostByRequestForResult(API_GET_STRANGER_INFO, map[string]interface{}{
		"user_id":  userId,
		"no_cache": noCache,
	}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 获取好友列表
// get_friend_list
func (cli *ChannelApiFriendClient) GetFriendList() (*model.FriendListResult, error) {
	var result model.FriendListResult
	if err := cli.PostForResult(API_GET_FRIEND_LIST, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 获取单向好友列表
// get_unidirectional_friend_list
func (cli *ChannelApiFriendClient) GetUnidirectionalFriendList() (*model.FriendListResult, error) {
	var result model.FriendListResult
	if err := cli.PostForResult(API_GET_UNIDIRECTIONAL_FRIEND_LIST, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
