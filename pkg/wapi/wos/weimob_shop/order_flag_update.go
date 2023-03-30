package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderFlagUpdate
// @id 1816
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1816?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=订单标记)
func (client *Client) OrderFlagUpdate(request *OrderFlagUpdateRequest) (response *OrderFlagUpdateResponse, err error) {
	rpcResponse := CreateOrderFlagUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderFlagUpdateRequest struct {
	*api.BaseRequest

	OrderNos    []int64 `json:"orderNos,omitempty"`
	FlagRank    int64   `json:"flagRank,omitempty"`
	FlagContent string  `json:"flagContent,omitempty"`
}

type OrderFlagUpdateResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateOrderFlagUpdateRequest() (request *OrderFlagUpdateRequest) {
	request = &OrderFlagUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "order/flag/update", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderFlagUpdateResponse() (response *api.BaseResponse[OrderFlagUpdateResponse]) {
	response = api.CreateResponse[OrderFlagUpdateResponse](&OrderFlagUpdateResponse{})
	return
}
