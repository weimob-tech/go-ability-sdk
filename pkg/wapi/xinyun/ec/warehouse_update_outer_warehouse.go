package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// WarehouseUpdateOuterWarehouse
// @id 1041
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改外部仓库接口)
func (client *Client) WarehouseUpdateOuterWarehouse(request *WarehouseUpdateOuterWarehouseRequest) (response *WarehouseUpdateOuterWarehouseResponse, err error) {
	rpcResponse := CreateWarehouseUpdateOuterWarehouseResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WarehouseUpdateOuterWarehouseRequest struct {
	*api.BaseRequest

	WarehouseId   int64  `json:"warehouseId,omitempty"`
	WarehouseCode string `json:"warehouseCode,omitempty"`
	Name          string `json:"name,omitempty"`
	AdminName     string `json:"adminName,omitempty"`
	Phone         string `json:"phone,omitempty"`
	Remark        string `json:"remark,omitempty"`
}

type WarehouseUpdateOuterWarehouseResponse struct {
}

func CreateWarehouseUpdateOuterWarehouseRequest() (request *WarehouseUpdateOuterWarehouseRequest) {
	request = &WarehouseUpdateOuterWarehouseRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "warehouse/updateOuterWarehouse", "api")
	request.Method = api.POST
	return
}

func CreateWarehouseUpdateOuterWarehouseResponse() (response *api.BaseResponse[WarehouseUpdateOuterWarehouseResponse]) {
	response = api.CreateResponse[WarehouseUpdateOuterWarehouseResponse](&WarehouseUpdateOuterWarehouseResponse{})
	return
}
