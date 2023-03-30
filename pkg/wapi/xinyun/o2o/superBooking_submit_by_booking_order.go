package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SuperBookingSubmitByBookingOrder
// @id 1839
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=代预订订单下单)
func (client *Client) SuperBookingSubmitByBookingOrder(request *SuperBookingSubmitByBookingOrderRequest) (response *SuperBookingSubmitByBookingOrderResponse, err error) {
	rpcResponse := CreateSuperBookingSubmitByBookingOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SuperBookingSubmitByBookingOrderRequest struct {
	*api.BaseRequest

	BookingWid    int64  `json:"bookingWid,omitempty"`
	BookingName   string `json:"bookingName,omitempty"`
	BookingPhone  int64  `json:"bookingPhone,omitempty"`
	BookingSex    int    `json:"bookingSex,omitempty"`
	DinningRegion string `json:"dinningRegion,omitempty"`
	ThirdOrderNo  string `json:"thirdOrderNo,omitempty"`
	StoreId       string `json:"storeId,omitempty"`
	BookingTime   int64  `json:"bookingTime,omitempty"`
	DinningNum    int    `json:"dinningNum,omitempty"`
	DinningName   string `json:"dinningName,omitempty"`
	DinningPhone  string `json:"dinningPhone,omitempty"`
	DinningSex    int    `json:"dinningSex,omitempty"`
	Remark        string `json:"remark,omitempty"`
}

type SuperBookingSubmitByBookingOrderResponse struct {
}

func CreateSuperBookingSubmitByBookingOrderRequest() (request *SuperBookingSubmitByBookingOrderRequest) {
	request = &SuperBookingSubmitByBookingOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "superBooking/submitByBookingOrder", "api")
	request.Method = api.POST
	return
}

func CreateSuperBookingSubmitByBookingOrderResponse() (response *api.BaseResponse[SuperBookingSubmitByBookingOrderResponse]) {
	response = api.CreateResponse[SuperBookingSubmitByBookingOrderResponse](&SuperBookingSubmitByBookingOrderResponse{})
	return
}
