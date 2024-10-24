package mock

import "github.com/stretchr/testify/mock"

type EventBridge struct {
	mock.Mock
}

func (e *EventBridge) Create() error {
	ret := e.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
	// return errors.New("Invalid")
}
