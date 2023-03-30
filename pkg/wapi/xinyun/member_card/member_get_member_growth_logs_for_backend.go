package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberGetMemberGrowthLogsForBackend
// @id 1282
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取成长值记录)
func (client *Client) MemberGetMemberGrowthLogsForBackend(request *MemberGetMemberGrowthLogsForBackendRequest) (response *MemberGetMemberGrowthLogsForBackendResponse, err error) {
	rpcResponse := CreateMemberGetMemberGrowthLogsForBackendResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberGetMemberGrowthLogsForBackendRequest struct {
	*api.BaseRequest

	Wid         int64 `json:"wid,omitempty"`
	ChangeType  int   `json:"changeType,omitempty"`
	ChannelType int   `json:"channelType,omitempty"`
	StartTime   int   `json:"startTime,omitempty"`
	EndTime     int   `json:"endTime,omitempty"`
	Page        int   `json:"page,omitempty"`
	PageSize    int   `json:"pageSize,omitempty"`
}

type MemberGetMemberGrowthLogsForBackendResponse struct {
}

func CreateMemberGetMemberGrowthLogsForBackendRequest() (request *MemberGetMemberGrowthLogsForBackendRequest) {
	request = &MemberGetMemberGrowthLogsForBackendRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "member/getMemberGrowthLogsForBackend", "api")
	request.Method = api.POST
	return
}

func CreateMemberGetMemberGrowthLogsForBackendResponse() (response *api.BaseResponse[MemberGetMemberGrowthLogsForBackendResponse]) {
	response = api.CreateResponse[MemberGetMemberGrowthLogsForBackendResponse](&MemberGetMemberGrowthLogsForBackendResponse{})
	return
}
