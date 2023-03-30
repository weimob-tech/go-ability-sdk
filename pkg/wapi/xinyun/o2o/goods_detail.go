package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsDetail
// @id 62
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询菜品详情)
func (client *Client) GoodsDetail(request *GoodsDetailRequest) (response *GoodsDetailResponse, err error) {
	rpcResponse := CreateGoodsDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsDetailRequest struct {
	*api.BaseRequest

	GoodsId    int64 `json:"goodsId,omitempty"`
	MerchantId int64 `json:"merchantId,omitempty"`
	StoreId    int64 `json:"storeId,omitempty"`
	SkuId      int64 `json:"skuId,omitempty"`
}

type GoodsDetailResponse struct {
}

func CreateGoodsDetailRequest() (request *GoodsDetailRequest) {
	request = &GoodsDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goods/detail", "api")
	request.Method = api.POST
	return
}

func CreateGoodsDetailResponse() (response *api.BaseResponse[GoodsDetailResponse]) {
	response = api.CreateResponse[GoodsDetailResponse](&GoodsDetailResponse{})
	return
}
