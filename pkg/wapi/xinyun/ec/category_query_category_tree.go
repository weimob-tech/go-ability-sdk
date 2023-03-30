package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CategoryQueryCategoryTree
// @id 584
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取一级类目)
func (client *Client) CategoryQueryCategoryTree(request *CategoryQueryCategoryTreeRequest) (response *CategoryQueryCategoryTreeResponse, err error) {
	rpcResponse := CreateCategoryQueryCategoryTreeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CategoryQueryCategoryTreeRequest struct {
	*api.BaseRequest
}

type CategoryQueryCategoryTreeResponse struct {
	CategoryList []CategoryQueryCategoryTreeResponseCategoryList `json:"categoryList,omitempty"`
}

func CreateCategoryQueryCategoryTreeRequest() (request *CategoryQueryCategoryTreeRequest) {
	request = &CategoryQueryCategoryTreeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "category/queryCategoryTree", "api")
	request.Method = api.POST
	return
}

func CreateCategoryQueryCategoryTreeResponse() (response *api.BaseResponse[CategoryQueryCategoryTreeResponse]) {
	response = api.CreateResponse[CategoryQueryCategoryTreeResponse](&CategoryQueryCategoryTreeResponse{})
	return
}

type CategoryQueryCategoryTreeResponseCategoryList struct {
	CategoryId int64  `json:"categoryId,omitempty"`
	Title      string `json:"title,omitempty"`
}
