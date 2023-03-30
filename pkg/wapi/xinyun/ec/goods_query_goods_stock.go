package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsQueryGoodsStock
// @id 1402
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询商品库存-可查占用库存（微商城）)
func (client *Client) GoodsQueryGoodsStock(request *GoodsQueryGoodsStockRequest) (response *GoodsQueryGoodsStockResponse, err error) {
	rpcResponse := CreateGoodsQueryGoodsStockResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsQueryGoodsStockRequest struct {
	*api.BaseRequest

	GoodsId int64 `json:"goodsId,omitempty"`
}

type GoodsQueryGoodsStockResponse struct {
}

func CreateGoodsQueryGoodsStockRequest() (request *GoodsQueryGoodsStockRequest) {
	request = &GoodsQueryGoodsStockRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goods/queryGoodsStock", "api")
	request.Method = api.POST
	return
}

func CreateGoodsQueryGoodsStockResponse() (response *api.BaseResponse[GoodsQueryGoodsStockResponse]) {
	response = api.CreateResponse[GoodsQueryGoodsStockResponse](&GoodsQueryGoodsStockResponse{})
	return
}
