package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TradeManageQueryEnabledPaymentConfig
// @id 2106
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=财务查询支付配置)
func (client *Client) TradeManageQueryEnabledPaymentConfig(request *TradeManageQueryEnabledPaymentConfigRequest) (response *TradeManageQueryEnabledPaymentConfigResponse, err error) {
	rpcResponse := CreateTradeManageQueryEnabledPaymentConfigResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TradeManageQueryEnabledPaymentConfigRequest struct {
	*api.BaseRequest

	BusinessCode    string `json:"businessCode,omitempty"`
	CooperationId   string `json:"cooperationId,omitempty"`
	CooperationType string `json:"cooperationType,omitempty"`
	ChannelType     int    `json:"channelType,omitempty"`
	OriginalId      string `json:"originalId,omitempty"`
}

type TradeManageQueryEnabledPaymentConfigResponse struct {
}

func CreateTradeManageQueryEnabledPaymentConfigRequest() (request *TradeManageQueryEnabledPaymentConfigRequest) {
	request = &TradeManageQueryEnabledPaymentConfigRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "tradeManage/queryEnabledPaymentConfig", "api")
	request.Method = api.POST
	return
}

func CreateTradeManageQueryEnabledPaymentConfigResponse() (response *api.BaseResponse[TradeManageQueryEnabledPaymentConfigResponse]) {
	response = api.CreateResponse[TradeManageQueryEnabledPaymentConfigResponse](&TradeManageQueryEnabledPaymentConfigResponse{})
	return
}
