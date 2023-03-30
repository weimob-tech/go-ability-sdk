package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// BalanceRecharge
// @id 1666
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1666?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=余额充值)
func (client *Client) BalanceRecharge(request *BalanceRechargeRequest) (response *BalanceRechargeResponse, err error) {
	rpcResponse := CreateBalanceRechargeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type BalanceRechargeRequest struct {
	*api.BaseRequest

	SettleRule    []BalanceRechargeRequestSettleRule `json:"settleRule,omitempty"`
	RequestId     string                             `json:"requestId,omitempty"`
	RequestType   string                             `json:"requestType,omitempty"`
	OutTransNo    string                             `json:"outTransNo,omitempty"`
	OutTransType  string                             `json:"outTransType,omitempty"`
	CapitalAmount string                             `json:"capitalAmount,omitempty"`
	GrantsAmount  string                             `json:"grantsAmount,omitempty"`
	ChangeType    string                             `json:"changeType,omitempty"`
	OccurVid      int64                              `json:"occurVid,omitempty"`
	Remark        string                             `json:"remark,omitempty"`
	SourceVid     int64                              `json:"sourceVid,omitempty"`
	Vid           int64                              `json:"vid,omitempty"`
	OperatorWid   int64                              `json:"operatorWid,omitempty"`
	BalancePlanId int64                              `json:"balancePlanId,omitempty"`
	Wid           int64                              `json:"wid,omitempty"`
}

type BalanceRechargeResponse struct {
	Result     bool   `json:"result,omitempty"`
	BizTransId string `json:"bizTransId,omitempty"`
}

func CreateBalanceRechargeRequest() (request *BalanceRechargeRequest) {
	request = &BalanceRechargeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "balance/recharge", "apigw")
	request.Method = api.POST
	return
}

func CreateBalanceRechargeResponse() (response *api.BaseResponse[BalanceRechargeResponse]) {
	response = api.CreateResponse[BalanceRechargeResponse](&BalanceRechargeResponse{})
	return
}

type BalanceRechargeRequestSettleRule struct {
	Vid  int64 `json:"vid,omitempty"`
	Rate int64 `json:"rate,omitempty"`
}
