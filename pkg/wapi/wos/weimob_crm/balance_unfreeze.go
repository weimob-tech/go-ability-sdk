package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// BalanceUnfreeze
// @id 1679
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1679?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=解冻余额)
func (client *Client) BalanceUnfreeze(request *BalanceUnfreezeRequest) (response *BalanceUnfreezeResponse, err error) {
	rpcResponse := CreateBalanceUnfreezeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type BalanceUnfreezeRequest struct {
	*api.BaseRequest

	RequestId     string `json:"requestId,omitempty"`
	RequestType   string `json:"requestType,omitempty"`
	OutTransNo    string `json:"outTransNo,omitempty"`
	Vid           int64  `json:"vid,omitempty"`
	OperatorWid   int64  `json:"operatorWid,omitempty"`
	UserType      int64  `json:"userType,omitempty"`
	ChannelType   int64  `json:"channelType,omitempty"`
	BalancePlanId int64  `json:"balancePlanId,omitempty"`
	Wid           int64  `json:"wid,omitempty"`
	UserTransId   int64  `json:"userTransId,omitempty"`
}

type BalanceUnfreezeResponse struct {
}

func CreateBalanceUnfreezeRequest() (request *BalanceUnfreezeRequest) {
	request = &BalanceUnfreezeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "balance/unfreeze", "apigw")
	request.Method = api.POST
	return
}

func CreateBalanceUnfreezeResponse() (response *api.BaseResponse[BalanceUnfreezeResponse]) {
	response = api.CreateResponse[BalanceUnfreezeResponse](&BalanceUnfreezeResponse{})
	return
}
