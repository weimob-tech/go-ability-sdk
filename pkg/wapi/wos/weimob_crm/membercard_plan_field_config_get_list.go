package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembercardPlanFieldConfigGetList
// @id 1741
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1741?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询会员方案开卡字段)
func (client *Client) MembercardPlanFieldConfigGetList(request *MembercardPlanFieldConfigGetListRequest) (response *MembercardPlanFieldConfigGetListResponse, err error) {
	rpcResponse := CreateMembercardPlanFieldConfigGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembercardPlanFieldConfigGetListRequest struct {
	*api.BaseRequest

	Vid              int64 `json:"vid,omitempty"`
	Channel          int64 `json:"channel,omitempty"`
	MembershipPlanId int64 `json:"membershipPlanId,omitempty"`
}

type MembercardPlanFieldConfigGetListResponse struct {
	Fields []MembercardPlanFieldConfigGetListResponseFields `json:"fields,omitempty"`
}

func CreateMembercardPlanFieldConfigGetListRequest() (request *MembercardPlanFieldConfigGetListRequest) {
	request = &MembercardPlanFieldConfigGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "membercard/plan/field/config/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateMembercardPlanFieldConfigGetListResponse() (response *api.BaseResponse[MembercardPlanFieldConfigGetListResponse]) {
	response = api.CreateResponse[MembercardPlanFieldConfigGetListResponse](&MembercardPlanFieldConfigGetListResponse{})
	return
}

type MembercardPlanFieldConfigGetListResponseFields struct {
	Channel int64  `json:"channel,omitempty"`
	Name    string `json:"name,omitempty"`
	Tips    string `json:"tips,omitempty"`
}
