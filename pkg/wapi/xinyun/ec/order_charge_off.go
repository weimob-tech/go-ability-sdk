package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderChargeOff
// @id 986
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=自提订单核销)
func (client *Client) OrderChargeOff(request *OrderChargeOffRequest) (response *OrderChargeOffResponse, err error) {
	rpcResponse := CreateOrderChargeOffResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderChargeOffRequest struct {
	*api.BaseRequest

	StoreId        int64  `json:"storeId,omitempty"`
	SiteId         int64  `json:"siteId,omitempty"`
	OrderNo        int64  `json:"orderNo,omitempty"`
	SelfPickupCode string `json:"selfPickupCode,omitempty"`
	SceneType      int    `json:"sceneType,omitempty"`
}

type OrderChargeOffResponse struct {
}

func CreateOrderChargeOffRequest() (request *OrderChargeOffRequest) {
	request = &OrderChargeOffRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "order/chargeOff", "api")
	request.Method = api.POST
	return
}

func CreateOrderChargeOffResponse() (response *api.BaseResponse[OrderChargeOffResponse]) {
	response = api.CreateResponse[OrderChargeOffResponse](&OrderChargeOffResponse{})
	return
}
