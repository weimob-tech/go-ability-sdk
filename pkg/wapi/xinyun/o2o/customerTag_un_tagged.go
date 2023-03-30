package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerTagUnTagged
// @id 1278
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量取消打标)
func (client *Client) CustomerTagUnTagged(request *CustomerTagUnTaggedRequest) (response *CustomerTagUnTaggedResponse, err error) {
	rpcResponse := CreateCustomerTagUnTaggedResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerTagUnTaggedRequest struct {
	*api.BaseRequest

	TagIdList int64 `json:"tagIdList,omitempty"`
	WidList   int64 `json:"widList,omitempty"`
}

type CustomerTagUnTaggedResponse struct {
}

func CreateCustomerTagUnTaggedRequest() (request *CustomerTagUnTaggedRequest) {
	request = &CustomerTagUnTaggedRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "customerTag/unTagged", "api")
	request.Method = api.POST
	return
}

func CreateCustomerTagUnTaggedResponse() (response *api.BaseResponse[CustomerTagUnTaggedResponse]) {
	response = api.CreateResponse[CustomerTagUnTaggedResponse](&CustomerTagUnTaggedResponse{})
	return
}
