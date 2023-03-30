package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StockInventoryInsert
// @id 2286
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2286?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=创建采购、换货、其他入库单或者创建出库单)
func (client *Client) StockInventoryInsert(request *StockInventoryInsertRequest) (response *StockInventoryInsertResponse, err error) {
	rpcResponse := CreateStockInventoryInsertResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StockInventoryInsertRequest struct {
	*api.BaseRequest

	BasicInfo       StockInventoryInsertRequestBasicInfo         `json:"basicInfo,omitempty"`
	ReceiptItemList []StockInventoryInsertRequestReceiptItemList `json:"receiptItemList,omitempty"`
	OperateType     int64                                        `json:"operateType,omitempty"`
	BizType         int64                                        `json:"bizType,omitempty"`
	OutReceiptId    int64                                        `json:"outReceiptId,omitempty"`
	SupplierId      int64                                        `json:"supplierId,omitempty"`
	Remark          string                                       `json:"remark,omitempty"`
	Wid             int64                                        `json:"wid,omitempty"`
}

type StockInventoryInsertResponse struct {
	BusinessReceiptId int64 `json:"businessReceiptId,omitempty"`
	InventoryId       int64 `json:"inventoryId,omitempty"`
	ReferId           int64 `json:"referId,omitempty"`
	CreateResult      int64 `json:"createResult,omitempty"`
	TotalCount        int64 `json:"totalCount,omitempty"`
	SuccessCount      int64 `json:"successCount,omitempty"`
	FailCount         int64 `json:"failCount,omitempty"`
}

func CreateStockInventoryInsertRequest() (request *StockInventoryInsertRequest) {
	request = &StockInventoryInsertRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "stock/inventory/insert", "apigw")
	request.Method = api.POST
	return
}

func CreateStockInventoryInsertResponse() (response *api.BaseResponse[StockInventoryInsertResponse]) {
	response = api.CreateResponse[StockInventoryInsertResponse](&StockInventoryInsertResponse{})
	return
}

type StockInventoryInsertRequestBasicInfo struct {
	Vid   int64 `json:"vid,omitempty"`
	BosId int64 `json:"bosId,omitempty"`
}

type StockInventoryInsertRequestReceiptItemList struct {
	ItemSkuId int64  `json:"itemSkuId,omitempty"`
	Quantity  string `json:"quantity,omitempty"`
	ItemId    int64  `json:"itemId,omitempty"`
}
