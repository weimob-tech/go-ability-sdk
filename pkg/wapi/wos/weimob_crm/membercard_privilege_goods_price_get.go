package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembercardPrivilegeGoodsPriceGet
// @id 2150
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2150?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取商品预计开卡优惠)
func (client *Client) MembercardPrivilegeGoodsPriceGet(request *MembercardPrivilegeGoodsPriceGetRequest) (response *MembercardPrivilegeGoodsPriceGetResponse, err error) {
	rpcResponse := CreateMembercardPrivilegeGoodsPriceGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembercardPrivilegeGoodsPriceGetRequest struct {
	*api.BaseRequest

	GoodsList []MembercardPrivilegeGoodsPriceGetRequestGoodsList `json:"goodsList,omitempty"`
	Wid       int64                                              `json:"wid,omitempty"`
}

type MembercardPrivilegeGoodsPriceGetResponse struct {
	GoodsList          []MembercardPrivilegeGoodsPriceGetResponseGoodsList `json:"goodsList,omitempty"`
	HasValidMemberCard bool                                                `json:"hasValidMemberCard,omitempty"`
	MembershipPlanId   int64                                               `json:"membershipPlanId,omitempty"`
}

func CreateMembercardPrivilegeGoodsPriceGetRequest() (request *MembercardPrivilegeGoodsPriceGetRequest) {
	request = &MembercardPrivilegeGoodsPriceGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "membercard/privilege/goods/price/get", "apigw")
	request.Method = api.POST
	return
}

func CreateMembercardPrivilegeGoodsPriceGetResponse() (response *api.BaseResponse[MembercardPrivilegeGoodsPriceGetResponse]) {
	response = api.CreateResponse[MembercardPrivilegeGoodsPriceGetResponse](&MembercardPrivilegeGoodsPriceGetResponse{})
	return
}

type MembercardPrivilegeGoodsPriceGetRequestGoodsList struct {
	ItemId    int64  `json:"itemId,omitempty"`
	GoodsId   int64  `json:"goodsId,omitempty"`
	SkuId     int64  `json:"skuId,omitempty"`
	SalePrice string `json:"salePrice,omitempty"`
}

type MembercardPrivilegeGoodsPriceGetResponseGoodsList struct {
	MatchRuleList []MembercardPrivilegeGoodsPriceGetResponseMatchRuleList `json:"matchRuleList,omitempty"`
	ItemId        int64                                                   `json:"itemId,omitempty"`
	GoodsId       int64                                                   `json:"goodsId,omitempty"`
	SkuId         int64                                                   `json:"skuId,omitempty"`
	SalePrice     string                                                  `json:"salePrice,omitempty"`
}

type MembercardPrivilegeGoodsPriceGetResponseMatchRuleList struct {
	DiscountId     int64  `json:"discountId,omitempty"`
	RuleId         int64  `json:"ruleId,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	DiscountAmount string `json:"discountAmount,omitempty"`
}
