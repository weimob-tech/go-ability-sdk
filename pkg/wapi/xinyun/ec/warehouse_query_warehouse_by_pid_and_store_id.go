package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// WarehouseQueryWarehouseByPidAndStoreId
// @id 1084
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=门店获取仓库信息)
func (client *Client) WarehouseQueryWarehouseByPidAndStoreId(request *WarehouseQueryWarehouseByPidAndStoreIdRequest) (response *WarehouseQueryWarehouseByPidAndStoreIdResponse, err error) {
	rpcResponse := CreateWarehouseQueryWarehouseByPidAndStoreIdResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WarehouseQueryWarehouseByPidAndStoreIdRequest struct {
	*api.BaseRequest

	StoreId int64 `json:"storeId,omitempty"`
}

type WarehouseQueryWarehouseByPidAndStoreIdResponse struct {
}

func CreateWarehouseQueryWarehouseByPidAndStoreIdRequest() (request *WarehouseQueryWarehouseByPidAndStoreIdRequest) {
	request = &WarehouseQueryWarehouseByPidAndStoreIdRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "warehouse/queryWarehouseByPidAndStoreId", "api")
	request.Method = api.POST
	return
}

func CreateWarehouseQueryWarehouseByPidAndStoreIdResponse() (response *api.BaseResponse[WarehouseQueryWarehouseByPidAndStoreIdResponse]) {
	response = api.CreateResponse[WarehouseQueryWarehouseByPidAndStoreIdResponse](&WarehouseQueryWarehouseByPidAndStoreIdResponse{})
	return
}
