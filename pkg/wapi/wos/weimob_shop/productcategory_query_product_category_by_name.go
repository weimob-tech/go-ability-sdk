package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ProductcategoryQueryProductCategoryByName
// @id 1267
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1267?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=品类名称查询品类信息)
func (client *Client) ProductcategoryQueryProductCategoryByName(request *ProductcategoryQueryProductCategoryByNameRequest) (response *ProductcategoryQueryProductCategoryByNameResponse, err error) {
	rpcResponse := CreateProductcategoryQueryProductCategoryByNameResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ProductcategoryQueryProductCategoryByNameRequest struct {
	*api.BaseRequest

	Input ProductcategoryQueryProductCategoryByNameRequestInput `json:"input,omitempty"`
}

type ProductcategoryQueryProductCategoryByNameResponse struct {
	Output ProductcategoryQueryProductCategoryByNameResponseOutput `json:"output,omitempty"`
}

func CreateProductcategoryQueryProductCategoryByNameRequest() (request *ProductcategoryQueryProductCategoryByNameRequest) {
	request = &ProductcategoryQueryProductCategoryByNameRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "productcategory/queryProductCategoryByName", "apigw")
	request.Method = api.POST
	return
}

func CreateProductcategoryQueryProductCategoryByNameResponse() (response *api.BaseResponse[ProductcategoryQueryProductCategoryByNameResponse]) {
	response = api.CreateResponse[ProductcategoryQueryProductCategoryByNameResponse](&ProductcategoryQueryProductCategoryByNameResponse{})
	return
}

type ProductcategoryQueryProductCategoryByNameRequestInput struct {
}

type ProductcategoryQueryProductCategoryByNameResponseOutput struct {
}
