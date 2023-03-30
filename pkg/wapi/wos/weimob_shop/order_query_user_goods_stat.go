package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderQueryUserGoodsStat
// @id 1076
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1076?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=产品融合订单统计用户购买商品数量)
func (client *Client) OrderQueryUserGoodsStat(request *OrderQueryUserGoodsStatRequest) (response *OrderQueryUserGoodsStatResponse, err error) {
	rpcResponse := CreateOrderQueryUserGoodsStatResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderQueryUserGoodsStatRequest struct {
	*api.BaseRequest

	Obj string `json:"obj,omitempty"`
}

type OrderQueryUserGoodsStatResponse struct {
	Obj string `json:"obj,omitempty"`
}

func CreateOrderQueryUserGoodsStatRequest() (request *OrderQueryUserGoodsStatRequest) {
	request = &OrderQueryUserGoodsStatRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "order/queryUserGoodsStat", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderQueryUserGoodsStatResponse() (response *api.BaseResponse[OrderQueryUserGoodsStatResponse]) {
	response = api.CreateResponse[OrderQueryUserGoodsStatResponse](&OrderQueryUserGoodsStatResponse{})
	return
}
