package wmpay

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PaymentQueryPayStatus
// @id 278
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=支付状态查询（预留接口）)
func (client *Client) PaymentQueryPayStatus(request *PaymentQueryPayStatusRequest) (response *PaymentQueryPayStatusResponse, err error) {
	rpcResponse := CreatePaymentQueryPayStatusResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PaymentQueryPayStatusRequest struct {
	*api.BaseRequest

	AppId       string `json:"app_id,omitempty"`
	ChannelType string `json:"channel_type,omitempty"`
	TradeId     string `json:"trade_id,omitempty"`
	TraceNo     string `json:"trace_no,omitempty"`
}

type PaymentQueryPayStatusResponse struct {
}

func CreatePaymentQueryPayStatusRequest() (request *PaymentQueryPayStatusRequest) {
	request = &PaymentQueryPayStatusRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wmpay", "1_0", "payment/queryPayStatus", "api")
	request.Method = api.POST
	return
}

func CreatePaymentQueryPayStatusResponse() (response *api.BaseResponse[PaymentQueryPayStatusResponse]) {
	response = api.CreateResponse[PaymentQueryPayStatusResponse](&PaymentQueryPayStatusResponse{})
	return
}
