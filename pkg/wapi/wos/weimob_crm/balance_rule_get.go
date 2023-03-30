package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// BalanceRuleGet
// @id 1494
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1494?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取余额方案规则)
func (client *Client) BalanceRuleGet(request *BalanceRuleGetRequest) (response *BalanceRuleGetResponse, err error) {
	rpcResponse := CreateBalanceRuleGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type BalanceRuleGetRequest struct {
	*api.BaseRequest

	Vid           int64 `json:"vid,omitempty"`
	BalancePlanId int64 `json:"balancePlanId,omitempty"`
}

type BalanceRuleGetResponse struct {
	UseRule         BalanceRuleGetResponseUseRule      `json:"useRule,omitempty"`
	RechargeRule    BalanceRuleGetResponseRechargeRule `json:"rechargeRule,omitempty"`
	BasicRule       BalanceRuleGetResponseBasicRule    `json:"basicRule,omitempty"`
	BalancePlanId   int64                              `json:"balancePlanId,omitempty"`
	BalancePlanName string                             `json:"balancePlanName,omitempty"`
	Instruction     string                             `json:"instruction,omitempty"`
}

func CreateBalanceRuleGetRequest() (request *BalanceRuleGetRequest) {
	request = &BalanceRuleGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "balance/rule/get", "apigw")
	request.Method = api.POST
	return
}

func CreateBalanceRuleGetResponse() (response *api.BaseResponse[BalanceRuleGetResponse]) {
	response = api.CreateResponse[BalanceRuleGetResponse](&BalanceRuleGetResponse{})
	return
}

type BalanceRuleGetResponseUseRule struct {
	MerchantPay   int64 `json:"merchantPay,omitempty"`
	UserPay       int64 `json:"userPay,omitempty"`
	IsOpenVerify  int64 `json:"isOpenVerify,omitempty"`
	DefaultVerify int64 `json:"defaultVerify,omitempty"`
	UseStatus     int64 `json:"useStatus,omitempty"`
}

type BalanceRuleGetResponseRechargeRule struct {
	RechargeType       int64   `json:"rechargeType,omitempty"`
	RechargeAmountList []int64 `json:"rechargeAmountList,omitempty"`
	CustomRechargeType int64   `json:"customRechargeType,omitempty"`
	RechargeLimit      int64   `json:"rechargeLimit,omitempty"`
	MinRechargeLimit   int64   `json:"minRechargeLimit,omitempty"`
	MinRechargeAmount  string  `json:"minRechargeAmount,omitempty"`
	MaxRechargeLimit   int64   `json:"maxRechargeLimit,omitempty"`
	MaxRechargeAmount  string  `json:"maxRechargeAmount,omitempty"`
	PayeeType          int64   `json:"payeeType,omitempty"`
	TargetVid          []int64 `json:"targetVid,omitempty"`
}

type BalanceRuleGetResponseBasicRule struct {
	AvailableCrowd   int64  `json:"availableCrowd,omitempty"`
	BalanceLimitType int64  `json:"balanceLimitType,omitempty"`
	BalanceLimit     string `json:"balanceLimit,omitempty"`
	UpdateTime       string `json:"updateTime,omitempty"`
}
