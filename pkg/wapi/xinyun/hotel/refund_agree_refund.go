package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RefundAgreeRefund
// @id 2337
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=同意售后)
func (client *Client) RefundAgreeRefund(request *RefundAgreeRefundRequest) (response *RefundAgreeRefundResponse, err error) {
	rpcResponse := CreateRefundAgreeRefundResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RefundAgreeRefundRequest struct {
	*api.BaseRequest

	RefundNo     string `json:"refundNo,omitempty"`
	OperatorName string `json:"operatorName,omitempty"`
}

type RefundAgreeRefundResponse struct {
}

func CreateRefundAgreeRefundRequest() (request *RefundAgreeRefundRequest) {
	request = &RefundAgreeRefundRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "refund/agreeRefund", "api")
	request.Method = api.POST
	return
}

func CreateRefundAgreeRefundResponse() (response *api.BaseResponse[RefundAgreeRefundResponse]) {
	response = api.CreateResponse[RefundAgreeRefundResponse](&RefundAgreeRefundResponse{})
	return
}
