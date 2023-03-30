package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsPriceUpdate
// @id 1638
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1638?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新商品价格)
func (client *Client) GoodsPriceUpdate(request *GoodsPriceUpdateRequest) (response *GoodsPriceUpdateResponse, err error) {
	rpcResponse := CreateGoodsPriceUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsPriceUpdateRequest struct {
	*api.BaseRequest

	SkuList   []GoodsPriceUpdateRequestSkuList `json:"skuList,omitempty"`
	BasicInfo GoodsPriceUpdateRequestBasicInfo `json:"basicInfo,omitempty"`
	GoodsId   int64                            `json:"goodsId,omitempty"`
}

type GoodsPriceUpdateResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateGoodsPriceUpdateRequest() (request *GoodsPriceUpdateRequest) {
	request = &GoodsPriceUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/price/update", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsPriceUpdateResponse() (response *api.BaseResponse[GoodsPriceUpdateResponse]) {
	response = api.CreateResponse[GoodsPriceUpdateResponse](&GoodsPriceUpdateResponse{})
	return
}

type GoodsPriceUpdateRequestSkuList struct {
	SkuId              int64 `json:"skuId,omitempty"`
	SalePrice          int64 `json:"salePrice,omitempty"`
	MarketPrice        int64 `json:"marketPrice,omitempty"`
	CostPrice          int64 `json:"costPrice,omitempty"`
	AdviseSalePriceMin int64 `json:"adviseSalePriceMin,omitempty"`
	AdviseSalePriceMax int64 `json:"adviseSalePriceMax,omitempty"`
}

type GoodsPriceUpdateRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
