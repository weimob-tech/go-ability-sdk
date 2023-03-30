package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PointRuleGet
// @id 1490
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1490?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取积分方案规则)
func (client *Client) PointRuleGet(request *PointRuleGetRequest) (response *PointRuleGetResponse, err error) {
	rpcResponse := CreatePointRuleGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PointRuleGetRequest struct {
	*api.BaseRequest

	Vid         int64 `json:"vid,omitempty"`
	PointPlanId int64 `json:"pointPlanId,omitempty"`
}

type PointRuleGetResponse struct {
	ExpiryRule    PointRuleGetResponseExpiryRule `json:"expiryRule,omitempty"`
	PointUnit     PointRuleGetResponsePointUnit  `json:"pointUnit,omitempty"`
	DeductRule    PointRuleGetResponseDeductRule `json:"deductRule,omitempty"`
	BasicRule     PointRuleGetResponseBasicRule  `json:"basicRule,omitempty"`
	PointPlanId   int64                          `json:"pointPlanId,omitempty"`
	PointPlanName string                         `json:"pointPlanName,omitempty"`
	Instruction   string                         `json:"instruction,omitempty"`
	DeductFlag    int64                          `json:"deductFlag,omitempty"`
}

func CreatePointRuleGetRequest() (request *PointRuleGetRequest) {
	request = &PointRuleGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "point/rule/get", "apigw")
	request.Method = api.POST
	return
}

func CreatePointRuleGetResponse() (response *api.BaseResponse[PointRuleGetResponse]) {
	response = api.CreateResponse[PointRuleGetResponse](&PointRuleGetResponse{})
	return
}

type PointRuleGetResponseExpiryRule struct {
	Type  string `json:"type,omitempty"`
	Year  int64  `json:"year,omitempty"`
	Month int64  `json:"month,omitempty"`
	Day   int64  `json:"day,omitempty"`
}

type PointRuleGetResponsePointUnit struct {
	Unit       int64 `json:"unit,omitempty"`
	UnitAmount int64 `json:"unitAmount,omitempty"`
}

type PointRuleGetResponseDeductRule struct {
	DeductItem                int64   `json:"deductItem,omitempty"`
	DeductGoodsItem           int64   `json:"deductGoodsItem,omitempty"`
	Freight                   int64   `json:"freight,omitempty"`
	ApplicableScope           int64   `json:"applicableScope,omitempty"`
	IsExcludeGoods            int64   `json:"isExcludeGoods,omitempty"`
	IsOpenVerify              int64   `json:"isOpenVerify,omitempty"`
	DefaultVerify             int64   `json:"defaultVerify,omitempty"`
	DefaultUseFlag            int64   `json:"defaultUseFlag,omitempty"`
	MinDeductPointRequire     int64   `json:"minDeductPointRequire,omitempty"`
	MinDeductPoint            int64   `json:"minDeductPoint,omitempty"`
	MaxDeductPointRequire     int64   `json:"maxDeductPointRequire,omitempty"`
	MaxDeductPoint            int64   `json:"maxDeductPoint,omitempty"`
	DeductItemAmtRatioRequire int64   `json:"deductItemAmtRatioRequire,omitempty"`
	DeductItemAmtRatio        int64   `json:"deductItemAmtRatio,omitempty"`
	ExcludeGoods              []int64 `json:"excludeGoods,omitempty"`
	ApplicableGoods           []int64 `json:"applicableGoods,omitempty"`
}

type PointRuleGetResponseBasicRule struct {
	SourceVid         int64  `json:"sourceVid,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	AvailableCrowd    int64  `json:"availableCrowd,omitempty"`
	UpdateTime        string `json:"updateTime,omitempty"`
}
