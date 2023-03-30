package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerTagTagged
// @id 1277
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量打标)
func (client *Client) CustomerTagTagged(request *CustomerTagTaggedRequest) (response *CustomerTagTaggedResponse, err error) {
	rpcResponse := CreateCustomerTagTaggedResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerTagTaggedRequest struct {
	*api.BaseRequest

	TagIdList int64 `json:"tagIdList,omitempty"`
	WidList   int64 `json:"widList,omitempty"`
}

type CustomerTagTaggedResponse struct {
}

func CreateCustomerTagTaggedRequest() (request *CustomerTagTaggedRequest) {
	request = &CustomerTagTaggedRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "customerTag/tagged", "api")
	request.Method = api.POST
	return
}

func CreateCustomerTagTaggedResponse() (response *api.BaseResponse[CustomerTagTaggedResponse]) {
	response = api.CreateResponse[CustomerTagTaggedResponse](&CustomerTagTaggedResponse{})
	return
}
