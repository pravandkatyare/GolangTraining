package service

import eventbridge "28test/eventBridge"

type Service struct {
	bridger eventbridge.Bridger
}

func NewService(bridger eventbridge.Bridger) *Service {
	return &Service{bridger: bridger}
}

func (s *Service) createEvent() error {
	return s.bridger.Create()
}
