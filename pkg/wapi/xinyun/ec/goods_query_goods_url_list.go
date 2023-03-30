package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsQueryGoodsUrlList
// @id 988
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=根据商品id列表查询h5链接)
func (client *Client) GoodsQueryGoodsUrlList(request *GoodsQueryGoodsUrlListRequest) (response *GoodsQueryGoodsUrlListResponse, err error) {
	rpcResponse := CreateGoodsQueryGoodsUrlListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsQueryGoodsUrlListRequest struct {
	*api.BaseRequest

	StoreId     int64   `json:"storeId,omitempty"`
	GoodsIdList []int64 `json:"goodsIdList,omitempty"`
}

type GoodsQueryGoodsUrlListResponse struct {
}

func CreateGoodsQueryGoodsUrlListRequest() (request *GoodsQueryGoodsUrlListRequest) {
	request = &GoodsQueryGoodsUrlListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goods/queryGoodsUrlList", "api")
	request.Method = api.POST
	return
}

func CreateGoodsQueryGoodsUrlListResponse() (response *api.BaseResponse[GoodsQueryGoodsUrlListResponse]) {
	response = api.CreateResponse[GoodsQueryGoodsUrlListResponse](&GoodsQueryGoodsUrlListResponse{})
	return
}
