package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsTagUpdateGoodsTag
// @id 361
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新商品标签)
func (client *Client) GoodsTagUpdateGoodsTag(request *GoodsTagUpdateGoodsTagRequest) (response *GoodsTagUpdateGoodsTagResponse, err error) {
	rpcResponse := CreateGoodsTagUpdateGoodsTagResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsTagUpdateGoodsTagRequest struct {
	*api.BaseRequest

	Title string `json:"title,omitempty"`
	TagId int64  `json:"tagId,omitempty"`
}

type GoodsTagUpdateGoodsTagResponse struct {
}

func CreateGoodsTagUpdateGoodsTagRequest() (request *GoodsTagUpdateGoodsTagRequest) {
	request = &GoodsTagUpdateGoodsTagRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goodsTag/updateGoodsTag", "api")
	request.Method = api.POST
	return
}

func CreateGoodsTagUpdateGoodsTagResponse() (response *api.BaseResponse[GoodsTagUpdateGoodsTagResponse]) {
	response = api.CreateResponse[GoodsTagUpdateGoodsTagResponse](&GoodsTagUpdateGoodsTagResponse{})
	return
}
