package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberScrollMemberAmountLogsForBackend
// @id 2543
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=滚动查询余额记录)
func (client *Client) MemberScrollMemberAmountLogsForBackendV2_0(request *MemberScrollMemberAmountLogsForBackendRequestV2_0) (response *MemberScrollMemberAmountLogsForBackendResponseV2_0, err error) {
	rpcResponse := CreateMemberScrollMemberAmountLogsForBackendResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberScrollMemberAmountLogsForBackendRequestV2_0 struct {
	*api.BaseRequest

	Cursor string `json:"cursor,omitempty"`
	Size   int    `json:"size,omitempty"`
	Pid    int64  `json:"pid,omitempty"`
}

type MemberScrollMemberAmountLogsForBackendResponseV2_0 struct {
}

func CreateMemberScrollMemberAmountLogsForBackendRequestV2_0() (request *MemberScrollMemberAmountLogsForBackendRequestV2_0) {
	request = &MemberScrollMemberAmountLogsForBackendRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "2_0", "member/scrollMemberAmountLogsForBackend", "api")
	request.Method = api.POST
	return
}

func CreateMemberScrollMemberAmountLogsForBackendResponseV2_0() (response *api.BaseResponse[MemberScrollMemberAmountLogsForBackendResponseV2_0]) {
	response = api.CreateResponse[MemberScrollMemberAmountLogsForBackendResponseV2_0](&MemberScrollMemberAmountLogsForBackendResponseV2_0{})
	return
}
