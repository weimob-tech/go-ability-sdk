package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberFreeze
// @id 984
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量冻结会员)
func (client *Client) MemberFreeze(request *MemberFreezeRequest) (response *MemberFreezeResponse, err error) {
	rpcResponse := CreateMemberFreezeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberFreezeRequest struct {
	*api.BaseRequest

	WidList []int            `json:"widList,omitempty"`
	MidList []map[string]any `json:"midList,omitempty"`
}

type MemberFreezeResponse struct {
	Wid int64 `json:"wid,omitempty"`
}

func CreateMemberFreezeRequest() (request *MemberFreezeRequest) {
	request = &MemberFreezeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "member/freeze", "api")
	request.Method = api.POST
	return
}

func CreateMemberFreezeResponse() (response *api.BaseResponse[MemberFreezeResponse]) {
	response = api.CreateResponse[MemberFreezeResponse](&MemberFreezeResponse{})
	return
}
