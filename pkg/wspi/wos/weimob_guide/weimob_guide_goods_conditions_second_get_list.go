package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideGoodsConditionsSecondGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideGoodsConditionsSecondGetListRequest, PaasWeimobGuideGoodsConditionsSecondGetListResponse]
}

func (s PaasWeimobGuideGoodsConditionsSecondGetListService) NewRequest() spi.InvocationRequest[PaasWeimobGuideGoodsConditionsSecondGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideGoodsConditionsSecondGetListRequest]{
		Params: &PaasWeimobGuideGoodsConditionsSecondGetListRequest{},
	}
}

func (s PaasWeimobGuideGoodsConditionsSecondGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideGoodsConditionsSecondGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideGoodsConditionsSecondGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideGoodsConditionsSecondGetListRequest]))
}

type PaasWeimobGuideGoodsConditionsSecondGetListRequest struct {
	Type int64 `json:"type,omitempty"`
	Id   int64 `json:"id,omitempty"`
}

type PaasWeimobGuideGoodsConditionsSecondGetListResponse struct {
	Id       int64  `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Type     int64  `json:"type,omitempty"`
	ParentId int64  `json:"parentId,omitempty"`
}

func CreatePaasWeimobGuideGoodsConditionsSecondGetListResponse() *PaasWeimobGuideGoodsConditionsSecondGetListResponse {
	return &PaasWeimobGuideGoodsConditionsSecondGetListResponse{}
}

// OnPaasWeimobGuideGoodsConditionsSecondGetListServiceInvocation
// @id 722
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/722?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询商品二级搜索条件)
func (s *Service) OnPaasWeimobGuideGoodsConditionsSecondGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideGoodsConditionsSecondGetListRequest, PaasWeimobGuideGoodsConditionsSecondGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideGoodsConditionsSecondGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideGoodsConditionsSecondGetListRequest, PaasWeimobGuideGoodsConditionsSecondGetListResponse](invocation),
		"PaasWeimobGuideGoodsConditionsSecondGetListService",
		"weimobGuide.goods.conditions.second.getList",
	)
	return s
}
