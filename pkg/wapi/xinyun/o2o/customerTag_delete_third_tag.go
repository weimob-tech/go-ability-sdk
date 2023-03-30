package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerTagDeleteThirdTag
// @id 1829
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=给第三方调用的删除标签)
func (client *Client) CustomerTagDeleteThirdTag(request *CustomerTagDeleteThirdTagRequest) (response *CustomerTagDeleteThirdTagResponse, err error) {
	rpcResponse := CreateCustomerTagDeleteThirdTagResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerTagDeleteThirdTagRequest struct {
	*api.BaseRequest

	Data CustomerTagDeleteThirdTagRequestData `json:"data,omitempty"`
}

type CustomerTagDeleteThirdTagResponse struct {
}

func CreateCustomerTagDeleteThirdTagRequest() (request *CustomerTagDeleteThirdTagRequest) {
	request = &CustomerTagDeleteThirdTagRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "customerTag/deleteThirdTag", "api")
	request.Method = api.POST
	return
}

func CreateCustomerTagDeleteThirdTagResponse() (response *api.BaseResponse[CustomerTagDeleteThirdTagResponse]) {
	response = api.CreateResponse[CustomerTagDeleteThirdTagResponse](&CustomerTagDeleteThirdTagResponse{})
	return
}

type CustomerTagDeleteThirdTagRequestData struct {
	Pid           int64  `json:"pid,omitempty"`
	TagId         string `json:"tagId,omitempty"`
	ThirdSourceId int    `json:"thirdSourceId,omitempty"`
}
