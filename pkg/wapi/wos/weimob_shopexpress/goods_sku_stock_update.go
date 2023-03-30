package weimob_shopexpress

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsSkuStockUpdate
// @id 2109
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2109?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=增量更新商品库存)
func (client *Client) GoodsSkuStockUpdate(request *GoodsSkuStockUpdateRequest) (response *GoodsSkuStockUpdateResponse, err error) {
	rpcResponse := CreateGoodsSkuStockUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsSkuStockUpdateRequest struct {
	*api.BaseRequest

	BizId     string `json:"bizId,omitempty"`
	Count     int64  `json:"count,omitempty"`
	GoodsCode string `json:"goodsCode,omitempty"`
	OpName    string `json:"opName,omitempty"`
	SkuCode   string `json:"skuCode,omitempty"`
}

type GoodsSkuStockUpdateResponse struct {
}

func CreateGoodsSkuStockUpdateRequest() (request *GoodsSkuStockUpdateRequest) {
	request = &GoodsSkuStockUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shopexpress", "v2.0", "goods/sku/stock/update", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsSkuStockUpdateResponse() (response *api.BaseResponse[GoodsSkuStockUpdateResponse]) {
	response = api.CreateResponse[GoodsSkuStockUpdateResponse](&GoodsSkuStockUpdateResponse{})
	return
}
