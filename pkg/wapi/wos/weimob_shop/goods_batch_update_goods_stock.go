package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsBatchUpdateGoodsStock
// @id 1261
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1261?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=批量更新库存【覆盖】)
func (client *Client) GoodsBatchUpdateGoodsStock(request *GoodsBatchUpdateGoodsStockRequest) (response *GoodsBatchUpdateGoodsStockResponse, err error) {
	rpcResponse := CreateGoodsBatchUpdateGoodsStockResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsBatchUpdateGoodsStockRequest struct {
	*api.BaseRequest

	Input GoodsBatchUpdateGoodsStockRequestInput `json:"input,omitempty"`
}

type GoodsBatchUpdateGoodsStockResponse struct {
	OutPut GoodsBatchUpdateGoodsStockResponseOutPut `json:"outPut,omitempty"`
}

func CreateGoodsBatchUpdateGoodsStockRequest() (request *GoodsBatchUpdateGoodsStockRequest) {
	request = &GoodsBatchUpdateGoodsStockRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/batchUpdateGoodsStock", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsBatchUpdateGoodsStockResponse() (response *api.BaseResponse[GoodsBatchUpdateGoodsStockResponse]) {
	response = api.CreateResponse[GoodsBatchUpdateGoodsStockResponse](&GoodsBatchUpdateGoodsStockResponse{})
	return
}

type GoodsBatchUpdateGoodsStockRequestInput struct {
}

type GoodsBatchUpdateGoodsStockResponseOutPut struct {
}
