package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FinancePaymentWeixinpayMessageGet
// @id 1791
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1791?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=支付交易报文)
func (client *Client) FinancePaymentWeixinpayMessageGet(request *FinancePaymentWeixinpayMessageGetRequest) (response *FinancePaymentWeixinpayMessageGetResponse, err error) {
	rpcResponse := CreateFinancePaymentWeixinpayMessageGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FinancePaymentWeixinpayMessageGetRequest struct {
	*api.BaseRequest

	InteractId string `json:"interactId,omitempty"`
}

type FinancePaymentWeixinpayMessageGetResponse struct {
	InteractId      string `json:"interactId,omitempty"`
	RequestMessage  string `json:"requestMessage,omitempty"`
	ChannelTrxNo    string `json:"channelTrxNo,omitempty"`
	ChannelCode     string `json:"channelCode,omitempty"`
	ResponseMessage string `json:"responseMessage,omitempty"`
}

func CreateFinancePaymentWeixinpayMessageGetRequest() (request *FinancePaymentWeixinpayMessageGetRequest) {
	request = &FinancePaymentWeixinpayMessageGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "finance/payment/weixinpay/message/get", "apigw")
	request.Method = api.POST
	return
}

func CreateFinancePaymentWeixinpayMessageGetResponse() (response *api.BaseResponse[FinancePaymentWeixinpayMessageGetResponse]) {
	response = api.CreateResponse[FinancePaymentWeixinpayMessageGetResponse](&FinancePaymentWeixinpayMessageGetResponse{})
	return
}
