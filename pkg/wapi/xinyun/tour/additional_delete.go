package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AdditionalDelete
// @id 1030
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除附加产品)
func (client *Client) AdditionalDelete(request *AdditionalDeleteRequest) (response *AdditionalDeleteResponse, err error) {
	rpcResponse := CreateAdditionalDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AdditionalDeleteRequest struct {
	*api.BaseRequest

	AdditionalCode string `json:"additionalCode,omitempty"`
}

type AdditionalDeleteResponse struct {
}

func CreateAdditionalDeleteRequest() (request *AdditionalDeleteRequest) {
	request = &AdditionalDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "additional/delete", "api")
	request.Method = api.POST
	return
}

func CreateAdditionalDeleteResponse() (response *api.BaseResponse[AdditionalDeleteResponse]) {
	response = api.CreateResponse[AdditionalDeleteResponse](&AdditionalDeleteResponse{})
	return
}
