package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// WarehouseAddOuterWarehouse
// @id 1040
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增外部仓库接口)
func (client *Client) WarehouseAddOuterWarehouse(request *WarehouseAddOuterWarehouseRequest) (response *WarehouseAddOuterWarehouseResponse, err error) {
	rpcResponse := CreateWarehouseAddOuterWarehouseResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WarehouseAddOuterWarehouseRequest struct {
	*api.BaseRequest

	WarehouseCode         string `json:"warehouseCode,omitempty"`
	Name                  string `json:"name,omitempty"`
	AdminName             string `json:"adminName,omitempty"`
	Phone                 string `json:"phone,omitempty"`
	Remark                string `json:"remark,omitempty"`
	WarehouseDeliveryType int    `json:"warehouseDeliveryType,omitempty"`
}

type WarehouseAddOuterWarehouseResponse struct {
}

func CreateWarehouseAddOuterWarehouseRequest() (request *WarehouseAddOuterWarehouseRequest) {
	request = &WarehouseAddOuterWarehouseRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "warehouse/addOuterWarehouse", "api")
	request.Method = api.POST
	return
}

func CreateWarehouseAddOuterWarehouseResponse() (response *api.BaseResponse[WarehouseAddOuterWarehouseResponse]) {
	response = api.CreateResponse[WarehouseAddOuterWarehouseResponse](&WarehouseAddOuterWarehouseResponse{})
	return
}
