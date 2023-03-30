package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsGetGoodsInfo
// @id 503
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=商品信息)
func (client *Client) GoodsGetGoodsInfo(request *GoodsGetGoodsInfoRequest) (response *GoodsGetGoodsInfoResponse, err error) {
	rpcResponse := CreateGoodsGetGoodsInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsGetGoodsInfoRequest struct {
	*api.BaseRequest

	IsNeedStoreId bool  `json:"isNeedStoreId,omitempty"`
	GoodsId       int64 `json:"goodsId,omitempty"`
}

type GoodsGetGoodsInfoResponse struct {
}

func CreateGoodsGetGoodsInfoRequest() (request *GoodsGetGoodsInfoRequest) {
	request = &GoodsGetGoodsInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "goods/getGoodsInfo", "api")
	request.Method = api.POST
	return
}

func CreateGoodsGetGoodsInfoResponse() (response *api.BaseResponse[GoodsGetGoodsInfoResponse]) {
	response = api.CreateResponse[GoodsGetGoodsInfoResponse](&GoodsGetGoodsInfoResponse{})
	return
}
