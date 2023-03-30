package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SuperBookingModifyOrder
// @id 1679
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改预订订单信息)
func (client *Client) SuperBookingModifyOrder(request *SuperBookingModifyOrderRequest) (response *SuperBookingModifyOrderResponse, err error) {
	rpcResponse := CreateSuperBookingModifyOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SuperBookingModifyOrderRequest struct {
	*api.BaseRequest

	StoreId       int64  `json:"storeId,omitempty"`
	OrderNo       string `json:"orderNo,omitempty"`
	DinningNum    int    `json:"dinningNum,omitempty"`
	BookingTime   int64  `json:"bookingTime,omitempty"`
	DinningRegion string `json:"dinningRegion,omitempty"`
}

type SuperBookingModifyOrderResponse struct {
}

func CreateSuperBookingModifyOrderRequest() (request *SuperBookingModifyOrderRequest) {
	request = &SuperBookingModifyOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "superBooking/modifyOrder", "api")
	request.Method = api.POST
	return
}

func CreateSuperBookingModifyOrderResponse() (response *api.BaseResponse[SuperBookingModifyOrderResponse]) {
	response = api.CreateResponse[SuperBookingModifyOrderResponse](&SuperBookingModifyOrderResponse{})
	return
}
