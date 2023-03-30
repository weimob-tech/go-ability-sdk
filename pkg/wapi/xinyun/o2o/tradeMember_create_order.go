package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TradeMemberCreateOrder
// @id 1898
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=储值接口)
func (client *Client) TradeMemberCreateOrder(request *TradeMemberCreateOrderRequest) (response *TradeMemberCreateOrderResponse, err error) {
	rpcResponse := CreateTradeMemberCreateOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TradeMemberCreateOrderRequest struct {
	*api.BaseRequest

	StoreId     int    `json:"storeId,omitempty"`
	ConsumerWid int    `json:"consumerWid,omitempty"`
	OperateName string `json:"operateName,omitempty"`
	OperateId   int    `json:"operateId,omitempty"`
	Amount      int64  `json:"amount,omitempty"`
	Comment     string `json:"comment,omitempty"`
	PayWay      int    `json:"payWay,omitempty"`
	ThirdNo     string `json:"thirdNo,omitempty"`
}

type TradeMemberCreateOrderResponse struct {
}

func CreateTradeMemberCreateOrderRequest() (request *TradeMemberCreateOrderRequest) {
	request = &TradeMemberCreateOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "tradeMember/createOrder", "api")
	request.Method = api.POST
	return
}

func CreateTradeMemberCreateOrderResponse() (response *api.BaseResponse[TradeMemberCreateOrderResponse]) {
	response = api.CreateResponse[TradeMemberCreateOrderResponse](&TradeMemberCreateOrderResponse{})
	return
}
