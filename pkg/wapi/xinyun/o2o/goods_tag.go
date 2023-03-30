package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsTag
// @id 74
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=菜品打标签（增加/删除）-按参数列表更新)
func (client *Client) GoodsTag(request *GoodsTagRequest) (response *GoodsTagResponse, err error) {
	rpcResponse := CreateGoodsTagResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsTagRequest struct {
	*api.BaseRequest

	GoodsId    int64 `json:"goodsId,omitempty"`
	GoodsTagId int64 `json:"goodsTagId,omitempty"`
	StoreId    int64 `json:"storeId,omitempty"`
	MerchantId int64 `json:"merchantId,omitempty"`
}

type GoodsTagResponse struct {
}

func CreateGoodsTagRequest() (request *GoodsTagRequest) {
	request = &GoodsTagRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goods/tag", "api")
	request.Method = api.POST
	return
}

func CreateGoodsTagResponse() (response *api.BaseResponse[GoodsTagResponse]) {
	response = api.CreateResponse[GoodsTagResponse](&GoodsTagResponse{})
	return
}
