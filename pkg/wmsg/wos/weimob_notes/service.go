package weimob_notes

import (
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
)

type Service struct {
	msg.WosService
}

func NewService() (service *Service, err error) {
	service = &Service{}
	err = service.Setup()
	return
}
