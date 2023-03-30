package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderGetBillItems
// @id 1596
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取支付方式)
func (client *Client) OrderGetBillItems(request *OrderGetBillItemsRequest) (response *OrderGetBillItemsResponse, err error) {
	rpcResponse := CreateOrderGetBillItemsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderGetBillItemsRequest struct {
	*api.BaseRequest

	OrderNo string `json:"orderNo,omitempty"`
	StoreId int64  `json:"storeId,omitempty"`
}

type OrderGetBillItemsResponse struct {
}

func CreateOrderGetBillItemsRequest() (request *OrderGetBillItemsRequest) {
	request = &OrderGetBillItemsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "order/getBillItems", "api")
	request.Method = api.POST
	return
}

func CreateOrderGetBillItemsResponse() (response *api.BaseResponse[OrderGetBillItemsResponse]) {
	response = api.CreateResponse[OrderGetBillItemsResponse](&OrderGetBillItemsResponse{})
	return
}
