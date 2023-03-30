package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreDeskSyn
// @id 290
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=同步桌台)
func (client *Client) StoreDeskSyn(request *StoreDeskSynRequest) (response *StoreDeskSynResponse, err error) {
	rpcResponse := CreateStoreDeskSynResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreDeskSynRequest struct {
	*api.BaseRequest

	StoreId int64  `json:"storeId,omitempty"`
	ThirdId string `json:"thirdId,omitempty"`
	Name    string `json:"name,omitempty"`
	Number  int    `json:"number,omitempty"`
	Depict  string `json:"depict,omitempty"`
	BizType int    `json:"bizType,omitempty"`
}

type StoreDeskSynResponse struct {
}

func CreateStoreDeskSynRequest() (request *StoreDeskSynRequest) {
	request = &StoreDeskSynRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "storeDesk/syn", "api")
	request.Method = api.POST
	return
}

func CreateStoreDeskSynResponse() (response *api.BaseResponse[StoreDeskSynResponse]) {
	response = api.CreateResponse[StoreDeskSynResponse](&StoreDeskSynResponse{})
	return
}
