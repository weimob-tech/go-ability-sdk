package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SuperBookingOptOrder
// @id 1678
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=操作预订订单)
func (client *Client) SuperBookingOptOrder(request *SuperBookingOptOrderRequest) (response *SuperBookingOptOrderResponse, err error) {
	rpcResponse := CreateSuperBookingOptOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SuperBookingOptOrderRequest struct {
	*api.BaseRequest

	StoreId int64  `json:"storeId,omitempty"`
	OrderNo string `json:"orderNo,omitempty"`
	Status  int    `json:"status,omitempty"`
}

type SuperBookingOptOrderResponse struct {
}

func CreateSuperBookingOptOrderRequest() (request *SuperBookingOptOrderRequest) {
	request = &SuperBookingOptOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "superBooking/optOrder", "api")
	request.Method = api.POST
	return
}

func CreateSuperBookingOptOrderResponse() (response *api.BaseResponse[SuperBookingOptOrderResponse]) {
	response = api.CreateResponse[SuperBookingOptOrderResponse](&SuperBookingOptOrderResponse{})
	return
}
