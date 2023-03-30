package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AdditionalQueryByCondition
// @id 1031
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=限定条件查询附加产品列表)
func (client *Client) AdditionalQueryByCondition(request *AdditionalQueryByConditionRequest) (response *AdditionalQueryByConditionResponse, err error) {
	rpcResponse := CreateAdditionalQueryByConditionResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AdditionalQueryByConditionRequest struct {
	*api.BaseRequest

	Name      string `json:"name,omitempty"`
	PageIndex int    `json:"pageIndex,omitempty"`
	PageSize  int    `json:"pageSize,omitempty"`
}

type AdditionalQueryByConditionResponse struct {
}

func CreateAdditionalQueryByConditionRequest() (request *AdditionalQueryByConditionRequest) {
	request = &AdditionalQueryByConditionRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "additional/queryByCondition", "api")
	request.Method = api.POST
	return
}

func CreateAdditionalQueryByConditionResponse() (response *api.BaseResponse[AdditionalQueryByConditionResponse]) {
	response = api.CreateResponse[AdditionalQueryByConditionResponse](&AdditionalQueryByConditionResponse{})
	return
}
