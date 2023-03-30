package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberGetMemberListByWids
// @id 1449
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量查询客户详情)
func (client *Client) MemberGetMemberListByWids(request *MemberGetMemberListByWidsRequest) (response *MemberGetMemberListByWidsResponse, err error) {
	rpcResponse := CreateMemberGetMemberListByWidsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberGetMemberListByWidsRequest struct {
	*api.BaseRequest

	WidList        []int `json:"widList,omitempty"`
	IsNeedExtInfo  bool  `json:"isNeedExtInfo,omitempty"`
	IsNeedTagsInfo bool  `json:"isNeedTagsInfo,omitempty"`
}

type MemberGetMemberListByWidsResponse struct {
}

func CreateMemberGetMemberListByWidsRequest() (request *MemberGetMemberListByWidsRequest) {
	request = &MemberGetMemberListByWidsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "member/getMemberListByWids", "api")
	request.Method = api.POST
	return
}

func CreateMemberGetMemberListByWidsResponse() (response *api.BaseResponse[MemberGetMemberListByWidsResponse]) {
	response = api.CreateResponse[MemberGetMemberListByWidsResponse](&MemberGetMemberListByWidsResponse{})
	return
}
