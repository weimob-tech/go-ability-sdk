package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ProductCreate
// @id 2659
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=创建产品)
func (client *Client) ProductCreate(request *ProductCreateRequest) (response *ProductCreateResponse, err error) {
	rpcResponse := CreateProductCreateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ProductCreateRequest struct {
	*api.BaseRequest

	UserWid int64   `json:"userWid,omitempty"`
	List    []int64 `json:"list,omitempty"`
}

type ProductCreateResponse struct {
}

func CreateProductCreateRequest() (request *ProductCreateRequest) {
	request = &ProductCreateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "product/create", "api")
	request.Method = api.POST
	return
}

func CreateProductCreateResponse() (response *api.BaseResponse[ProductCreateResponse]) {
	response = api.CreateResponse[ProductCreateResponse](&ProductCreateResponse{})
	return
}
