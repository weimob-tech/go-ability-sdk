package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StockUpdate
// @id 1649
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1649?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=覆盖/增量更新库存)
func (client *Client) StockUpdate(request *StockUpdateRequest) (response *StockUpdateResponse, err error) {
	rpcResponse := CreateStockUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StockUpdateRequest struct {
	*api.BaseRequest

	BasicInfo        StockUpdateRequestBasicInfo   `json:"basicInfo,omitempty"`
	GoodsList        []StockUpdateRequestGoodsList `json:"goodsList,omitempty"`
	WarehouseId      int64                         `json:"warehouseId,omitempty"`
	QuantityEditType int64                         `json:"quantityEditType,omitempty"`
	WarehouseCode    string                        `json:"warehouseCode,omitempty"`
}

type StockUpdateResponse struct {
	FailList []StockUpdateResponseFailList `json:"failList,omitempty"`
}

func CreateStockUpdateRequest() (request *StockUpdateRequest) {
	request = &StockUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "stock/update", "apigw")
	request.Method = api.POST
	return
}

func CreateStockUpdateResponse() (response *api.BaseResponse[StockUpdateResponse]) {
	response = api.CreateResponse[StockUpdateResponse](&StockUpdateResponse{})
	return
}

type StockUpdateRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type StockUpdateRequestGoodsList struct {
	SkuList []StockUpdateRequestSkuList `json:"skuList,omitempty"`
	GoodsId int64                       `json:"goodsId,omitempty"`
}

type StockUpdateRequestSkuList struct {
	SkuId    int64 `json:"skuId,omitempty"`
	StockNum int64 `json:"stockNum,omitempty"`
}

type StockUpdateResponseFailList struct {
	GoodsId  int64   `json:"goodsId,omitempty"`
	Message  string  `json:"message,omitempty"`
	SkuIdSet []int64 `json:"skuIdSet,omitempty"`
}
