package event

type OnebotEventType string

const (
	OnebotEventTypeMessage   OnebotEventType = "message"
	OnebotEventTypeNotice    OnebotEventType = "notice"
	OnebotEventTypeRequest   OnebotEventType = "request"
	OnebotEventTypeMetaEvent OnebotEventType = "meta_event"
)
