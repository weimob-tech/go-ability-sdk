package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembercardPrivilegeGetList
// @id 1507
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1507?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取会员权益列表)
func (client *Client) MembercardPrivilegeGetList(request *MembercardPrivilegeGetListRequest) (response *MembercardPrivilegeGetListResponse, err error) {
	rpcResponse := CreateMembercardPrivilegeGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembercardPrivilegeGetListRequest struct {
	*api.BaseRequest

	Wid            int64   `json:"wid,omitempty"`
	Vid            int64   `json:"vid,omitempty"`
	PrivilegeTypes []int64 `json:"privilegeTypes,omitempty"`
	VidType        int64   `json:"vidType,omitempty"`
}

type MembercardPrivilegeGetListResponse struct {
	PrivilegeList      MembercardPrivilegeGetListResponsePrivilegeList `json:"privilegeList,omitempty"`
	HasValidMemberCard bool                                            `json:"hasValidMemberCard,omitempty"`
}

func CreateMembercardPrivilegeGetListRequest() (request *MembercardPrivilegeGetListRequest) {
	request = &MembercardPrivilegeGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "membercard/privilege/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateMembercardPrivilegeGetListResponse() (response *api.BaseResponse[MembercardPrivilegeGetListResponse]) {
	response = api.CreateResponse[MembercardPrivilegeGetListResponse](&MembercardPrivilegeGetListResponse{})
	return
}

type MembercardPrivilegeGetListResponsePrivilegeList struct {
	RuleList      []MembercardPrivilegeGetListResponseRuleList `json:"ruleList,omitempty"`
	PrivilegeId   int64                                        `json:"privilegeId,omitempty"`
	PrivilegeType int64                                        `json:"privilegeType,omitempty"`
	PrivilegeName string                                       `json:"privilegeName,omitempty"`
}

type MembercardPrivilegeGetListResponseRuleList struct {
	RuleId   int64  `json:"ruleId,omitempty"`
	RuleName string `json:"ruleName,omitempty"`
}
