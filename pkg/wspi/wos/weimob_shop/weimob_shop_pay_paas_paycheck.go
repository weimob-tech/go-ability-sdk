package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobShopPayPaasPaycheckService struct {
	handler spi.WosInvocationHandler[PaasWeimobShopPayPaasPaycheckRequest, PaasWeimobShopPayPaasPaycheckResponse]
}

func (s PaasWeimobShopPayPaasPaycheckService) NewRequest() spi.InvocationRequest[PaasWeimobShopPayPaasPaycheckRequest] {
	return &spi.WosInvocationRequest[PaasWeimobShopPayPaasPaycheckRequest]{
		Params: &PaasWeimobShopPayPaasPaycheckRequest{},
	}
}

func (s PaasWeimobShopPayPaasPaycheckService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobShopPayPaasPaycheckRequest],
) (
	spi.InvocationResponse[PaasWeimobShopPayPaasPaycheckResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobShopPayPaasPaycheckRequest]))
}

type PaasWeimobShopPayPaasPaycheckRequest struct {
	BosId     int64  `json:"bosId,omitempty"`
	ActionKey string `json:"actionKey,omitempty"`
}

type PaasWeimobShopPayPaasPaycheckResponse struct {
	IsPaas bool `json:"isPaas,omitempty"`
}

func CreatePaasWeimobShopPayPaasPaycheckResponse() *PaasWeimobShopPayPaasPaycheckResponse {
	return &PaasWeimobShopPayPaasPaycheckResponse{}
}

// OnPaasWeimobShopPayPaasPaycheckServiceInvocation
// @id 1180
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1180?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=支付校验)
func (s *Service) OnPaasWeimobShopPayPaasPaycheckServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobShopPayPaasPaycheckRequest, PaasWeimobShopPayPaasPaycheckResponse],
) (service *Service) {
	var invocation = &PaasWeimobShopPayPaasPaycheckService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobShopPayPaasPaycheckRequest, PaasWeimobShopPayPaasPaycheckResponse](invocation),
		"PaasWeimobShopPayPaasPaycheckService",
		"weimobShop.pay.paasPaycheck",
	)
	return s
}
