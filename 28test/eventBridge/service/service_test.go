package service

import (
	"28test/eventBridge/mock"
	"fmt"
	"testing"
)

func TestUnit_Create(t *testing.T) {
	m := &mock.EventBridge{}

	s := NewService(m)

	m.On("Create").Return(nil)

	fmt.Println("s.createEvent():", s.createEvent())
}
