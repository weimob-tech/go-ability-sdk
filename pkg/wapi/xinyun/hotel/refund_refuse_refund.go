package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RefundRefuseRefund
// @id 2338
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=拒绝售后)
func (client *Client) RefundRefuseRefund(request *RefundRefuseRefundRequest) (response *RefundRefuseRefundResponse, err error) {
	rpcResponse := CreateRefundRefuseRefundResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RefundRefuseRefundRequest struct {
	*api.BaseRequest

	RefundNo     string `json:"refundNo,omitempty"`
	RefuseInfo   string `json:"refuseInfo,omitempty"`
	OperatorName string `json:"operatorName,omitempty"`
}

type RefundRefuseRefundResponse struct {
}

func CreateRefundRefuseRefundRequest() (request *RefundRefuseRefundRequest) {
	request = &RefundRefuseRefundRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "refund/refuseRefund", "api")
	request.Method = api.POST
	return
}

func CreateRefundRefuseRefundResponse() (response *api.BaseResponse[RefundRefuseRefundResponse]) {
	response = api.CreateResponse[RefundRefuseRefundResponse](&RefundRefuseRefundResponse{})
	return
}
