package entity

type GenericEvent struct {
	EventType string
}

func (g GenericEvent) Type() string {
	return g.EventType
}

type Event interface {
	Type() string
}
