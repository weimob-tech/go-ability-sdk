package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderQuerySimpleOrderListByOrderNos
// @id 1073
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1073?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=产品融合批量查询订单)
func (client *Client) OrderQuerySimpleOrderListByOrderNos(request *OrderQuerySimpleOrderListByOrderNosRequest) (response *OrderQuerySimpleOrderListByOrderNosResponse, err error) {
	rpcResponse := CreateOrderQuerySimpleOrderListByOrderNosResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderQuerySimpleOrderListByOrderNosRequest struct {
	*api.BaseRequest

	Obj string `json:"obj,omitempty"`
}

type OrderQuerySimpleOrderListByOrderNosResponse struct {
	Obj string `json:"obj,omitempty"`
}

func CreateOrderQuerySimpleOrderListByOrderNosRequest() (request *OrderQuerySimpleOrderListByOrderNosRequest) {
	request = &OrderQuerySimpleOrderListByOrderNosRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "order/querySimpleOrderListByOrderNos", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderQuerySimpleOrderListByOrderNosResponse() (response *api.BaseResponse[OrderQuerySimpleOrderListByOrderNosResponse]) {
	response = api.CreateResponse[OrderQuerySimpleOrderListByOrderNosResponse](&OrderQuerySimpleOrderListByOrderNosResponse{})
	return
}
