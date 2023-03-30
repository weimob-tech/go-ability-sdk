package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsDeleteGoods
// @id 1808
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除商品)
func (client *Client) GoodsDeleteGoods(request *GoodsDeleteGoodsRequest) (response *GoodsDeleteGoodsResponse, err error) {
	rpcResponse := CreateGoodsDeleteGoodsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsDeleteGoodsRequest struct {
	*api.BaseRequest

	GoodsIdList []int `json:"goodsIdList,omitempty"`
}

type GoodsDeleteGoodsResponse struct {
}

func CreateGoodsDeleteGoodsRequest() (request *GoodsDeleteGoodsRequest) {
	request = &GoodsDeleteGoodsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goods/deleteGoods", "api")
	request.Method = api.POST
	return
}

func CreateGoodsDeleteGoodsResponse() (response *api.BaseResponse[GoodsDeleteGoodsResponse]) {
	response = api.CreateResponse[GoodsDeleteGoodsResponse](&GoodsDeleteGoodsResponse{})
	return
}
