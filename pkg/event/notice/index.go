package notice

import (
	"encoding/json"

	"github.com/dezhishen/onebot-sdk/pkg/event/base"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type OnebotNoticeEventCli interface {
	base.OnebotBaseEventCli
	ListenNoticeGroupUpload(handler func(data model.EventNoticeGroupUpload) error)
	ListenNoticeGroupAdmin(handler func(data model.EventNoticeGroupAdmin) error)
	ListenNoticeGroupDecrease(handler func(data model.EventNoticeGroupDecrease) error)
	ListenNoticeGroupIncrease(handler func(data model.EventNoticeGroupIncrease) error)
	ListenNoticeGroupBan(handler func(data model.EventNoticeGroupBan) error)
	ListenNoticeFriendAdd(handler func(data model.EventNoticeFriendAdd) error)
	ListenNoticeGroupRecall(handler func(data model.EventNoticeGroupRecall) error)
	ListenNoticeFriendRecall(handler func(data model.EventNoticeFriendRecall) error)
	ListenNoticeGroupNotifyPoke(handler func(data model.EventNoticeGroupNotifyPoke) error)
	ListenNoticeGroupNotifyHonor(handler func(data model.EventNoticeGroupNotifyHonor) error)
	ListenNoticeGroupNotifyLuckyKing(handler func(data model.EventNoticeGroupNotifyLuckyKing) error)
}

type websocketNoticeEventCli struct {
	noticeGroupUploadHandlers          []func(data model.EventNoticeGroupUpload) error
	noticeGroupAdminHandlers           []func(data model.EventNoticeGroupAdmin) error
	noticeGroupDecreaseHandlers        []func(data model.EventNoticeGroupDecrease) error
	noticeGroupIncreaseHandlers        []func(data model.EventNoticeGroupIncrease) error
	noticeGroupBanHandlers             []func(data model.EventNoticeGroupBan) error
	noticeGroupRecallHandlers          []func(data model.EventNoticeGroupRecall) error
	noticeGroupNotifyPokeHandlers      []func(data model.EventNoticeGroupNotifyPoke) error
	noticeGroupNotifyLuckyKingHandlers []func(data model.EventNoticeGroupNotifyLuckyKing) error
	noticeGroupNotifyHonorHandlers     []func(data model.EventNoticeGroupNotifyHonor) error
	noticeFriendAddHandlers            []func(data model.EventNoticeFriendAdd) error
	noticeFriendRecallHandlers         []func(data model.EventNoticeFriendRecall) error
}

func NewOnebotNoticeEventCli() (OnebotNoticeEventCli, error) {
	return &websocketNoticeEventCli{}, nil
}

func (c *websocketNoticeEventCli) EventType() base.OnebotEventType {
	return base.OnebotEventTypeNotice
}

func (c *websocketNoticeEventCli) Handler(data []byte) error {
	return c.noticeHandler(data)
}

func (c *websocketNoticeEventCli) noticeHandler(data []byte) error {
	var notice model.EventNoticeBase
	err := json.Unmarshal(data, &notice)
	if err != nil {
		return err
	}
	if notice.NoticeType == "group_upload" {
		var relNotice model.EventNoticeGroupUpload
		err = json.Unmarshal(data, &relNotice)
		if err != nil {
			return err
		}
		for _, e := range c.noticeGroupUploadHandlers {
			err = e(relNotice)
			if err != nil {
				return err
			}
		}
	} else if notice.NoticeType == "group_admin" {
		var relNotice model.EventNoticeGroupAdmin
		err = json.Unmarshal(data, &relNotice)
		if err != nil {
			return err
		}
		for _, e := range c.noticeGroupAdminHandlers {
			err = e(relNotice)
			if err != nil {
				return err
			}
		}
	} else if notice.NoticeType == "group_decrease" {
		var relNotice model.EventNoticeGroupDecrease
		err = json.Unmarshal(data, &relNotice)
		if err != nil {
			return err
		}
		for _, e := range c.noticeGroupDecreaseHandlers {
			err = e(relNotice)
			if err != nil {
				return err
			}
		}
	} else if notice.NoticeType == "group_increase" {
		var relNotice model.EventNoticeGroupIncrease
		err = json.Unmarshal(data, &relNotice)
		if err != nil {
			return err
		}
		for _, e := range c.noticeGroupIncreaseHandlers {
			err = e(relNotice)
			if err != nil {
				return err
			}
		}
	} else if notice.NoticeType == "group_ban" {
		var relNotice model.EventNoticeGroupBan
		err = json.Unmarshal(data, &relNotice)
		if err != nil {
			return err
		}
		for _, e := range c.noticeGroupBanHandlers {
			err = e(relNotice)
			if err != nil {
				return err
			}
		}
	} else if notice.NoticeType == "friend_add" {
		var relNotice model.EventNoticeFriendAdd
		err = json.Unmarshal(data, &relNotice)
		if err != nil {
			return err
		}
		for _, e := range c.noticeFriendAddHandlers {
			err = e(relNotice)
			if err != nil {
				return err
			}
		}
	} else if notice.NoticeType == "group_recall" {
		var relNotice model.EventNoticeGroupRecall
		err = json.Unmarshal(data, &relNotice)
		if err != nil {
			return err
		}
		for _, e := range c.noticeGroupRecallHandlers {
			err = e(relNotice)
			if err != nil {
				return err
			}
		}
	} else if notice.NoticeType == "friend_recall" {
		var relNotice model.EventNoticeFriendRecall
		err = json.Unmarshal(data, &relNotice)
		if err != nil {
			return err
		}
		for _, e := range c.noticeFriendRecallHandlers {
			err = e(relNotice)
			if err != nil {
				return err
			}
		}
	} else if notice.NoticeType == "notify" {
		var notify model.EventNoticeNotifyBase
		err = json.Unmarshal(data, &notify)
		if err != nil {
			return err
		}
		if notify.SubType == "honor" {
			var relNotice model.EventNoticeGroupNotifyHonor
			err = json.Unmarshal(data, &relNotice)
			if err != nil {
				return err
			}
			for _, e := range c.noticeGroupNotifyHonorHandlers {
				err = e(relNotice)
				if err != nil {
					return err
				}
			}
		} else if notify.SubType == "lucky_king" {
			var relNotice model.EventNoticeGroupNotifyLuckyKing
			err = json.Unmarshal(data, &relNotice)
			if err != nil {
				return err
			}
			for _, e := range c.noticeGroupNotifyLuckyKingHandlers {
				err = e(relNotice)
				if err != nil {
					return err
				}
			}
		} else if notify.SubType == "poke" {
			var relNotice model.EventNoticeGroupNotifyPoke
			err = json.Unmarshal(data, &relNotice)
			if err != nil {
				return err
			}
			for _, e := range c.noticeGroupNotifyPokeHandlers {
				err = e(relNotice)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (c *websocketNoticeEventCli) ListenNoticeGroupUpload(handler func(data model.EventNoticeGroupUpload) error) {
	c.noticeGroupUploadHandlers = append(c.noticeGroupUploadHandlers, handler)
}

func (c *websocketNoticeEventCli) ListenNoticeGroupAdmin(handler func(data model.EventNoticeGroupAdmin) error) {
	c.noticeGroupAdminHandlers = append(c.noticeGroupAdminHandlers, handler)
}

func (c *websocketNoticeEventCli) ListenNoticeGroupDecrease(handler func(data model.EventNoticeGroupDecrease) error) {
	c.noticeGroupDecreaseHandlers = append(c.noticeGroupDecreaseHandlers, handler)
}

func (c *websocketNoticeEventCli) ListenNoticeGroupIncrease(handler func(data model.EventNoticeGroupIncrease) error) {
	c.noticeGroupIncreaseHandlers = append(c.noticeGroupIncreaseHandlers, handler)
}

func (c *websocketNoticeEventCli) ListenNoticeGroupBan(handler func(data model.EventNoticeGroupBan) error) {
	c.noticeGroupBanHandlers = append(c.noticeGroupBanHandlers, handler)
}

func (c *websocketNoticeEventCli) ListenNoticeFriendAdd(handler func(data model.EventNoticeFriendAdd) error) {
	c.noticeFriendAddHandlers = append(c.noticeFriendAddHandlers, handler)
}

func (c *websocketNoticeEventCli) ListenNoticeGroupRecall(handler func(data model.EventNoticeGroupRecall) error) {
	c.noticeGroupRecallHandlers = append(c.noticeGroupRecallHandlers, handler)
}

func (c *websocketNoticeEventCli) ListenNoticeFriendRecall(handler func(data model.EventNoticeFriendRecall) error) {
	c.noticeFriendRecallHandlers = append(c.noticeFriendRecallHandlers, handler)
}

func (c *websocketNoticeEventCli) ListenNoticeGroupNotifyPoke(handler func(data model.EventNoticeGroupNotifyPoke) error) {
	c.noticeGroupNotifyPokeHandlers = append(c.noticeGroupNotifyPokeHandlers, handler)
}

func (c *websocketNoticeEventCli) ListenNoticeGroupNotifyHonor(handler func(data model.EventNoticeGroupNotifyHonor) error) {
	c.noticeGroupNotifyHonorHandlers = append(c.noticeGroupNotifyHonorHandlers, handler)
}

func (c *websocketNoticeEventCli) ListenNoticeGroupNotifyLuckyKing(handler func(data model.EventNoticeGroupNotifyLuckyKing) error) {
	c.noticeGroupNotifyLuckyKingHandlers = append(c.noticeGroupNotifyLuckyKingHandlers, handler)
}
