package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// BalanceConfirm
// @id 1835
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1835?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=余额锁定确认)
func (client *Client) BalanceConfirm(request *BalanceConfirmRequest) (response *BalanceConfirmResponse, err error) {
	rpcResponse := CreateBalanceConfirmResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type BalanceConfirmRequest struct {
	*api.BaseRequest

	RequestId   int64  `json:"requestId,omitempty"`
	RequestType string `json:"requestType,omitempty"`
	OutTransNo  string `json:"outTransNo,omitempty"`
	Amount      string `json:"amount,omitempty"`
	ChangeType  string `json:"changeType,omitempty"`
	Remark      string `json:"remark,omitempty"`
	UserTransId int64  `json:"userTransId,omitempty"`
	Wid         int64  `json:"wid,omitempty"`
	Vid         int64  `json:"vid,omitempty"`
	OperatorWid int64  `json:"operatorWid,omitempty"`
	ChannelType string `json:"channelType,omitempty"`
	PlanId      int64  `json:"planId,omitempty"`
}

type BalanceConfirmResponse struct {
}

func CreateBalanceConfirmRequest() (request *BalanceConfirmRequest) {
	request = &BalanceConfirmRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "balance/confirm", "apigw")
	request.Method = api.POST
	return
}

func CreateBalanceConfirmResponse() (response *api.BaseResponse[BalanceConfirmResponse]) {
	response = api.CreateResponse[BalanceConfirmResponse](&BalanceConfirmResponse{})
	return
}
