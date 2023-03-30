package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ProductSkuList
// @id 3443
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取产品列表（SKU）)
func (client *Client) ProductSkuList(request *ProductSkuListRequest) (response *ProductSkuListResponse, err error) {
	rpcResponse := CreateProductSkuListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ProductSkuListRequest struct {
	*api.BaseRequest
}

type ProductSkuListResponse struct {
}

func CreateProductSkuListRequest() (request *ProductSkuListRequest) {
	request = &ProductSkuListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "product/skuList", "api")
	request.Method = api.POST
	return
}

func CreateProductSkuListResponse() (response *api.BaseResponse[ProductSkuListResponse]) {
	response = api.CreateResponse[ProductSkuListResponse](&ProductSkuListResponse{})
	return
}
