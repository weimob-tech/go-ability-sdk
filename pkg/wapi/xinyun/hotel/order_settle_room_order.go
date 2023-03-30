package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderSettleRoomOrder
// @id 523
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=办理结算)
func (client *Client) OrderSettleRoomOrder(request *OrderSettleRoomOrderRequest) (response *OrderSettleRoomOrderResponse, err error) {
	rpcResponse := CreateOrderSettleRoomOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderSettleRoomOrderRequest struct {
	*api.BaseRequest

	StoreId              int64   `json:"storeId,omitempty"`
	OrderNo              string  `json:"orderNo,omitempty"`
	BillItemCode         int     `json:"billItemCode,omitempty"`
	UsePledge            int     `json:"usePledge,omitempty"`
	PledgeAmount         float64 `json:"pledgeAmount,omitempty"`
	UseMemberDiscount    int     `json:"useMemberDiscount,omitempty"`
	MemberDiscountAmount float64 `json:"memberDiscountAmount,omitempty"`
	RemarkForPledge      string  `json:"remarkForPledge,omitempty"`
	Amount               float64 `json:"amount,omitempty"`
}

type OrderSettleRoomOrderResponse struct {
}

func CreateOrderSettleRoomOrderRequest() (request *OrderSettleRoomOrderRequest) {
	request = &OrderSettleRoomOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "order/settleRoomOrder", "api")
	request.Method = api.POST
	return
}

func CreateOrderSettleRoomOrderResponse() (response *api.BaseResponse[OrderSettleRoomOrderResponse]) {
	response = api.CreateResponse[OrderSettleRoomOrderResponse](&OrderSettleRoomOrderResponse{})
	return
}
