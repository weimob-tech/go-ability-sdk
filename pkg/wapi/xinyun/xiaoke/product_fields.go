package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ProductFields
// @id 2673
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询产品字段列表)
func (client *Client) ProductFields(request *ProductFieldsRequest) (response *ProductFieldsResponse, err error) {
	rpcResponse := CreateProductFieldsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ProductFieldsRequest struct {
	*api.BaseRequest

	UserWid int64 `json:"userWid,omitempty"`
}

type ProductFieldsResponse struct {
}

func CreateProductFieldsRequest() (request *ProductFieldsRequest) {
	request = &ProductFieldsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "product/fields", "api")
	request.Method = api.POST
	return
}

func CreateProductFieldsResponse() (response *api.BaseResponse[ProductFieldsResponse]) {
	response = api.CreateResponse[ProductFieldsResponse](&ProductFieldsResponse{})
	return
}
