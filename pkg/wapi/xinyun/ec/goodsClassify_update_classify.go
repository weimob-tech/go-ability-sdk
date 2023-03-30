package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsClassifyUpdateClassify
// @id 359
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改分组)
func (client *Client) GoodsClassifyUpdateClassify(request *GoodsClassifyUpdateClassifyRequest) (response *GoodsClassifyUpdateClassifyResponse, err error) {
	rpcResponse := CreateGoodsClassifyUpdateClassifyResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsClassifyUpdateClassifyRequest struct {
	*api.BaseRequest

	Title      string `json:"title,omitempty"`
	ClassifyId int64  `json:"classifyId,omitempty"`
	ImageUrl   string `json:"imageUrl,omitempty"`
	ParentId   int64  `json:"parentId,omitempty"`
}

type GoodsClassifyUpdateClassifyResponse struct {
}

func CreateGoodsClassifyUpdateClassifyRequest() (request *GoodsClassifyUpdateClassifyRequest) {
	request = &GoodsClassifyUpdateClassifyRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goodsClassify/updateClassify", "api")
	request.Method = api.POST
	return
}

func CreateGoodsClassifyUpdateClassifyResponse() (response *api.BaseResponse[GoodsClassifyUpdateClassifyResponse]) {
	response = api.CreateResponse[GoodsClassifyUpdateClassifyResponse](&GoodsClassifyUpdateClassifyResponse{})
	return
}
