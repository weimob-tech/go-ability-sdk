package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberUseMemberPointAmountOffLine
// @id 675
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=使用积分余额（扣减）)
func (client *Client) MemberUseMemberPointAmountOffLine(request *MemberUseMemberPointAmountOffLineRequest) (response *MemberUseMemberPointAmountOffLineResponse, err error) {
	rpcResponse := CreateMemberUseMemberPointAmountOffLineResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberUseMemberPointAmountOffLineRequest struct {
	*api.BaseRequest

	Mid         int64   `json:"mid,omitempty"`
	Point       int     `json:"point,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	Reason      string  `json:"reason,omitempty"`
	StoreId     int64   `json:"storeId,omitempty"`
	ChannelType int     `json:"channelType,omitempty"`
	AttachId    string  `json:"attachId,omitempty"`
	RequestId   string  `json:"requestId,omitempty"`
	Wid         int64   `json:"wid,omitempty"`
}

type MemberUseMemberPointAmountOffLineResponse struct {
}

func CreateMemberUseMemberPointAmountOffLineRequest() (request *MemberUseMemberPointAmountOffLineRequest) {
	request = &MemberUseMemberPointAmountOffLineRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "member/useMemberPointAmountOffLine", "api")
	request.Method = api.POST
	return
}

func CreateMemberUseMemberPointAmountOffLineResponse() (response *api.BaseResponse[MemberUseMemberPointAmountOffLineResponse]) {
	response = api.CreateResponse[MemberUseMemberPointAmountOffLineResponse](&MemberUseMemberPointAmountOffLineResponse{})
	return
}
