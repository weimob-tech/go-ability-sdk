package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembershipPrepareByRound
// @id 1307
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=传入任务数据)
func (client *Client) MembershipPrepareByRound(request *MembershipPrepareByRoundRequest) (response *MembershipPrepareByRoundResponse, err error) {
	rpcResponse := CreateMembershipPrepareByRoundResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembershipPrepareByRoundRequest struct {
	*api.BaseRequest

	WidList []int  `json:"widList,omitempty"`
	TagId   int64  `json:"tagId,omitempty"`
	Round   string `json:"round,omitempty"`
}

type MembershipPrepareByRoundResponse struct {
}

func CreateMembershipPrepareByRoundRequest() (request *MembershipPrepareByRoundRequest) {
	request = &MembershipPrepareByRoundRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "membership/prepareByRound", "api")
	request.Method = api.POST
	return
}

func CreateMembershipPrepareByRoundResponse() (response *api.BaseResponse[MembershipPrepareByRoundResponse]) {
	response = api.CreateResponse[MembershipPrepareByRoundResponse](&MembershipPrepareByRoundResponse{})
	return
}
