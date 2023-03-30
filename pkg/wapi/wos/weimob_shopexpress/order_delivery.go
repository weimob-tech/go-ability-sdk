package weimob_shopexpress

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderDelivery
// @id 1967
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1967?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=订单发货)
func (client *Client) OrderDelivery(request *OrderDeliveryRequest) (response *OrderDeliveryResponse, err error) {
	rpcResponse := CreateOrderDeliveryResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderDeliveryRequest struct {
	*api.BaseRequest

	GoodsList          []OrderDeliveryRequestGoodsList `json:"goodsList,omitempty"`
	ExpressCompany     string                          `json:"expressCompany,omitempty"`
	ExpressCompanyCode string                          `json:"expressCompanyCode,omitempty"`
	ExpressNo          string                          `json:"expressNo,omitempty"`
	ExpressUrl         string                          `json:"expressUrl,omitempty"`
	OperatorName       string                          `json:"operatorName,omitempty"`
	OrderNo            string                          `json:"orderNo,omitempty"`
}

type OrderDeliveryResponse struct {
}

func CreateOrderDeliveryRequest() (request *OrderDeliveryRequest) {
	request = &OrderDeliveryRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shopexpress", "v2.0", "order/delivery", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderDeliveryResponse() (response *api.BaseResponse[OrderDeliveryResponse]) {
	response = api.CreateResponse[OrderDeliveryResponse](&OrderDeliveryResponse{})
	return
}

type OrderDeliveryRequestGoodsList struct {
	Count       int64  `json:"count,omitempty"`
	OrderItemNo string `json:"orderItemNo,omitempty"`
}
