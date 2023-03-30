package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ProductEdit
// @id 2657
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=编辑产品)
func (client *Client) ProductEdit(request *ProductEditRequest) (response *ProductEditResponse, err error) {
	rpcResponse := CreateProductEditResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ProductEditRequest struct {
	*api.BaseRequest

	UserWid         int64  `json:"userWid,omitempty"`
	ProductUniqueNo string `json:"productUniqueNo,omitempty"`
}

type ProductEditResponse struct {
}

func CreateProductEditRequest() (request *ProductEditRequest) {
	request = &ProductEditRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "product/edit", "api")
	request.Method = api.POST
	return
}

func CreateProductEditResponse() (response *api.BaseResponse[ProductEditResponse]) {
	response = api.CreateResponse[ProductEditResponse](&ProductEditResponse{})
	return
}
