package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderCheckInRoom
// @id 518
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=办理入住)
func (client *Client) OrderCheckInRoom(request *OrderCheckInRoomRequest) (response *OrderCheckInRoomResponse, err error) {
	rpcResponse := CreateOrderCheckInRoomResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderCheckInRoomRequest struct {
	*api.BaseRequest

	CheckInDetails []OrderCheckInRoomRequestCheckInDetails `json:"checkInDetails,omitempty"`
	StoreId        int64                                   `json:"storeId,omitempty"`
	BookingNo      string                                  `json:"bookingNo,omitempty"`
}

type OrderCheckInRoomResponse struct {
}

func CreateOrderCheckInRoomRequest() (request *OrderCheckInRoomRequest) {
	request = &OrderCheckInRoomRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "order/checkInRoom", "api")
	request.Method = api.POST
	return
}

func CreateOrderCheckInRoomResponse() (response *api.BaseResponse[OrderCheckInRoomResponse]) {
	response = api.CreateResponse[OrderCheckInRoomResponse](&OrderCheckInRoomResponse{})
	return
}

type OrderCheckInRoomRequestCheckInDetails struct {
	TogetherDetails   []OrderCheckInRoomRequestTogetherDetails `json:"togetherDetails,omitempty"`
	RecordNo          string                                   `json:"recordNo,omitempty"`
	RoomId            int64                                    `json:"roomId,omitempty"`
	RoomCode          string                                   `json:"roomCode,omitempty"`
	GuesetName        string                                   `json:"guesetName,omitempty"`
	GuestPhone        string                                   `json:"guestPhone,omitempty"`
	GuestDocumentType int                                      `json:"guestDocumentType,omitempty"`
	GuestDocumentNo   string                                   `json:"guestDocumentNo,omitempty"`
	Mark              string                                   `json:"mark,omitempty"`
}

type OrderCheckInRoomRequestTogetherDetails struct {
	GuestName         string `json:"guestName,omitempty"`
	GuestPhone        string `json:"guestPhone,omitempty"`
	GuestDocumentType int    `json:"guestDocumentType,omitempty"`
	GuestDocumentNo   string `json:"guestDocumentNo,omitempty"`
}
