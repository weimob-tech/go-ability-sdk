package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembercardGetList
// @id 1661
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1661?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取会员卡列表)
func (client *Client) MembercardGetList(request *MembercardGetListRequest) (response *MembercardGetListResponse, err error) {
	rpcResponse := CreateMembercardGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembercardGetListRequest struct {
	*api.BaseRequest

	Wid              int64   `json:"wid,omitempty"`
	MembershipPlanId int64   `json:"membershipPlanId,omitempty"`
	CustomCardNoList []int64 `json:"customCardNoList,omitempty"`
	QueryType        []int64 `json:"queryType,omitempty"`
}

type MembercardGetListResponse struct {
	ObtainTime            int64  `json:"obtainTime,omitempty"`
	CustomCardNo          string `json:"customCardNo,omitempty"`
	EffectiveTime         int64  `json:"effectiveTime,omitempty"`
	MembershipCardChannel int64  `json:"membershipCardChannel,omitempty"`
	GrowthValue           int64  `json:"growthValue,omitempty"`
	MembershipPlanId      int64  `json:"membershipPlanId,omitempty"`
	ExpireTime            int64  `json:"expireTime,omitempty"`
	MembershipCardScene   int64  `json:"membershipCardScene,omitempty"`
	LevelId               int64  `json:"levelId,omitempty"`
	Status                int64  `json:"status,omitempty"`
	MembershipCardSceneId int64  `json:"membershipCardSceneId,omitempty"`
	Wid                   int64  `json:"wid,omitempty"`
	Vid                   int64  `json:"vid,omitempty"`
}

func CreateMembercardGetListRequest() (request *MembercardGetListRequest) {
	request = &MembercardGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "membercard/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateMembercardGetListResponse() (response *api.BaseResponse[MembercardGetListResponse]) {
	response = api.CreateResponse[MembercardGetListResponse](&MembercardGetListResponse{})
	return
}
