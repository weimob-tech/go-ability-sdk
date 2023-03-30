package weimob_shopexpress

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsStockGet
// @id 1966
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1966?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取商品库存)
func (client *Client) GoodsStockGet(request *GoodsStockGetRequest) (response *GoodsStockGetResponse, err error) {
	rpcResponse := CreateGoodsStockGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsStockGetRequest struct {
	*api.BaseRequest

	GoodsCode string `json:"goodsCode,omitempty"`
	SkuCode   string `json:"skuCode,omitempty"`
}

type GoodsStockGetResponse struct {
	GoodsCode  string `json:"goodsCode,omitempty"`
	SkuCode    string `json:"skuCode,omitempty"`
	StockCount int64  `json:"stockCount,omitempty"`
}

func CreateGoodsStockGetRequest() (request *GoodsStockGetRequest) {
	request = &GoodsStockGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shopexpress", "v2.0", "goods/stock/get", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsStockGetResponse() (response *api.BaseResponse[GoodsStockGetResponse]) {
	response = api.CreateResponse[GoodsStockGetResponse](&GoodsStockGetResponse{})
	return
}
