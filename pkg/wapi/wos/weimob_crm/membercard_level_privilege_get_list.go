package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembercardLevelPrivilegeGetList
// @id 1773
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1773?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询会员卡设置的等级权益列表)
func (client *Client) MembercardLevelPrivilegeGetList(request *MembercardLevelPrivilegeGetListRequest) (response *MembercardLevelPrivilegeGetListResponse, err error) {
	rpcResponse := CreateMembercardLevelPrivilegeGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembercardLevelPrivilegeGetListRequest struct {
	*api.BaseRequest

	MembershipPlanId int64  `json:"membershipPlanId,omitempty"`
	PrivilegeType    string `json:"privilegeType,omitempty"`
}

type MembercardLevelPrivilegeGetListResponse struct {
	LevelPrivilege []MembercardLevelPrivilegeGetListResponseLevelPrivilege `json:"levelPrivilege,omitempty"`
}

func CreateMembercardLevelPrivilegeGetListRequest() (request *MembercardLevelPrivilegeGetListRequest) {
	request = &MembercardLevelPrivilegeGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "membercard/level/privilege/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateMembercardLevelPrivilegeGetListResponse() (response *api.BaseResponse[MembercardLevelPrivilegeGetListResponse]) {
	response = api.CreateResponse[MembercardLevelPrivilegeGetListResponse](&MembercardLevelPrivilegeGetListResponse{})
	return
}

type MembercardLevelPrivilegeGetListResponseLevelPrivilege struct {
	PrivilegeList []MembercardLevelPrivilegeGetListResponsePrivilegeList `json:"privilegeList,omitempty"`
	LevelId       int64                                                  `json:"levelId,omitempty"`
	LevelName     string                                                 `json:"levelName,omitempty"`
}

type MembercardLevelPrivilegeGetListResponsePrivilegeList struct {
	PrivilegeId   int64  `json:"privilegeId,omitempty"`
	PrivilegeName string `json:"privilegeName,omitempty"`
	PrivilegeType string `json:"privilegeType,omitempty"`
	IconType      int64  `json:"iconType,omitempty"`
	PrivilegeIcon string `json:"privilegeIcon,omitempty"`
	RuleId        string `json:"ruleId,omitempty"`
	RuleName      string `json:"ruleName,omitempty"`
}
