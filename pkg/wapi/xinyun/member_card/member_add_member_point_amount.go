package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberAddMemberPointAmount
// @id 673
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=增加会员积分余额)
func (client *Client) MemberAddMemberPointAmount(request *MemberAddMemberPointAmountRequest) (response *MemberAddMemberPointAmountResponse, err error) {
	rpcResponse := CreateMemberAddMemberPointAmountResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberAddMemberPointAmountRequest struct {
	*api.BaseRequest

	Mid             int64   `json:"mid,omitempty"`
	Amount          float64 `json:"amount,omitempty"`
	AddAmountReason string  `json:"addAmountReason,omitempty"`
	Point           int     `json:"point,omitempty"`
	AddPointReason  string  `json:"addPointReason,omitempty"`
	ChannelType     int     `json:"channelType,omitempty"`
	StoreId         int64   `json:"storeId,omitempty"`
	AttachId        string  `json:"attachId,omitempty"`
	RequestId       string  `json:"requestId,omitempty"`
	Wid             int64   `json:"wid,omitempty"`
}

type MemberAddMemberPointAmountResponse struct {
}

func CreateMemberAddMemberPointAmountRequest() (request *MemberAddMemberPointAmountRequest) {
	request = &MemberAddMemberPointAmountRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "member/addMemberPointAmount", "api")
	request.Method = api.POST
	return
}

func CreateMemberAddMemberPointAmountResponse() (response *api.BaseResponse[MemberAddMemberPointAmountResponse]) {
	response = api.CreateResponse[MemberAddMemberPointAmountResponse](&MemberAddMemberPointAmountResponse{})
	return
}
