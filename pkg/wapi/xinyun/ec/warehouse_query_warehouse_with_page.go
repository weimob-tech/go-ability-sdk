package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// WarehouseQueryWarehouseWithPage
// @id 1043
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取商家仓库列表)
func (client *Client) WarehouseQueryWarehouseWithPage(request *WarehouseQueryWarehouseWithPageRequest) (response *WarehouseQueryWarehouseWithPageResponse, err error) {
	rpcResponse := CreateWarehouseQueryWarehouseWithPageResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WarehouseQueryWarehouseWithPageRequest struct {
	*api.BaseRequest

	QueryParameter WarehouseQueryWarehouseWithPageRequestQueryParameter `json:"queryParameter,omitempty"`
	PageNum        int                                                  `json:"pageNum,omitempty"`
	PageSize       int                                                  `json:"pageSize,omitempty"`
	SearchContent  string                                               `json:"searchContent,omitempty"`
	SearchType     int64                                                `json:"searchType,omitempty"`
}

type WarehouseQueryWarehouseWithPageResponse struct {
}

func CreateWarehouseQueryWarehouseWithPageRequest() (request *WarehouseQueryWarehouseWithPageRequest) {
	request = &WarehouseQueryWarehouseWithPageRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "warehouse/queryWarehouseWithPage", "api")
	request.Method = api.POST
	return
}

func CreateWarehouseQueryWarehouseWithPageResponse() (response *api.BaseResponse[WarehouseQueryWarehouseWithPageResponse]) {
	response = api.CreateResponse[WarehouseQueryWarehouseWithPageResponse](&WarehouseQueryWarehouseWithPageResponse{})
	return
}

type WarehouseQueryWarehouseWithPageRequestQueryParameter struct {
}
