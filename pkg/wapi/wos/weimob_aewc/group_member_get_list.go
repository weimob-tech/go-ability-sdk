package weimob_aewc

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GroupMemberGetList
// @id 2218
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2218?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询指定客户群的群成员)
func (client *Client) GroupMemberGetList(request *GroupMemberGetListRequest) (response *GroupMemberGetListResponse, err error) {
	rpcResponse := CreateGroupMemberGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GroupMemberGetListRequest struct {
	*api.BaseRequest

	ChatId   string `json:"chatId,omitempty"`
	PageNum  int64  `json:"pageNum,omitempty"`
	PageSize int64  `json:"pageSize,omitempty"`
}

type GroupMemberGetListResponse struct {
	Results    []GroupMemberGetListResponseResults `json:"results,omitempty"`
	TotalCount int64                               `json:"totalCount,omitempty"`
	TotalPage  int64                               `json:"totalPage,omitempty"`
	PageSize   int64                               `json:"pageSize,omitempty"`
	PageNum    int64                               `json:"pageNum,omitempty"`
}

func CreateGroupMemberGetListRequest() (request *GroupMemberGetListRequest) {
	request = &GroupMemberGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_aewc", "v2.0", "group/member/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGroupMemberGetListResponse() (response *api.BaseResponse[GroupMemberGetListResponse]) {
	response = api.CreateResponse[GroupMemberGetListResponse](&GroupMemberGetListResponse{})
	return
}

type GroupMemberGetListResponseResults struct {
	Name          string `json:"name,omitempty"`
	Type          int64  `json:"type,omitempty"`
	JoinGroupTime int64  `json:"joinGroupTime,omitempty"`
	JoinScene     int64  `json:"joinScene,omitempty"`
	InviteName    string `json:"inviteName,omitempty"`
}
