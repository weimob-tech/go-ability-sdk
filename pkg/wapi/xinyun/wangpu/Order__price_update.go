package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderPriceUpdate
// @id 32
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改订单金额)
func (client *Client) OrderPriceUpdate(request *OrderPriceUpdateRequest) (response *OrderPriceUpdateResponse, err error) {
	rpcResponse := CreateOrderPriceUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderPriceUpdateRequest struct {
	*api.BaseRequest

	OrderDetails      []map[string]any `json:"order_details,omitempty"`
	OrderNo           string           `json:"order_no,omitempty"`
	ShopId            string           `json:"shop_id,omitempty"`
	UpdateMan         string           `json:"update_man,omitempty"`
	NewDeliveryAmount float64          `json:"new_delivery_amount,omitempty"`
	OrderDetailId     int64            `json:"order_detail_id,omitempty"`
	OldPrice          float64          `json:"old_price,omitempty"`
	OldQty            int              `json:"old_qty,omitempty"`
	NewPrice          float64          `json:"new_price,omitempty"`
	NewQty            int              `json:"new_qty,omitempty"`
}

type OrderPriceUpdateResponse struct {
}

func CreateOrderPriceUpdateRequest() (request *OrderPriceUpdateRequest) {
	request = &OrderPriceUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "Order/PriceUpdate", "api")
	request.Method = api.POST
	return
}

func CreateOrderPriceUpdateResponse() (response *api.BaseResponse[OrderPriceUpdateResponse]) {
	response = api.CreateResponse[OrderPriceUpdateResponse](&OrderPriceUpdateResponse{})
	return
}
