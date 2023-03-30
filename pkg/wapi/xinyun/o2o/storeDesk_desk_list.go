package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreDeskDeskList
// @id 80
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询桌台列表)
func (client *Client) StoreDeskDeskList(request *StoreDeskDeskListRequest) (response *StoreDeskDeskListResponse, err error) {
	rpcResponse := CreateStoreDeskDeskListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreDeskDeskListRequest struct {
	*api.BaseRequest

	Name     string `json:"name,omitempty"`
	Status   int    `json:"status,omitempty"`
	StoreId  int64  `json:"storeId,omitempty"`
	PageNum  int    `json:"pageNum,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
	BizType  int    `json:"bizType,omitempty"`
}

type StoreDeskDeskListResponse struct {
}

func CreateStoreDeskDeskListRequest() (request *StoreDeskDeskListRequest) {
	request = &StoreDeskDeskListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "storeDesk/deskList", "api")
	request.Method = api.POST
	return
}

func CreateStoreDeskDeskListResponse() (response *api.BaseResponse[StoreDeskDeskListResponse]) {
	response = api.CreateResponse[StoreDeskDeskListResponse](&StoreDeskDeskListResponse{})
	return
}
