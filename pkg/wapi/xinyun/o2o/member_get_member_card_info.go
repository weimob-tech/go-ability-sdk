package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberGetMemberCardInfo
// @id 2071
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查会员卡信息)
func (client *Client) MemberGetMemberCardInfo(request *MemberGetMemberCardInfoRequest) (response *MemberGetMemberCardInfoResponse, err error) {
	rpcResponse := CreateMemberGetMemberCardInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberGetMemberCardInfoRequest struct {
	*api.BaseRequest

	Source               int    `json:"source,omitempty"`
	Wid                  int64  `json:"wid,omitempty"`
	MemberCardTemplateId int64  `json:"memberCardTemplateId,omitempty"`
	Code                 string `json:"code,omitempty"`
	Mid                  int64  `json:"mid,omitempty"`
}

type MemberGetMemberCardInfoResponse struct {
}

func CreateMemberGetMemberCardInfoRequest() (request *MemberGetMemberCardInfoRequest) {
	request = &MemberGetMemberCardInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "member/getMemberCardInfo", "api")
	request.Method = api.POST
	return
}

func CreateMemberGetMemberCardInfoResponse() (response *api.BaseResponse[MemberGetMemberCardInfoResponse]) {
	response = api.CreateResponse[MemberGetMemberCardInfoResponse](&MemberGetMemberCardInfoResponse{})
	return
}
