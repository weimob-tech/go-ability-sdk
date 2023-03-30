package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DishesGoodsStoreStockBatchAdjust
// @id 2096
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量调整门店库存)
func (client *Client) DishesGoodsStoreStockBatchAdjust(request *DishesGoodsStoreStockBatchAdjustRequest) (response *DishesGoodsStoreStockBatchAdjustResponse, err error) {
	rpcResponse := CreateDishesGoodsStoreStockBatchAdjustResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DishesGoodsStoreStockBatchAdjustRequest struct {
	*api.BaseRequest

	GoodsIds    []map[string]any `json:"goodsIds,omitempty"`
	BrandId     int64            `json:"brandId,omitempty"`
	StoreId     int64            `json:"storeId,omitempty"`
	MerchantWid int64            `json:"merchantWid,omitempty"`
	Stock       int              `json:"stock,omitempty"`
}

type DishesGoodsStoreStockBatchAdjustResponse struct {
}

func CreateDishesGoodsStoreStockBatchAdjustRequest() (request *DishesGoodsStoreStockBatchAdjustRequest) {
	request = &DishesGoodsStoreStockBatchAdjustRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "dishesGoods/storeStockBatchAdjust", "api")
	request.Method = api.POST
	return
}

func CreateDishesGoodsStoreStockBatchAdjustResponse() (response *api.BaseResponse[DishesGoodsStoreStockBatchAdjustResponse]) {
	response = api.CreateResponse[DishesGoodsStoreStockBatchAdjustResponse](&DishesGoodsStoreStockBatchAdjustResponse{})
	return
}
