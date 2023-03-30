package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// BalanceGet
// @id 1496
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1496?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询用户余额)
func (client *Client) BalanceGet(request *BalanceGetRequest) (response *BalanceGetResponse, err error) {
	rpcResponse := CreateBalanceGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type BalanceGetRequest struct {
	*api.BaseRequest

	IsStatistics  bool  `json:"isStatistics,omitempty"`
	Vid           int64 `json:"vid,omitempty"`
	BalancePlanId int64 `json:"balancePlanId,omitempty"`
	Wid           int64 `json:"wid,omitempty"`
}

type BalanceGetResponse struct {
	Wid                int64  `json:"wid,omitempty"`
	AcctStatus         string `json:"acctStatus,omitempty"`
	CustomerStatus     int64  `json:"customerStatus,omitempty"`
	AvailableAmount    string `json:"availableAmount,omitempty"`
	TotalAmount        string `json:"totalAmount,omitempty"`
	FrozenAmount       string `json:"frozenAmount,omitempty"`
	TotalIssueAmount   string `json:"totalIssueAmount,omitempty"`
	TotalCapitalAmount string `json:"totalCapitalAmount,omitempty"`
	TotalGrantsAmount  string `json:"totalGrantsAmount,omitempty"`
	RechargeTimes      int64  `json:"rechargeTimes,omitempty"`
}

func CreateBalanceGetRequest() (request *BalanceGetRequest) {
	request = &BalanceGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "balance/get", "apigw")
	request.Method = api.POST
	return
}

func CreateBalanceGetResponse() (response *api.BaseResponse[BalanceGetResponse]) {
	response = api.CreateResponse[BalanceGetResponse](&BalanceGetResponse{})
	return
}
