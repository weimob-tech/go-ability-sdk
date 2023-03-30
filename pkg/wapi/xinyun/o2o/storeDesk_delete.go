package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreDeskDelete
// @id 56
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除桌台)
func (client *Client) StoreDeskDelete(request *StoreDeskDeleteRequest) (response *StoreDeskDeleteResponse, err error) {
	rpcResponse := CreateStoreDeskDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreDeskDeleteRequest struct {
	*api.BaseRequest

	DeskId int64 `json:"deskId,omitempty"`
}

type StoreDeskDeleteResponse struct {
}

func CreateStoreDeskDeleteRequest() (request *StoreDeskDeleteRequest) {
	request = &StoreDeskDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "storeDesk/delete", "api")
	request.Method = api.POST
	return
}

func CreateStoreDeskDeleteResponse() (response *api.BaseResponse[StoreDeskDeleteResponse]) {
	response = api.CreateResponse[StoreDeskDeleteResponse](&StoreDeskDeleteResponse{})
	return
}
