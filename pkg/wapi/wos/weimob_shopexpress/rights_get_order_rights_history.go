package weimob_shopexpress

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsGetOrderRightsHistory
// @id 2151
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2151?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=根据订单编号查询售后记录)
func (client *Client) RightsGetOrderRightsHistory(request *RightsGetOrderRightsHistoryRequest) (response *RightsGetOrderRightsHistoryResponse, err error) {
	rpcResponse := CreateRightsGetOrderRightsHistoryResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsGetOrderRightsHistoryRequest struct {
	*api.BaseRequest

	OrderNo string `json:"orderNo,omitempty"`
	Uid     int64  `json:"uid,omitempty"`
}

type RightsGetOrderRightsHistoryResponse struct {
	RightsHistoryList []RightsGetOrderRightsHistoryResponseRightsHistoryList `json:"rightsHistoryList,omitempty"`
}

func CreateRightsGetOrderRightsHistoryRequest() (request *RightsGetOrderRightsHistoryRequest) {
	request = &RightsGetOrderRightsHistoryRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shopexpress", "v2.0", "rights/getOrderRightsHistory", "apigw")
	request.Method = api.POST
	return
}

func CreateRightsGetOrderRightsHistoryResponse() (response *api.BaseResponse[RightsGetOrderRightsHistoryResponse]) {
	response = api.CreateResponse[RightsGetOrderRightsHistoryResponse](&RightsGetOrderRightsHistoryResponse{})
	return
}

type RightsGetOrderRightsHistoryResponseRightsHistoryList struct {
	AftersaleNo  string `json:"aftersaleNo,omitempty"`
	RefundStatus int64  `json:"refundStatus,omitempty"`
	RefundType   int64  `json:"refundType,omitempty"`
	TotalAmount  int64  `json:"totalAmount,omitempty"`
	RealAmount   int64  `json:"realAmount,omitempty"`
	AutoTime     string `json:"autoTime,omitempty"`
	CreateTime   string `json:"createTime,omitempty"`
}
