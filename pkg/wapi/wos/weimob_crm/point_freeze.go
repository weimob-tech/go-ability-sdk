package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PointFreeze
// @id 1674
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1674?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=冻结积分)
func (client *Client) PointFreeze(request *PointFreezeRequest) (response *PointFreezeResponse, err error) {
	rpcResponse := CreatePointFreezeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PointFreezeRequest struct {
	*api.BaseRequest

	ChangeType   string `json:"changeType,omitempty"`
	RequestId    string `json:"requestId,omitempty"`
	RequestType  string `json:"requestType,omitempty"`
	TotalAmount  string `json:"totalAmount,omitempty"`
	PointAmount  string `json:"pointAmount,omitempty"`
	OutTransNo   string `json:"outTransNo,omitempty"`
	OutTransType string `json:"outTransType,omitempty"`
	Point        string `json:"point,omitempty"`
	OccurVid     int64  `json:"occurVid,omitempty"`
	Remark       string `json:"remark,omitempty"`
	Vid          int64  `json:"vid,omitempty"`
	OperatorWid  int64  `json:"operatorWid,omitempty"`
	UserType     int64  `json:"userType,omitempty"`
	ChannelType  string `json:"channelType,omitempty"`
	PointPlanId  int64  `json:"pointPlanId,omitempty"`
	Wid          int64  `json:"wid,omitempty"`
}

type PointFreezeResponse struct {
	BosId       int64  `json:"bosId,omitempty"`
	Wid         int64  `json:"wid,omitempty"`
	Point       string `json:"point,omitempty"`
	ChangeType  string `json:"changeType,omitempty"`
	UserTransId int64  `json:"userTransId,omitempty"`
	OutTransNo  string `json:"outTransNo,omitempty"`
}

func CreatePointFreezeRequest() (request *PointFreezeRequest) {
	request = &PointFreezeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "point/freeze", "apigw")
	request.Method = api.POST
	return
}

func CreatePointFreezeResponse() (response *api.BaseResponse[PointFreezeResponse]) {
	response = api.CreateResponse[PointFreezeResponse](&PointFreezeResponse{})
	return
}
