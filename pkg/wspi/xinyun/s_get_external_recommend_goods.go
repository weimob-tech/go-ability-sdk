package xinyun

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasGetExternalRecommendGoodsService struct {
	handler spi.XinyunInvocationHandler[PaasGetExternalRecommendGoodsRequest, PaasGetExternalRecommendGoodsResponse]
}

func (s PaasGetExternalRecommendGoodsService) NewRequest() spi.InvocationRequest[PaasGetExternalRecommendGoodsRequest] {
	return &spi.XinyunInvocationRequest[PaasGetExternalRecommendGoodsRequest]{
		Params: &PaasGetExternalRecommendGoodsRequest{},
	}
}

func (s PaasGetExternalRecommendGoodsService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasGetExternalRecommendGoodsRequest],
) (
	spi.InvocationResponse[PaasGetExternalRecommendGoodsResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.XinyunInvocationRequest[PaasGetExternalRecommendGoodsRequest]))
}

type PaasGetExternalRecommendGoodsRequest struct {
	GoodsIdList []int64 `json:"goodsIdList,omitempty"`
	SceneId     int64   `json:"sceneId,omitempty"`
	Pid         int64   `json:"pid,omitempty"`
	Wid         int64   `json:"wid,omitempty"`
}

type PaasGetExternalRecommendGoodsResponse struct {
	RecommendGoodsIdList []int64 `json:"recommendGoodsIdList,omitempty"`
	SceneSwitchOn        bool    `json:"sceneSwitchOn,omitempty"`
	DisplayCount         int64   `json:"displayCount,omitempty"`
}

func CreatePaasGetExternalRecommendGoodsResponse() *PaasGetExternalRecommendGoodsResponse {
	return &PaasGetExternalRecommendGoodsResponse{}
}

// OnPaasGetExternalRecommendGoodsServiceInvocation
// @id 1853
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=获取外部系统推荐商品)
func (s *Service) OnPaasGetExternalRecommendGoodsServiceInvocation(
	handler spi.XinyunInvocationHandler[PaasGetExternalRecommendGoodsRequest, PaasGetExternalRecommendGoodsResponse],
) (service *Service) {
	var invocation = &PaasGetExternalRecommendGoodsService{handler}
	s.Register(
		spi.WrapInvoker[PaasGetExternalRecommendGoodsRequest, PaasGetExternalRecommendGoodsResponse](invocation),
		"PaasGetExternalRecommendGoodsService",
		"sGetExternalRecommendGoods",
	)
	return s
}
