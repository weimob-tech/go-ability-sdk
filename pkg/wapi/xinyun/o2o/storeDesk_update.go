package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreDeskUpdate
// @id 55
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改桌台信息)
func (client *Client) StoreDeskUpdate(request *StoreDeskUpdateRequest) (response *StoreDeskUpdateResponse, err error) {
	rpcResponse := CreateStoreDeskUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreDeskUpdateRequest struct {
	*api.BaseRequest

	Id      int64  `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	StoreId int64  `json:"storeId,omitempty"`
	Status  int    `json:"status,omitempty"`
}

type StoreDeskUpdateResponse struct {
}

func CreateStoreDeskUpdateRequest() (request *StoreDeskUpdateRequest) {
	request = &StoreDeskUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "storeDesk/update", "api")
	request.Method = api.POST
	return
}

func CreateStoreDeskUpdateResponse() (response *api.BaseResponse[StoreDeskUpdateResponse]) {
	response = api.CreateResponse[StoreDeskUpdateResponse](&StoreDeskUpdateResponse{})
	return
}
