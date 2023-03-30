package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StockWarehouseGetList
// @id 1593
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1593?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取仓库列表)
func (client *Client) StockWarehouseGetList(request *StockWarehouseGetListRequest) (response *StockWarehouseGetListResponse, err error) {
	rpcResponse := CreateStockWarehouseGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StockWarehouseGetListRequest struct {
	*api.BaseRequest

	QueryParameter StockWarehouseGetListRequestQueryParameter `json:"queryParameter,omitempty"`
	BasicInfo      StockWarehouseGetListRequestBasicInfo      `json:"basicInfo,omitempty"`
	PageNum        int64                                      `json:"pageNum,omitempty"`
	PageSize       int64                                      `json:"pageSize,omitempty"`
}

type StockWarehouseGetListResponse struct {
	PageList   []StockWarehouseGetListResponsePageList `json:"pageList,omitempty"`
	PageSize   int64                                   `json:"pageSize,omitempty"`
	TotalCount int64                                   `json:"totalCount,omitempty"`
	PageNum    int64                                   `json:"pageNum,omitempty"`
}

func CreateStockWarehouseGetListRequest() (request *StockWarehouseGetListRequest) {
	request = &StockWarehouseGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "stock/warehouse/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateStockWarehouseGetListResponse() (response *api.BaseResponse[StockWarehouseGetListResponse]) {
	response = api.CreateResponse[StockWarehouseGetListResponse](&StockWarehouseGetListResponse{})
	return
}

type StockWarehouseGetListRequestQueryParameter struct {
	WarehouseType int64  `json:"warehouseType,omitempty"`
	SearchType    int64  `json:"searchType,omitempty"`
	Search        string `json:"search,omitempty"`
}

type StockWarehouseGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type StockWarehouseGetListResponsePageList struct {
	DeliveryType        int64   `json:"deliveryType,omitempty"`
	WarehouseName       string  `json:"warehouseName,omitempty"`
	WarehouseCode       string  `json:"warehouseCode,omitempty"`
	WarehouseId         int64   `json:"warehouseId,omitempty"`
	SubDeliveryTypeList []int64 `json:"subDeliveryTypeList,omitempty"`
	Address             string  `json:"address,omitempty"`
	ManagerName         string  `json:"managerName,omitempty"`
	Phone               string  `json:"phone,omitempty"`
	Remark              string  `json:"remark,omitempty"`
}
