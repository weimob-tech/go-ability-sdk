package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideCrowdDetailGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideCrowdDetailGetListRequest, PaasWeimobGuideCrowdDetailGetListResponse]
}

func (s PaasWeimobGuideCrowdDetailGetListService) NewRequest() spi.InvocationRequest[PaasWeimobGuideCrowdDetailGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideCrowdDetailGetListRequest]{
		Params: &PaasWeimobGuideCrowdDetailGetListRequest{},
	}
}

func (s PaasWeimobGuideCrowdDetailGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideCrowdDetailGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideCrowdDetailGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideCrowdDetailGetListRequest]))
}

type PaasWeimobGuideCrowdDetailGetListRequest struct {
	CrowIdList        []int64 `json:"crowIdList,omitempty"`
	BosId             int64   `json:"bosId,omitempty"`
	VidType           int64   `json:"vidType,omitempty"`
	Vid               int64   `json:"vid,omitempty"`
	GuiderWid         int64   `json:"guiderWid,omitempty"`
	ProductId         int64   `json:"productId,omitempty"`
	ProductInstanceId int64   `json:"productInstanceId,omitempty"`
	GuiderId          int64   `json:"guiderId,omitempty"`
}

type PaasWeimobGuideCrowdDetailGetListResponse struct {
	CustomCrowd []PaasWeimobGuideCrowdDetailGetListResponseCustomCrowd `json:"customCrowd,omitempty"`
}
type PaasWeimobGuideCrowdDetailGetListResponseCustomCrowd struct {
	CrowdId     int64  `json:"crowdId,omitempty"`
	CrowdName   string `json:"crowdName,omitempty"`
	Description string `json:"description,omitempty"`
	CategoryId  int64  `json:"categoryId,omitempty"`
}

func CreatePaasWeimobGuideCrowdDetailGetListResponse() *PaasWeimobGuideCrowdDetailGetListResponse {
	return &PaasWeimobGuideCrowdDetailGetListResponse{}
}

// OnPaasWeimobGuideCrowdDetailGetListServiceInvocation
// @id 726
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/726?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询人群信息)
func (s *Service) OnPaasWeimobGuideCrowdDetailGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideCrowdDetailGetListRequest, PaasWeimobGuideCrowdDetailGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideCrowdDetailGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideCrowdDetailGetListRequest, PaasWeimobGuideCrowdDetailGetListResponse](invocation),
		"PaasWeimobGuideCrowdDetailGetListService",
		"weimobGuide.crowd.detail.getList",
	)
	return s
}
