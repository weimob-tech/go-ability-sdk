package weimob_cdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerMembercardGetList
// @id 1717
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1717?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询客户会员卡列表（治理）)
func (client *Client) CustomerMembercardGetList(request *CustomerMembercardGetListRequest) (response *CustomerMembercardGetListResponse, err error) {
	rpcResponse := CreateCustomerMembercardGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerMembercardGetListRequest struct {
	*api.BaseRequest

	ChannelId         string `json:"channelId,omitempty"`
	ChannelType       string `json:"channelType,omitempty"`
	Ukey              string `json:"ukey,omitempty"`
	UkeyType          string `json:"ukeyType,omitempty"`
	Status            int64  `json:"status,omitempty"`
	SchannelId        string `json:"schannelId,omitempty"`
	SchannelType      string `json:"schannelType,omitempty"`
	PageNum           int64  `json:"pageNum,omitempty"`
	PageSize          int64  `json:"pageSize,omitempty"`
	VidType           int64  `json:"vidType,omitempty"`
	Vid               int64  `json:"vid,omitempty"`
	BosId             int64  `json:"bosId,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	MerchantId        int64  `json:"merchantId,omitempty"`
	Tcode             string `json:"tcode,omitempty"`
}

type CustomerMembercardGetListResponse struct {
	List     []CustomerMembercardGetListResponselist `json:"list,omitempty"`
	PageSize int64                                   `json:"pageSize,omitempty"`
	Total    int64                                   `json:"total,omitempty"`
	PageNum  int64                                   `json:"pageNum,omitempty"`
}

func CreateCustomerMembercardGetListRequest() (request *CustomerMembercardGetListRequest) {
	request = &CustomerMembercardGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_cdp", "v2.0", "customer/membercard/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerMembercardGetListResponse() (response *api.BaseResponse[CustomerMembercardGetListResponse]) {
	response = api.CreateResponse[CustomerMembercardGetListResponse](&CustomerMembercardGetListResponse{})
	return
}

type CustomerMembercardGetListResponselist struct {
	MemberLevelId    int64  `json:"memberLevelId,omitempty"`
	UpdateTime       string `json:"updateTime,omitempty"`
	ExpireTime       string `json:"expireTime,omitempty"`
	MemberCardTitle  string `json:"memberCardTitle,omitempty"`
	CreateTime       string `json:"createTime,omitempty"`
	MemberCardNumber string `json:"memberCardNumber,omitempty"`
	EffectTime       string `json:"effectTime,omitempty"`
	MemberLevel      string `json:"memberLevel,omitempty"`
	RankRuleType     int64  `json:"rankRuleType,omitempty"`
	MemberStatus     int64  `json:"memberStatus,omitempty"`
	SchannelId       string `json:"schannelId,omitempty"`
	MemberCardType   int64  `json:"memberCardType,omitempty"`
	UkeyType         string `json:"ukeyType,omitempty"`
	Ukey             string `json:"ukey,omitempty"`
	ChannelId        string `json:"channelId,omitempty"`
	Vid              string `json:"vid,omitempty"`
	CardNumber       string `json:"cardNumber,omitempty"`
	VidType          string `json:"vidType,omitempty"`
	SchannelType     string `json:"schannelType,omitempty"`
	ChannelType      string `json:"channelType,omitempty"`
	BosId            string `json:"bosId,omitempty"`
	RegisterTime     string `json:"registerTime,omitempty"`
	PlanId           int64  `json:"planId,omitempty"`
}
