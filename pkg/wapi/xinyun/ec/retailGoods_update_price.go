package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RetailGoodsUpdatePrice
// @id 541
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改商品价格（智慧零售）)
func (client *Client) RetailGoodsUpdatePrice(request *RetailGoodsUpdatePriceRequest) (response *RetailGoodsUpdatePriceResponse, err error) {
	rpcResponse := CreateRetailGoodsUpdatePriceResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RetailGoodsUpdatePriceRequest struct {
	*api.BaseRequest

	StoreId                   int64   `json:"storeId,omitempty"`
	GoodsId                   int64   `json:"goodsId,omitempty"`
	SkuList                   []int64 `json:"skuList,omitempty"`
	UpdateStoreSalePriceRange bool    `json:"updateStoreSalePriceRange,omitempty"`
}

type RetailGoodsUpdatePriceResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateRetailGoodsUpdatePriceRequest() (request *RetailGoodsUpdatePriceRequest) {
	request = &RetailGoodsUpdatePriceRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "retailGoods/updatePrice", "api")
	request.Method = api.POST
	return
}

func CreateRetailGoodsUpdatePriceResponse() (response *api.BaseResponse[RetailGoodsUpdatePriceResponse]) {
	response = api.CreateResponse[RetailGoodsUpdatePriceResponse](&RetailGoodsUpdatePriceResponse{})
	return
}
