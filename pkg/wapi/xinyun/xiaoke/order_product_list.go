package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderProductList
// @id 2652
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询订单下产品列表)
func (client *Client) OrderProductList(request *OrderProductListRequest) (response *OrderProductListResponse, err error) {
	rpcResponse := CreateOrderProductListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderProductListRequest struct {
	*api.BaseRequest

	UserWid  int64  `json:"userWid,omitempty"`
	Key      string `json:"key,omitempty"`
	PageNum  int    `json:"pageNum,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
}

type OrderProductListResponse struct {
}

func CreateOrderProductListRequest() (request *OrderProductListRequest) {
	request = &OrderProductListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "order/productList", "api")
	request.Method = api.POST
	return
}

func CreateOrderProductListResponse() (response *api.BaseResponse[OrderProductListResponse]) {
	response = api.CreateResponse[OrderProductListResponse](&OrderProductListResponse{})
	return
}
