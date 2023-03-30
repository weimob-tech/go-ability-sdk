package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsClassifyAddClassify
// @id 358
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=添加分组)
func (client *Client) GoodsClassifyAddClassify(request *GoodsClassifyAddClassifyRequest) (response *GoodsClassifyAddClassifyResponse, err error) {
	rpcResponse := CreateGoodsClassifyAddClassifyResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsClassifyAddClassifyRequest struct {
	*api.BaseRequest

	Title    string `json:"title,omitempty"`
	ParentId int64  `json:"parentId,omitempty"`
	ImageUrl string `json:"imageUrl,omitempty"`
}

type GoodsClassifyAddClassifyResponse struct {
}

func CreateGoodsClassifyAddClassifyRequest() (request *GoodsClassifyAddClassifyRequest) {
	request = &GoodsClassifyAddClassifyRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goodsClassify/addClassify", "api")
	request.Method = api.POST
	return
}

func CreateGoodsClassifyAddClassifyResponse() (response *api.BaseResponse[GoodsClassifyAddClassifyResponse]) {
	response = api.CreateResponse[GoodsClassifyAddClassifyResponse](&GoodsClassifyAddClassifyResponse{})
	return
}
