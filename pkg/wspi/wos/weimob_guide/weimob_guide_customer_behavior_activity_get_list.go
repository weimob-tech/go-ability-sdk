package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideCustomerBehaviorActivityGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideCustomerBehaviorActivityGetListRequest, PaasWeimobGuideCustomerBehaviorActivityGetListResponse]
}

func (s PaasWeimobGuideCustomerBehaviorActivityGetListService) NewRequest() spi.InvocationRequest[PaasWeimobGuideCustomerBehaviorActivityGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideCustomerBehaviorActivityGetListRequest]{
		Params: &PaasWeimobGuideCustomerBehaviorActivityGetListRequest{},
	}
}

func (s PaasWeimobGuideCustomerBehaviorActivityGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideCustomerBehaviorActivityGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideCustomerBehaviorActivityGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideCustomerBehaviorActivityGetListRequest]))
}

type PaasWeimobGuideCustomerBehaviorActivityGetListRequest struct {
	FilterList        PaasWeimobGuideCustomerBehaviorActivityGetListRequestFilterList `json:"filterList,omitempty"`
	GuiderWid         int64                                                           `json:"guiderWid,omitempty"`
	StartTime         int64                                                           `json:"startTime,omitempty"`
	EndTime           int64                                                           `json:"endTime,omitempty"`
	BosId             int64                                                           `json:"bosId,omitempty"`
	Vid               int64                                                           `json:"vid,omitempty"`
	ProductInstanceId int64                                                           `json:"productInstanceId,omitempty"`
	Order             string                                                          `json:"order,omitempty"`
	QueryUserSize     int64                                                           `json:"queryUserSize,omitempty"`
	QueryRecordSize   int64                                                           `json:"queryRecordSize,omitempty"`
	VidType           int64                                                           `json:"vidType,omitempty"`
	ProductId         int64                                                           `json:"productId,omitempty"`
}
type PaasWeimobGuideCustomerBehaviorActivityGetListRequestFilterList struct {
	En       string `json:"en,omitempty"`
	PageName string `json:"pageName,omitempty"`
}

type PaasWeimobGuideCustomerBehaviorActivityGetListResponse struct {
	OldBehaviorInfo    PaasWeimobGuideCustomerBehaviorActivityGetListResponseOldBehaviorInfo    `json:"oldBehaviorInfo,omitempty"`
	BehaviorRecordList PaasWeimobGuideCustomerBehaviorActivityGetListResponseBehaviorRecordList `json:"behaviorRecordList,omitempty"`
	IsOld              bool                                                                     `json:"isOld,omitempty"`
}
type PaasWeimobGuideCustomerBehaviorActivityGetListResponseOldBehaviorInfo struct {
}
type PaasWeimobGuideCustomerBehaviorActivityGetListResponseBehaviorRecordList struct {
	CustomerWid PaasWeimobGuideCustomerBehaviorActivityGetListResponseCustomer_wid `json:"customer_wid,omitempty"`
}
type PaasWeimobGuideCustomerBehaviorActivityGetListResponseCustomer_wid struct {
	BosId         int64  `json:"bosId,omitempty"`
	Ukey          string `json:"ukey,omitempty"`
	UkeyType      string `json:"ukeyType,omitempty"`
	PageName      string `json:"pageName,omitempty"`
	Vid           int64  `json:"vid,omitempty"`
	VidType       int64  `json:"vidType,omitempty"`
	En            string `json:"en,omitempty"`
	GoodsId       string `json:"goodsId,omitempty"`
	OrderId       string `json:"orderId,omitempty"`
	Keyword       string `json:"keyword,omitempty"`
	ActivityName  string `json:"activityName,omitempty"`
	VisitDuration string `json:"visitDuration,omitempty"`
}

func CreatePaasWeimobGuideCustomerBehaviorActivityGetListResponse() *PaasWeimobGuideCustomerBehaviorActivityGetListResponse {
	return &PaasWeimobGuideCustomerBehaviorActivityGetListResponse{}
}

// OnPaasWeimobGuideCustomerBehaviorActivityGetListServiceInvocation
// @id 728
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/728?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询客户行为记录)
func (s *Service) OnPaasWeimobGuideCustomerBehaviorActivityGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideCustomerBehaviorActivityGetListRequest, PaasWeimobGuideCustomerBehaviorActivityGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideCustomerBehaviorActivityGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideCustomerBehaviorActivityGetListRequest, PaasWeimobGuideCustomerBehaviorActivityGetListResponse](invocation),
		"PaasWeimobGuideCustomerBehaviorActivityGetListService",
		"weimobGuide.customer.behavior.activity.getList",
	)
	return s
}
