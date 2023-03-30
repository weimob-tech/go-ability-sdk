package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StockWarehouseGetListByVid
// @id 2272
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2272?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=门店获取仓库信息)
func (client *Client) StockWarehouseGetListByVid(request *StockWarehouseGetListByVidRequest) (response *StockWarehouseGetListByVidResponse, err error) {
	rpcResponse := CreateStockWarehouseGetListByVidResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StockWarehouseGetListByVidRequest struct {
	*api.BaseRequest

	BasicInfo StockWarehouseGetListByVidRequestBasicInfo `json:"basicInfo,omitempty"`
}

type StockWarehouseGetListByVidResponse struct {
	WarehouseCode        string `json:"warehouseCode,omitempty"`
	WarehouseChannelType int64  `json:"warehouseChannelType,omitempty"`
	Name                 string `json:"name,omitempty"`
	WarehouseId          int64  `json:"warehouseId,omitempty"`
	Type                 int64  `json:"type,omitempty"`
	Sort                 int64  `json:"sort,omitempty"`
}

func CreateStockWarehouseGetListByVidRequest() (request *StockWarehouseGetListByVidRequest) {
	request = &StockWarehouseGetListByVidRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "stock/warehouse/getListByVid", "apigw")
	request.Method = api.POST
	return
}

func CreateStockWarehouseGetListByVidResponse() (response *api.BaseResponse[StockWarehouseGetListByVidResponse]) {
	response = api.CreateResponse[StockWarehouseGetListByVidResponse](&StockWarehouseGetListByVidResponse{})
	return
}

type StockWarehouseGetListByVidRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
