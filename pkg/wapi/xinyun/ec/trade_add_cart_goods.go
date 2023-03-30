package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TradeAddCartGoods
// @id 1824
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=加入购物车)
func (client *Client) TradeAddCartGoods(request *TradeAddCartGoodsRequest) (response *TradeAddCartGoodsResponse, err error) {
	rpcResponse := CreateTradeAddCartGoodsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TradeAddCartGoodsRequest struct {
	*api.BaseRequest

	Params    TradeAddCartGoodsRequestParams `json:"params,omitempty"`
	Sign      string                         `json:"sign,omitempty"`
	Timestamp string                         `json:"timestamp,omitempty"`
}

type TradeAddCartGoodsResponse struct {
}

func CreateTradeAddCartGoodsRequest() (request *TradeAddCartGoodsRequest) {
	request = &TradeAddCartGoodsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "trade/addCartGoods", "api")
	request.Method = api.POST
	return
}

func CreateTradeAddCartGoodsResponse() (response *api.BaseResponse[TradeAddCartGoodsResponse]) {
	response = api.CreateResponse[TradeAddCartGoodsResponse](&TradeAddCartGoodsResponse{})
	return
}

type TradeAddCartGoodsRequestParams struct {
}
