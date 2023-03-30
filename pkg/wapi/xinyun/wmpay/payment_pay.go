package wmpay

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PaymentPay
// @id 277
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=支付请求（预留接口）)
func (client *Client) PaymentPay(request *PaymentPayRequest) (response *PaymentPayResponse, err error) {
	rpcResponse := CreatePaymentPayResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PaymentPayRequest struct {
	*api.BaseRequest

	AppId       string `json:"app_id,omitempty"`
	TraceNo     string `json:"trace_no,omitempty"`
	ProductNo   string `json:"product_no,omitempty"`
	RequestTime int64  `json:"request_time,omitempty"`
	TradeDesc   string `json:"trade_desc,omitempty"`
	OrderNo     string `json:"order_no,omitempty"`
	TradeAmount int64  `json:"trade_amount,omitempty"`
	Currency    string `json:"currency,omitempty"`
	OpenId      string `json:"open_id,omitempty"`
	ClientIp    string `json:"client_ip,omitempty"`
	ChannelType string `json:"channel_type,omitempty"`
}

type PaymentPayResponse struct {
}

func CreatePaymentPayRequest() (request *PaymentPayRequest) {
	request = &PaymentPayRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wmpay", "1_0", "payment/pay", "api")
	request.Method = api.POST
	return
}

func CreatePaymentPayResponse() (response *api.BaseResponse[PaymentPayResponse]) {
	response = api.CreateResponse[PaymentPayResponse](&PaymentPayResponse{})
	return
}
