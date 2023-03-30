package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsUpdateGoodsDesc
// @id 375
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新商品文描)
func (client *Client) GoodsUpdateGoodsDesc(request *GoodsUpdateGoodsDescRequest) (response *GoodsUpdateGoodsDescResponse, err error) {
	rpcResponse := CreateGoodsUpdateGoodsDescResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsUpdateGoodsDescRequest struct {
	*api.BaseRequest

	GoodsId   int64  `json:"goodsId,omitempty"`
	GoodsDesc string `json:"goodsDesc,omitempty"`
}

type GoodsUpdateGoodsDescResponse struct {
}

func CreateGoodsUpdateGoodsDescRequest() (request *GoodsUpdateGoodsDescRequest) {
	request = &GoodsUpdateGoodsDescRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goods/updateGoodsDesc", "api")
	request.Method = api.POST
	return
}

func CreateGoodsUpdateGoodsDescResponse() (response *api.BaseResponse[GoodsUpdateGoodsDescResponse]) {
	response = api.CreateResponse[GoodsUpdateGoodsDescResponse](&GoodsUpdateGoodsDescResponse{})
	return
}
