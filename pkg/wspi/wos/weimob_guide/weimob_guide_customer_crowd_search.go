package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideCrowdSearchService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideCrowdSearchRequest, PaasWeimobGuideCrowdSearchResponse]
}

func (s PaasWeimobGuideCrowdSearchService) NewRequest() spi.InvocationRequest[PaasWeimobGuideCrowdSearchRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideCrowdSearchRequest]{
		Params: &PaasWeimobGuideCrowdSearchRequest{},
	}
}

func (s PaasWeimobGuideCrowdSearchService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideCrowdSearchRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideCrowdSearchResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideCrowdSearchRequest]))
}

type PaasWeimobGuideCrowdSearchRequest struct {
	BosId             int64   `json:"bosId,omitempty"`
	Vid               int64   `json:"vid,omitempty"`
	ProductInstanceId int64   `json:"productInstanceId,omitempty"`
	VidType           int64   `json:"vidType,omitempty"`
	GuiderWid         int64   `json:"guiderWid,omitempty"`
	CrowdIds          []int64 `json:"crowdIds,omitempty"`
	SearchKey         string  `json:"searchKey,omitempty"`
	CategoryId        int64   `json:"categoryId,omitempty"`
	CrowdName         string  `json:"crowdName,omitempty"`
}

type PaasWeimobGuideCrowdSearchResponse struct {
	CustomCrowd []PaasWeimobGuideCrowdSearchResponseCustomCrowd `json:"customCrowd,omitempty"`
}
type PaasWeimobGuideCrowdSearchResponseCustomCrowd struct {
	CrowdId      int64  `json:"crowdId,omitempty"`
	Description  string `json:"description,omitempty"`
	CrowdName    string `json:"crowdName,omitempty"`
	CrowdType    int64  `json:"crowdType,omitempty"`
	CategoryId   int64  `json:"categoryId,omitempty"`
	CategoryName string `json:"categoryName,omitempty"`
	CoverNumber  int64  `json:"coverNumber,omitempty"`
	Status       int64  `json:"status,omitempty"`
	CreateTime   string `json:"createTime,omitempty"`
	UpdateTime   string `json:"updateTime,omitempty"`
}

func CreatePaasWeimobGuideCrowdSearchResponse() *PaasWeimobGuideCrowdSearchResponse {
	return &PaasWeimobGuideCrowdSearchResponse{}
}

// OnPaasWeimobGuideCrowdSearchServiceInvocation
// @id 743
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/743?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=搜索人群信息)
func (s *Service) OnPaasWeimobGuideCrowdSearchServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideCrowdSearchRequest, PaasWeimobGuideCrowdSearchResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideCrowdSearchService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideCrowdSearchRequest, PaasWeimobGuideCrowdSearchResponse](invocation),
		"PaasWeimobGuideCrowdSearchService",
		"weimobGuide.customer.crowd.search",
	)
	return s
}
