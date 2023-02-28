package friendopt

import "github.com/dezhishen/onebot-sdk/pkg/channel"

type ChannelApiFriendOptClient struct {
	channel.ApiChannel
}

func NewChannelApiFriendOptClient(channel channel.ApiChannel) (OnebotApiFriendOptClient, error) {
	return &ChannelApiFriendOptClient{
		channel,
	}, nil
} // 删除好友
// delete_friend
// user_id: 对方 QQ 号
func (cli *ChannelApiFriendOptClient) DeleteFriend(userId int64) error {
	return cli.PostByRequest(API_DELETE_FRIEND, map[string]interface{}{
		"user_id": userId,
	})
}

// 删除单向好友
// delete_unidirectional_friend
// user_id: 对方 QQ 号
func (cli *ChannelApiFriendOptClient) DeleteUnidirectionalFriend(userId int64) error {
	return cli.PostByRequest(API_DELETE_UNIDIRECTIONAL_FRIEND, map[string]interface{}{
		"user_id": userId,
	})
}
