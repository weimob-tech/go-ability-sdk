package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StockItemSkuGetList
// @id 2302
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2302?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询门店库存列表)
func (client *Client) StockItemSkuGetList(request *StockItemSkuGetListRequest) (response *StockItemSkuGetListResponse, err error) {
	rpcResponse := CreateStockItemSkuGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StockItemSkuGetListRequest struct {
	*api.BaseRequest

	BasicInfo      StockItemSkuGetListRequestBasicInfo      `json:"basicInfo,omitempty"`
	QueryParameter StockItemSkuGetListRequestQueryParameter `json:"queryParameter,omitempty"`
	PageNum        int64                                    `json:"pageNum,omitempty"`
	PageSize       int64                                    `json:"pageSize,omitempty"`
}

type StockItemSkuGetListResponse struct {
	PageList   []StockItemSkuGetListResponsePageList `json:"pageList,omitempty"`
	TatalCount int64                                 `json:"tatalCount,omitempty"`
	PageSize   int64                                 `json:"pageSize,omitempty"`
	PageNum    int64                                 `json:"pageNum,omitempty"`
}

func CreateStockItemSkuGetListRequest() (request *StockItemSkuGetListRequest) {
	request = &StockItemSkuGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "stock/item/sku/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateStockItemSkuGetListResponse() (response *api.BaseResponse[StockItemSkuGetListResponse]) {
	response = api.CreateResponse[StockItemSkuGetListResponse](&StockItemSkuGetListResponse{})
	return
}

type StockItemSkuGetListRequestBasicInfo struct {
	Vid   int64 `json:"vid,omitempty"`
	BosId int64 `json:"bosId,omitempty"`
}

type StockItemSkuGetListRequestQueryParameter struct {
	Search     string `json:"search,omitempty"`
	SearchType int64  `json:"searchType,omitempty"`
	ClassifyId int64  `json:"classifyId,omitempty"`
}

type StockItemSkuGetListResponsePageList struct {
	StoreWarehouse StockItemSkuGetListResponseStoreWarehouse `json:"storeWarehouse,omitempty"`
	ItemId         int64                                     `json:"itemId,omitempty"`
	ItemSkuTitle   string                                    `json:"itemSkuTitle,omitempty"`
	SpecDetail     string                                    `json:"specDetail,omitempty"`
	ProductCode    string                                    `json:"productCode,omitempty"`
	UnitName       string                                    `json:"unitName,omitempty"`
	SourceVidName  string                                    `json:"sourceVidName,omitempty"`
	ItemSkuId      int64                                     `json:"itemSkuId,omitempty"`
	ImageUrl       string                                    `json:"imageUrl,omitempty"`
	CostPrice      string                                    `json:"costPrice,omitempty"`
	ClassifyName   string                                    `json:"classifyName,omitempty"`
}

type StockItemSkuGetListResponseStoreWarehouse struct {
	OccupyStockNum        string `json:"occupyStockNum,omitempty"`
	WarehouseId           int64  `json:"warehouseId,omitempty"`
	AllowOversoldStockNum string `json:"allowOversoldStockNum,omitempty"`
	TotalStockNum         string `json:"totalStockNum,omitempty"`
	UsableStockNum        string `json:"usableStockNum,omitempty"`
	TransportStockNum     string `json:"transportStockNum,omitempty"`
	DefectiveStockNum     string `json:"defectiveStockNum,omitempty"`
}
