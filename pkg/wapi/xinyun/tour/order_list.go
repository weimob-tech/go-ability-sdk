package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderList
// @id 1001
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询线路订单列表)
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

	OrderStatus int    `json:"orderStatus,omitempty"`
	OrderSource int    `json:"orderSource,omitempty"`
	TimeType    int    `json:"timeType,omitempty"`
	StartTime   int    `json:"startTime,omitempty"`
	EndTime     int    `json:"endTime,omitempty"`
	Keyword     int    `json:"keyword,omitempty"`
	Value       string `json:"value,omitempty"`
	PageIndex   int    `json:"pageIndex,omitempty"`
	PageSize    int    `json:"pageSize,omitempty"`
}

type OrderListResponse struct {
}

func CreateOrderListRequest() (request *OrderListRequest) {
	request = &OrderListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "order/list", "api")
	request.Method = api.POST
	return
}

func CreateOrderListResponse() (response *api.BaseResponse[OrderListResponse]) {
	response = api.CreateResponse[OrderListResponse](&OrderListResponse{})
	return
}
