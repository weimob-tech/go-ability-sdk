package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StockWarehouseGet
// @id 2291
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2291?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取仓库详情接口)
func (client *Client) StockWarehouseGet(request *StockWarehouseGetRequest) (response *StockWarehouseGetResponse, err error) {
	rpcResponse := CreateStockWarehouseGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StockWarehouseGetRequest struct {
	*api.BaseRequest

	BasicInfo   StockWarehouseGetRequestBasicInfo `json:"basicInfo,omitempty"`
	WarehouseId int64                             `json:"warehouseId,omitempty"`
}

type StockWarehouseGetResponse struct {
	WarehousePropertyList StockWarehouseGetResponseWarehousePropertyList `json:"warehousePropertyList,omitempty"`
	StoreId               int64                                          `json:"storeId,omitempty"`
	WarehouseId           int64                                          `json:"warehouseId,omitempty"`
	WarehouseCode         string                                         `json:"warehouseCode,omitempty"`
	WarehouseName         string                                         `json:"warehouseName,omitempty"`
	WarehouseType         int64                                          `json:"warehouseType,omitempty"`
	Address               string                                         `json:"address,omitempty"`
	ManagerName           string                                         `json:"managerName,omitempty"`
	Phone                 int64                                          `json:"phone,omitempty"`
	Remark                string                                         `json:"remark,omitempty"`
	DeliveryType          int64                                          `json:"deliveryType,omitempty"`
}

func CreateStockWarehouseGetRequest() (request *StockWarehouseGetRequest) {
	request = &StockWarehouseGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "stock/warehouse/get", "apigw")
	request.Method = api.POST
	return
}

func CreateStockWarehouseGetResponse() (response *api.BaseResponse[StockWarehouseGetResponse]) {
	response = api.CreateResponse[StockWarehouseGetResponse](&StockWarehouseGetResponse{})
	return
}

type StockWarehouseGetRequestBasicInfo struct {
	Vid   int64 `json:"vid,omitempty"`
	BosId int64 `json:"bosId,omitempty"`
}

type StockWarehouseGetResponseWarehousePropertyList struct {
	ExtKey   string `json:"extKey,omitempty"`
	ExtValue string `json:"extValue,omitempty"`
}
