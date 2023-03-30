package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StockCreateInventoryOrder
// @id 549
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=创建采购、其他入库单或者创建出库单)
func (client *Client) StockCreateInventoryOrder(request *StockCreateInventoryOrderRequest) (response *StockCreateInventoryOrderResponse, err error) {
	rpcResponse := CreateStockCreateInventoryOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StockCreateInventoryOrderRequest struct {
	*api.BaseRequest

	OrderItemList []StockCreateInventoryOrderRequestOrderItemList `json:"orderItemList,omitempty"`
	StoreId       int64                                           `json:"storeId,omitempty"`
	InventoryType int                                             `json:"inventoryType,omitempty"`
	ReferType     int                                             `json:"referType,omitempty"`
	ReferId       int64                                           `json:"referId,omitempty"`
	SupplierId    int64                                           `json:"supplierId,omitempty"`
	Remark        string                                          `json:"remark,omitempty"`
}

type StockCreateInventoryOrderResponse struct {
}

func CreateStockCreateInventoryOrderRequest() (request *StockCreateInventoryOrderRequest) {
	request = &StockCreateInventoryOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "stock/createInventoryOrder", "api")
	request.Method = api.POST
	return
}

func CreateStockCreateInventoryOrderResponse() (response *api.BaseResponse[StockCreateInventoryOrderResponse]) {
	response = api.CreateResponse[StockCreateInventoryOrderResponse](&StockCreateInventoryOrderResponse{})
	return
}

type StockCreateInventoryOrderRequestOrderItemList struct {
	SingleProductId  int64 `json:"singleProductId,omitempty"`
	StockChangeNum   int   `json:"stockChangeNum,omitempty"`
	OriginalStockNum int   `json:"originalStockNum,omitempty"`
}
