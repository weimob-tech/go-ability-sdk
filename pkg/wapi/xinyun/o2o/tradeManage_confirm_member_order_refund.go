package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TradeManageConfirmMemberOrderRefund
// @id 2105
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=会员储值订单确认退款)
func (client *Client) TradeManageConfirmMemberOrderRefund(request *TradeManageConfirmMemberOrderRefundRequest) (response *TradeManageConfirmMemberOrderRefundResponse, err error) {
	rpcResponse := CreateTradeManageConfirmMemberOrderRefundResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TradeManageConfirmMemberOrderRefundRequest struct {
	*api.BaseRequest

	StoreId     int64  `json:"storeId,omitempty"`
	OperateName string `json:"operateName,omitempty"`
	OperateId   int64  `json:"operateId,omitempty"`
	Amount      int64  `json:"amount,omitempty"`
	Comment     string `json:"comment,omitempty"`
	OrderNo     string `json:"orderNo,omitempty"`
}

type TradeManageConfirmMemberOrderRefundResponse struct {
}

func CreateTradeManageConfirmMemberOrderRefundRequest() (request *TradeManageConfirmMemberOrderRefundRequest) {
	request = &TradeManageConfirmMemberOrderRefundRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "tradeManage/confirmMemberOrderRefund", "api")
	request.Method = api.POST
	return
}

func CreateTradeManageConfirmMemberOrderRefundResponse() (response *api.BaseResponse[TradeManageConfirmMemberOrderRefundResponse]) {
	response = api.CreateResponse[TradeManageConfirmMemberOrderRefundResponse](&TradeManageConfirmMemberOrderRefundResponse{})
	return
}
