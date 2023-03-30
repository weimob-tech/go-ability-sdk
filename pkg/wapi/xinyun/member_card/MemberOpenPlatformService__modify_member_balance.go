package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberOpenPlatformServiceModifyMemberBalance
// @id 169
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改会员余额)
func (client *Client) MemberOpenPlatformServiceModifyMemberBalance(request *MemberOpenPlatformServiceModifyMemberBalanceRequest) (response *MemberOpenPlatformServiceModifyMemberBalanceResponse, err error) {
	rpcResponse := CreateMemberOpenPlatformServiceModifyMemberBalanceResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberOpenPlatformServiceModifyMemberBalanceRequest struct {
	*api.BaseRequest

	CustomerNos  []string `json:"customer_nos,omitempty"`
	Amount       float64  `json:"amount,omitempty"`
	Password     string   `json:"password,omitempty"`
	RechargeType int      `json:"recharge_type,omitempty"`
	Opreater     string   `json:"opreater,omitempty"`
	Remark       string   `json:"remark,omitempty"`
	VerifyId     string   `json:"verify_id,omitempty"`
	Mulshopid    int64    `json:"mulshopid,omitempty"`
}

type MemberOpenPlatformServiceModifyMemberBalanceResponse struct {
}

func CreateMemberOpenPlatformServiceModifyMemberBalanceRequest() (request *MemberOpenPlatformServiceModifyMemberBalanceRequest) {
	request = &MemberOpenPlatformServiceModifyMemberBalanceRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "MemberOpenPlatformService/ModifyMemberBalance", "api")
	request.Method = api.POST
	return
}

func CreateMemberOpenPlatformServiceModifyMemberBalanceResponse() (response *api.BaseResponse[MemberOpenPlatformServiceModifyMemberBalanceResponse]) {
	response = api.CreateResponse[MemberOpenPlatformServiceModifyMemberBalanceResponse](&MemberOpenPlatformServiceModifyMemberBalanceResponse{})
	return
}
