package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PointGet
// @id 1497
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1497?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询用户积分)
func (client *Client) PointGet(request *PointGetRequest) (response *PointGetResponse, err error) {
	rpcResponse := CreatePointGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PointGetRequest struct {
	*api.BaseRequest

	Vid int64 `json:"vid,omitempty"`
	Wid int64 `json:"wid,omitempty"`
}

type PointGetResponse struct {
	AcctId          int64  `json:"acctId,omitempty"`
	AcctStatus      string `json:"acctStatus,omitempty"`
	CustomerStatus  int64  `json:"customerStatus,omitempty"`
	AvailablePoint  string `json:"availablePoint,omitempty"`
	TotalPoint      string `json:"totalPoint,omitempty"`
	FrozenPoint     string `json:"frozenPoint,omitempty"`
	TotalIssuePoint string `json:"totalIssuePoint,omitempty"`
}

func CreatePointGetRequest() (request *PointGetRequest) {
	request = &PointGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "point/get", "apigw")
	request.Method = api.POST
	return
}

func CreatePointGetResponse() (response *api.BaseResponse[PointGetResponse]) {
	response = api.CreateResponse[PointGetResponse](&PointGetResponse{})
	return
}
