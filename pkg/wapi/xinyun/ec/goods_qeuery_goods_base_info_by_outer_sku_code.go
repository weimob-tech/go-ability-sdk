package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsQeueryGoodsBaseInfoByOuterSkuCode
// @id 1421
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=根据商家编码查询商品信息)
func (client *Client) GoodsQeueryGoodsBaseInfoByOuterSkuCode(request *GoodsQeueryGoodsBaseInfoByOuterSkuCodeRequest) (response *GoodsQeueryGoodsBaseInfoByOuterSkuCodeResponse, err error) {
	rpcResponse := CreateGoodsQeueryGoodsBaseInfoByOuterSkuCodeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsQeueryGoodsBaseInfoByOuterSkuCodeRequest struct {
	*api.BaseRequest

	StoreId      int64 `json:"storeId,omitempty"`
	OuterSkuCode int64 `json:"outerSkuCode,omitempty"`
}

type GoodsQeueryGoodsBaseInfoByOuterSkuCodeResponse struct {
	GoodsId      int64  `json:"goodsId,omitempty"`
	SkuId        int64  `json:"skuId,omitempty"`
	OuterSkuCode string `json:"outerSkuCode,omitempty"`
	StoreId      int64  `json:"storeId,omitempty"`
	Title        string `json:"title,omitempty"`
	ImageUrl     string `json:"imageUrl,omitempty"`
	ValidateId   int64  `json:"validateId,omitempty"`
}

func CreateGoodsQeueryGoodsBaseInfoByOuterSkuCodeRequest() (request *GoodsQeueryGoodsBaseInfoByOuterSkuCodeRequest) {
	request = &GoodsQeueryGoodsBaseInfoByOuterSkuCodeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goods/qeueryGoodsBaseInfoByOuterSkuCode", "api")
	request.Method = api.POST
	return
}

func CreateGoodsQeueryGoodsBaseInfoByOuterSkuCodeResponse() (response *api.BaseResponse[GoodsQeueryGoodsBaseInfoByOuterSkuCodeResponse]) {
	response = api.CreateResponse[GoodsQeueryGoodsBaseInfoByOuterSkuCodeResponse](&GoodsQeueryGoodsBaseInfoByOuterSkuCodeResponse{})
	return
}
