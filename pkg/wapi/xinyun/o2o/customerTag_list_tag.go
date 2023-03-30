package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerTagListTag
// @id 1280
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=根据标签名查询手动标签列表)
func (client *Client) CustomerTagListTag(request *CustomerTagListTagRequest) (response *CustomerTagListTagResponse, err error) {
	rpcResponse := CreateCustomerTagListTagResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerTagListTagRequest struct {
	*api.BaseRequest

	PageIndex int    `json:"pageIndex,omitempty"`
	PageSize  int    `json:"pageSize,omitempty"`
	TagName   string `json:"tagName,omitempty"`
}

type CustomerTagListTagResponse struct {
}

func CreateCustomerTagListTagRequest() (request *CustomerTagListTagRequest) {
	request = &CustomerTagListTagRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "customerTag/listTag", "api")
	request.Method = api.POST
	return
}

func CreateCustomerTagListTagResponse() (response *api.BaseResponse[CustomerTagListTagResponse]) {
	response = api.CreateResponse[CustomerTagListTagResponse](&CustomerTagListTagResponse{})
	return
}
