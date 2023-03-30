package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CategoryGet
// @id 6
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取商品类目)
func (client *Client) CategoryGet(request *CategoryGetRequest) (response *CategoryGetResponse, err error) {
	rpcResponse := CreateCategoryGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CategoryGetRequest struct {
	*api.BaseRequest

	CategoryPid string `json:"category_pid,omitempty"`
}

type CategoryGetResponse struct {
}

func CreateCategoryGetRequest() (request *CategoryGetRequest) {
	request = &CategoryGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "Category/Get", "api")
	request.Method = api.POST
	return
}

func CreateCategoryGetResponse() (response *api.BaseResponse[CategoryGetResponse]) {
	response = api.CreateResponse[CategoryGetResponse](&CategoryGetResponse{})
	return
}
