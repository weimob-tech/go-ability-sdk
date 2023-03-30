package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberMemberDeposit
// @id 1672
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=余额充值V2)
func (client *Client) MemberMemberDepositV2_0(request *MemberMemberDepositRequestV2_0) (response *MemberMemberDepositResponseV2_0, err error) {
	rpcResponse := CreateMemberMemberDepositResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberMemberDepositRequestV2_0 struct {
	*api.BaseRequest

	Wid                int64   `json:"wid,omitempty"`
	Mid                int64   `json:"mid,omitempty"`
	OperaterName       string  `json:"operaterName,omitempty"`
	Comment            string  `json:"comment,omitempty"`
	Amount             float64 `json:"amount,omitempty"`
	StoreId            int64   `json:"storeId,omitempty"`
	ChannelType        int     `json:"channelType,omitempty"`
	BizViewFilter3Type int     `json:"bizViewFilter3Type,omitempty"`
	AttachId           string  `json:"attachId,omitempty"`
	RequestId          string  `json:"requestId,omitempty"`
}

type MemberMemberDepositResponseV2_0 struct {
}

func CreateMemberMemberDepositRequestV2_0() (request *MemberMemberDepositRequestV2_0) {
	request = &MemberMemberDepositRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "2_0", "member/memberDeposit", "api")
	request.Method = api.POST
	return
}

func CreateMemberMemberDepositResponseV2_0() (response *api.BaseResponse[MemberMemberDepositResponseV2_0]) {
	response = api.CreateResponse[MemberMemberDepositResponseV2_0](&MemberMemberDepositResponseV2_0{})
	return
}
