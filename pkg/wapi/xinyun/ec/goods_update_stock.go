package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsUpdateStock
// @id 1291
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=多商品批量库存更新)
func (client *Client) GoodsUpdateStock(request *GoodsUpdateStockRequest) (response *GoodsUpdateStockResponse, err error) {
	rpcResponse := CreateGoodsUpdateStockResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsUpdateStockRequest struct {
	*api.BaseRequest

	UpdateGoodsStockList []GoodsUpdateStockRequestUpdateGoodsStockList `json:"updateGoodsStockList,omitempty"`
	WarehouseId          int64                                         `json:"warehouseId,omitempty"`
	WarehouseCode        string                                        `json:"warehouseCode,omitempty"`
	StoreId              int64                                         `json:"storeId,omitempty"`
}

type GoodsUpdateStockResponse struct {
}

func CreateGoodsUpdateStockRequest() (request *GoodsUpdateStockRequest) {
	request = &GoodsUpdateStockRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goods/updateStock", "api")
	request.Method = api.POST
	return
}

func CreateGoodsUpdateStockResponse() (response *api.BaseResponse[GoodsUpdateStockResponse]) {
	response = api.CreateResponse[GoodsUpdateStockResponse](&GoodsUpdateStockResponse{})
	return
}

type GoodsUpdateStockRequestUpdateGoodsStockList struct {
	SkuId    int64 `json:"skuId,omitempty"`
	StockNum int   `json:"stockNum,omitempty"`
}
