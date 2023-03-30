package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsDeliveryExchange
// @id 1618
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=换货发货)
func (client *Client) RightsDeliveryExchange(request *RightsDeliveryExchangeRequest) (response *RightsDeliveryExchangeResponse, err error) {
	rpcResponse := CreateRightsDeliveryExchangeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsDeliveryExchangeRequest struct {
	*api.BaseRequest

	DeliveryExchangeVo []RightsDeliveryExchangeRequestDeliveryExchangeVo `json:"deliveryExchangeVo,omitempty"`
	ExchangeNo         int64                                             `json:"exchangeNo,omitempty"`
	OperationId        string                                            `json:"operationId,omitempty"`
	OperationName      string                                            `json:"operationName,omitempty"`
	OperationPhone     string                                            `json:"operationPhone,omitempty"`
	DeliveryMethod     string                                            `json:"deliveryMethod,omitempty"`
}

type RightsDeliveryExchangeResponse struct {
	Success    bool   `json:"success,omitempty"`
	FailReason string `json:"failReason,omitempty"`
}

func CreateRightsDeliveryExchangeRequest() (request *RightsDeliveryExchangeRequest) {
	request = &RightsDeliveryExchangeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "rights/deliveryExchange", "api")
	request.Method = api.POST
	return
}

func CreateRightsDeliveryExchangeResponse() (response *api.BaseResponse[RightsDeliveryExchangeResponse]) {
	response = api.CreateResponse[RightsDeliveryExchangeResponse](&RightsDeliveryExchangeResponse{})
	return
}

type RightsDeliveryExchangeRequestDeliveryExchangeVo struct {
	DeliveryNo          string `json:"deliveryNo,omitempty"`
	DeliveryCompanyCode string `json:"deliveryCompanyCode,omitempty"`
	DeliveryCompanyName string `json:"deliveryCompanyName,omitempty"`
}
