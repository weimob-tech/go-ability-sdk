package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsClassifyDeleteClassify
// @id 1807
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除商品分组)
func (client *Client) GoodsClassifyDeleteClassify(request *GoodsClassifyDeleteClassifyRequest) (response *GoodsClassifyDeleteClassifyResponse, err error) {
	rpcResponse := CreateGoodsClassifyDeleteClassifyResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsClassifyDeleteClassifyRequest struct {
	*api.BaseRequest

	ClassifyId int64 `json:"classifyId,omitempty"`
}

type GoodsClassifyDeleteClassifyResponse struct {
}

func CreateGoodsClassifyDeleteClassifyRequest() (request *GoodsClassifyDeleteClassifyRequest) {
	request = &GoodsClassifyDeleteClassifyRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goodsClassify/deleteClassify", "api")
	request.Method = api.POST
	return
}

func CreateGoodsClassifyDeleteClassifyResponse() (response *api.BaseResponse[GoodsClassifyDeleteClassifyResponse]) {
	response = api.CreateResponse[GoodsClassifyDeleteClassifyResponse](&GoodsClassifyDeleteClassifyResponse{})
	return
}
