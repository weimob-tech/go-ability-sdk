package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsUpdateGoodsImage
// @id 376
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新商品图片)
func (client *Client) GoodsUpdateGoodsImage(request *GoodsUpdateGoodsImageRequest) (response *GoodsUpdateGoodsImageResponse, err error) {
	rpcResponse := CreateGoodsUpdateGoodsImageResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsUpdateGoodsImageRequest struct {
	*api.BaseRequest

	GoodsId   int64    `json:"goodsId,omitempty"`
	ImageList []string `json:"imageList,omitempty"`
}

type GoodsUpdateGoodsImageResponse struct {
}

func CreateGoodsUpdateGoodsImageRequest() (request *GoodsUpdateGoodsImageRequest) {
	request = &GoodsUpdateGoodsImageRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goods/updateGoodsImage", "api")
	request.Method = api.POST
	return
}

func CreateGoodsUpdateGoodsImageResponse() (response *api.BaseResponse[GoodsUpdateGoodsImageResponse]) {
	response = api.CreateResponse[GoodsUpdateGoodsImageResponse](&GoodsUpdateGoodsImageResponse{})
	return
}
