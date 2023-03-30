package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StockInventoryDetailGetList
// @id 2287
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2287?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询出入库单据详情)
func (client *Client) StockInventoryDetailGetList(request *StockInventoryDetailGetListRequest) (response *StockInventoryDetailGetListResponse, err error) {
	rpcResponse := CreateStockInventoryDetailGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StockInventoryDetailGetListRequest struct {
	*api.BaseRequest

	BasicInfo         StockInventoryDetailGetListRequestBasicInfo `json:"basicInfo,omitempty"`
	ReceiptId         int64                                       `json:"receiptId,omitempty"`
	BusinessReceiptId int64                                       `json:"businessReceiptId,omitempty"`
	SearchType        int64                                       `json:"searchType,omitempty"`
	Search            string                                      `json:"search,omitempty"`
	Sort              int64                                       `json:"sort,omitempty"`
	PageNum           int64                                       `json:"pageNum,omitempty"`
	PageSize          int64                                       `json:"pageSize,omitempty"`
}

type StockInventoryDetailGetListResponse struct {
	PageList       []StockInventoryDetailGetListResponsePageList `json:"pageList,omitempty"`
	PageNum        int64                                         `json:"pageNum,omitempty"`
	PageSize       int64                                         `json:"pageSize,omitempty"`
	TotalCount     int64                                         `json:"totalCount,omitempty"`
	TotalQuantity  string                                        `json:"totalQuantity,omitempty"`
	TotalRemainNum string                                        `json:"totalRemainNum,omitempty"`
}

func CreateStockInventoryDetailGetListRequest() (request *StockInventoryDetailGetListRequest) {
	request = &StockInventoryDetailGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "stock/inventory/detail/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateStockInventoryDetailGetListResponse() (response *api.BaseResponse[StockInventoryDetailGetListResponse]) {
	response = api.CreateResponse[StockInventoryDetailGetListResponse](&StockInventoryDetailGetListResponse{})
	return
}

type StockInventoryDetailGetListRequestBasicInfo struct {
	Vid   int64 `json:"vid,omitempty"`
	BosId int64 `json:"bosId,omitempty"`
}

type StockInventoryDetailGetListResponsePageList struct {
	Title       string `json:"title,omitempty"`
	SpecDetail  string `json:"specDetail,omitempty"`
	ProductCode string `json:"productCode,omitempty"`
	ImgUrl      string `json:"imgUrl,omitempty"`
	ItemId      int64  `json:"itemId,omitempty"`
	ItemSkuId   int64  `json:"itemSkuId,omitempty"`
	Quantity    string `json:"quantity,omitempty"`
	UnitName    string `json:"unitName,omitempty"`
	RemainNum   int64  `json:"remainNum,omitempty"`
}
