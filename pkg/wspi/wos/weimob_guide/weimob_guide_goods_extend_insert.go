package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideGoodsExtendInsertService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideGoodsExtendInsertRequest, PaasWeimobGuideGoodsExtendInsertResponse]
}

func (s PaasWeimobGuideGoodsExtendInsertService) NewRequest() spi.InvocationRequest[PaasWeimobGuideGoodsExtendInsertRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideGoodsExtendInsertRequest]{
		Params: &PaasWeimobGuideGoodsExtendInsertRequest{},
	}
}

func (s PaasWeimobGuideGoodsExtendInsertService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideGoodsExtendInsertRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideGoodsExtendInsertResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideGoodsExtendInsertRequest]))
}

type PaasWeimobGuideGoodsExtendInsertRequest struct {
	GoodsExtList []PaasWeimobGuideGoodsExtendInsertRequestGoodsExtList `json:"goodsExtList,omitempty"`
}
type PaasWeimobGuideGoodsExtendInsertRequestGoodsExtList struct {
	GoodsId int64   `json:"goodsId,omitempty"`
	Vid     int64   `json:"vid,omitempty"`
	Key     string  `json:"key,omitempty"`
	Value   []int64 `json:"value,omitempty"`
}

type PaasWeimobGuideGoodsExtendInsertResponse struct {
	Id int64 `json:"id,omitempty"`
}

func CreatePaasWeimobGuideGoodsExtendInsertResponse() *PaasWeimobGuideGoodsExtendInsertResponse {
	return &PaasWeimobGuideGoodsExtendInsertResponse{}
}

// OnPaasWeimobGuideGoodsExtendInsertServiceInvocation
// @id 737
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/737?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=标记商品为导购商品)
func (s *Service) OnPaasWeimobGuideGoodsExtendInsertServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideGoodsExtendInsertRequest, PaasWeimobGuideGoodsExtendInsertResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideGoodsExtendInsertService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideGoodsExtendInsertRequest, PaasWeimobGuideGoodsExtendInsertResponse](invocation),
		"PaasWeimobGuideGoodsExtendInsertService",
		"weimobGuide.goods.extend.insert",
	)
	return s
}
