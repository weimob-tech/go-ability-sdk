package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerTagThirdUntagged
// @id 1828
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=给第三方调用的删除标签和Wid关联)
func (client *Client) CustomerTagThirdUntagged(request *CustomerTagThirdUntaggedRequest) (response *CustomerTagThirdUntaggedResponse, err error) {
	rpcResponse := CreateCustomerTagThirdUntaggedResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerTagThirdUntaggedRequest struct {
	*api.BaseRequest

	Data CustomerTagThirdUntaggedRequestData `json:"data,omitempty"`
}

type CustomerTagThirdUntaggedResponse struct {
}

func CreateCustomerTagThirdUntaggedRequest() (request *CustomerTagThirdUntaggedRequest) {
	request = &CustomerTagThirdUntaggedRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "customerTag/thirdUntagged", "api")
	request.Method = api.POST
	return
}

func CreateCustomerTagThirdUntaggedResponse() (response *api.BaseResponse[CustomerTagThirdUntaggedResponse]) {
	response = api.CreateResponse[CustomerTagThirdUntaggedResponse](&CustomerTagThirdUntaggedResponse{})
	return
}

type CustomerTagThirdUntaggedRequestData struct {
	Pid           int64  `json:"pid,omitempty"`
	TagId         string `json:"tagId,omitempty"`
	ThirdSourceId int    `json:"thirdSourceId,omitempty"`
}
