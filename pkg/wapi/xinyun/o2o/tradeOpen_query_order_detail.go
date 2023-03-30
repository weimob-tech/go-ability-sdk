package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TradeOpenQueryOrderDetail
// @id 1452
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询订单基础信息)
func (client *Client) TradeOpenQueryOrderDetail(request *TradeOpenQueryOrderDetailRequest) (response *TradeOpenQueryOrderDetailResponse, err error) {
	rpcResponse := CreateTradeOpenQueryOrderDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TradeOpenQueryOrderDetailRequest struct {
	*api.BaseRequest
}

type TradeOpenQueryOrderDetailResponse struct {
}

func CreateTradeOpenQueryOrderDetailRequest() (request *TradeOpenQueryOrderDetailRequest) {
	request = &TradeOpenQueryOrderDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "tradeOpen/queryOrderDetail", "api")
	request.Method = api.POST
	return
}

func CreateTradeOpenQueryOrderDetailResponse() (response *api.BaseResponse[TradeOpenQueryOrderDetailResponse]) {
	response = api.CreateResponse[TradeOpenQueryOrderDetailResponse](&TradeOpenQueryOrderDetailResponse{})
	return
}
