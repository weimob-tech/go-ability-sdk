package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberMemberDeposit
// @id 676
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=余额充值)
func (client *Client) MemberMemberDeposit(request *MemberMemberDepositRequest) (response *MemberMemberDepositResponse, err error) {
	rpcResponse := CreateMemberMemberDepositResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberMemberDepositRequest struct {
	*api.BaseRequest

	Wid          int64   `json:"wid,omitempty"`
	OperaterName string  `json:"operaterName,omitempty"`
	Comment      string  `json:"comment,omitempty"`
	Amount       float64 `json:"amount,omitempty"`
	StoreId      int64   `json:"storeId,omitempty"`
	ChannelType  int     `json:"channelType,omitempty"`
	AttachId     string  `json:"attachId,omitempty"`
	RequestId    string  `json:"requestId,omitempty"`
	Mid          int64   `json:"mid,omitempty"`
}

type MemberMemberDepositResponse struct {
}

func CreateMemberMemberDepositRequest() (request *MemberMemberDepositRequest) {
	request = &MemberMemberDepositRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "member/memberDeposit", "api")
	request.Method = api.POST
	return
}

func CreateMemberMemberDepositResponse() (response *api.BaseResponse[MemberMemberDepositResponse]) {
	response = api.CreateResponse[MemberMemberDepositResponse](&MemberMemberDepositResponse{})
	return
}
