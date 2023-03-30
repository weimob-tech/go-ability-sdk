package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PointConsume
// @id 1676
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1676?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=消耗积分)
func (client *Client) PointConsume(request *PointConsumeRequest) (response *PointConsumeResponse, err error) {
	rpcResponse := CreatePointConsumeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PointConsumeRequest struct {
	*api.BaseRequest

	RequestId    string `json:"requestId,omitempty"`
	RequestType  string `json:"requestType,omitempty"`
	TotalAmount  string `json:"totalAmount,omitempty"`
	PointAmount  string `json:"pointAmount,omitempty"`
	OutTransNo   string `json:"outTransNo,omitempty"`
	OutTransType string `json:"outTransType,omitempty"`
	Point        string `json:"point,omitempty"`
	ChangeType   string `json:"changeType,omitempty"`
	OccurVid     int64  `json:"occurVid,omitempty"`
	Vid          int64  `json:"vid,omitempty"`
	OperatorWid  int64  `json:"operatorWid,omitempty"`
	UserType     int64  `json:"userType,omitempty"`
	ChannelType  int64  `json:"channelType,omitempty"`
	PointPlanId  int64  `json:"pointPlanId,omitempty"`
	Wid          int64  `json:"wid,omitempty"`
}

type PointConsumeResponse struct {
	BosId       int64  `json:"bosId,omitempty"`
	Wid         int64  `json:"wid,omitempty"`
	Point       string `json:"point,omitempty"`
	ChangeType  string `json:"changeType,omitempty"`
	UserTransId int64  `json:"userTransId,omitempty"`
	OutTransNo  string `json:"outTransNo,omitempty"`
}

func CreatePointConsumeRequest() (request *PointConsumeRequest) {
	request = &PointConsumeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "point/consume", "apigw")
	request.Method = api.POST
	return
}

func CreatePointConsumeResponse() (response *api.BaseResponse[PointConsumeResponse]) {
	response = api.CreateResponse[PointConsumeResponse](&PointConsumeResponse{})
	return
}
