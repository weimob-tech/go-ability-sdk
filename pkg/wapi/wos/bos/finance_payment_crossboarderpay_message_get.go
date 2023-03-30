package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FinancePaymentCrossboarderpayMessageGet
// @id 1790
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1790?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=跨境支付报文)
func (client *Client) FinancePaymentCrossboarderpayMessageGet(request *FinancePaymentCrossboarderpayMessageGetRequest) (response *FinancePaymentCrossboarderpayMessageGetResponse, err error) {
	rpcResponse := CreateFinancePaymentCrossboarderpayMessageGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FinancePaymentCrossboarderpayMessageGetRequest struct {
	*api.BaseRequest

	ChannelMerchantNo string `json:"channelMerchantNo,omitempty"`
	InteractId        string `json:"interactId,omitempty"`
}

type FinancePaymentCrossboarderpayMessageGetResponse struct {
	RequestMessage  string `json:"requestMessage,omitempty"`
	ResponseMessage string `json:"responseMessage,omitempty"`
	NoticeMessage   string `json:"noticeMessage,omitempty"`
}

func CreateFinancePaymentCrossboarderpayMessageGetRequest() (request *FinancePaymentCrossboarderpayMessageGetRequest) {
	request = &FinancePaymentCrossboarderpayMessageGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "finance/payment/crossboarderpay/message/get", "apigw")
	request.Method = api.POST
	return
}

func CreateFinancePaymentCrossboarderpayMessageGetResponse() (response *api.BaseResponse[FinancePaymentCrossboarderpayMessageGetResponse]) {
	response = api.CreateResponse[FinancePaymentCrossboarderpayMessageGetResponse](&FinancePaymentCrossboarderpayMessageGetResponse{})
	return
}
