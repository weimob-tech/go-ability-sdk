package wmpay

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PaymentRefund
// @id 280
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=退款申请)
func (client *Client) PaymentRefund(request *PaymentRefundRequest) (response *PaymentRefundResponse, err error) {
	rpcResponse := CreatePaymentRefundResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PaymentRefundRequest struct {
	*api.BaseRequest

	ChannelType     string `json:"channel_type,omitempty"`
	AppId           string `json:"app_id,omitempty"`
	TraceNo         string `json:"trace_no,omitempty"`
	RefundNo        string `json:"refund_no,omitempty"`
	OriginalTraceNo string `json:"original_trace_no,omitempty"`
	TradeId         string `json:"trade_id,omitempty"`
	RefundAmount    int64  `json:"refund_amount,omitempty"`
	RefundDesc      string `json:"refund_desc,omitempty"`
	RequestTime     int64  `json:"request_time,omitempty"`
}

type PaymentRefundResponse struct {
}

func CreatePaymentRefundRequest() (request *PaymentRefundRequest) {
	request = &PaymentRefundRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wmpay", "1_0", "payment/refund", "api")
	request.Method = api.POST
	return
}

func CreatePaymentRefundResponse() (response *api.BaseResponse[PaymentRefundResponse]) {
	response = api.CreateResponse[PaymentRefundResponse](&PaymentRefundResponse{})
	return
}
