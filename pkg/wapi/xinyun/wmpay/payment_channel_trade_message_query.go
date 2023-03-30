package wmpay

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PaymentChannelTradeMessageQuery
// @id 1871
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=支付通道交易报文查询)
func (client *Client) PaymentChannelTradeMessageQuery(request *PaymentChannelTradeMessageQueryRequest) (response *PaymentChannelTradeMessageQueryResponse, err error) {
	rpcResponse := CreatePaymentChannelTradeMessageQueryResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PaymentChannelTradeMessageQueryRequest struct {
	*api.BaseRequest

	InteractId string `json:"interactId,omitempty"`
}

type PaymentChannelTradeMessageQueryResponse struct {
}

func CreatePaymentChannelTradeMessageQueryRequest() (request *PaymentChannelTradeMessageQueryRequest) {
	request = &PaymentChannelTradeMessageQueryRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wmpay", "1_0", "payment/channelTradeMessageQuery", "api")
	request.Method = api.POST
	return
}

func CreatePaymentChannelTradeMessageQueryResponse() (response *api.BaseResponse[PaymentChannelTradeMessageQueryResponse]) {
	response = api.CreateResponse[PaymentChannelTradeMessageQueryResponse](&PaymentChannelTradeMessageQueryResponse{})
	return
}
