package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberGetMemberInterestsInfo
// @id 2072
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查会员权益信息)
func (client *Client) MemberGetMemberInterestsInfo(request *MemberGetMemberInterestsInfoRequest) (response *MemberGetMemberInterestsInfoResponse, err error) {
	rpcResponse := CreateMemberGetMemberInterestsInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberGetMemberInterestsInfoRequest struct {
	*api.BaseRequest

	Wid          string `json:"wid,omitempty"`
	Code         string `json:"code,omitempty"`
	Mid          int64  `json:"mid,omitempty"`
	Phone        string `json:"phone,omitempty"`
	IsNeedDetail bool   `json:"isNeedDetail,omitempty"`
	Source       int    `json:"source,omitempty"`
	Type         int    `json:"type,omitempty"`
}

type MemberGetMemberInterestsInfoResponse struct {
}

func CreateMemberGetMemberInterestsInfoRequest() (request *MemberGetMemberInterestsInfoRequest) {
	request = &MemberGetMemberInterestsInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "member/getMemberInterestsInfo", "api")
	request.Method = api.POST
	return
}

func CreateMemberGetMemberInterestsInfoResponse() (response *api.BaseResponse[MemberGetMemberInterestsInfoResponse]) {
	response = api.CreateResponse[MemberGetMemberInterestsInfoResponse](&MemberGetMemberInterestsInfoResponse{})
	return
}
