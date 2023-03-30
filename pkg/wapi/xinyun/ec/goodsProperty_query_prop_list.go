package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsPropertyQueryPropList
// @id 1169
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询商品属性列表)
func (client *Client) GoodsPropertyQueryPropList(request *GoodsPropertyQueryPropListRequest) (response *GoodsPropertyQueryPropListResponse, err error) {
	rpcResponse := CreateGoodsPropertyQueryPropListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsPropertyQueryPropListRequest struct {
	*api.BaseRequest

	QueryValue bool `json:"queryValue,omitempty"`
}

type GoodsPropertyQueryPropListResponse struct {
}

func CreateGoodsPropertyQueryPropListRequest() (request *GoodsPropertyQueryPropListRequest) {
	request = &GoodsPropertyQueryPropListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goodsProperty/queryPropList", "api")
	request.Method = api.POST
	return
}

func CreateGoodsPropertyQueryPropListResponse() (response *api.BaseResponse[GoodsPropertyQueryPropListResponse]) {
	response = api.CreateResponse[GoodsPropertyQueryPropListResponse](&GoodsPropertyQueryPropListResponse{})
	return
}
