package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderQueryOrderList
// @id 69
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询订单列表)
func (client *Client) OrderQueryOrderList(request *OrderQueryOrderListRequest) (response *OrderQueryOrderListResponse, err error) {
	rpcResponse := CreateOrderQueryOrderListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderQueryOrderListRequest struct {
	*api.BaseRequest

	StoreId     int64 `json:"storeId,omitempty"`
	OrderStatus int   `json:"orderStatus,omitempty"`
	Page        int   `json:"page,omitempty"`
	PageSize    int   `json:"pageSize,omitempty"`
	StartTime   int64 `json:"startTime,omitempty"`
	EndTime     int64 `json:"endTime,omitempty"`
	PayStatus   int   `json:"payStatus,omitempty"`
	PageNum     int   `json:"pageNum,omitempty"`
}

type OrderQueryOrderListResponse struct {
}

func CreateOrderQueryOrderListRequest() (request *OrderQueryOrderListRequest) {
	request = &OrderQueryOrderListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "order/queryOrderList", "api")
	request.Method = api.POST
	return
}

func CreateOrderQueryOrderListResponse() (response *api.BaseResponse[OrderQueryOrderListResponse]) {
	response = api.CreateResponse[OrderQueryOrderListResponse](&OrderQueryOrderListResponse{})
	return
}
