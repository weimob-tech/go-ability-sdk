package kld

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// KLDMemberCardConsumeMember
// @id 329
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=会员线下消费)
func (client *Client) KLDMemberCardConsumeMember(request *KLDMemberCardConsumeMemberRequest) (response *KLDMemberCardConsumeMemberResponse, err error) {
	rpcResponse := CreateKLDMemberCardConsumeMemberResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type KLDMemberCardConsumeMemberRequest struct {
	*api.BaseRequest

	CashCouponSnNoList  []string `json:"cashCouponSnNoList,omitempty"`
	MemberCardNo        string   `json:"memberCardNo,omitempty"`
	TotalAmount         string   `json:"totalAmount,omitempty"`
	StoreId             int64    `json:"storeId,omitempty"`
	NoDiscountAmount    int64    `json:"noDiscountAmount,omitempty"`
	IsUseMemberDiscount bool     `json:"isUseMemberDiscount,omitempty"`
	SnNo                string   `json:"snNo,omitempty"`
	RealAmount          int64    `json:"realAmount,omitempty"`
	BlanceAmount        int64    `json:"blanceAmount,omitempty"`
	CashAmount          int64    `json:"cashAmount,omitempty"`
	Points              int64    `json:"points,omitempty"`
	ChangeAmount        int64    `json:"changeAmount,omitempty"`
	AdminPassword       string   `json:"adminPassword,omitempty"`
	Operator            string   `json:"operator,omitempty"`
	Remark              string   `json:"remark,omitempty"`
}

type KLDMemberCardConsumeMemberResponse struct {
}

func CreateKLDMemberCardConsumeMemberRequest() (request *KLDMemberCardConsumeMemberRequest) {
	request = &KLDMemberCardConsumeMemberRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("kld", "1_0", "KLDMemberCard/ConsumeMember", "api")
	request.Method = api.POST
	return
}

func CreateKLDMemberCardConsumeMemberResponse() (response *api.BaseResponse[KLDMemberCardConsumeMemberResponse]) {
	response = api.CreateResponse[KLDMemberCardConsumeMemberResponse](&KLDMemberCardConsumeMemberResponse{})
	return
}
