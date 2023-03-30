package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderCheckOutRoom
// @id 519
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=办理退房)
func (client *Client) OrderCheckOutRoom(request *OrderCheckOutRoomRequest) (response *OrderCheckOutRoomResponse, err error) {
	rpcResponse := CreateOrderCheckOutRoomResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderCheckOutRoomRequest struct {
	*api.BaseRequest

	StoreId   int64    `json:"storeId,omitempty"`
	BookingNo string   `json:"bookingNo,omitempty"`
	RecordNos []string `json:"recordNos,omitempty"`
}

type OrderCheckOutRoomResponse struct {
}

func CreateOrderCheckOutRoomRequest() (request *OrderCheckOutRoomRequest) {
	request = &OrderCheckOutRoomRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "order/checkOutRoom", "api")
	request.Method = api.POST
	return
}

func CreateOrderCheckOutRoomResponse() (response *api.BaseResponse[OrderCheckOutRoomResponse]) {
	response = api.CreateResponse[OrderCheckOutRoomResponse](&OrderCheckOutRoomResponse{})
	return
}
