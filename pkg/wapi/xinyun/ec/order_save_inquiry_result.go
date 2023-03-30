package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderSaveInquiryResult
// @id 1878
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=问诊单结果回调)
func (client *Client) OrderSaveInquiryResult(request *OrderSaveInquiryResultRequest) (response *OrderSaveInquiryResultResponse, err error) {
	rpcResponse := CreateOrderSaveInquiryResultResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderSaveInquiryResultRequest struct {
	*api.BaseRequest
}

type OrderSaveInquiryResultResponse struct {
}

func CreateOrderSaveInquiryResultRequest() (request *OrderSaveInquiryResultRequest) {
	request = &OrderSaveInquiryResultRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "order/saveInquiryResult", "api")
	request.Method = api.POST
	return
}

func CreateOrderSaveInquiryResultResponse() (response *api.BaseResponse[OrderSaveInquiryResultResponse]) {
	response = api.CreateResponse[OrderSaveInquiryResultResponse](&OrderSaveInquiryResultResponse{})
	return
}
