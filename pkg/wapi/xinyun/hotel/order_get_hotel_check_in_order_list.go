package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderGetHotelCheckInOrderList
// @id 514
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=酒店住宿单列表)
func (client *Client) OrderGetHotelCheckInOrderList(request *OrderGetCheckInOrderListRequest) (response *OrderGetCheckInOrderListResponse, err error) {
	rpcResponse := CreateOrderGetCheckInOrderListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderGetCheckInOrderListRequest struct {
	*api.BaseRequest

	StoreId       int64  `json:"storeId,omitempty"`
	PageIndex     int    `json:"pageIndex,omitempty"`
	PageSize      int    `json:"pageSize,omitempty"`
	SelectStoreId int64  `json:"selectStoreId,omitempty"`
	QuickStatus   int    `json:"quickStatus,omitempty"`
	QueryStatus   int    `json:"queryStatus,omitempty"`
	QueryTimeType int    `json:"queryTimeType,omitempty"`
	BeginTime     int    `json:"beginTime,omitempty"`
	EndTime       int    `json:"endTime,omitempty"`
	KeyWordType   int    `json:"keyWordType,omitempty"`
	QueryKeyWord  string `json:"queryKeyWord,omitempty"`
}

type OrderGetCheckInOrderListResponse struct {
}

func CreateOrderGetCheckInOrderListRequest() (request *OrderGetCheckInOrderListRequest) {
	request = &OrderGetCheckInOrderListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "order/getHotelCheckInOrderList", "api")
	request.Method = api.POST
	return
}

func CreateOrderGetCheckInOrderListResponse() (response *api.BaseResponse[OrderGetCheckInOrderListResponse]) {
	response = api.CreateResponse[OrderGetCheckInOrderListResponse](&OrderGetCheckInOrderListResponse{})
	return
}
