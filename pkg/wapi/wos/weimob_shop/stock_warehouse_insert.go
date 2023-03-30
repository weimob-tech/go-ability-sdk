package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StockWarehouseInsert
// @id 1591
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1591?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=添加仓库)
func (client *Client) StockWarehouseInsert(request *StockWarehouseInsertRequest) (response *StockWarehouseInsertResponse, err error) {
	rpcResponse := CreateStockWarehouseInsertResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StockWarehouseInsertRequest struct {
	*api.BaseRequest

	BasicInfo           StockWarehouseInsertRequestBasicInfo `json:"basicInfo,omitempty"`
	Phone               string                               `json:"phone,omitempty"`
	Remark              string                               `json:"remark,omitempty"`
	Address             string                               `json:"address,omitempty"`
	ManagerName         string                               `json:"managerName,omitempty"`
	WarehouseName       string                               `json:"warehouseName,omitempty"`
	WarehouseCode       string                               `json:"warehouseCode,omitempty"`
	SubDeliveryTypeList []int64                              `json:"subDeliveryTypeList,omitempty"`
	DeliveryType        int64                                `json:"deliveryType,omitempty"`
}

type StockWarehouseInsertResponse struct {
	Id int64 `json:"id,omitempty"`
}

func CreateStockWarehouseInsertRequest() (request *StockWarehouseInsertRequest) {
	request = &StockWarehouseInsertRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "stock/warehouse/insert", "apigw")
	request.Method = api.POST
	return
}

func CreateStockWarehouseInsertResponse() (response *api.BaseResponse[StockWarehouseInsertResponse]) {
	response = api.CreateResponse[StockWarehouseInsertResponse](&StockWarehouseInsertResponse{})
	return
}

type StockWarehouseInsertRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
