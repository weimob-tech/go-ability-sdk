package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberGetMemberDetailNew
// @id 1834
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询会员信息)
func (client *Client) MemberGetMemberDetailNew(request *MemberGetMemberDetailNewRequest) (response *MemberGetMemberDetailNewResponse, err error) {
	rpcResponse := CreateMemberGetMemberDetailNewResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberGetMemberDetailNewRequest struct {
	*api.BaseRequest

	Wid      int64  `json:"wid,omitempty"`
	Phone    string `json:"phone,omitempty"`
	CardNo   string `json:"cardNo,omitempty"`
	CardType int    `json:"cardType,omitempty"`
}

type MemberGetMemberDetailNewResponse struct {
}

func CreateMemberGetMemberDetailNewRequest() (request *MemberGetMemberDetailNewRequest) {
	request = &MemberGetMemberDetailNewRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "member/getMemberDetailNew", "api")
	request.Method = api.POST
	return
}

func CreateMemberGetMemberDetailNewResponse() (response *api.BaseResponse[MemberGetMemberDetailNewResponse]) {
	response = api.CreateResponse[MemberGetMemberDetailNewResponse](&MemberGetMemberDetailNewResponse{})
	return
}
