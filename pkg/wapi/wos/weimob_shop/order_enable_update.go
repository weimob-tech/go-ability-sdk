package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderEnableUpdate
// @id 1813
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1813?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=订单审核)
func (client *Client) OrderEnableUpdate(request *OrderEnableUpdateRequest) (response *OrderEnableUpdateResponse, err error) {
	rpcResponse := CreateOrderEnableUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderEnableUpdateRequest struct {
	*api.BaseRequest

	OrderNo int64 `json:"orderNo,omitempty"`
}

type OrderEnableUpdateResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateOrderEnableUpdateRequest() (request *OrderEnableUpdateRequest) {
	request = &OrderEnableUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "order/enable/update", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderEnableUpdateResponse() (response *api.BaseResponse[OrderEnableUpdateResponse]) {
	response = api.CreateResponse[OrderEnableUpdateResponse](&OrderEnableUpdateResponse{})
	return
}
