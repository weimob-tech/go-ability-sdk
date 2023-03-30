package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// WarehouseQueryWarehouseDetail
// @id 1042
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取仓库详情接口)
func (client *Client) WarehouseQueryWarehouseDetail(request *WarehouseQueryWarehouseDetailRequest) (response *WarehouseQueryWarehouseDetailResponse, err error) {
	rpcResponse := CreateWarehouseQueryWarehouseDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WarehouseQueryWarehouseDetailRequest struct {
	*api.BaseRequest

	WarehouseId int64 `json:"warehouseId,omitempty"`
}

type WarehouseQueryWarehouseDetailResponse struct {
	WarehouseId   int64  `json:"warehouseId,omitempty"`
	WarehouseCode string `json:"warehouseCode,omitempty"`
	Name          string `json:"name,omitempty"`
	AdminName     string `json:"adminName,omitempty"`
	Phone         string `json:"phone,omitempty"`
	Remark        string `json:"remark,omitempty"`
}

func CreateWarehouseQueryWarehouseDetailRequest() (request *WarehouseQueryWarehouseDetailRequest) {
	request = &WarehouseQueryWarehouseDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "warehouse/queryWarehouseDetail", "api")
	request.Method = api.POST
	return
}

func CreateWarehouseQueryWarehouseDetailResponse() (response *api.BaseResponse[WarehouseQueryWarehouseDetailResponse]) {
	response = api.CreateResponse[WarehouseQueryWarehouseDetailResponse](&WarehouseQueryWarehouseDetailResponse{})
	return
}
