package weimob_crm

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCrmGoodsQuerySimpleBaseGoodsInfoListService struct {
	handler spi.WosInvocationHandler[PaasWeimobCrmGoodsQuerySimpleBaseGoodsInfoListRequest, PaasWeimobCrmGoodsQuerySimpleBaseGoodsInfoListResponse]
}

func (s PaasWeimobCrmGoodsQuerySimpleBaseGoodsInfoListService) NewRequest() spi.InvocationRequest[PaasWeimobCrmGoodsQuerySimpleBaseGoodsInfoListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCrmGoodsQuerySimpleBaseGoodsInfoListRequest]{
		Params: &PaasWeimobCrmGoodsQuerySimpleBaseGoodsInfoListRequest{},
	}
}

func (s PaasWeimobCrmGoodsQuerySimpleBaseGoodsInfoListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCrmGoodsQuerySimpleBaseGoodsInfoListRequest],
) (
	spi.InvocationResponse[PaasWeimobCrmGoodsQuerySimpleBaseGoodsInfoListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCrmGoodsQuerySimpleBaseGoodsInfoListRequest]))
}

type PaasWeimobCrmGoodsQuerySimpleBaseGoodsInfoListRequest struct {
}

type PaasWeimobCrmGoodsQuerySimpleBaseGoodsInfoListResponse struct {
}

func CreatePaasWeimobCrmGoodsQuerySimpleBaseGoodsInfoListResponse() *PaasWeimobCrmGoodsQuerySimpleBaseGoodsInfoListResponse {
	return &PaasWeimobCrmGoodsQuerySimpleBaseGoodsInfoListResponse{}
}

// OnPaasWeimobCrmGoodsQuerySimpleBaseGoodsInfoListServiceInvocation
// @id 507
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/507?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询简单商品信息)
func (s *Service) OnPaasWeimobCrmGoodsQuerySimpleBaseGoodsInfoListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCrmGoodsQuerySimpleBaseGoodsInfoListRequest, PaasWeimobCrmGoodsQuerySimpleBaseGoodsInfoListResponse],
) (service *Service) {
	var invocation = &PaasWeimobCrmGoodsQuerySimpleBaseGoodsInfoListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCrmGoodsQuerySimpleBaseGoodsInfoListRequest, PaasWeimobCrmGoodsQuerySimpleBaseGoodsInfoListResponse](invocation),
		"PaasWeimobCrmGoodsQuerySimpleBaseGoodsInfoListService",
		"weimobCrm.goods.querySimpleBaseGoodsInfoList",
	)
	return s
}
