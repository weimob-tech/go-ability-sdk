package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembershipActivateByRound
// @id 1308
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=激活任务)
func (client *Client) MembershipActivateByRound(request *MembershipActivateByRoundRequest) (response *MembershipActivateByRoundResponse, err error) {
	rpcResponse := CreateMembershipActivateByRoundResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembershipActivateByRoundRequest struct {
	*api.BaseRequest

	TagId int64  `json:"tagId,omitempty"`
	Round string `json:"round,omitempty"`
}

type MembershipActivateByRoundResponse struct {
}

func CreateMembershipActivateByRoundRequest() (request *MembershipActivateByRoundRequest) {
	request = &MembershipActivateByRoundRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "membership/activateByRound", "api")
	request.Method = api.POST
	return
}

func CreateMembershipActivateByRoundResponse() (response *api.BaseResponse[MembershipActivateByRoundResponse]) {
	response = api.CreateResponse[MembershipActivateByRoundResponse](&MembershipActivateByRoundResponse{})
	return
}
