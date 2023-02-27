package base

type OnebotEventType string

const (
	OnebotEventTypeMessage   OnebotEventType = "message"
	OnebotEventTypeNotice    OnebotEventType = "notice"
	OnebotEventTypeRequest   OnebotEventType = "request"
	OnebotEventTypeMetaEvent OnebotEventType = "meta_event"
)

type OnebotBaseEventCli interface {
	EventType() OnebotEventType
	Handler(data []byte) error
}
