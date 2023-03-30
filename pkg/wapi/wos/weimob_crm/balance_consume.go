package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// BalanceConsume
// @id 1680
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1680?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=消耗余额)
func (client *Client) BalanceConsume(request *BalanceConsumeRequest) (response *BalanceConsumeResponse, err error) {
	rpcResponse := CreateBalanceConsumeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type BalanceConsumeRequest struct {
	*api.BaseRequest

	ChangeType    string `json:"changeType,omitempty"`
	RequestId     string `json:"requestId,omitempty"`
	RequestType   string `json:"requestType,omitempty"`
	OutTransNo    string `json:"outTransNo,omitempty"`
	OutTransType  string `json:"outTransType,omitempty"`
	Amount        string `json:"amount,omitempty"`
	OccurVid      int64  `json:"occurVid,omitempty"`
	Remark        string `json:"remark,omitempty"`
	Vid           int64  `json:"vid,omitempty"`
	OperatorWid   int64  `json:"operatorWid,omitempty"`
	UserType      int64  `json:"userType,omitempty"`
	ChannelType   int64  `json:"channelType,omitempty"`
	BalancePlanId int64  `json:"balancePlanId,omitempty"`
	Wid           int64  `json:"wid,omitempty"`
	UserTransId   int64  `json:"userTransId,omitempty"`
}

type BalanceConsumeResponse struct {
	Amount      string `json:"amount,omitempty"`
	Wid         int64  `json:"wid,omitempty"`
	ChangeType  string `json:"changeType,omitempty"`
	BosId       int64  `json:"bosId,omitempty"`
	UserTransId int64  `json:"userTransId,omitempty"`
}

func CreateBalanceConsumeRequest() (request *BalanceConsumeRequest) {
	request = &BalanceConsumeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "balance/consume", "apigw")
	request.Method = api.POST
	return
}

func CreateBalanceConsumeResponse() (response *api.BaseResponse[BalanceConsumeResponse]) {
	response = api.CreateResponse[BalanceConsumeResponse](&BalanceConsumeResponse{})
	return
}
