package xinyun

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"

type Service struct {
	spi.XinyunService
}

func NewService() (service *Service, err error) {
	service = &Service{}
	err = service.Setup()
	return
}
