package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StockWarehouseUpdate
// @id 1592
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1592?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=修改仓库)
func (client *Client) StockWarehouseUpdate(request *StockWarehouseUpdateRequest) (response *StockWarehouseUpdateResponse, err error) {
	rpcResponse := CreateStockWarehouseUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StockWarehouseUpdateRequest struct {
	*api.BaseRequest

	BasicInfo           StockWarehouseUpdateRequestBasicInfo `json:"basicInfo,omitempty"`
	ManagerName         string                               `json:"managerName,omitempty"`
	WarehouseId         int64                                `json:"warehouseId,omitempty"`
	Phone               string                               `json:"phone,omitempty"`
	WarehouseName       string                               `json:"warehouseName,omitempty"`
	Remark              string                               `json:"remark,omitempty"`
	WarehouseCode       string                               `json:"warehouseCode,omitempty"`
	Address             string                               `json:"address,omitempty"`
	DeliveryType        int64                                `json:"deliveryType,omitempty"`
	SubDeliveryTypeList []int64                              `json:"subDeliveryTypeList,omitempty"`
}

type StockWarehouseUpdateResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateStockWarehouseUpdateRequest() (request *StockWarehouseUpdateRequest) {
	request = &StockWarehouseUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "stock/warehouse/update", "apigw")
	request.Method = api.POST
	return
}

func CreateStockWarehouseUpdateResponse() (response *api.BaseResponse[StockWarehouseUpdateResponse]) {
	response = api.CreateResponse[StockWarehouseUpdateResponse](&StockWarehouseUpdateResponse{})
	return
}

type StockWarehouseUpdateRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
