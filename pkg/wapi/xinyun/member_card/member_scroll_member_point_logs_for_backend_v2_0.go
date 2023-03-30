package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberScrollMemberPointLogsForBackend
// @id 2544
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=滚动查询商户积分)
func (client *Client) MemberScrollMemberPointLogsForBackendV2_0(request *MemberScrollMemberPointLogsForBackendRequestV2_0) (response *MemberScrollMemberPointLogsForBackendResponseV2_0, err error) {
	rpcResponse := CreateMemberScrollMemberPointLogsForBackendResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberScrollMemberPointLogsForBackendRequestV2_0 struct {
	*api.BaseRequest

	Cursor string `json:"cursor,omitempty"`
	Size   int    `json:"size,omitempty"`
	Pid    int64  `json:"pid,omitempty"`
}

type MemberScrollMemberPointLogsForBackendResponseV2_0 struct {
}

func CreateMemberScrollMemberPointLogsForBackendRequestV2_0() (request *MemberScrollMemberPointLogsForBackendRequestV2_0) {
	request = &MemberScrollMemberPointLogsForBackendRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "2_0", "member/scrollMemberPointLogsForBackend", "api")
	request.Method = api.POST
	return
}

func CreateMemberScrollMemberPointLogsForBackendResponseV2_0() (response *api.BaseResponse[MemberScrollMemberPointLogsForBackendResponseV2_0]) {
	response = api.CreateResponse[MemberScrollMemberPointLogsForBackendResponseV2_0](&MemberScrollMemberPointLogsForBackendResponseV2_0{})
	return
}
