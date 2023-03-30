package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsQueryBaseGoodsListByGoodsId
// @id 1067
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=根据商品id列表查询基本信息)
func (client *Client) GoodsQueryBaseGoodsListByGoodsId(request *GoodsQueryBaseGoodsListByGoodsIdRequest) (response *GoodsQueryBaseGoodsListByGoodsIdResponse, err error) {
	rpcResponse := CreateGoodsQueryBaseGoodsListByGoodsIdResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsQueryBaseGoodsListByGoodsIdRequest struct {
	*api.BaseRequest

	GoodsIdList []int `json:"goodsIdList,omitempty"`
}

type GoodsQueryBaseGoodsListByGoodsIdResponse struct {
}

func CreateGoodsQueryBaseGoodsListByGoodsIdRequest() (request *GoodsQueryBaseGoodsListByGoodsIdRequest) {
	request = &GoodsQueryBaseGoodsListByGoodsIdRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goods/queryBaseGoodsListByGoodsId", "api")
	request.Method = api.POST
	return
}

func CreateGoodsQueryBaseGoodsListByGoodsIdResponse() (response *api.BaseResponse[GoodsQueryBaseGoodsListByGoodsIdResponse]) {
	response = api.CreateResponse[GoodsQueryBaseGoodsListByGoodsIdResponse](&GoodsQueryBaseGoodsListByGoodsIdResponse{})
	return
}
