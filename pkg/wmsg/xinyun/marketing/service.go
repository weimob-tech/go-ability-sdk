package marketing

import (
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
)

type Service struct {
	msg.XinyunService
}

func NewService() (service *Service, err error) {
	service = &Service{}
	err = service.Setup()
	return
}
