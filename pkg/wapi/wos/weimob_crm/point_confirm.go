package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PointConfirm
// @id 1834
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1834?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=积分锁定确认)
func (client *Client) PointConfirm(request *PointConfirmRequest) (response *PointConfirmResponse, err error) {
	rpcResponse := CreatePointConfirmResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PointConfirmRequest struct {
	*api.BaseRequest

	RequestId   string `json:"requestId,omitempty"`
	RequestType string `json:"requestType,omitempty"`
	UserTransId int64  `json:"userTransId,omitempty"`
	OutTransNo  string `json:"outTransNo,omitempty"`
	Point       string `json:"point,omitempty"`
	Remark      string `json:"remark,omitempty"`
	Wid         int64  `json:"wid,omitempty"`
	Vid         int64  `json:"vid,omitempty"`
	OperatorWid int64  `json:"operatorWid,omitempty"`
	ChannelType int64  `json:"channelType,omitempty"`
	PlanId      int64  `json:"planId,omitempty"`
}

type PointConfirmResponse struct {
}

func CreatePointConfirmRequest() (request *PointConfirmRequest) {
	request = &PointConfirmRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "point/confirm", "apigw")
	request.Method = api.POST
	return
}

func CreatePointConfirmResponse() (response *api.BaseResponse[PointConfirmResponse]) {
	response = api.CreateResponse[PointConfirmResponse](&PointConfirmResponse{})
	return
}
