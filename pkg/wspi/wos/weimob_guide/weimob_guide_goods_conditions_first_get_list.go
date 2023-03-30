package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideGoodsConditionsFirstGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideGoodsConditionsFirstGetListRequest, PaasWeimobGuideGoodsConditionsFirstGetListResponse]
}

func (s PaasWeimobGuideGoodsConditionsFirstGetListService) NewRequest() spi.InvocationRequest[PaasWeimobGuideGoodsConditionsFirstGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideGoodsConditionsFirstGetListRequest]{
		Params: &PaasWeimobGuideGoodsConditionsFirstGetListRequest{},
	}
}

func (s PaasWeimobGuideGoodsConditionsFirstGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideGoodsConditionsFirstGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideGoodsConditionsFirstGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideGoodsConditionsFirstGetListRequest]))
}

type PaasWeimobGuideGoodsConditionsFirstGetListRequest struct {
}

type PaasWeimobGuideGoodsConditionsFirstGetListResponse struct {
	ItemList []PaasWeimobGuideGoodsConditionsFirstGetListResponseItemList `json:"itemList,omitempty"`
	Type     int64                                                        `json:"type,omitempty"`
	Title    string                                                       `json:"title,omitempty"`
}
type PaasWeimobGuideGoodsConditionsFirstGetListResponseItemList struct {
	Id     int64  `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Type   int64  `json:"type,omitempty"`
	HasSub bool   `json:"hasSub,omitempty"`
}

func CreatePaasWeimobGuideGoodsConditionsFirstGetListResponse() *PaasWeimobGuideGoodsConditionsFirstGetListResponse {
	return &PaasWeimobGuideGoodsConditionsFirstGetListResponse{}
}

// OnPaasWeimobGuideGoodsConditionsFirstGetListServiceInvocation
// @id 723
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/723?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询商品一级搜索条件)
func (s *Service) OnPaasWeimobGuideGoodsConditionsFirstGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideGoodsConditionsFirstGetListRequest, PaasWeimobGuideGoodsConditionsFirstGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideGoodsConditionsFirstGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideGoodsConditionsFirstGetListRequest, PaasWeimobGuideGoodsConditionsFirstGetListResponse](invocation),
		"PaasWeimobGuideGoodsConditionsFirstGetListService",
		"weimobGuide.goods.conditions.first.getList",
	)
	return s
}
