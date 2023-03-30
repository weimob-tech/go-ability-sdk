package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpQueryPointReductRule
// @id 1819
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询积分抵现规则)
func (client *Client) MbpQueryPointReductRule(request *MbpQueryPointReductRuleRequest) (response *MbpQueryPointReductRuleResponse, err error) {
	rpcResponse := CreateMbpQueryPointReductRuleResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpQueryPointReductRuleRequest struct {
	*api.BaseRequest

	Pid     int64 `json:"pid,omitempty"`
	StoreId int64 `json:"storeId,omitempty"`
}

type MbpQueryPointReductRuleResponse struct {
	PointName             string `json:"pointName,omitempty"`
	CanUsePointFlag       int    `json:"canUsePointFlag,omitempty"`
	CrowdType             int    `json:"crowdType,omitempty"`
	MinReductUnit         int    `json:"minReductUnit,omitempty"`
	ReductRuleAmount      int64  `json:"reductRuleAmount,omitempty"`
	VerifyCustomerFlag    int    `json:"verifyCustomerFlag,omitempty"`
	VerifyDefaultStatus   int    `json:"verifyDefaultStatus,omitempty"`
	DefaultSettlementFlag int    `json:"defaultSettlementFlag,omitempty"`
	RuleScope             int    `json:"ruleScope,omitempty"`
	ReductItem            int    `json:"reductItem,omitempty"`
	GoodsType             int    `json:"goodsType,omitempty"`
}

func CreateMbpQueryPointReductRuleRequest() (request *MbpQueryPointReductRuleRequest) {
	request = &MbpQueryPointReductRuleRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "mbp/queryPointReductRule", "api")
	request.Method = api.POST
	return
}

func CreateMbpQueryPointReductRuleResponse() (response *api.BaseResponse[MbpQueryPointReductRuleResponse]) {
	response = api.CreateResponse[MbpQueryPointReductRuleResponse](&MbpQueryPointReductRuleResponse{})
	return
}
