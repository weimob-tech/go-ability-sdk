package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsCancelRightsOrder
// @id 381
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=拒绝售后)
func (client *Client) RightsCancelRightsOrder(request *RightsCancelRightsOrderRequest) (response *RightsCancelRightsOrderResponse, err error) {
	rpcResponse := CreateRightsCancelRightsOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsCancelRightsOrderRequest struct {
	*api.BaseRequest

	Id            int64  `json:"id,omitempty"`
	CancelType    int    `json:"cancelType,omitempty"`
	RefusedReason string `json:"refusedReason,omitempty"`
}

type RightsCancelRightsOrderResponse struct {
}

func CreateRightsCancelRightsOrderRequest() (request *RightsCancelRightsOrderRequest) {
	request = &RightsCancelRightsOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "rights/cancelRightsOrder", "api")
	request.Method = api.POST
	return
}

func CreateRightsCancelRightsOrderResponse() (response *api.BaseResponse[RightsCancelRightsOrderResponse]) {
	response = api.CreateResponse[RightsCancelRightsOrderResponse](&RightsCancelRightsOrderResponse{})
	return
}
