package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideGoodsTagGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideGoodsTagGetListRequest, PaasWeimobGuideGoodsTagGetListResponse]
}

func (s PaasWeimobGuideGoodsTagGetListService) NewRequest() spi.InvocationRequest[PaasWeimobGuideGoodsTagGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideGoodsTagGetListRequest]{
		Params: &PaasWeimobGuideGoodsTagGetListRequest{},
	}
}

func (s PaasWeimobGuideGoodsTagGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideGoodsTagGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideGoodsTagGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideGoodsTagGetListRequest]))
}

type PaasWeimobGuideGoodsTagGetListRequest struct {
	PageNum  int64 `json:"pageNum,omitempty"`
	PageSize int64 `json:"pageSize,omitempty"`
}

type PaasWeimobGuideGoodsTagGetListResponse struct {
	PageList   []PaasWeimobGuideGoodsTagGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                            `json:"pageNum,omitempty"`
	PageSize   int64                                            `json:"pageSize,omitempty"`
	TotalCount int64                                            `json:"totalCount,omitempty"`
}
type PaasWeimobGuideGoodsTagGetListResponsePageList struct {
	TagId         int64  `json:"tagId,omitempty"`
	Name          string `json:"name,omitempty"`
	GoodsNum      int64  `json:"goodsNum,omitempty"`
	GoodsApplyNum int64  `json:"goodsApplyNum,omitempty"`
}

func CreatePaasWeimobGuideGoodsTagGetListResponse() *PaasWeimobGuideGoodsTagGetListResponse {
	return &PaasWeimobGuideGoodsTagGetListResponse{}
}

// OnPaasWeimobGuideGoodsTagGetListServiceInvocation
// @id 739
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/739?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询商品标签列表)
func (s *Service) OnPaasWeimobGuideGoodsTagGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideGoodsTagGetListRequest, PaasWeimobGuideGoodsTagGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideGoodsTagGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideGoodsTagGetListRequest, PaasWeimobGuideGoodsTagGetListResponse](invocation),
		"PaasWeimobGuideGoodsTagGetListService",
		"weimobGuide.goods.tag.getList",
	)
	return s
}
