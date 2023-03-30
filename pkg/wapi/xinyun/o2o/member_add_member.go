package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberAddMember
// @id 1050
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增会员)
func (client *Client) MemberAddMember(request *MemberAddMemberRequest) (response *MemberAddMemberResponse, err error) {
	rpcResponse := CreateMemberAddMemberResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberAddMemberRequest struct {
	*api.BaseRequest

	AvailablePoint int64  `json:"availablePoint,omitempty"`
	Birthday       string `json:"birthday,omitempty"`
	DepositAmount  int64  `json:"depositAmount,omitempty"`
	GiveAmount     int64  `json:"giveAmount,omitempty"`
	Growth         int64  `json:"growth,omitempty"`
	Name           string `json:"name,omitempty"`
	Phone          string `json:"phone,omitempty"`
	Sex            int64  `json:"sex,omitempty"`
}

type MemberAddMemberResponse struct {
}

func CreateMemberAddMemberRequest() (request *MemberAddMemberRequest) {
	request = &MemberAddMemberRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "member/addMember", "api")
	request.Method = api.POST
	return
}

func CreateMemberAddMemberResponse() (response *api.BaseResponse[MemberAddMemberResponse]) {
	response = api.CreateResponse[MemberAddMemberResponse](&MemberAddMemberResponse{})
	return
}
