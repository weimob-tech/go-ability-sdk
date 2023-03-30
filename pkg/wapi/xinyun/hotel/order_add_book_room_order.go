package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderAddBookRoomOrder
// @id 516
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=办理预定)
func (client *Client) OrderAddBookRoomOrder(request *OrderAddBookRoomOrderRequest) (response *OrderAddBookRoomOrderResponse, err error) {
	rpcResponse := CreateOrderAddBookRoomOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderAddBookRoomOrderRequest struct {
	*api.BaseRequest

	BookingDetailInfos []OrderAddBookRoomOrderRequestBookingDetailInfos `json:"bookingDetailInfos,omitempty"`
	StoreId            int64                                            `json:"storeId,omitempty"`
	BookingName        string                                           `json:"bookingName,omitempty"`
	BookingPhone       string                                           `json:"bookingPhone,omitempty"`
	Mark               string                                           `json:"mark,omitempty"`
}

type OrderAddBookRoomOrderResponse struct {
}

func CreateOrderAddBookRoomOrderRequest() (request *OrderAddBookRoomOrderRequest) {
	request = &OrderAddBookRoomOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "order/addBookRoomOrder", "api")
	request.Method = api.POST
	return
}

func CreateOrderAddBookRoomOrderResponse() (response *api.BaseResponse[OrderAddBookRoomOrderResponse]) {
	response = api.CreateResponse[OrderAddBookRoomOrderResponse](&OrderAddBookRoomOrderResponse{})
	return
}

type OrderAddBookRoomOrderRequestBookingDetailInfos struct {
	RoomNos           []OrderAddBookRoomOrderRequestRoomNos `json:"roomNos,omitempty"`
	CheckInTime       int                                   `json:"checkInTime,omitempty"`
	ChekOutTime       int                                   `json:"chekOutTime,omitempty"`
	RoomTypeId        int64                                 `json:"roomTypeId,omitempty"`
	RoomTypeName      string                                `json:"roomTypeName,omitempty"`
	BookingNum        int                                   `json:"bookingNum,omitempty"`
	MealTicket        int                                   `json:"mealTicket,omitempty"`
	PriceDayList      []float64                             `json:"priceDayList,omitempty"`
	Days              int                                   `json:"days,omitempty"`
	GuestName         string                                `json:"guestName,omitempty"`
	GuestPhone        string                                `json:"guestPhone,omitempty"`
	GuestDocumentType int                                   `json:"guestDocumentType,omitempty"`
	GuestDocumentNo   string                                `json:"guestDocumentNo,omitempty"`
	Mark              string                                `json:"mark,omitempty"`
}

type OrderAddBookRoomOrderRequestRoomNos struct {
	RoomId   int64  `json:"roomId,omitempty"`
	RoomCode string `json:"roomCode,omitempty"`
}
