package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderUpdate
// @id 2322
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2322?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新订单)
func (client *Client) OrderUpdate(request *OrderUpdateRequest) (response *OrderUpdateResponse, err error) {
	rpcResponse := CreateOrderUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderUpdateRequest struct {
	*api.BaseRequest

	OrderBaseInfo    OrderUpdateRequestOrderBaseInfo      `json:"orderBaseInfo,omitempty"`
	AssociateBizList []OrderUpdateRequestAssociateBizList `json:"associateBizList,omitempty"`
}

type OrderUpdateResponse struct {
	OrderNo    int64  `json:"orderNo,omitempty"`
	OutOrderNo string `json:"outOrderNo,omitempty"`
}

func CreateOrderUpdateRequest() (request *OrderUpdateRequest) {
	request = &OrderUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "order/update", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderUpdateResponse() (response *api.BaseResponse[OrderUpdateResponse]) {
	response = api.CreateResponse[OrderUpdateResponse](&OrderUpdateResponse{})
	return
}

type OrderUpdateRequestOrderBaseInfo struct {
	OutOrderNo   string `json:"outOrderNo,omitempty"`
	ChannelType  int64  `json:"channelType,omitempty"`
	PlatformType int64  `json:"platformType,omitempty"`
}

type OrderUpdateRequestAssociateBizList struct {
	BizValue int64 `json:"bizValue,omitempty"`
}
