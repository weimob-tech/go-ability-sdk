package wmpay

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PaymentQueryRefund
// @id 281
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=退款状态查询)
func (client *Client) PaymentQueryRefund(request *PaymentQueryRefundRequest) (response *PaymentQueryRefundResponse, err error) {
	rpcResponse := CreatePaymentQueryRefundResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PaymentQueryRefundRequest struct {
	*api.BaseRequest

	ChannelType string `json:"channel_type,omitempty"`
	AppId       string `json:"app_id,omitempty"`
	RefundId    string `json:"refund_id,omitempty"`
	TraceNo     string `json:"trace_no,omitempty"`
}

type PaymentQueryRefundResponse struct {
}

func CreatePaymentQueryRefundRequest() (request *PaymentQueryRefundRequest) {
	request = &PaymentQueryRefundRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wmpay", "1_0", "payment/queryRefund", "api")
	request.Method = api.POST
	return
}

func CreatePaymentQueryRefundResponse() (response *api.BaseResponse[PaymentQueryRefundResponse]) {
	response = api.CreateResponse[PaymentQueryRefundResponse](&PaymentQueryRefundResponse{})
	return
}
