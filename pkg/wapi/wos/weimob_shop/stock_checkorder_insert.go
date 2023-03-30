package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StockCheckorderInsert
// @id 2303
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2303?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=新增盘点单)
func (client *Client) StockCheckorderInsert(request *StockCheckorderInsertRequest) (response *StockCheckorderInsertResponse, err error) {
	rpcResponse := CreateStockCheckorderInsertResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StockCheckorderInsertRequest struct {
	*api.BaseRequest

	BasicInfo          StockCheckorderInsertRequestBasicInfo            `json:"basicInfo,omitempty"`
	StockCheckItemList []StockCheckorderInsertRequestStockCheckItemList `json:"stockCheckItemList,omitempty"`
	Remark             string                                           `json:"remark,omitempty"`
	Wid                int64                                            `json:"wid,omitempty"`
	Status             int64                                            `json:"status,omitempty"`
}

type StockCheckorderInsertResponse struct {
	FailCount    int64 `json:"failCount,omitempty"`
	SuccessCount int64 `json:"successCount,omitempty"`
	CreateResult int64 `json:"createResult,omitempty"`
	TotalCount   int64 `json:"totalCount,omitempty"`
	ReceiptId    int64 `json:"receiptId,omitempty"`
}

func CreateStockCheckorderInsertRequest() (request *StockCheckorderInsertRequest) {
	request = &StockCheckorderInsertRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "stock/checkorder/insert", "apigw")
	request.Method = api.POST
	return
}

func CreateStockCheckorderInsertResponse() (response *api.BaseResponse[StockCheckorderInsertResponse]) {
	response = api.CreateResponse[StockCheckorderInsertResponse](&StockCheckorderInsertResponse{})
	return
}

type StockCheckorderInsertRequestBasicInfo struct {
	Vid   int64 `json:"vid,omitempty"`
	BosId int64 `json:"bosId,omitempty"`
}

type StockCheckorderInsertRequestStockCheckItemList struct {
	ItemId        int64  `json:"itemId,omitempty"`
	ItemSkuId     int64  `json:"itemSkuId,omitempty"`
	StockCheckNum string `json:"stockCheckNum,omitempty"`
}
