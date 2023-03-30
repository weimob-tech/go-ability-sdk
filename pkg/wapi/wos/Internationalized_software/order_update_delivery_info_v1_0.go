package Internationalized_software

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderUpdateDeliveryInfo
// @id 1196
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1196?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新订单物流信息)
func (client *Client) OrderUpdateDeliveryInfoV1_0(request *WeimobShopexpressOrderUpdateDeliveryInfoRequest) (response *WeimobShopexpressOrderUpdateDeliveryInfoResponse, err error) {
	rpcResponse := CreateWeimobShopexpressOrderUpdateDeliveryInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WeimobShopexpressOrderUpdateDeliveryInfoRequest struct {
	*api.BaseRequest

	Uid                int64  `json:"uid,omitempty"`
	OrderNo            string `json:"orderNo,omitempty"`
	ExpressNo          string `json:"expressNo,omitempty"`
	ExpressCompanyCode string `json:"expressCompanyCode,omitempty"`
	ExpressCompany     string `json:"expressCompany,omitempty"`
	ExpressUrl         string `json:"expressUrl,omitempty"`
	Aid                int64  `json:"aid,omitempty"`
}

type WeimobShopexpressOrderUpdateDeliveryInfoResponse struct {
}

func CreateWeimobShopexpressOrderUpdateDeliveryInfoRequest() (request *WeimobShopexpressOrderUpdateDeliveryInfoRequest) {
	request = &WeimobShopexpressOrderUpdateDeliveryInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("Internationalized_software", "v1.0", "order/updateDeliveryInfo", "apigw")
	request.Method = api.POST
	return
}

func CreateWeimobShopexpressOrderUpdateDeliveryInfoResponse() (response *api.BaseResponse[WeimobShopexpressOrderUpdateDeliveryInfoResponse]) {
	response = api.CreateResponse[WeimobShopexpressOrderUpdateDeliveryInfoResponse](&WeimobShopexpressOrderUpdateDeliveryInfoResponse{})
	return
}
