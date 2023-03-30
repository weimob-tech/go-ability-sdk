package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsWholeUpdateStock
// @id 377
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=全量更新商品库存)
func (client *Client) GoodsWholeUpdateStock(request *GoodsWholeUpdateStockRequest) (response *GoodsWholeUpdateStockResponse, err error) {
	rpcResponse := CreateGoodsWholeUpdateStockResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsWholeUpdateStockRequest struct {
	*api.BaseRequest

	SkuList       []GoodsWholeUpdateStockRequestSkuList `json:"skuList,omitempty"`
	GoodsId       int64                                 `json:"goodsId,omitempty"`
	StoreId       int64                                 `json:"storeId,omitempty"`
	WarehouseId   int64                                 `json:"warehouseId,omitempty"`
	WarehouseCode string                                `json:"warehouseCode,omitempty"`
}

type GoodsWholeUpdateStockResponse struct {
}

func CreateGoodsWholeUpdateStockRequest() (request *GoodsWholeUpdateStockRequest) {
	request = &GoodsWholeUpdateStockRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goods/wholeUpdateStock", "api")
	request.Method = api.POST
	return
}

func CreateGoodsWholeUpdateStockResponse() (response *api.BaseResponse[GoodsWholeUpdateStockResponse]) {
	response = api.CreateResponse[GoodsWholeUpdateStockResponse](&GoodsWholeUpdateStockResponse{})
	return
}

type GoodsWholeUpdateStockRequestSkuList struct {
	SkuId        int64 `json:"skuId,omitempty"`
	EditStockNum int   `json:"editStockNum,omitempty"`
}
