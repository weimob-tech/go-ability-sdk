package wmpay

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PaymentOverseaTradeMessageQuery
// @id 1798
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=境外通道交易报文查询)
func (client *Client) PaymentOverseaTradeMessageQuery(request *PaymentOverseaTradeMessageQueryRequest) (response *PaymentOverseaTradeMessageQueryResponse, err error) {
	rpcResponse := CreatePaymentOverseaTradeMessageQueryResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PaymentOverseaTradeMessageQueryRequest struct {
	*api.BaseRequest

	ChannelMerchantNo string `json:"channelMerchantNo,omitempty"`
	InteractId        string `json:"interactId,omitempty"`
}

type PaymentOverseaTradeMessageQueryResponse struct {
}

func CreatePaymentOverseaTradeMessageQueryRequest() (request *PaymentOverseaTradeMessageQueryRequest) {
	request = &PaymentOverseaTradeMessageQueryRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wmpay", "1_0", "payment/overseaTradeMessageQuery", "api")
	request.Method = api.POST
	return
}

func CreatePaymentOverseaTradeMessageQueryResponse() (response *api.BaseResponse[PaymentOverseaTradeMessageQueryResponse]) {
	response = api.CreateResponse[PaymentOverseaTradeMessageQueryResponse](&PaymentOverseaTradeMessageQueryResponse{})
	return
}
