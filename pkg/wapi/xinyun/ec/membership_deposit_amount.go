package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembershipDepositAmount
// @id 1835
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=会员充值接口_支持触发充值有礼)
func (client *Client) MembershipDepositAmount(request *MembershipDepositAmountRequest) (response *MembershipDepositAmountResponse, err error) {
	rpcResponse := CreateMembershipDepositAmountResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembershipDepositAmountRequest struct {
	*api.BaseRequest

	StoreId          int64   `json:"storeId,omitempty"`
	Wid              int64   `json:"wid,omitempty"`
	OperatorName     string  `json:"operatorName,omitempty"`
	OrderNo          string  `json:"orderNo,omitempty"`
	Amount           float64 `json:"amount,omitempty"`
	Comment          string  `json:"comment,omitempty"`
	ChannelType      int64   `json:"channelType,omitempty"`
	NeedRechargeGift int     `json:"needRechargeGift,omitempty"`
	RechargeType     string  `json:"rechargeType,omitempty"`
	OperatorId       int64   `json:"operatorId,omitempty"`
}

type MembershipDepositAmountResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateMembershipDepositAmountRequest() (request *MembershipDepositAmountRequest) {
	request = &MembershipDepositAmountRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "membership/depositAmount", "api")
	request.Method = api.POST
	return
}

func CreateMembershipDepositAmountResponse() (response *api.BaseResponse[MembershipDepositAmountResponse]) {
	response = api.CreateResponse[MembershipDepositAmountResponse](&MembershipDepositAmountResponse{})
	return
}
