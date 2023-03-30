package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderQuerySimpleOrderList
// @id 1066
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1066?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=产品融合订单列表)
func (client *Client) OrderQuerySimpleOrderList(request *OrderQuerySimpleOrderListRequest) (response *OrderQuerySimpleOrderListResponse, err error) {
	rpcResponse := CreateOrderQuerySimpleOrderListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderQuerySimpleOrderListRequest struct {
	*api.BaseRequest

	Obj string `json:"obj,omitempty"`
}

type OrderQuerySimpleOrderListResponse struct {
	Obj string `json:"obj,omitempty"`
}

func CreateOrderQuerySimpleOrderListRequest() (request *OrderQuerySimpleOrderListRequest) {
	request = &OrderQuerySimpleOrderListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "order/querySimpleOrderList", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderQuerySimpleOrderListResponse() (response *api.BaseResponse[OrderQuerySimpleOrderListResponse]) {
	response = api.CreateResponse[OrderQuerySimpleOrderListResponse](&OrderQuerySimpleOrderListResponse{})
	return
}
