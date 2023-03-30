package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberGetUserInfoList
// @id 2074
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查用户信息列表)
func (client *Client) MemberGetUserInfoList(request *MemberGetUserInfoListRequest) (response *MemberGetUserInfoListResponse, err error) {
	rpcResponse := CreateMemberGetUserInfoListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberGetUserInfoListRequest struct {
	*api.BaseRequest

	HomeStoreIds []map[string]any `json:"homeStoreIds,omitempty"`
	UserType     int              `json:"userType,omitempty"`
	StartTime    int              `json:"startTime,omitempty"`
	EndTime      int              `json:"endTime,omitempty"`
	PageSize     int              `json:"pageSize,omitempty"`
	Page         int              `json:"page,omitempty"`
}

type MemberGetUserInfoListResponse struct {
}

func CreateMemberGetUserInfoListRequest() (request *MemberGetUserInfoListRequest) {
	request = &MemberGetUserInfoListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "member/getUserInfoList", "api")
	request.Method = api.POST
	return
}

func CreateMemberGetUserInfoListResponse() (response *api.BaseResponse[MemberGetUserInfoListResponse]) {
	response = api.CreateResponse[MemberGetUserInfoListResponse](&MemberGetUserInfoListResponse{})
	return
}
