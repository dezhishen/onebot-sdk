package api

// type onebotApiMessageClient interface {
// 	// 发送消息
// 	SendMsg(msg *model.MsgForSend) (*model.SendMessageResult, error)
// 	// 发送私信
// 	SendPrivateMsg(msg *model.PrivateMsg) (*model.SendMessageResult, error)
// 	// 发送群消息
// 	SendGroupMsg(msg *model.GroupMsg) (*model.SendMessageResult, error)
// 	// 撤回消息
// 	DelMsg(id int64) error
// 	// 获取消息
// 	GetMsg(id int64) (*model.MessageDataResult, error)
// 	// 获取合并转发消息
// 	GetForwardMsg(id int64) (*model.ForwardMessageDataResult, error)
// 	// 发送合并转发消息
// 	SendGroupForwardMsgByRawMsg(groupId, uin int64, name string, msg []*model.MessageSegment) (*model.SendForwardMessageDataResult, error)
// 	// 发送合并转发消息
// 	SendGroupForwardMsg(groupId int64, messages []*model.MessageSegment) (*model.SendForwardMessageDataResult, error)
// }

// type httpOnebotApiMessageClient struct {
// 	*client.HttpCli
// }

// func NewOnebotApiMessageClient(conf *config.OnebotConfig) (onebotApiMessageClient, error) {
// 	return &httpOnebotApiMessageClient{
// 		client.NewHttpCli(conf),
// 	}, nil
// }

// // 发送消息
// func (cli *httpOnebotApiMessageClient) SendMsg(msg *model.MsgForSend) (*model.SendMessageResult, error) {
// 	if msg.MessageType == model.PrivateMessageType {
// 		return cli.SendPrivateMsg(&model.PrivateMsg{
// 			UserId:     msg.UserId,
// 			GroupId:    msg.GroupId,
// 			Message:    msg.Message,
// 			AutoEscape: msg.AutoEscape,
// 		})
// 	}
// 	return cli.SendGroupMsg(
// 		&model.GroupMsg{
// 			GroupId:    msg.GroupId,
// 			Message:    msg.Message,
// 			AutoEscape: msg.AutoEscape,
// 		},
// 	)
// }

// // 发送私信
// func (cli *httpOnebotApiMessageClient) SendPrivateMsg(msg *model.PrivateMsg) (*model.SendMessageResult, error) {
// 	var result model.SendMessageResult
// 	url := "/send_private_msg"
// 	if err := cli.PostWithRequsetForResult(url, msg, &result); err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }

// // 发送群消息
// func (cli *httpOnebotApiMessageClient) SendGroupMsg(msg *model.GroupMsg) (*model.SendMessageResult, error) {
// 	var result model.SendMessageResult
// 	url := "/send_group_msg"
// 	if err := cli.PostWithRequsetForResult(url, msg, &result); err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }

// func (cli *httpOnebotApiMessageClient) DelMsg(id int64) error {
// 	req := make(map[string]int64)
// 	req["message_id"] = id
// 	url := "/delete_msg"
// 	err := cli.PostWithRequest(
// 		url,
// 		req,
// 	)
// 	return err
// }

// func (cli *httpOnebotApiMessageClient) GetMsg(id int64) (*model.MessageDataResult, error) {
// 	req := make(map[string]int64)
// 	req["message_id"] = id
// 	var result model.MessageDataResult
// 	url := "/get_msg"
// 	if err := cli.PostWithRequsetForResult(url, req, &result); err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }

// func (cli *httpOnebotApiMessageClient) GetForwardMsg(id int64) (*model.ForwardMessageDataResult, error) {
// 	req := make(map[string]int64)
// 	req["message_id"] = id
// 	url := "/get_forward_msg"
// 	var result model.ForwardMessageDataResult
// 	if err := cli.PostWithRequsetForResult(url, req, &result); err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }

// func (cli *httpOnebotApiMessageClient) SendGroupForwardMsgByRawMsg(groupId, uin int64, name string, msg []*model.MessageSegment) (*model.SendForwardMessageDataResult, error) {
// 	var forwardMsg []*model.MessageSegment
// 	forwardMsg = append(forwardMsg, &model.MessageSegment{
// 		Type: "node",
// 		Data: &model.MessageElementNode{
// 			Uin:     uin,
// 			Name:    name,
// 			Content: msg,
// 		},
// 	})
// 	return cli.SendGroupForwardMsg(groupId, forwardMsg)
// }

// func (cli *httpOnebotApiMessageClient) SendGroupForwardMsg(groupId int64, messages []*model.MessageSegment) (*model.SendForwardMessageDataResult, error) {
// 	req := make(map[string]interface{})
// 	req["group_id"] = groupId
// 	req["messages"] = messages
// 	var result model.SendForwardMessageDataResult
// 	url := "/send_group_forward_msg"
// 	if err := cli.PostWithRequsetForResult(url, req, &result); err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }
