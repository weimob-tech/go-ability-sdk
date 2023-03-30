package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerTagCreateThirdTag
// @id 1826
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=给第三方调用的创建标签)
func (client *Client) CustomerTagCreateThirdTag(request *CustomerTagCreateThirdTagRequest) (response *CustomerTagCreateThirdTagResponse, err error) {
	rpcResponse := CreateCustomerTagCreateThirdTagResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerTagCreateThirdTagRequest struct {
	*api.BaseRequest

	Request CustomerTagCreateThirdTagRequestRequest `json:"request,omitempty"`
}

type CustomerTagCreateThirdTagResponse struct {
}

func CreateCustomerTagCreateThirdTagRequest() (request *CustomerTagCreateThirdTagRequest) {
	request = &CustomerTagCreateThirdTagRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "customerTag/createThirdTag", "api")
	request.Method = api.POST
	return
}

func CreateCustomerTagCreateThirdTagResponse() (response *api.BaseResponse[CustomerTagCreateThirdTagResponse]) {
	response = api.CreateResponse[CustomerTagCreateThirdTagResponse](&CustomerTagCreateThirdTagResponse{})
	return
}

type CustomerTagCreateThirdTagRequestRequest struct {
	Data CustomerTagCreateThirdTagRequestData `json:"data,omitempty"`
	Pid  int64                                `json:"pid,omitempty"`
	Bid  int64                                `json:"bid,omitempty"`
}

type CustomerTagCreateThirdTagRequestData struct {
	Pid           int64  `json:"pid,omitempty"`
	TagName       string `json:"tagName,omitempty"`
	TagId         string `json:"tagId,omitempty"`
	ThirdSourceId int    `json:"thirdSourceId,omitempty"`
	TagType       string `json:"tagType,omitempty"`
}
