package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightWayRefund
// @id 2161
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=线路维权)
func (client *Client) RightWayRefund(request *RightWayRefundRequest) (response *RightWayRefundResponse, err error) {
	rpcResponse := CreateRightWayRefundResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightWayRefundRequest struct {
	*api.BaseRequest

	OrderNo      string  `json:"orderNo,omitempty"`
	AllRefund    int     `json:"allRefund,omitempty"`
	RefundsMoney float64 `json:"refundsMoney,omitempty"`
	RightReason  string  `json:"rightReason,omitempty"`
}

type RightWayRefundResponse struct {
}

func CreateRightWayRefundRequest() (request *RightWayRefundRequest) {
	request = &RightWayRefundRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "right/wayRefund", "api")
	request.Method = api.POST
	return
}

func CreateRightWayRefundResponse() (response *api.BaseResponse[RightWayRefundResponse]) {
	response = api.CreateResponse[RightWayRefundResponse](&RightWayRefundResponse{})
	return
}
