package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsClassifyAppendClassify
// @id 1342
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量修改商品分组)
func (client *Client) GoodsClassifyAppendClassify(request *GoodsClassifyAppendClassifyRequest) (response *GoodsClassifyAppendClassifyResponse, err error) {
	rpcResponse := CreateGoodsClassifyAppendClassifyResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsClassifyAppendClassifyRequest struct {
	*api.BaseRequest
}

type GoodsClassifyAppendClassifyResponse struct {
}

func CreateGoodsClassifyAppendClassifyRequest() (request *GoodsClassifyAppendClassifyRequest) {
	request = &GoodsClassifyAppendClassifyRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goodsClassify/appendClassify", "api")
	request.Method = api.POST
	return
}

func CreateGoodsClassifyAppendClassifyResponse() (response *api.BaseResponse[GoodsClassifyAppendClassifyResponse]) {
	response = api.CreateResponse[GoodsClassifyAppendClassifyResponse](&GoodsClassifyAppendClassifyResponse{})
	return
}
