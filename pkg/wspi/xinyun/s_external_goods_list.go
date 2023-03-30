package xinyun

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasExternalGoodsListService struct {
	handler spi.XinyunInvocationHandler[PaasExternalGoodsListRequest, PaasExternalGoodsListResponse]
}

func (s PaasExternalGoodsListService) NewRequest() spi.InvocationRequest[PaasExternalGoodsListRequest] {
	return &spi.XinyunInvocationRequest[PaasExternalGoodsListRequest]{
		Params: &PaasExternalGoodsListRequest{},
	}
}

func (s PaasExternalGoodsListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasExternalGoodsListRequest],
) (
	spi.InvocationResponse[PaasExternalGoodsListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.XinyunInvocationRequest[PaasExternalGoodsListRequest]))
}

type PaasExternalGoodsListRequest struct {
}

type PaasExternalGoodsListResponse struct {
	GoodsIdList []int64 `json:"goodsIdList,omitempty"`
}

func CreatePaasExternalGoodsListResponse() *PaasExternalGoodsListResponse {
	return &PaasExternalGoodsListResponse{}
}

// OnPaasExternalGoodsListServiceInvocation
// @id 1920
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=获取外部商品搜索结果)
func (s *Service) OnPaasExternalGoodsListServiceInvocation(
	handler spi.XinyunInvocationHandler[PaasExternalGoodsListRequest, PaasExternalGoodsListResponse],
) (service *Service) {
	var invocation = &PaasExternalGoodsListService{handler}
	s.Register(
		spi.WrapInvoker[PaasExternalGoodsListRequest, PaasExternalGoodsListResponse](invocation),
		"PaasExternalGoodsListService",
		"sExternalGoodsList",
	)
	return s
}
