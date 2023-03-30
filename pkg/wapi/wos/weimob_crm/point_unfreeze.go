package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PointUnfreeze
// @id 1675
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1675?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=解冻积分)
func (client *Client) PointUnfreeze(request *PointUnfreezeRequest) (response *PointUnfreezeResponse, err error) {
	rpcResponse := CreatePointUnfreezeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PointUnfreezeRequest struct {
	*api.BaseRequest

	RequestId   string `json:"requestId,omitempty"`
	RequestType string `json:"requestType,omitempty"`
	UserTransId int64  `json:"userTransId,omitempty"`
	OutTransNo  string `json:"outTransNo,omitempty"`
	Vid         int64  `json:"vid,omitempty"`
	OperatorWid int64  `json:"operatorWid,omitempty"`
	UserType    int64  `json:"userType,omitempty"`
	ChannelType int64  `json:"channelType,omitempty"`
	PointPlanId int64  `json:"pointPlanId,omitempty"`
	Wid         int64  `json:"wid,omitempty"`
}

type PointUnfreezeResponse struct {
}

func CreatePointUnfreezeRequest() (request *PointUnfreezeRequest) {
	request = &PointUnfreezeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "point/unfreeze", "apigw")
	request.Method = api.POST
	return
}

func CreatePointUnfreezeResponse() (response *api.BaseResponse[PointUnfreezeResponse]) {
	response = api.CreateResponse[PointUnfreezeResponse](&PointUnfreezeResponse{})
	return
}
