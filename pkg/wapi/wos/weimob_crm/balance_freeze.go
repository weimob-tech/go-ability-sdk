package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// BalanceFreeze
// @id 1678
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1678?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=冻结余额)
func (client *Client) BalanceFreeze(request *BalanceFreezeRequest) (response *BalanceFreezeResponse, err error) {
	rpcResponse := CreateBalanceFreezeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type BalanceFreezeRequest struct {
	*api.BaseRequest

	ChangeType    string `json:"changeType,omitempty"`
	RequestId     string `json:"requestId,omitempty"`
	RequestType   string `json:"requestType,omitempty"`
	OutTransNo    string `json:"outTransNo,omitempty"`
	OutTransType  string `json:"outTransType,omitempty"`
	Amount        string `json:"amount,omitempty"`
	OccurVid      int64  `json:"occurVid,omitempty"`
	Mode          string `json:"mode,omitempty"`
	Remark        string `json:"remark,omitempty"`
	Vid           int64  `json:"vid,omitempty"`
	OperatorWid   int64  `json:"operatorWid,omitempty"`
	UserType      int64  `json:"userType,omitempty"`
	ChannelType   int64  `json:"channelType,omitempty"`
	BalancePlanId int64  `json:"balancePlanId,omitempty"`
	Wid           int64  `json:"wid,omitempty"`
}

type BalanceFreezeResponse struct {
	BosId       int64  `json:"bosId,omitempty"`
	Wid         int64  `json:"wid,omitempty"`
	Amount      string `json:"amount,omitempty"`
	ChangeType  string `json:"changeType,omitempty"`
	UserTransId int64  `json:"userTransId,omitempty"`
}

func CreateBalanceFreezeRequest() (request *BalanceFreezeRequest) {
	request = &BalanceFreezeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "balance/freeze", "apigw")
	request.Method = api.POST
	return
}

func CreateBalanceFreezeResponse() (response *api.BaseResponse[BalanceFreezeResponse]) {
	response = api.CreateResponse[BalanceFreezeResponse](&BalanceFreezeResponse{})
	return
}
