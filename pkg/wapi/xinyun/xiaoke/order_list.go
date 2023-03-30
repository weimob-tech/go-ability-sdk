package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderList
// @id 2676
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取订单列表)
func (client *Client) OrderList(request *OrderListRequest) (response *OrderListResponse, err error) {
	rpcResponse := CreateOrderListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderListRequest struct {
	*api.BaseRequest

	UserWid        int64 `json:"userWid,omitempty"`
	PageNum        int   `json:"pageNum,omitempty"`
	PageSize       int   `json:"pageSize,omitempty"`
	ResourcesScope int   `json:"resourcesScope,omitempty"`
	OrderByRule    int   `json:"orderByRule,omitempty"`
}

type OrderListResponse struct {
	TotalPage  int   `json:"totalPage,omitempty"`
	PageSize   int   `json:"pageSize,omitempty"`
	TotalCount int64 `json:"totalCount,omitempty"`
	PageNum    int64 `json:"pageNum,omitempty"`
}

func CreateOrderListRequest() (request *OrderListRequest) {
	request = &OrderListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "order/list", "api")
	request.Method = api.POST
	return
}

func CreateOrderListResponse() (response *api.BaseResponse[OrderListResponse]) {
	response = api.CreateResponse[OrderListResponse](&OrderListResponse{})
	return
}
