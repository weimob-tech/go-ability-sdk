package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreDeskAdd
// @id 54
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=添加桌台)
func (client *Client) StoreDeskAdd(request *StoreDeskAddRequest) (response *StoreDeskAddResponse, err error) {
	rpcResponse := CreateStoreDeskAddResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreDeskAddRequest struct {
	*api.BaseRequest

	Name    string `json:"name,omitempty"`
	StoreId int64  `json:"storeId,omitempty"`
	Status  int    `json:"status,omitempty"`
	BizType int    `json:"bizType,omitempty"`
}

type StoreDeskAddResponse struct {
}

func CreateStoreDeskAddRequest() (request *StoreDeskAddRequest) {
	request = &StoreDeskAddRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "storeDesk/add", "api")
	request.Method = api.POST
	return
}

func CreateStoreDeskAddResponse() (response *api.BaseResponse[StoreDeskAddResponse]) {
	response = api.CreateResponse[StoreDeskAddResponse](&StoreDeskAddResponse{})
	return
}
