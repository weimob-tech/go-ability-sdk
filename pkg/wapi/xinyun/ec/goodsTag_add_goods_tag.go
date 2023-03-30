package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsTagAddGoodsTag
// @id 362
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=添加商品标签)
func (client *Client) GoodsTagAddGoodsTag(request *GoodsTagAddGoodsTagRequest) (response *GoodsTagAddGoodsTagResponse, err error) {
	rpcResponse := CreateGoodsTagAddGoodsTagResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsTagAddGoodsTagRequest struct {
	*api.BaseRequest

	Title string `json:"title,omitempty"`
}

type GoodsTagAddGoodsTagResponse struct {
}

func CreateGoodsTagAddGoodsTagRequest() (request *GoodsTagAddGoodsTagRequest) {
	request = &GoodsTagAddGoodsTagRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goodsTag/addGoodsTag", "api")
	request.Method = api.POST
	return
}

func CreateGoodsTagAddGoodsTagResponse() (response *api.BaseResponse[GoodsTagAddGoodsTagResponse]) {
	response = api.CreateResponse[GoodsTagAddGoodsTagResponse](&GoodsTagAddGoodsTagResponse{})
	return
}
