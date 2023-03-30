package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerTagThirdTagged
// @id 1827
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=给第三方调用的建立标签和Wid关联)
func (client *Client) CustomerTagThirdTagged(request *CustomerTagThirdTaggedRequest) (response *CustomerTagThirdTaggedResponse, err error) {
	rpcResponse := CreateCustomerTagThirdTaggedResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerTagThirdTaggedRequest struct {
	*api.BaseRequest

	Data CustomerTagThirdTaggedRequestData `json:"data,omitempty"`
}

type CustomerTagThirdTaggedResponse struct {
}

func CreateCustomerTagThirdTaggedRequest() (request *CustomerTagThirdTaggedRequest) {
	request = &CustomerTagThirdTaggedRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "customerTag/thirdTagged", "api")
	request.Method = api.POST
	return
}

func CreateCustomerTagThirdTaggedResponse() (response *api.BaseResponse[CustomerTagThirdTaggedResponse]) {
	response = api.CreateResponse[CustomerTagThirdTaggedResponse](&CustomerTagThirdTaggedResponse{})
	return
}

type CustomerTagThirdTaggedRequestData struct {
	Pid           int64  `json:"pid,omitempty"`
	TagIdList     string `json:"tagIdList,omitempty"`
	MobileList    string `json:"mobileList,omitempty"`
	ThirdSourceId int    `json:"thirdSourceId,omitempty"`
}
