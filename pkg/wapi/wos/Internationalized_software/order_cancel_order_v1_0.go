package Internationalized_software

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderCancelOrder
// @id 1195
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1195?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商家关闭订单)
func (client *Client) OrderCancelOrderV1_0(request *WeimobShopexpressOrderCancelOrderRequest) (response *WeimobShopexpressOrderCancelOrderResponse, err error) {
	rpcResponse := CreateWeimobShopexpressOrderCancelOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WeimobShopexpressOrderCancelOrderRequest struct {
	*api.BaseRequest

	Uid     int64  `json:"uid,omitempty"`
	OrderNo string `json:"orderNo,omitempty"`
	Aid     int64  `json:"aid,omitempty"`
}

type WeimobShopexpressOrderCancelOrderResponse struct {
}

func CreateWeimobShopexpressOrderCancelOrderRequest() (request *WeimobShopexpressOrderCancelOrderRequest) {
	request = &WeimobShopexpressOrderCancelOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("Internationalized_software", "v1.0", "order/cancelOrder", "apigw")
	request.Method = api.POST
	return
}

func CreateWeimobShopexpressOrderCancelOrderResponse() (response *api.BaseResponse[WeimobShopexpressOrderCancelOrderResponse]) {
	response = api.CreateResponse[WeimobShopexpressOrderCancelOrderResponse](&WeimobShopexpressOrderCancelOrderResponse{})
	return
}
