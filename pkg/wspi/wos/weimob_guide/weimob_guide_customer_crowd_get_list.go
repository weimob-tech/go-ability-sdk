package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideCustomerCrowdGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideCustomerCrowdGetListRequest, PaasWeimobGuideCustomerCrowdGetListResponse]
}

func (s PaasWeimobGuideCustomerCrowdGetListService) NewRequest() spi.InvocationRequest[PaasWeimobGuideCustomerCrowdGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideCustomerCrowdGetListRequest]{
		Params: &PaasWeimobGuideCustomerCrowdGetListRequest{},
	}
}

func (s PaasWeimobGuideCustomerCrowdGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideCustomerCrowdGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideCustomerCrowdGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideCustomerCrowdGetListRequest]))
}

type PaasWeimobGuideCustomerCrowdGetListRequest struct {
	WidList           []PaasWeimobGuideCustomerCrowdGetListRequestWidList `json:"widList,omitempty"`
	BosId             int64                                               `json:"bosId,omitempty"`
	Vid               int64                                               `json:"vid,omitempty"`
	ProductInstanceId int64                                               `json:"productInstanceId,omitempty"`
	VidType           int64                                               `json:"vidType,omitempty"`
}
type PaasWeimobGuideCustomerCrowdGetListRequestWidList struct {
}

type PaasWeimobGuideCustomerCrowdGetListResponse struct {
	List []PaasWeimobGuideCustomerCrowdGetListResponselist `json:"list,omitempty"`
}
type PaasWeimobGuideCustomerCrowdGetListResponselist struct {
	CrowdIdList []PaasWeimobGuideCustomerCrowdGetListResponseCrowdIdList `json:"crowdIdList,omitempty"`
	Wid         int64                                                    `json:"wid,omitempty"`
}
type PaasWeimobGuideCustomerCrowdGetListResponseCrowdIdList struct {
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

func CreatePaasWeimobGuideCustomerCrowdGetListResponse() *PaasWeimobGuideCustomerCrowdGetListResponse {
	return &PaasWeimobGuideCustomerCrowdGetListResponse{}
}

// OnPaasWeimobGuideCustomerCrowdGetListServiceInvocation
// @id 746
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/746?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询客户（批量）所属人群信息)
func (s *Service) OnPaasWeimobGuideCustomerCrowdGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideCustomerCrowdGetListRequest, PaasWeimobGuideCustomerCrowdGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideCustomerCrowdGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideCustomerCrowdGetListRequest, PaasWeimobGuideCustomerCrowdGetListResponse](invocation),
		"PaasWeimobGuideCustomerCrowdGetListService",
		"weimobGuide.customer.crowd.getList",
	)
	return s
}
