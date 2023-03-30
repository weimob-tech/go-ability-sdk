package weimob_shopexpress

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderDeliveryUpdate
// @id 1963
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1963?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=订单发货信息修改)
func (client *Client) OrderDeliveryUpdate(request *OrderDeliveryUpdateRequest) (response *OrderDeliveryUpdateResponse, err error) {
	rpcResponse := CreateOrderDeliveryUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderDeliveryUpdateRequest struct {
	*api.BaseRequest

	DeliveryOrderCode  string `json:"deliveryOrderCode,omitempty"`
	ExpressCompany     string `json:"expressCompany,omitempty"`
	ExpressCompanyCode string `json:"expressCompanyCode,omitempty"`
	ExpressNo          string `json:"expressNo,omitempty"`
	ExpressUrl         string `json:"expressUrl,omitempty"`
	OperatorName       string `json:"operatorName,omitempty"`
	OrderNo            string `json:"orderNo,omitempty"`
}

type OrderDeliveryUpdateResponse struct {
}

func CreateOrderDeliveryUpdateRequest() (request *OrderDeliveryUpdateRequest) {
	request = &OrderDeliveryUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shopexpress", "v2.0", "order/delivery/update", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderDeliveryUpdateResponse() (response *api.BaseResponse[OrderDeliveryUpdateResponse]) {
	response = api.CreateResponse[OrderDeliveryUpdateResponse](&OrderDeliveryUpdateResponse{})
	return
}
