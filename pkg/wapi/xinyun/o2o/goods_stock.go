package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsStock
// @id 67
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新菜品库存)
func (client *Client) GoodsStock(request *GoodsStockRequest) (response *GoodsStockResponse, err error) {
	rpcResponse := CreateGoodsStockResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsStockRequest struct {
	*api.BaseRequest

	GoodsId    int64 `json:"goodsId,omitempty"`
	StockNum   int   `json:"stockNum,omitempty"`
	MerchantId int64 `json:"merchantId,omitempty"`
	StoreId    int64 `json:"storeId,omitempty"`
}

type GoodsStockResponse struct {
}

func CreateGoodsStockRequest() (request *GoodsStockRequest) {
	request = &GoodsStockRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goods/stock", "api")
	request.Method = api.POST
	return
}

func CreateGoodsStockResponse() (response *api.BaseResponse[GoodsStockResponse]) {
	response = api.CreateResponse[GoodsStockResponse](&GoodsStockResponse{})
	return
}
