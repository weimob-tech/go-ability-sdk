package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembershipCreate
// @id 1463
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=创建会员)
func (client *Client) MembershipCreate(request *MembershipCreateRequest) (response *MembershipCreateResponse, err error) {
	rpcResponse := CreateMembershipCreateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembershipCreateRequest struct {
	*api.BaseRequest
}

type MembershipCreateResponse struct {
}

func CreateMembershipCreateRequest() (request *MembershipCreateRequest) {
	request = &MembershipCreateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "membership/create", "api")
	request.Method = api.POST
	return
}

func CreateMembershipCreateResponse() (response *api.BaseResponse[MembershipCreateResponse]) {
	response = api.CreateResponse[MembershipCreateResponse](&MembershipCreateResponse{})
	return
}
