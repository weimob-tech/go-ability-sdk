package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TradeMerchantUpdateOrder
// @id 1285
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=操作订单)
func (client *Client) TradeMerchantUpdateOrder(request *TradeMerchantUpdateOrderRequest) (response *TradeMerchantUpdateOrderResponse, err error) {
	rpcResponse := CreateTradeMerchantUpdateOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TradeMerchantUpdateOrderRequest struct {
	*api.BaseRequest

	OptType int    `json:"optType,omitempty"`
	OrderNo string `json:"orderNo,omitempty"`
}

type TradeMerchantUpdateOrderResponse struct {
}

func CreateTradeMerchantUpdateOrderRequest() (request *TradeMerchantUpdateOrderRequest) {
	request = &TradeMerchantUpdateOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "tradeMerchant/updateOrder", "api")
	request.Method = api.POST
	return
}

func CreateTradeMerchantUpdateOrderResponse() (response *api.BaseResponse[TradeMerchantUpdateOrderResponse]) {
	response = api.CreateResponse[TradeMerchantUpdateOrderResponse](&TradeMerchantUpdateOrderResponse{})
	return
}
