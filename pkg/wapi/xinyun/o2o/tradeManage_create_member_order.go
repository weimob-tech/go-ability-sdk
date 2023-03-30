package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TradeManageCreateMemberOrder
// @id 2104
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=会员储值订单上传订单)
func (client *Client) TradeManageCreateMemberOrder(request *TradeManageCreateMemberOrderRequest) (response *TradeManageCreateMemberOrderResponse, err error) {
	rpcResponse := CreateTradeManageCreateMemberOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TradeManageCreateMemberOrderRequest struct {
	*api.BaseRequest

	StoreId     int64  `json:"storeId,omitempty"`
	ConsumerWid int64  `json:"consumerWid,omitempty"`
	OperateName string `json:"operateName,omitempty"`
	OperateId   int64  `json:"operateId,omitempty"`
	Amount      int64  `json:"amount,omitempty"`
	Comment     string `json:"comment,omitempty"`
	PayWay      int    `json:"payWay,omitempty"`
	ThirdNo     string `json:"thirdNo,omitempty"`
}

type TradeManageCreateMemberOrderResponse struct {
}

func CreateTradeManageCreateMemberOrderRequest() (request *TradeManageCreateMemberOrderRequest) {
	request = &TradeManageCreateMemberOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "tradeManage/createMemberOrder", "api")
	request.Method = api.POST
	return
}

func CreateTradeManageCreateMemberOrderResponse() (response *api.BaseResponse[TradeManageCreateMemberOrderResponse]) {
	response = api.CreateResponse[TradeManageCreateMemberOrderResponse](&TradeManageCreateMemberOrderResponse{})
	return
}
