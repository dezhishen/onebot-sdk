package event

import (
	"encoding/json"

	"github.com/dezhishen/onebot-sdk/pkg/model"
)

var groupUploadHandlers []func(data model.EventGroupUpload) error
var groupAdminHandlers []func(data model.EventGroupAdmin) error
var groupDecreaseHandlers []func(data model.EventGroupDecrease) error
var groupIncreaseHandlers []func(data model.EventGroupIncrease) error
var groupBanHandlers []func(data model.EventGroupBan) error
var groupRecallHandlers []func(data model.EventGroupRecall) error
var groupNotifyPokeHandlers []func(data model.EventGroupNotifyPoke) error
var groupNotifyLuckyKingHandlers []func(data model.EventGroupNotifyLuckyKing) error
var groupNotifyHonorHandlers []func(data model.EventGroupNotifyHonor) error
var friendAddHandlers []func(data model.EventFriendAdd) error
var friendRecallHandlers []func(data model.EventFriendRecall) error

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
		var relNotice model.EventGroupUpload
		err = json.Unmarshal(data, &relNotice)
		if err != nil {
			return err
		}
		for _, e := range groupUploadHandlers {
			err = e(relNotice)
			if err != nil {
				return err
			}
		}
	} else if notice.NoticeType == "group_admin" {
		var relNotice model.EventGroupAdmin
		err = json.Unmarshal(data, &relNotice)
		if err != nil {
			return err
		}
		for _, e := range groupAdminHandlers {
			err = e(relNotice)
			if err != nil {
				return err
			}
		}
	} else if notice.NoticeType == "group_decrease" {
		var relNotice model.EventGroupDecrease
		err = json.Unmarshal(data, &relNotice)
		if err != nil {
			return err
		}
		for _, e := range groupDecreaseHandlers {
			err = e(relNotice)
			if err != nil {
				return err
			}
		}
	} else if notice.NoticeType == "group_increase" {
		var relNotice model.EventGroupIncrease
		err = json.Unmarshal(data, &relNotice)
		if err != nil {
			return err
		}
		for _, e := range groupIncreaseHandlers {
			err = e(relNotice)
			if err != nil {
				return err
			}
		}
	} else if notice.NoticeType == "group_ban" {
		var relNotice model.EventGroupBan
		err = json.Unmarshal(data, &relNotice)
		if err != nil {
			return err
		}
		for _, e := range groupBanHandlers {
			err = e(relNotice)
			if err != nil {
				return err
			}
		}
	} else if notice.NoticeType == "friend_add" {
		var relNotice model.EventFriendAdd
		err = json.Unmarshal(data, &relNotice)
		if err != nil {
			return err
		}
		for _, e := range friendAddHandlers {
			err = e(relNotice)
			if err != nil {
				return err
			}
		}
	} else if notice.NoticeType == "group_recall" {
		var relNotice model.EventGroupRecall
		err = json.Unmarshal(data, &relNotice)
		if err != nil {
			return err
		}
		for _, e := range groupRecallHandlers {
			err = e(relNotice)
			if err != nil {
				return err
			}
		}
	} else if notice.NoticeType == "friend_recall" {
		var relNotice model.EventFriendRecall
		err = json.Unmarshal(data, &relNotice)
		if err != nil {
			return err
		}
		for _, e := range friendRecallHandlers {
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
			var relNotice model.EventGroupNotifyHonor
			err = json.Unmarshal(data, &relNotice)
			if err != nil {
				return err
			}
			for _, e := range groupNotifyHonorHandlers {
				err = e(relNotice)
				if err != nil {
					return err
				}
			}
		} else if notify.SubType == "lucky_king" {
			var relNotice model.EventGroupNotifyLuckyKing
			err = json.Unmarshal(data, &relNotice)
			if err != nil {
				return err
			}
			for _, e := range groupNotifyLuckyKingHandlers {
				err = e(relNotice)
				if err != nil {
					return err
				}
			}
		} else if notify.SubType == "poke" {
			var relNotice model.EventGroupNotifyPoke
			err = json.Unmarshal(data, &relNotice)
			if err != nil {
				return err
			}
			for _, e := range groupNotifyPokeHandlers {
				err = e(relNotice)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func ListenGroupUpload(handler func(data model.EventGroupUpload) error) {
	groupUploadHandlers = append(groupUploadHandlers, handler)
}

func ListenGroupAdmin(handler func(data model.EventGroupAdmin) error) {
	groupAdminHandlers = append(groupAdminHandlers, handler)
}

func ListenGroupDecrease(handler func(data model.EventGroupDecrease) error) {
	groupDecreaseHandlers = append(groupDecreaseHandlers, handler)
}

func ListenGroupIncrease(handler func(data model.EventGroupIncrease) error) {
	groupIncreaseHandlers = append(groupIncreaseHandlers, handler)
}

func ListenGroupBan(handler func(data model.EventGroupBan) error) {
	groupBanHandlers = append(groupBanHandlers, handler)
}

func ListenFriendAdd(handler func(data model.EventFriendAdd) error) {
	friendAddHandlers = append(friendAddHandlers, handler)
}

func ListenGroupRecall(handler func(data model.EventGroupRecall) error) {
	groupRecallHandlers = append(groupRecallHandlers, handler)
}

func ListenFriendRecall(handler func(data model.EventFriendRecall) error) {
	friendRecallHandlers = append(friendRecallHandlers, handler)
}

func ListenGroupNotifyPoke(handler func(data model.EventGroupNotifyPoke) error) {
	groupNotifyPokeHandlers = append(groupNotifyPokeHandlers, handler)
}

func ListenGroupNotifyHonor(handler func(data model.EventGroupNotifyHonor) error) {
	groupNotifyHonorHandlers = append(groupNotifyHonorHandlers, handler)
}

func ListenGroupNotifyLuckyKing(handler func(data model.EventGroupNotifyLuckyKing) error) {
	groupNotifyLuckyKingHandlers = append(groupNotifyLuckyKingHandlers, handler)
}
