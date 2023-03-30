package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PointRefund
// @id 1677
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1677?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=回退积分)
func (client *Client) PointRefund(request *PointRefundRequest) (response *PointRefundResponse, err error) {
	rpcResponse := CreatePointRefundResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PointRefundRequest struct {
	*api.BaseRequest

	RefundRequestId   string `json:"refundRequestId,omitempty"`
	RefundRequestType string `json:"refundRequestType,omitempty"`
	Point             string `json:"point,omitempty"`
	UserTransId       int64  `json:"userTransId,omitempty"`
	RequestId         string `json:"requestId,omitempty"`
	RequestType       string `json:"requestType,omitempty"`
	OutTransNo        string `json:"outTransNo,omitempty"`
	ChangeType        string `json:"changeType,omitempty"`
	Remark            string `json:"remark,omitempty"`
	OccurVid          int64  `json:"occurVid,omitempty"`
	Vid               int64  `json:"vid,omitempty"`
	OperatorWid       int64  `json:"operatorWid,omitempty"`
	UserType          int64  `json:"userType,omitempty"`
	ChannelType       int64  `json:"channelType,omitempty"`
	PointPlanId       int64  `json:"pointPlanId,omitempty"`
	Wid               int64  `json:"wid,omitempty"`
}

type PointRefundResponse struct {
	RefundTransId int64 `json:"refundTransId,omitempty"`
}

func CreatePointRefundRequest() (request *PointRefundRequest) {
	request = &PointRefundRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "point/refund", "apigw")
	request.Method = api.POST
	return
}

func CreatePointRefundResponse() (response *api.BaseResponse[PointRefundResponse]) {
	response = api.CreateResponse[PointRefundResponse](&PointRefundResponse{})
	return
}
