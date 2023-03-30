package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreQueryCategory
// @id 52
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询门店类目)
func (client *Client) StoreQueryCategory(request *StoreQueryCategoryRequest) (response *StoreQueryCategoryResponse, err error) {
	rpcResponse := CreateStoreQueryCategoryResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreQueryCategoryRequest struct {
	*api.BaseRequest

	Token string `json:"token,omitempty"`
}

type StoreQueryCategoryResponse struct {
}

func CreateStoreQueryCategoryRequest() (request *StoreQueryCategoryRequest) {
	request = &StoreQueryCategoryRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "store/queryCategory", "api")
	request.Method = api.POST
	return
}

func CreateStoreQueryCategoryResponse() (response *api.BaseResponse[StoreQueryCategoryResponse]) {
	response = api.CreateResponse[StoreQueryCategoryResponse](&StoreQueryCategoryResponse{})
	return
}
