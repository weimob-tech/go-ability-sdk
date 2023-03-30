package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PointPlanGetList
// @id 1491
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1491?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取积分方案列表)
func (client *Client) PointPlanGetList(request *PointPlanGetListRequest) (response *PointPlanGetListResponse, err error) {
	rpcResponse := CreatePointPlanGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PointPlanGetListRequest struct {
	*api.BaseRequest

	Vid int64 `json:"vid,omitempty"`
}

type PointPlanGetListResponse struct {
	PointPlanList []PointPlanGetListResponsePointPlanList `json:"pointPlanList,omitempty"`
}

func CreatePointPlanGetListRequest() (request *PointPlanGetListRequest) {
	request = &PointPlanGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "point/plan/getList", "apigw")
	request.Method = api.POST
	return
}

func CreatePointPlanGetListResponse() (response *api.BaseResponse[PointPlanGetListResponse]) {
	response = api.CreateResponse[PointPlanGetListResponse](&PointPlanGetListResponse{})
	return
}

type PointPlanGetListResponsePointPlanList struct {
	SourceVid         int64  `json:"sourceVid,omitempty"`
	PointPlanName     string `json:"pointPlanName,omitempty"`
	PointPlanId       int64  `json:"pointPlanId,omitempty"`
	UpdateTime        string `json:"updateTime,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
}
