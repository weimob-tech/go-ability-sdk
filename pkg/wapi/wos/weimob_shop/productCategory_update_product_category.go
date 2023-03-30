package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ProductCategoryUpdateProductCategory
// @id 1259
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1259?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新品类)
func (client *Client) ProductCategoryUpdateProductCategory(request *ProductCategoryUpdateProductCategoryRequest) (response *ProductCategoryUpdateProductCategoryResponse, err error) {
	rpcResponse := CreateProductCategoryUpdateProductCategoryResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ProductCategoryUpdateProductCategoryRequest struct {
	*api.BaseRequest
}

type ProductCategoryUpdateProductCategoryResponse struct {
	OutPut ProductCategoryUpdateProductCategoryResponseOutPut `json:"outPut,omitempty"`
}

func CreateProductCategoryUpdateProductCategoryRequest() (request *ProductCategoryUpdateProductCategoryRequest) {
	request = &ProductCategoryUpdateProductCategoryRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "productCategory/updateProductCategory", "apigw")
	request.Method = api.POST
	return
}

func CreateProductCategoryUpdateProductCategoryResponse() (response *api.BaseResponse[ProductCategoryUpdateProductCategoryResponse]) {
	response = api.CreateResponse[ProductCategoryUpdateProductCategoryResponse](&ProductCategoryUpdateProductCategoryResponse{})
	return
}

type ProductCategoryUpdateProductCategoryResponseOutPut struct {
}
