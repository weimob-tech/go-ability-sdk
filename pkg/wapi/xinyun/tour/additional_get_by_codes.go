package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AdditionalGetByCodes
// @id 1032
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=根据code批量查询附加产品列表)
func (client *Client) AdditionalGetByCodes(request *AdditionalGetByCodesRequest) (response *AdditionalGetByCodesResponse, err error) {
	rpcResponse := CreateAdditionalGetByCodesResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AdditionalGetByCodesRequest struct {
	*api.BaseRequest

	AdditionalCodes []string `json:"additionalCodes,omitempty"`
}

type AdditionalGetByCodesResponse struct {
}

func CreateAdditionalGetByCodesRequest() (request *AdditionalGetByCodesRequest) {
	request = &AdditionalGetByCodesRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "additional/getByCodes", "api")
	request.Method = api.POST
	return
}

func CreateAdditionalGetByCodesResponse() (response *api.BaseResponse[AdditionalGetByCodesResponse]) {
	response = api.CreateResponse[AdditionalGetByCodesResponse](&AdditionalGetByCodesResponse{})
	return
}
