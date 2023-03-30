package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsUpdateGoodsTitle
// @id 374
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新商品名称)
func (client *Client) GoodsUpdateGoodsTitle(request *GoodsUpdateGoodsTitleRequest) (response *GoodsUpdateGoodsTitleResponse, err error) {
	rpcResponse := CreateGoodsUpdateGoodsTitleResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsUpdateGoodsTitleRequest struct {
	*api.BaseRequest

	GoodsId int64  `json:"goodsId,omitempty"`
	Title   string `json:"title,omitempty"`
}

type GoodsUpdateGoodsTitleResponse struct {
}

func CreateGoodsUpdateGoodsTitleRequest() (request *GoodsUpdateGoodsTitleRequest) {
	request = &GoodsUpdateGoodsTitleRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goods/updateGoodsTitle", "api")
	request.Method = api.POST
	return
}

func CreateGoodsUpdateGoodsTitleResponse() (response *api.BaseResponse[GoodsUpdateGoodsTitleResponse]) {
	response = api.CreateResponse[GoodsUpdateGoodsTitleResponse](&GoodsUpdateGoodsTitleResponse{})
	return
}
