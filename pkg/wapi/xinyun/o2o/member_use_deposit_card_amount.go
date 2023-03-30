package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberUseDepositCardAmount
// @id 1337
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=使用储值卡余额)
func (client *Client) MemberUseDepositCardAmount(request *MemberUseDepositCardAmountRequest) (response *MemberUseDepositCardAmountResponse, err error) {
	rpcResponse := CreateMemberUseDepositCardAmountResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberUseDepositCardAmountRequest struct {
	*api.BaseRequest

	Amount         float64 `json:"amount,omitempty"`
	CardTemplateId int64   `json:"cardTemplateId,omitempty"`
	ConsumerWid    int64   `json:"consumerWid,omitempty"`
	Reason         string  `json:"reason,omitempty"`
	RequestId      string  `json:"requestId,omitempty"`
	StoreId        int64   `json:"storeId,omitempty"`
	ThirdOrderNo   string  `json:"thirdOrderNo,omitempty"`
}

type MemberUseDepositCardAmountResponse struct {
}

func CreateMemberUseDepositCardAmountRequest() (request *MemberUseDepositCardAmountRequest) {
	request = &MemberUseDepositCardAmountRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "member/useDepositCardAmount", "api")
	request.Method = api.POST
	return
}

func CreateMemberUseDepositCardAmountResponse() (response *api.BaseResponse[MemberUseDepositCardAmountResponse]) {
	response = api.CreateResponse[MemberUseDepositCardAmountResponse](&MemberUseDepositCardAmountResponse{})
	return
}
