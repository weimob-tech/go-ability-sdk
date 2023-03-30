package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StockAddStockCheckOrder
// @id 561
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增盘点单)
func (client *Client) StockAddStockCheckOrder(request *StockAddStockCheckOrderRequest) (response *StockAddStockCheckOrderResponse, err error) {
	rpcResponse := CreateStockAddStockCheckOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StockAddStockCheckOrderRequest struct {
	*api.BaseRequest

	Id            int64   `json:"id,omitempty"`
	Status        int     `json:"status,omitempty"`
	Remark        string  `json:"remark,omitempty"`
	OrderItemList []int64 `json:"orderItemList,omitempty"`
	StoreId       string  `json:"storeId,omitempty"`
}

type StockAddStockCheckOrderResponse struct {
}

func CreateStockAddStockCheckOrderRequest() (request *StockAddStockCheckOrderRequest) {
	request = &StockAddStockCheckOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "stock/addStockCheckOrder", "api")
	request.Method = api.POST
	return
}

func CreateStockAddStockCheckOrderResponse() (response *api.BaseResponse[StockAddStockCheckOrderResponse]) {
	response = api.CreateResponse[StockAddStockCheckOrderResponse](&StockAddStockCheckOrderResponse{})
	return
}
