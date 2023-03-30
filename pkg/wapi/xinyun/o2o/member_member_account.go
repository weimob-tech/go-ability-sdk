package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberMemberAccount
// @id 1415
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询客户可用余额)
func (client *Client) MemberMemberAccount(request *MemberMemberAccountRequest) (response *MemberMemberAccountResponse, err error) {
	rpcResponse := CreateMemberMemberAccountResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberMemberAccountRequest struct {
	*api.BaseRequest

	ConsumerWid string `json:"consumerWid,omitempty"`
}

type MemberMemberAccountResponse struct {
}

func CreateMemberMemberAccountRequest() (request *MemberMemberAccountRequest) {
	request = &MemberMemberAccountRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "member/memberAccount", "api")
	request.Method = api.POST
	return
}

func CreateMemberMemberAccountResponse() (response *api.BaseResponse[MemberMemberAccountResponse]) {
	response = api.CreateResponse[MemberMemberAccountResponse](&MemberMemberAccountResponse{})
	return
}
