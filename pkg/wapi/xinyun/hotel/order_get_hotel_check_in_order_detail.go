package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderGetHotelCheckInOrderDetail
// @id 515
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=酒店住宿单详情)
func (client *Client) OrderGetHotelCheckInOrderDetail(request *OrderGetCheckInOrderDetailRequest) (response *OrderGetCheckInOrderDetailResponse, err error) {
	rpcResponse := CreateOrderGetCheckInOrderDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderGetCheckInOrderDetailRequest struct {
	*api.BaseRequest

	StoreId   int64  `json:"storeId,omitempty"`
	BookingNo string `json:"bookingNo,omitempty"`
}

type OrderGetCheckInOrderDetailResponse struct {
}

func CreateOrderGetCheckInOrderDetailRequest() (request *OrderGetCheckInOrderDetailRequest) {
	request = &OrderGetCheckInOrderDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "order/getHotelCheckInOrderDetail", "api")
	request.Method = api.POST
	return
}

func CreateOrderGetCheckInOrderDetailResponse() (response *api.BaseResponse[OrderGetCheckInOrderDetailResponse]) {
	response = api.CreateResponse[OrderGetCheckInOrderDetailResponse](&OrderGetCheckInOrderDetailResponse{})
	return
}
