package event

import (
	"encoding/json"

	"github.com/dezhishen/onebot-sdk/pkg/model"
)

var noticeGroupUploadHandlers []func(data model.EventNoticeGroupUpload) error
var noticeGroupAdminHandlers []func(data model.EventNoticeGroupAdmin) error
var noticeGroupDecreaseHandlers []func(data model.EventNoticeGroupDecrease) error
var noticeGroupIncreaseHandlers []func(data model.EventNoticeGroupIncrease) error
var noticeGroupBanHandlers []func(data model.EventNoticeGroupBan) error
var noticeGroupRecallHandlers []func(data model.EventNoticeGroupRecall) error
var noticeGroupNotifyPokeHandlers []func(data model.EventNoticeGroupNotifyPoke) error
var noticeGroupNotifyLuckyKingHandlers []func(data model.EventNoticeGroupNotifyLuckyKing) error
var noticeGroupNotifyHonorHandlers []func(data model.EventNoticeGroupNotifyHonor) error
var noticeFriendAddHandlers []func(data model.EventNoticeFriendAdd) error
var noticeFriendRecallHandlers []func(data model.EventNoticeFriendRecall) error

func init() {
	setHandler(
		EventTypeNotice,
		noticeHandler,
	)
}

func noticeHandler(data []byte) error {
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
		for _, e := range noticeGroupUploadHandlers {
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
		for _, e := range noticeGroupAdminHandlers {
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
		for _, e := range noticeGroupDecreaseHandlers {
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
		for _, e := range noticeGroupIncreaseHandlers {
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
		for _, e := range noticeGroupBanHandlers {
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
		for _, e := range noticeFriendAddHandlers {
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
		for _, e := range noticeGroupRecallHandlers {
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
		for _, e := range noticeFriendRecallHandlers {
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
			for _, e := range noticeGroupNotifyHonorHandlers {
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
			for _, e := range noticeGroupNotifyLuckyKingHandlers {
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
			for _, e := range noticeGroupNotifyPokeHandlers {
				err = e(relNotice)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func ListenNoticeGroupUpload(handler func(data model.EventNoticeGroupUpload) error) {
	noticeGroupUploadHandlers = append(noticeGroupUploadHandlers, handler)
}

func ListenNoticeGroupAdmin(handler func(data model.EventNoticeGroupAdmin) error) {
	noticeGroupAdminHandlers = append(noticeGroupAdminHandlers, handler)
}

func ListenNoticeGroupDecrease(handler func(data model.EventNoticeGroupDecrease) error) {
	noticeGroupDecreaseHandlers = append(noticeGroupDecreaseHandlers, handler)
}

func ListenNoticeGroupIncrease(handler func(data model.EventNoticeGroupIncrease) error) {
	noticeGroupIncreaseHandlers = append(noticeGroupIncreaseHandlers, handler)
}

func ListenNoticeGroupBan(handler func(data model.EventNoticeGroupBan) error) {
	noticeGroupBanHandlers = append(noticeGroupBanHandlers, handler)
}

func ListenNoticeFriendAdd(handler func(data model.EventNoticeFriendAdd) error) {
	noticeFriendAddHandlers = append(noticeFriendAddHandlers, handler)
}

func ListenNoticeGroupRecall(handler func(data model.EventNoticeGroupRecall) error) {
	noticeGroupRecallHandlers = append(noticeGroupRecallHandlers, handler)
}

func ListenNoticeFriendRecall(handler func(data model.EventNoticeFriendRecall) error) {
	noticeFriendRecallHandlers = append(noticeFriendRecallHandlers, handler)
}

func ListenNoticeGroupNotifyPoke(handler func(data model.EventNoticeGroupNotifyPoke) error) {
	noticeGroupNotifyPokeHandlers = append(noticeGroupNotifyPokeHandlers, handler)
}

func ListenNoticeGroupNotifyHonor(handler func(data model.EventNoticeGroupNotifyHonor) error) {
	noticeGroupNotifyHonorHandlers = append(noticeGroupNotifyHonorHandlers, handler)
}

func ListenNoticeGroupNotifyLuckyKing(handler func(data model.EventNoticeGroupNotifyLuckyKing) error) {
	noticeGroupNotifyLuckyKingHandlers = append(noticeGroupNotifyLuckyKingHandlers, handler)
}
