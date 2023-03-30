package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberUnfreeze
// @id 985
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量解冻会员)
func (client *Client) MemberUnfreeze(request *MemberUnfreezeRequest) (response *MemberUnfreezeResponse, err error) {
	rpcResponse := CreateMemberUnfreezeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberUnfreezeRequest struct {
	*api.BaseRequest

	WidList []int            `json:"widList,omitempty"`
	MidList []map[string]any `json:"midList,omitempty"`
}

type MemberUnfreezeResponse struct {
	Wid int64 `json:"wid,omitempty"`
}

func CreateMemberUnfreezeRequest() (request *MemberUnfreezeRequest) {
	request = &MemberUnfreezeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "member/unfreeze", "api")
	request.Method = api.POST
	return
}

func CreateMemberUnfreezeResponse() (response *api.BaseResponse[MemberUnfreezeResponse]) {
	response = api.CreateResponse[MemberUnfreezeResponse](&MemberUnfreezeResponse{})
	return
}
