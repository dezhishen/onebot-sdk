package api

// import (
// 	"github.com/dezhishen/onebot-sdk/pkg/client"
// 	"github.com/dezhishen/onebot-sdk/pkg/config"
// 	"github.com/dezhishen/onebot-sdk/pkg/model"
// )

// type onebotApiFriendClient interface {
// 	GetFriendList() (*model.FriendListResult, error)
// 	SetFriendAddRequest(flag string, approve bool, remark string) error
// 	SendLike(userId int64, times int64) error
// }

// type httpOnebotApiFriendClient struct {
// 	*client.HttpCli
// }

// func NewOnebotApiFriendClient(conf *config.OnebotConfig) (onebotApiFriendClient, error) {
// 	return &httpOnebotApiFriendClient{
// 		client.NewHttpCli(conf),
// 	}, nil
// }

// func (cli *httpOnebotApiFriendClient) SendLike(userId int64, times int64) error {
// 	req := make(map[string]int64)
// 	req["user_id"] = userId
// 	req["times"] = times
// 	url := "/send_like"
// 	return cli.PostWithRequest(url, req)
// }

// // 处理加好友请求
// func (cli *httpOnebotApiFriendClient) SetFriendAddRequest(flag string, approve bool, remark string) error {
// 	req := make(map[string]interface{})
// 	req["flag"] = flag
// 	req["approve"] = approve
// 	req["remark"] = remark
// 	url := "/set_friend_add_request"
// 	return cli.PostWithRequest(url, req)
// }

// func (cli *httpOnebotApiFriendClient) GetFriendList() (*model.FriendListResult, error) {
// 	var result model.FriendListResult
// 	url := "/get_friend_list"
// 	if err := cli.PostForResult(url, &result); err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }
