package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideCustomerCrowdCategoryGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideCustomerCrowdCategoryGetListRequest, PaasWeimobGuideCustomerCrowdCategoryGetListResponse]
}

func (s PaasWeimobGuideCustomerCrowdCategoryGetListService) NewRequest() spi.InvocationRequest[PaasWeimobGuideCustomerCrowdCategoryGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideCustomerCrowdCategoryGetListRequest]{
		Params: &PaasWeimobGuideCustomerCrowdCategoryGetListRequest{},
	}
}

func (s PaasWeimobGuideCustomerCrowdCategoryGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideCustomerCrowdCategoryGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideCustomerCrowdCategoryGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideCustomerCrowdCategoryGetListRequest]))
}

type PaasWeimobGuideCustomerCrowdCategoryGetListRequest struct {
	BosId             int64   `json:"bosId,omitempty"`
	Vid               int64   `json:"vid,omitempty"`
	VidType           int64   `json:"vidType,omitempty"`
	ProductInstanceId int64   `json:"productInstanceId,omitempty"`
	GuiderWid         int64   `json:"guiderWid,omitempty"`
	CrowdIds          []int64 `json:"crowdIds,omitempty"`
	SearchKey         string  `json:"searchKey,omitempty"`
	CategoryId        int64   `json:"categoryId,omitempty"`
}

type PaasWeimobGuideCustomerCrowdCategoryGetListResponse struct {
	CategoryInfoList []PaasWeimobGuideCustomerCrowdCategoryGetListResponseCategoryInfoList `json:"categoryInfoList,omitempty"`
	Count            int64                                                                 `json:"count,omitempty"`
}
type PaasWeimobGuideCustomerCrowdCategoryGetListResponseCategoryInfoList struct {
	CrowdId      int64  `json:"crowdId,omitempty"`
	CrowdName    string `json:"crowdName,omitempty"`
	CrowdType    int64  `json:"crowdType,omitempty"`
	Description  string `json:"description,omitempty"`
	CategoryId   int64  `json:"categoryId,omitempty"`
	CategoryName string `json:"categoryName,omitempty"`
	CoverNumber  int64  `json:"coverNumber,omitempty"`
	Status       int64  `json:"status,omitempty"`
	CreateTime   string `json:"createTime,omitempty"`
	UpdateTime   string `json:"updateTime,omitempty"`
}

func CreatePaasWeimobGuideCustomerCrowdCategoryGetListResponse() *PaasWeimobGuideCustomerCrowdCategoryGetListResponse {
	return &PaasWeimobGuideCustomerCrowdCategoryGetListResponse{}
}

// OnPaasWeimobGuideCustomerCrowdCategoryGetListServiceInvocation
// @id 744
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/744?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询人群分类)
func (s *Service) OnPaasWeimobGuideCustomerCrowdCategoryGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideCustomerCrowdCategoryGetListRequest, PaasWeimobGuideCustomerCrowdCategoryGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideCustomerCrowdCategoryGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideCustomerCrowdCategoryGetListRequest, PaasWeimobGuideCustomerCrowdCategoryGetListResponse](invocation),
		"PaasWeimobGuideCustomerCrowdCategoryGetListService",
		"weimobGuide.customer.crowd.category.getList",
	)
	return s
}
