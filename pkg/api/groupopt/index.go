package groupopt

import "github.com/dezhishen/onebot-sdk/pkg/model"

type OnebotApiGroupOptClient interface {
	//set_group_name
	SetGroupName(groupId int64, groupName string) error

	//set_group_portrait
	SetGroupPortrait(groupId int64, file string, cache int) error
	//set_group_admin
	SetGroupAdmin(groupId int64, userId int64, enable bool) error

	// set_group_card
	// 设置群名片
	// groupId 群号
	// userId QQ 号
	// card 群名片内容，不填或空字符串表示删除群名片
	SetGroupCard(groupId int64, userId int64, card string) error
	// set_group_special_title
	// 设置群头衔
	// groupId 群号
	// userId QQ 号
	// specialTitle 头衔，不填或空字符串表示删除群头衔
	// duration 专属头衔有效期, 单位秒, -1 表示永久, 不过此项似乎没有效果, 可能是只有某些特殊的时间长度有效, 有待测试
	SetGroupSpecialTitle(groupId int64, userId int64, specialTitle string, duration uint32) error

	// set_group_ban
	// 禁言群成员
	// groupId 群号
	// userId QQ 号
	// duration 禁言时长，单位秒，0 表示取消禁言
	SetGroupBan(groupId int64, userId int64, duration uint32) error

	// set_group_whole_ban
	// 设置全群禁言
	// groupId 群号
	// enable 是否禁言
	SetGroupWholeBan(groupId int64, enable bool) error

	// set_group_anonymous_ban
	// 禁言群匿名成员
	// groupId 群号
	// anonymous 可选, 要禁言的匿名用户对象（群消息上报的 anonymous 字段）
	// anonymousFlag 可选, 要禁言的匿名用户的 flag（需从群消息上报的数据中获得）
	// 上面的 anonymous 和 anonymous_flag 两者任选其一传入即可, 若都传入, 则使用 anonymous。
	// duration 禁言时长，单位秒，不能超过 30 天
	SetGroupAnonymousBan(groupId int64, anonymous interface{}, anonymousFlag string, duration uint32) error

	// set_essence_msg
	// 设置精华消息
	// message_id 消息 ID
	SetEssenceMsg(messageId int64) error
	// delete_essence_msg
	// 删除精华消息
	// message_id 消息 ID
	DeleteEssenceMsg(messageId int64) error

	// send_group_sign
	// 发送群签到
	// groupId 群号
	SendGroupSign(groupId int64) error

	// set_group_anonymous
	// 设置群匿名
	// groupId 群号
	// enable 是否允许匿名聊天
	SetGroupAnonymous(groupId int64, enable bool) error
	// _send_group_notice
	// 发送群公告
	// groupId 群号
	// content 公告内容
	// image 图片文件路径（可选）
	SendGroupNotice(groupId int64, content string, image string) error
	// _get_group_notice
	// 获取群公告
	// groupId 群号
	GetGroupNotice(groupId int64) (*model.GroupNoticeResult, error)
	// set_group_kick
	// 群组踢人
	// groupId 群号
	// userId QQ 号
	// rejectAddRequest 是否拒绝此人的加群请求
	SetGroupKick(groupId int64, userId int64, rejectAddRequest bool) error

	// set_group_leave
	// 退出群组
	// groupId 群号
	// isDismiss 是否解散，如果登录号是群主，则仅在此项为 true 时能够解散
	SetGroupLeave(groupId int64, isDismiss bool) error
}
