package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembercardPrivilegeGoodsGet
// @id 1749
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1749?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询折扣商品详细信息)
func (client *Client) MembercardPrivilegeGoodsGet(request *MembercardPrivilegeGoodsGetRequest) (response *MembercardPrivilegeGoodsGetResponse, err error) {
	rpcResponse := CreateMembercardPrivilegeGoodsGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembercardPrivilegeGoodsGetRequest struct {
	*api.BaseRequest

	RuleId      int64   `json:"ruleId,omitempty"`
	GoodsIdList []int64 `json:"goodsIdList,omitempty"`
	Wid         int64   `json:"wid,omitempty"`
	VidType     int64   `json:"vidType,omitempty"`
	Vid         int64   `json:"vid,omitempty"`
}

type MembercardPrivilegeGoodsGetResponse struct {
	DiscountGoodsList  []MembercardPrivilegeGoodsGetResponseDiscountGoodsList `json:"discountGoodsList,omitempty"`
	RuleId             int64                                                  `json:"ruleId,omitempty"`
	DefaultDiscount    int64                                                  `json:"defaultDiscount,omitempty"`
	MemberDiscountType int64                                                  `json:"memberDiscountType,omitempty"`
	IgnoreChangeType   int64                                                  `json:"ignoreChangeType,omitempty"`
}

func CreateMembercardPrivilegeGoodsGetRequest() (request *MembercardPrivilegeGoodsGetRequest) {
	request = &MembercardPrivilegeGoodsGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "membercard/privilege/goods/get", "apigw")
	request.Method = api.POST
	return
}

func CreateMembercardPrivilegeGoodsGetResponse() (response *api.BaseResponse[MembercardPrivilegeGoodsGetResponse]) {
	response = api.CreateResponse[MembercardPrivilegeGoodsGetResponse](&MembercardPrivilegeGoodsGetResponse{})
	return
}

type MembercardPrivilegeGoodsGetResponseDiscountGoodsList struct {
	SkuDiscountList  []MembercardPrivilegeGoodsGetResponseSkuDiscountList `json:"skuDiscountList,omitempty"`
	IgnoreChangeType int64                                                `json:"ignoreChangeType,omitempty"`
	GoodsId          int64                                                `json:"goodsId,omitempty"`
	DiscountType     int64                                                `json:"discountType,omitempty"`
}

type MembercardPrivilegeGoodsGetResponseSkuDiscountList struct {
	Discount int64 `json:"discount,omitempty"`
	SkuId    int64 `json:"skuId,omitempty"`
}
