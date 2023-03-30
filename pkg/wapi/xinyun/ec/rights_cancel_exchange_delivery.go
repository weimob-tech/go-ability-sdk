package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsCancelExchangeDelivery
// @id 1619
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=换货取消发货)
func (client *Client) RightsCancelExchangeDelivery(request *RightsCancelExchangeDeliveryRequest) (response *RightsCancelExchangeDeliveryResponse, err error) {
	rpcResponse := CreateRightsCancelExchangeDeliveryResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsCancelExchangeDeliveryRequest struct {
	*api.BaseRequest

	ExchangeNo     int    `json:"exchangeNo,omitempty"`
	OperationId    string `json:"operationId,omitempty"`
	OpreationPhone string `json:"opreationPhone,omitempty"`
	OperationName  string `json:"operationName,omitempty"`
}

type RightsCancelExchangeDeliveryResponse struct {
}

func CreateRightsCancelExchangeDeliveryRequest() (request *RightsCancelExchangeDeliveryRequest) {
	request = &RightsCancelExchangeDeliveryRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "rights/cancelExchangeDelivery", "api")
	request.Method = api.POST
	return
}

func CreateRightsCancelExchangeDeliveryResponse() (response *api.BaseResponse[RightsCancelExchangeDeliveryResponse]) {
	response = api.CreateResponse[RightsCancelExchangeDeliveryResponse](&RightsCancelExchangeDeliveryResponse{})
	return
}
