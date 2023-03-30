package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuidePointRuleGetService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuidePointRuleGetRequest, PaasWeimobGuidePointRuleGetResponse]
}

func (s PaasWeimobGuidePointRuleGetService) NewRequest() spi.InvocationRequest[PaasWeimobGuidePointRuleGetRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuidePointRuleGetRequest]{
		Params: &PaasWeimobGuidePointRuleGetRequest{},
	}
}

func (s PaasWeimobGuidePointRuleGetService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuidePointRuleGetRequest],
) (
	spi.InvocationResponse[PaasWeimobGuidePointRuleGetResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuidePointRuleGetRequest]))
}

type PaasWeimobGuidePointRuleGetRequest struct {
	PointPlanId int64  `json:"pointPlanId,omitempty"`
	Fields      string `json:"fields,omitempty"`
}

type PaasWeimobGuidePointRuleGetResponse struct {
	ExpiryRule    PaasWeimobGuidePointRuleGetResponseExpiryRule `json:"expiryRule,omitempty"`
	DeductRule    PaasWeimobGuidePointRuleGetResponseDeductRule `json:"deductRule,omitempty"`
	PointUnit     PaasWeimobGuidePointRuleGetResponsePointUnit  `json:"pointUnit,omitempty"`
	BasicRule     PaasWeimobGuidePointRuleGetResponseBasicRule  `json:"basicRule,omitempty"`
	PointPlanId   int64                                         `json:"pointPlanId,omitempty"`
	PointPlanName string                                        `json:"pointPlanName,omitempty"`
	Instruction   string                                        `json:"instruction,omitempty"`
}
type PaasWeimobGuidePointRuleGetResponseExpiryRule struct {
	Type  string `json:"type,omitempty"`
	Year  int64  `json:"year,omitempty"`
	Month int64  `json:"month,omitempty"`
	Day   int64  `json:"day,omitempty"`
}
type PaasWeimobGuidePointRuleGetResponseDeductRule struct {
	DeductItem                int64   `json:"deductItem,omitempty"`
	DeductGoodsItem           int64   `json:"deductGoodsItem,omitempty"`
	Freight                   int64   `json:"freight,omitempty"`
	ApplicableScope           int64   `json:"applicableScope,omitempty"`
	IsExcludeGoods            int64   `json:"isExcludeGoods,omitempty"`
	IsOpenVerify              int64   `json:"isOpenVerify,omitempty"`
	DefaultVerify             int64   `json:"defaultVerify,omitempty"`
	DefaultUseFlag            int64   `json:"defaultUseFlag,omitempty"`
	MinDeductPointRequire     int64   `json:"minDeductPointRequire,omitempty"`
	MaxDeductPointRequire     int64   `json:"maxDeductPointRequire,omitempty"`
	DeductItemAmtRatioRequire int64   `json:"deductItemAmtRatioRequire,omitempty"`
	ApplicableGoods           []int64 `json:"applicableGoods,omitempty"`
	MinDeductPoint            int64   `json:"minDeductPoint,omitempty"`
	MaxDeductPoint            int64   `json:"maxDeductPoint,omitempty"`
	DeductItemAmtRatio        int64   `json:"deductItemAmtRatio,omitempty"`
	ExcludeGoods              []int64 `json:"excludeGoods,omitempty"`
}
type PaasWeimobGuidePointRuleGetResponsePointUnit struct {
	Unit       int64  `json:"unit,omitempty"`
	UnitAmount string `json:"unitAmount,omitempty"`
}
type PaasWeimobGuidePointRuleGetResponseBasicRule struct {
	SourceVid         int64  `json:"sourceVid,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	AvailableCrowd    int64  `json:"availableCrowd,omitempty"`
	UpdateTime        string `json:"updateTime,omitempty"`
}

func CreatePaasWeimobGuidePointRuleGetResponse() *PaasWeimobGuidePointRuleGetResponse {
	return &PaasWeimobGuidePointRuleGetResponse{}
}

// OnPaasWeimobGuidePointRuleGetServiceInvocation
// @id 925
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/925?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询积分抵扣规则)
func (s *Service) OnPaasWeimobGuidePointRuleGetServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuidePointRuleGetRequest, PaasWeimobGuidePointRuleGetResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuidePointRuleGetService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuidePointRuleGetRequest, PaasWeimobGuidePointRuleGetResponse](invocation),
		"PaasWeimobGuidePointRuleGetService",
		"weimobGuide.point.rule.get",
	)
	return s
}
