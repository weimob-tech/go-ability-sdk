package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpQueryRechargeRule
// @id 1818
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询充值规则)
func (client *Client) MbpQueryRechargeRule(request *MbpQueryRechargeRuleRequest) (response *MbpQueryRechargeRuleResponse, err error) {
	rpcResponse := CreateMbpQueryRechargeRuleResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpQueryRechargeRuleRequest struct {
	*api.BaseRequest

	Pid     int64 `json:"pid,omitempty"`
	StoreId int64 `json:"storeId,omitempty"`
}

type MbpQueryRechargeRuleResponse struct {
	RechargeName        string `json:"rechargeName,omitempty"`
	CrowdType           int    `json:"crowdType,omitempty"`
	MaxBalanceLimitFlag int    `json:"maxBalanceLimitFlag,omitempty"`
	MaxBalance          int64  `json:"maxBalance,omitempty"`
	RechargeLimitFlag   int    `json:"rechargeLimitFlag,omitempty"`
	MinAmountFlag       int    `json:"minAmountFlag,omitempty"`
	MinAmount           int64  `json:"minAmount,omitempty"`
	MaxAmountFlag       int    `json:"maxAmountFlag,omitempty"`
	MaxAmount           int64  `json:"maxAmount,omitempty"`
	FixedAmountFlag     string `json:"fixedAmountFlag,omitempty"`
}

func CreateMbpQueryRechargeRuleRequest() (request *MbpQueryRechargeRuleRequest) {
	request = &MbpQueryRechargeRuleRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "mbp/queryRechargeRule", "api")
	request.Method = api.POST
	return
}

func CreateMbpQueryRechargeRuleResponse() (response *api.BaseResponse[MbpQueryRechargeRuleResponse]) {
	response = api.CreateResponse[MbpQueryRechargeRuleResponse](&MbpQueryRechargeRuleResponse{})
	return
}
