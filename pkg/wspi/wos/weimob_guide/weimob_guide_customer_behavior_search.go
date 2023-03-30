package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideCustomerBehaviorSearchService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideCustomerBehaviorSearchRequest, PaasWeimobGuideCustomerBehaviorSearchResponse]
}

func (s PaasWeimobGuideCustomerBehaviorSearchService) NewRequest() spi.InvocationRequest[PaasWeimobGuideCustomerBehaviorSearchRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideCustomerBehaviorSearchRequest]{
		Params: &PaasWeimobGuideCustomerBehaviorSearchRequest{},
	}
}

func (s PaasWeimobGuideCustomerBehaviorSearchService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideCustomerBehaviorSearchRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideCustomerBehaviorSearchResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideCustomerBehaviorSearchRequest]))
}

type PaasWeimobGuideCustomerBehaviorSearchRequest struct {
	QueryParameter PaasWeimobGuideCustomerBehaviorSearchRequestQueryParameter `json:"queryParameter,omitempty"`
	BasicInfo      PaasWeimobGuideCustomerBehaviorSearchRequestBasicInfo      `json:"basicInfo,omitempty"`
	PageSize       int64                                                      `json:"pageSize,omitempty"`
	Wid            int64                                                      `json:"wid,omitempty"`
	PageNum        int64                                                      `json:"pageNum,omitempty"`
	BizSceneType   int64                                                      `json:"bizSceneType,omitempty"`
}
type PaasWeimobGuideCustomerBehaviorSearchRequestQueryParameter struct {
	StartTime string `json:"startTime,omitempty"`
}
type PaasWeimobGuideCustomerBehaviorSearchRequestBasicInfo struct {
	Cid               int64  `json:"cid,omitempty"`
	Vid               int64  `json:"vid,omitempty"`
	Globalvid         int64  `json:"globalvid,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	MerchantId        int64  `json:"merchantId,omitempty"`
	ProductVersionId  int64  `json:"productVersionId,omitempty"`
	ProductId         int64  `json:"productId,omitempty"`
	BosId             int64  `json:"bosId,omitempty"`
	Tcode             string `json:"tcode,omitempty"`
	VidType           int64  `json:"vidType,omitempty"`
}

type PaasWeimobGuideCustomerBehaviorSearchResponse struct {
	PageList   []PaasWeimobGuideCustomerBehaviorSearchResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                                   `json:"pageNum,omitempty"`
	TotalCount int64                                                   `json:"totalCount,omitempty"`
	PageSize   int64                                                   `json:"pageSize,omitempty"`
}
type PaasWeimobGuideCustomerBehaviorSearchResponsePageList struct {
	Values    []PaasWeimobGuideCustomerBehaviorSearchResponseValues `json:"values,omitempty"`
	PageName  string                                                `json:"pageName,omitempty"`
	Content   string                                                `json:"content,omitempty"`
	Wid       int64                                                 `json:"wid,omitempty"`
	BizValue  string                                                `json:"bizValue,omitempty"`
	Timestamp string                                                `json:"timestamp,omitempty"`
	BosId     int64                                                 `json:"bosId,omitempty"`
	SessionId string                                                `json:"sessionId,omitempty"`
	BizType   int64                                                 `json:"bizType,omitempty"`
	Vid       int64                                                 `json:"vid,omitempty"`
}
type PaasWeimobGuideCustomerBehaviorSearchResponseValues struct {
	Value string `json:"value,omitempty"`
	Key   string `json:"key,omitempty"`
}

func CreatePaasWeimobGuideCustomerBehaviorSearchResponse() *PaasWeimobGuideCustomerBehaviorSearchResponse {
	return &PaasWeimobGuideCustomerBehaviorSearchResponse{}
}

// OnPaasWeimobGuideCustomerBehaviorSearchServiceInvocation
// @id 727
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/727?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询客户行为记录列表)
func (s *Service) OnPaasWeimobGuideCustomerBehaviorSearchServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideCustomerBehaviorSearchRequest, PaasWeimobGuideCustomerBehaviorSearchResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideCustomerBehaviorSearchService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideCustomerBehaviorSearchRequest, PaasWeimobGuideCustomerBehaviorSearchResponse](invocation),
		"PaasWeimobGuideCustomerBehaviorSearchService",
		"weimobGuide.customer.behavior.search",
	)
	return s
}
