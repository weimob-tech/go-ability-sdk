package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberUpdateMemberCodeByTemplate
// @id 1451
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新会员卡号)
func (client *Client) MemberUpdateMemberCodeByTemplateV2_0(request *MemberUpdateMemberCodeByTemplateRequestV2_0) (response *MemberUpdateMemberCodeByTemplateResponseV2_0, err error) {
	rpcResponse := CreateMemberUpdateMemberCodeByTemplateResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberUpdateMemberCodeByTemplateRequestV2_0 struct {
	*api.BaseRequest

	MemberCardTemplateId int64  `json:"memberCardTemplateId,omitempty"`
	Wid                  int64  `json:"wid,omitempty"`
	OldCode              string `json:"oldCode,omitempty"`
	NewCode              string `json:"newCode,omitempty"`
	Reason               int64  `json:"reason,omitempty"`
	RequestId            int64  `json:"requestId,omitempty"`
}

type MemberUpdateMemberCodeByTemplateResponseV2_0 struct {
}

func CreateMemberUpdateMemberCodeByTemplateRequestV2_0() (request *MemberUpdateMemberCodeByTemplateRequestV2_0) {
	request = &MemberUpdateMemberCodeByTemplateRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "2_0", "member/updateMemberCodeByTemplate", "api")
	request.Method = api.POST
	return
}

func CreateMemberUpdateMemberCodeByTemplateResponseV2_0() (response *api.BaseResponse[MemberUpdateMemberCodeByTemplateResponseV2_0]) {
	response = api.CreateResponse[MemberUpdateMemberCodeByTemplateResponseV2_0](&MemberUpdateMemberCodeByTemplateResponseV2_0{})
	return
}
