package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberGetMiniMemberInfo
// @id 2070
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查会员、会员卡信息复合接口)
func (client *Client) MemberGetMiniMemberInfo(request *MemberGetMiniMemberInfoRequest) (response *MemberGetMiniMemberInfoResponse, err error) {
	rpcResponse := CreateMemberGetMiniMemberInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberGetMiniMemberInfoRequest struct {
	*api.BaseRequest

	Wid int64 `json:"wid,omitempty"`
}

type MemberGetMiniMemberInfoResponse struct {
}

func CreateMemberGetMiniMemberInfoRequest() (request *MemberGetMiniMemberInfoRequest) {
	request = &MemberGetMiniMemberInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "member/getMiniMemberInfo", "api")
	request.Method = api.POST
	return
}

func CreateMemberGetMiniMemberInfoResponse() (response *api.BaseResponse[MemberGetMiniMemberInfoResponse]) {
	response = api.CreateResponse[MemberGetMiniMemberInfoResponse](&MemberGetMiniMemberInfoResponse{})
	return
}
