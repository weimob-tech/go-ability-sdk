package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsUpdateCostPrice
// @id 1420
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改商品的成本价)
func (client *Client) GoodsUpdateCostPrice(request *GoodsUpdateCostPriceRequest) (response *GoodsUpdateCostPriceResponse, err error) {
	rpcResponse := CreateGoodsUpdateCostPriceResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsUpdateCostPriceRequest struct {
	*api.BaseRequest

	UpdateGoodsCostPriceVoList []GoodsUpdateCostPriceRequestUpdateGoodsCostPriceVoList `json:"updateGoodsCostPriceVoList,omitempty"`
	StoreId                    int64                                                   `json:"storeId,omitempty"`
}

type GoodsUpdateCostPriceResponse struct {
}

func CreateGoodsUpdateCostPriceRequest() (request *GoodsUpdateCostPriceRequest) {
	request = &GoodsUpdateCostPriceRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goods/updateCostPrice", "api")
	request.Method = api.POST
	return
}

func CreateGoodsUpdateCostPriceResponse() (response *api.BaseResponse[GoodsUpdateCostPriceResponse]) {
	response = api.CreateResponse[GoodsUpdateCostPriceResponse](&GoodsUpdateCostPriceResponse{})
	return
}

type GoodsUpdateCostPriceRequestUpdateGoodsCostPriceVoList struct {
	SkuId     int64   `json:"skuId,omitempty"`
	CostPrice float64 `json:"costPrice,omitempty"`
}
