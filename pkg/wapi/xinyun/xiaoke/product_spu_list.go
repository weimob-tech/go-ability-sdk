package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ProductSpuList
// @id 3444
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取产品列表（SPU）)
func (client *Client) ProductSpuList(request *ProductSpuListRequest) (response *ProductSpuListResponse, err error) {
	rpcResponse := CreateProductSpuListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ProductSpuListRequest struct {
	*api.BaseRequest
}

type ProductSpuListResponse struct {
}

func CreateProductSpuListRequest() (request *ProductSpuListRequest) {
	request = &ProductSpuListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "product/spuList", "api")
	request.Method = api.POST
	return
}

func CreateProductSpuListResponse() (response *api.BaseResponse[ProductSpuListResponse]) {
	response = api.CreateResponse[ProductSpuListResponse](&ProductSpuListResponse{})
	return
}
