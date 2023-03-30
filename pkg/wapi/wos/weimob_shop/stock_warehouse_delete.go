package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StockWarehouseDelete
// @id 1861
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1861?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除外部仓库)
func (client *Client) StockWarehouseDelete(request *StockWarehouseDeleteRequest) (response *StockWarehouseDeleteResponse, err error) {
	rpcResponse := CreateStockWarehouseDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StockWarehouseDeleteRequest struct {
	*api.BaseRequest

	BasicInfo       StockWarehouseDeleteRequestBasicInfo `json:"basicInfo,omitempty"`
	WarehouseIdList []int64                              `json:"warehouseIdList,omitempty"`
}

type StockWarehouseDeleteResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateStockWarehouseDeleteRequest() (request *StockWarehouseDeleteRequest) {
	request = &StockWarehouseDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "stock/warehouse/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateStockWarehouseDeleteResponse() (response *api.BaseResponse[StockWarehouseDeleteResponse]) {
	response = api.CreateResponse[StockWarehouseDeleteResponse](&StockWarehouseDeleteResponse{})
	return
}

type StockWarehouseDeleteRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
