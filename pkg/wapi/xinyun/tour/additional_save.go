package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AdditionalSave
// @id 1029
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增/编辑附加产品)
func (client *Client) AdditionalSave(request *AdditionalSaveRequest) (response *AdditionalSaveResponse, err error) {
	rpcResponse := CreateAdditionalSaveResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AdditionalSaveRequest struct {
	*api.BaseRequest

	AdditionalCode string  `json:"additionalCode,omitempty"`
	Name           string  `json:"name,omitempty"`
	Intro          string  `json:"intro,omitempty"`
	Price          float64 `json:"price,omitempty"`
}

type AdditionalSaveResponse struct {
}

func CreateAdditionalSaveRequest() (request *AdditionalSaveRequest) {
	request = &AdditionalSaveRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "additional/save", "api")
	request.Method = api.POST
	return
}

func CreateAdditionalSaveResponse() (response *api.BaseResponse[AdditionalSaveResponse]) {
	response = api.CreateResponse[AdditionalSaveResponse](&AdditionalSaveResponse{})
	return
}
