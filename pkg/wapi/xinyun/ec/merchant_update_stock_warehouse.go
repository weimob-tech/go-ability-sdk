package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantUpdateStockWarehouse
// @id 2520
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更改门店仓库)
func (client *Client) MerchantUpdateStockWarehouse(request *MerchantUpdateStockWarehouseRequest) (response *MerchantUpdateStockWarehouseResponse, err error) {
	rpcResponse := CreateMerchantUpdateStockWarehouseResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantUpdateStockWarehouseRequest struct {
	*api.BaseRequest

	OnLineWarehouseList      []map[string]any                                   `json:"onLineWarehouseList,omitempty"`
	OffLineWarehouseList     []map[string]any                                   `json:"offLineWarehouseList,omitempty"`
	WarehouseInfoVO          MerchantUpdateStockWarehouseRequestWarehouseInfoVO `json:"WarehouseInfoVO,omitempty"`
	EcBizStoreId             int64                                              `json:"ecBizStoreId,omitempty"`
	OnLineWarehouseTypeList  []int                                              `json:"onLineWarehouseTypeList,omitempty"`
	OffLineWarehouseTypeList []int                                              `json:"offLineWarehouseTypeList,omitempty"`
}

type MerchantUpdateStockWarehouseResponse struct {
	OnlineStockFailList []map[string]any `json:"onlineStockFailList,omitempty"`
}

func CreateMerchantUpdateStockWarehouseRequest() (request *MerchantUpdateStockWarehouseRequest) {
	request = &MerchantUpdateStockWarehouseRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchant/updateStockWarehouse", "api")
	request.Method = api.POST
	return
}

func CreateMerchantUpdateStockWarehouseResponse() (response *api.BaseResponse[MerchantUpdateStockWarehouseResponse]) {
	response = api.CreateResponse[MerchantUpdateStockWarehouseResponse](&MerchantUpdateStockWarehouseResponse{})
	return
}

type MerchantUpdateStockWarehouseRequestWarehouseInfoVO struct {
	WarehouseId   int64  `json:"warehouseId,omitempty"`
	WarehouseName string `json:"warehouseName,omitempty"`
}
