package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"

type Service struct {
	spi.WosService
}

func NewService() (service *Service, err error) {
	service = &Service{}
	err = service.Setup()
	return
}
