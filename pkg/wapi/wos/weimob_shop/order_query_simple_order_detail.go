package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderQuerySimpleOrderDetail
// @id 1070
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1070?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=产品融合订单详情)
func (client *Client) OrderQuerySimpleOrderDetail(request *OrderQuerySimpleOrderDetailRequest) (response *OrderQuerySimpleOrderDetailResponse, err error) {
	rpcResponse := CreateOrderQuerySimpleOrderDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderQuerySimpleOrderDetailRequest struct {
	*api.BaseRequest

	Obj string `json:"obj,omitempty"`
}

type OrderQuerySimpleOrderDetailResponse struct {
	Obj string `json:"obj,omitempty"`
}

func CreateOrderQuerySimpleOrderDetailRequest() (request *OrderQuerySimpleOrderDetailRequest) {
	request = &OrderQuerySimpleOrderDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "order/querySimpleOrderDetail", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderQuerySimpleOrderDetailResponse() (response *api.BaseResponse[OrderQuerySimpleOrderDetailResponse]) {
	response = api.CreateResponse[OrderQuerySimpleOrderDetailResponse](&OrderQuerySimpleOrderDetailResponse{})
	return
}
