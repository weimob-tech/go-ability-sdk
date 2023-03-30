package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// BalanceRefund
// @id 1681
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1681?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=回退余额)
func (client *Client) BalanceRefund(request *BalanceRefundRequest) (response *BalanceRefundResponse, err error) {
	rpcResponse := CreateBalanceRefundResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type BalanceRefundRequest struct {
	*api.BaseRequest

	RefundRequestId   string `json:"refundRequestId,omitempty"`
	RefundRequestType string `json:"refundRequestType,omitempty"`
	RequestId         string `json:"requestId,omitempty"`
	RequestType       string `json:"requestType,omitempty"`
	UserTransId       int64  `json:"userTransId,omitempty"`
	OutTransNo        string `json:"outTransNo,omitempty"`
	Amount            string `json:"amount,omitempty"`
	ChangeType        string `json:"changeType,omitempty"`
	Remark            string `json:"remark,omitempty"`
	OccurVid          int64  `json:"occurVid,omitempty"`
	Vid               int64  `json:"vid,omitempty"`
	OperatorWid       int64  `json:"operatorWid,omitempty"`
	UserType          int64  `json:"userType,omitempty"`
	ChannelType       int64  `json:"channelType,omitempty"`
	BalancePlanId     int64  `json:"balancePlanId,omitempty"`
	Wid               int64  `json:"wid,omitempty"`
}

type BalanceRefundResponse struct {
	RefundTransId int64 `json:"refundTransId,omitempty"`
}

func CreateBalanceRefundRequest() (request *BalanceRefundRequest) {
	request = &BalanceRefundRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "balance/refund", "apigw")
	request.Method = api.POST
	return
}

func CreateBalanceRefundResponse() (response *api.BaseResponse[BalanceRefundResponse]) {
	response = api.CreateResponse[BalanceRefundResponse](&BalanceRefundResponse{})
	return
}
