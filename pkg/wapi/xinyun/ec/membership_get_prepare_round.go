package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembershipGetPrepareRound
// @id 1306
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取任务准备轮数)
func (client *Client) MembershipGetPrepareRound(request *MembershipGetPrepareRoundRequest) (response *MembershipGetPrepareRoundResponse, err error) {
	rpcResponse := CreateMembershipGetPrepareRoundResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembershipGetPrepareRoundRequest struct {
	*api.BaseRequest

	TagId int64 `json:"tagId,omitempty"`
}

type MembershipGetPrepareRoundResponse struct {
}

func CreateMembershipGetPrepareRoundRequest() (request *MembershipGetPrepareRoundRequest) {
	request = &MembershipGetPrepareRoundRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "membership/getPrepareRound", "api")
	request.Method = api.POST
	return
}

func CreateMembershipGetPrepareRoundResponse() (response *api.BaseResponse[MembershipGetPrepareRoundResponse]) {
	response = api.CreateResponse[MembershipGetPrepareRoundResponse](&MembershipGetPrepareRoundResponse{})
	return
}
