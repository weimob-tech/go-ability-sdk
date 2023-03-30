package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembercardGrowingcreditGetList
// @id 1767
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1767?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询成长值明细列表)
func (client *Client) MembercardGrowingcreditGetList(request *MembercardGrowingcreditGetListRequest) (response *MembercardGrowingcreditGetListResponse, err error) {
	rpcResponse := CreateMembercardGrowingcreditGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembercardGrowingcreditGetListRequest struct {
	*api.BaseRequest

	PageReqDTO       MembercardGrowingcreditGetListRequestPageReqDTO `json:"pageReqDTO,omitempty"`
	Vid              int64                                           `json:"vid,omitempty"`
	MembershipPlanId int64                                           `json:"membershipPlanId,omitempty"`
	Wid              int64                                           `json:"wid,omitempty"`
	EndTime          int64                                           `json:"endTime,omitempty"`
	BeginTime        int64                                           `json:"beginTime,omitempty"`
	Status           int64                                           `json:"status,omitempty"`
}

type MembercardGrowingcreditGetListResponse struct {
	GrowingcreditList []MembercardGrowingcreditGetListResponseGrowingcreditList `json:"growingcreditList,omitempty"`
	PageRespDTO       MembercardGrowingcreditGetListResponsePageRespDTO         `json:"PageRespDTO,omitempty"`
}

func CreateMembercardGrowingcreditGetListRequest() (request *MembercardGrowingcreditGetListRequest) {
	request = &MembercardGrowingcreditGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "membercard/growingcredit/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateMembercardGrowingcreditGetListResponse() (response *api.BaseResponse[MembercardGrowingcreditGetListResponse]) {
	response = api.CreateResponse[MembercardGrowingcreditGetListResponse](&MembercardGrowingcreditGetListResponse{})
	return
}

type MembercardGrowingcreditGetListRequestPageReqDTO struct {
	PageIndex int64 `json:"pageIndex,omitempty"`
	PageSize  int64 `json:"pageSize,omitempty"`
}

type MembercardGrowingcreditGetListResponseGrowingcreditList struct {
	GrowthBillId     string `json:"growthBillId,omitempty"`
	Vid              int64  `json:"vid,omitempty"`
	MembershipPlanId int64  `json:"membershipPlanId,omitempty"`
	Wid              int64  `json:"wid,omitempty"`
	ChangeValue      int64  `json:"changeValue,omitempty"`
	Reason           string `json:"reason,omitempty"`
	AvailableGrowth  int64  `json:"availableGrowth,omitempty"`
	ChangeTime       int64  `json:"changeTime,omitempty"`
}

type MembercardGrowingcreditGetListResponsePageRespDTO struct {
	TotalCount int64 `json:"totalCount,omitempty"`
	TotalPages int64 `json:"totalPages,omitempty"`
}
