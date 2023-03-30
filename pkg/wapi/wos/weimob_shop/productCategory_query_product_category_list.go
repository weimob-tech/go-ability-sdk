package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ProductCategoryQueryProductCategoryList
// @id 1258
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1258?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询品类列表)
func (client *Client) ProductCategoryQueryProductCategoryList(request *ProductCategoryQueryProductCategoryListRequest) (response *ProductCategoryQueryProductCategoryListResponse, err error) {
	rpcResponse := CreateProductCategoryQueryProductCategoryListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ProductCategoryQueryProductCategoryListRequest struct {
	*api.BaseRequest

	Input ProductCategoryQueryProductCategoryListRequestInput `json:"input,omitempty"`
}

type ProductCategoryQueryProductCategoryListResponse struct {
	Output ProductCategoryQueryProductCategoryListResponseOutput `json:"output,omitempty"`
}

func CreateProductCategoryQueryProductCategoryListRequest() (request *ProductCategoryQueryProductCategoryListRequest) {
	request = &ProductCategoryQueryProductCategoryListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "productCategory/queryProductCategoryList", "apigw")
	request.Method = api.POST
	return
}

func CreateProductCategoryQueryProductCategoryListResponse() (response *api.BaseResponse[ProductCategoryQueryProductCategoryListResponse]) {
	response = api.CreateResponse[ProductCategoryQueryProductCategoryListResponse](&ProductCategoryQueryProductCategoryListResponse{})
	return
}

type ProductCategoryQueryProductCategoryListRequestInput struct {
}

type ProductCategoryQueryProductCategoryListResponseOutput struct {
}
