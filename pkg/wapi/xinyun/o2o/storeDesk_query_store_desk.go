package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreDeskQueryStoreDesk
// @id 53
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询桌台信息)
func (client *Client) StoreDeskQueryStoreDesk(request *StoreDeskQueryStoreDeskRequest) (response *StoreDeskQueryStoreDeskResponse, err error) {
	rpcResponse := CreateStoreDeskQueryStoreDeskResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreDeskQueryStoreDeskRequest struct {
	*api.BaseRequest

	DeskId int64 `json:"deskId,omitempty"`
}

type StoreDeskQueryStoreDeskResponse struct {
}

func CreateStoreDeskQueryStoreDeskRequest() (request *StoreDeskQueryStoreDeskRequest) {
	request = &StoreDeskQueryStoreDeskRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "storeDesk/queryStoreDesk", "api")
	request.Method = api.POST
	return
}

func CreateStoreDeskQueryStoreDeskResponse() (response *api.BaseResponse[StoreDeskQueryStoreDeskResponse]) {
	response = api.CreateResponse[StoreDeskQueryStoreDeskResponse](&StoreDeskQueryStoreDeskResponse{})
	return
}
