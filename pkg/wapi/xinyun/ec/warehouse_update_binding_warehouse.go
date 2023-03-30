package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// WarehouseUpdateBindingWarehouse
// @id 2389
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=绑定解绑仓库)
func (client *Client) WarehouseUpdateBindingWarehouse(request *WarehouseUpdateBindingWarehouseRequest) (response *WarehouseUpdateBindingWarehouseResponse, err error) {
	rpcResponse := CreateWarehouseUpdateBindingWarehouseResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WarehouseUpdateBindingWarehouseRequest struct {
	*api.BaseRequest

	WarehouseIdList []int `json:"warehouseIdList,omitempty"`
	StoreId         int64 `json:"storeId,omitempty"`
	ChannelType     int   `json:"channelType,omitempty"`
}

type WarehouseUpdateBindingWarehouseResponse struct {
}

func CreateWarehouseUpdateBindingWarehouseRequest() (request *WarehouseUpdateBindingWarehouseRequest) {
	request = &WarehouseUpdateBindingWarehouseRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "warehouse/updateBindingWarehouse", "api")
	request.Method = api.POST
	return
}

func CreateWarehouseUpdateBindingWarehouseResponse() (response *api.BaseResponse[WarehouseUpdateBindingWarehouseResponse]) {
	response = api.CreateResponse[WarehouseUpdateBindingWarehouseResponse](&WarehouseUpdateBindingWarehouseResponse{})
	return
}
