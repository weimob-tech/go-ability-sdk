package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideCustomerTagGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideCustomerTagGetListRequest, PaasWeimobGuideCustomerTagGetListResponse]
}

func (s PaasWeimobGuideCustomerTagGetListService) NewRequest() spi.InvocationRequest[PaasWeimobGuideCustomerTagGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideCustomerTagGetListRequest]{
		Params: &PaasWeimobGuideCustomerTagGetListRequest{},
	}
}

func (s PaasWeimobGuideCustomerTagGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideCustomerTagGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideCustomerTagGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideCustomerTagGetListRequest]))
}

type PaasWeimobGuideCustomerTagGetListRequest struct {
	BosId             int64   `json:"bosId,omitempty"`
	ProductInstanceId int64   `json:"productInstanceId,omitempty"`
	Vid               int64   `json:"vid,omitempty"`
	VidType           int64   `json:"vidType,omitempty"`
	WidList           []int64 `json:"widList,omitempty"`
}

type PaasWeimobGuideCustomerTagGetListResponse struct {
	TagInfoMap PaasWeimobGuideCustomerTagGetListResponseTagInfoMap `json:"tagInfoMap,omitempty"`
}
type PaasWeimobGuideCustomerTagGetListResponseTagInfoMap struct {
	CustomerWid PaasWeimobGuideCustomerTagGetListResponseCustomer_wid `json:"customer_wid,omitempty"`
}
type PaasWeimobGuideCustomerTagGetListResponseCustomer_wid struct {
	UserTagAttrVoList []PaasWeimobGuideCustomerTagGetListResponseUserTagAttrVoList `json:"userTagAttrVoList,omitempty"`
	TagId             int64                                                        `json:"tagId,omitempty"`
	TagName           string                                                       `json:"tagName,omitempty"`
	UpdateTime        string                                                       `json:"updateTime,omitempty"`
}
type PaasWeimobGuideCustomerTagGetListResponseUserTagAttrVoList struct {
	AttrId     int64  `json:"attrId,omitempty"`
	AttrName   string `json:"attrName,omitempty"`
	StartValue string `json:"startValue,omitempty"`
	EndValue   string `json:"endValue,omitempty"`
}

func CreatePaasWeimobGuideCustomerTagGetListResponse() *PaasWeimobGuideCustomerTagGetListResponse {
	return &PaasWeimobGuideCustomerTagGetListResponse{}
}

// OnPaasWeimobGuideCustomerTagGetListServiceInvocation
// @id 747
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/747?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询客户标签)
func (s *Service) OnPaasWeimobGuideCustomerTagGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideCustomerTagGetListRequest, PaasWeimobGuideCustomerTagGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideCustomerTagGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideCustomerTagGetListRequest, PaasWeimobGuideCustomerTagGetListResponse](invocation),
		"PaasWeimobGuideCustomerTagGetListService",
		"weimobGuide.customer.tag.getList",
	)
	return s
}
