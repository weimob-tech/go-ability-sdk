package Internationalized_software

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderDeliveryOrder
// @id 1194
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1194?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=订单发货)
func (client *Client) OrderDeliveryOrderV1_0(request *WeimobShopexpressOrderDeliveryOrderRequest) (response *WeimobShopexpressOrderDeliveryOrderResponse, err error) {
	rpcResponse := CreateWeimobShopexpressOrderDeliveryOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WeimobShopexpressOrderDeliveryOrderRequest struct {
	*api.BaseRequest

	Uid                int64  `json:"uid,omitempty"`
	OrderNo            string `json:"orderNo,omitempty"`
	ExpressNo          string `json:"expressNo,omitempty"`
	ExpressCompanyCode string `json:"expressCompanyCode,omitempty"`
	ExpressCompany     string `json:"expressCompany,omitempty"`
	ExpressUrl         string `json:"expressUrl,omitempty"`
	Aid                int64  `json:"aid,omitempty"`
}

type WeimobShopexpressOrderDeliveryOrderResponse struct {
}

func CreateWeimobShopexpressOrderDeliveryOrderRequest() (request *WeimobShopexpressOrderDeliveryOrderRequest) {
	request = &WeimobShopexpressOrderDeliveryOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("Internationalized_software", "v1.0", "order/deliveryOrder", "apigw")
	request.Method = api.POST
	return
}

func CreateWeimobShopexpressOrderDeliveryOrderResponse() (response *api.BaseResponse[WeimobShopexpressOrderDeliveryOrderResponse]) {
	response = api.CreateResponse[WeimobShopexpressOrderDeliveryOrderResponse](&WeimobShopexpressOrderDeliveryOrderResponse{})
	return
}
