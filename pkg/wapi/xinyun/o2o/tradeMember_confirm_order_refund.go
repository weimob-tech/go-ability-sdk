package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TradeMemberConfirmOrderRefund
// @id 1899
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=储值退款接口)
func (client *Client) TradeMemberConfirmOrderRefund(request *TradeMemberConfirmOrderRefundRequest) (response *TradeMemberConfirmOrderRefundResponse, err error) {
	rpcResponse := CreateTradeMemberConfirmOrderRefundResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TradeMemberConfirmOrderRefundRequest struct {
	*api.BaseRequest

	StoreId     int    `json:"storeId,omitempty"`
	OperateName string `json:"operateName,omitempty"`
	OperateId   int    `json:"operateId,omitempty"`
	Amount      int64  `json:"amount,omitempty"`
	Comment     string `json:"comment,omitempty"`
	ThirdNo     string `json:"thirdNo,omitempty"`
	OrderNo     string `json:"orderNo,omitempty"`
}

type TradeMemberConfirmOrderRefundResponse struct {
}

func CreateTradeMemberConfirmOrderRefundRequest() (request *TradeMemberConfirmOrderRefundRequest) {
	request = &TradeMemberConfirmOrderRefundRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "tradeMember/confirmOrderRefund", "api")
	request.Method = api.POST
	return
}

func CreateTradeMemberConfirmOrderRefundResponse() (response *api.BaseResponse[TradeMemberConfirmOrderRefundResponse]) {
	response = api.CreateResponse[TradeMemberConfirmOrderRefundResponse](&TradeMemberConfirmOrderRefundResponse{})
	return
}
