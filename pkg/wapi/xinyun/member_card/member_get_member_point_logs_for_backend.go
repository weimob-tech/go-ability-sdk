package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberGetMemberPointLogsForBackend
// @id 715
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取积分记录)
func (client *Client) MemberGetMemberPointLogsForBackend(request *MemberGetMemberPointLogsForBackendRequest) (response *MemberGetMemberPointLogsForBackendResponse, err error) {
	rpcResponse := CreateMemberGetMemberPointLogsForBackendResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberGetMemberPointLogsForBackendRequest struct {
	*api.BaseRequest

	Mid         int64  `json:"mid,omitempty"`
	StartTime   int    `json:"startTime,omitempty"`
	EndTime     int    `json:"endTime,omitempty"`
	Page        int    `json:"page,omitempty"`
	PageSize    int    `json:"pageSize,omitempty"`
	ChannelType int    `json:"channelType,omitempty"`
	Wid         int64  `json:"wid,omitempty"`
	PointFlowId int64  `json:"pointFlowId,omitempty"`
	AttachId    string `json:"attachId,omitempty"`
	PointTid    int64  `json:"pointTid,omitempty"`
	StoreId     int64  `json:"storeId,omitempty"`
}

type MemberGetMemberPointLogsForBackendResponse struct {
}

func CreateMemberGetMemberPointLogsForBackendRequest() (request *MemberGetMemberPointLogsForBackendRequest) {
	request = &MemberGetMemberPointLogsForBackendRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "member/getMemberPointLogsForBackend", "api")
	request.Method = api.POST
	return
}

func CreateMemberGetMemberPointLogsForBackendResponse() (response *api.BaseResponse[MemberGetMemberPointLogsForBackendResponse]) {
	response = api.CreateResponse[MemberGetMemberPointLogsForBackendResponse](&MemberGetMemberPointLogsForBackendResponse{})
	return
}
