package eventbridge

type EventBridge struct {
	name string
}

type Bridger interface {
	Create() error
}

func (e *EventBridge) Create() error {
	return nil
}
