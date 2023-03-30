package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CategoryQueryChildrenCategory
// @id 351
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取二级类目)
func (client *Client) CategoryQueryChildrenCategory(request *CategoryQueryChildrenCategoryRequest) (response *CategoryQueryChildrenCategoryResponse, err error) {
	rpcResponse := CreateCategoryQueryChildrenCategoryResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CategoryQueryChildrenCategoryRequest struct {
	*api.BaseRequest

	CategoryId int64 `json:"categoryId,omitempty"`
}

type CategoryQueryChildrenCategoryResponse struct {
}

func CreateCategoryQueryChildrenCategoryRequest() (request *CategoryQueryChildrenCategoryRequest) {
	request = &CategoryQueryChildrenCategoryRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "category/queryChildrenCategory", "api")
	request.Method = api.POST
	return
}

func CreateCategoryQueryChildrenCategoryResponse() (response *api.BaseResponse[CategoryQueryChildrenCategoryResponse]) {
	response = api.CreateResponse[CategoryQueryChildrenCategoryResponse](&CategoryQueryChildrenCategoryResponse{})
	return
}
