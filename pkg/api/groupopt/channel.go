package groupopt

import (
	"github.com/dezhishen/onebot-sdk/pkg/channel"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type ChannelApiGroupOptClient struct {
	channel.ApiChannel
}

func NewChannelApiGroupOptClient(channel channel.ApiChannel) (OnebotApiGroupOptClient, error) {
	return &ChannelApiGroupOptClient{
		channel,
	}, nil
}

// 设置群名
// set_group_name
// groupId 群号
func (cli *ChannelApiGroupOptClient) SetGroupName(groupId int64, groupName string) error {
	return cli.PostByRequest(API_SET_GROUP_NAME, map[string]interface{}{
		"group_id":   groupId,
		"group_name": groupName,
	})
}

// 设置群头像
// set_group_portrait
// groupId 群号
// file 图片文件路径
// cache 缓存时间
func (cli *ChannelApiGroupOptClient) SetGroupPortrait(groupId int64, file string, cache int) error {
	return cli.PostByRequest(API_SET_GROUP_PORTRAIT, map[string]interface{}{
		"group_id": groupId,
		"file":     file,
		"cache":    cache,
	})
}

// set_group_admin
// 设置群管理员
// groupId 群号
// userId QQ 号
// enable true 为设置，false 为取消
func (cli *ChannelApiGroupOptClient) SetGroupAdmin(groupId int64, userId int64, enable bool) error {
	return cli.PostByRequest(API_SET_GROUP_ADMIN, map[string]interface{}{
		"group_id": groupId,
		"user_id":  userId,
		"enable":   enable,
	})
}

// 设置群名片
// set_group_card
// groupId 群号
// userId QQ 号
// card 群名片内容，不填或空字符串表示删除群名片
func (cli *ChannelApiGroupOptClient) SetGroupCard(groupId int64, userId int64, card string) error {
	return cli.PostByRequest(API_SET_GROUP_CARD, map[string]interface{}{
		"group_id": groupId,
		"user_id":  userId,
		"card":     card,
	})
}

// 设置群头衔
// set_group_special_title
// groupId 群号
// userId QQ 号
// specialTitle 头衔，不填或空字符串表示删除群头衔
// duration 专属头衔有效期, 单位秒, -1 表示永久, 不过此项似乎没有效果, 可能是只有某些特殊的时间长度有效, 有待测试
func (cli *ChannelApiGroupOptClient) SetGroupSpecialTitle(groupId int64, userId int64, specialTitle string, duration uint32) error {
	return cli.PostByRequest(API_SET_GROUP_SPECIAL_TITLE, map[string]interface{}{
		"group_id":      groupId,
		"user_id":       userId,
		"special_title": specialTitle,
		"duration":      duration,
	})
}

// 禁言群成员
// set_group_ban
// groupId 群号
// userId QQ 号
// duration 禁言时长，单位秒，0 表示取消禁言
func (cli *ChannelApiGroupOptClient) SetGroupBan(groupId int64, userId int64, duration uint32) error {
	return cli.PostByRequest(API_SET_GROUP_BAN, map[string]interface{}{
		"group_id": groupId,
		"user_id":  userId,
		"duration": duration,
	})
}

// 设置全群禁言
// set_group_whole_ban
// groupId 群号
// enable 是否禁言
func (cli *ChannelApiGroupOptClient) SetGroupWholeBan(groupId int64, enable bool) error {
	return cli.PostByRequest(API_SET_GROUP_WHOLE_BAN, map[string]interface{}{
		"group_id": groupId,
		"enable":   enable,
	})
}

// 禁言群匿名成员
// set_group_anonymous_ban
// groupId 群号
// anonymous 可选, 要禁言的匿名用户对象（群消息上报的 anonymous 字段）
// anonymousFlag 可选, 要禁言的匿名用户的 flag（需从群消息上报的数据中获得）
// 上面的 anonymous 和 anonymous_flag 两者任选其一传入即可, 若都传入, 则使用 anonymous。
// duration 禁言时长，单位秒，不能超过 30 天
func (cli *ChannelApiGroupOptClient) SetGroupAnonymousBan(groupId int64, anonymous interface{}, anonymousFlag string, duration uint32) error {
	return cli.PostByRequest(API_SET_GROUP_ANONYMOUS_BAN, map[string]interface{}{
		"group_id":       groupId,
		"anonymous":      anonymous,
		"anonymous_flag": anonymousFlag,
		"duration":       duration,
	})
}

// 设置精华消息
// set_essence_msg
// message_id 消息 ID
func (cli *ChannelApiGroupOptClient) SetEssenceMsg(messageId int64) error {
	return cli.PostByRequest(API_SET_ESSENCE_MSG, map[string]interface{}{
		"message_id": messageId,
	})
}

// 删除精华消息
// delete_essence_msg
// message_id 消息 ID
func (cli *ChannelApiGroupOptClient) DeleteEssenceMsg(messageId int64) error {
	return cli.PostByRequest(API_DELETE_ESSENCE_MSG, map[string]interface{}{
		"message_id": messageId,
	})
}

// 发送群签到
// send_group_sign
// groupId 群号
func (cli *ChannelApiGroupOptClient) SendGroupSign(groupId int64) error {
	return cli.PostByRequest(API_SEND_GROUP_SIGN, map[string]interface{}{
		"group_id": groupId,
	})
}

// 设置群匿名
// set_group_anonymous
// groupId 群号
// enable 是否允许匿名聊天
func (cli *ChannelApiGroupOptClient) SetGroupAnonymous(groupId int64, enable bool) error {
	return cli.PostByRequest(API_SET_GROUP_ANONYMOUS, map[string]interface{}{
		"group_id": groupId,
		"enable":   enable,
	})
}

// 发送群公告
// _send_group_notice
// groupId 群号
// content 公告内容
// image 图片文件路径（可选）
func (cli *ChannelApiGroupOptClient) SendGroupNotice(groupId int64, content string, image string) error {
	return cli.PostByRequest(API__SEND_GROUP_NOTICE, map[string]interface{}{
		"group_id": groupId,
		"content":  content,
		"image":    image,
	})
}

// 获取群公告
// _get_group_notice
// groupId 群号
func (cli *ChannelApiGroupOptClient) GetGroupNotice(groupId int64) (*model.GroupNoticeResult, error) {
	var result model.GroupNoticeResult
	if err := cli.PostByRequestForResult(API__GET_GROUP_NOTICE, map[string]interface{}{
		"group_id": groupId,
	}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// 群组踢人
// set_group_kick
// groupId 群号
// userId QQ 号
// rejectAddRequest 是否拒绝此人的加群请求
func (cli *ChannelApiGroupOptClient) SetGroupKick(groupId int64, userId int64, rejectAddRequest bool) error {
	return cli.PostByRequest(API_SET_GROUP_KICK, map[string]interface{}{
		"group_id":           groupId,
		"user_id":            userId,
		"reject_add_request": rejectAddRequest,
	})
}

// 退出群组
// set_group_leave
// groupId 群号
// isDismiss 是否解散，如果登录号是群主，则仅在此项为 true 时能够解散
func (cli *ChannelApiGroupOptClient) SetGroupLeave(groupId int64, isDismiss bool) error {
	return cli.PostByRequest(API_SET_GROUP_LEAVE, map[string]interface{}{
		"group_id":   groupId,
		"is_dismiss": isDismiss,
	})
}
