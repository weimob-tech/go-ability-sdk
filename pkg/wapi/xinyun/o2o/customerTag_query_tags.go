package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerTagQueryTags
// @id 1279
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=根据标签id查询标签列表)
func (client *Client) CustomerTagQueryTags(request *CustomerTagQueryTagsRequest) (response *CustomerTagQueryTagsResponse, err error) {
	rpcResponse := CreateCustomerTagQueryTagsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerTagQueryTagsRequest struct {
	*api.BaseRequest

	TagIdList int64 `json:"tagIdList,omitempty"`
}

type CustomerTagQueryTagsResponse struct {
}

func CreateCustomerTagQueryTagsRequest() (request *CustomerTagQueryTagsRequest) {
	request = &CustomerTagQueryTagsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "customerTag/queryTags", "api")
	request.Method = api.POST
	return
}

func CreateCustomerTagQueryTagsResponse() (response *api.BaseResponse[CustomerTagQueryTagsResponse]) {
	response = api.CreateResponse[CustomerTagQueryTagsResponse](&CustomerTagQueryTagsResponse{})
	return
}
