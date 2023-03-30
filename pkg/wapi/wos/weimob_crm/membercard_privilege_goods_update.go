package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembercardPrivilegeGoodsUpdate
// @id 1750
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1750?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新折扣商品信息)
func (client *Client) MembercardPrivilegeGoodsUpdate(request *MembercardPrivilegeGoodsUpdateRequest) (response *MembercardPrivilegeGoodsUpdateResponse, err error) {
	rpcResponse := CreateMembercardPrivilegeGoodsUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembercardPrivilegeGoodsUpdateRequest struct {
	*api.BaseRequest

	GoodsList        []MembercardPrivilegeGoodsUpdateRequestGoodsList `json:"goodsList,omitempty"`
	RuleId           int64                                            `json:"ruleId,omitempty"`
	DiscountType     int64                                            `json:"discountType,omitempty"`
	Wid              int64                                            `json:"wid,omitempty"`
	VidType          int64                                            `json:"vidType,omitempty"`
	Vid              int64                                            `json:"vid,omitempty"`
	IgnoreChangeType int64                                            `json:"ignoreChangeType,omitempty"`
}

type MembercardPrivilegeGoodsUpdateResponse struct {
	FailInfoList  []MembercardPrivilegeGoodsUpdateResponseFailInfoList `json:"failInfoList,omitempty"`
	SuccessIdList []int64                                              `json:"successIdList,omitempty"`
}

func CreateMembercardPrivilegeGoodsUpdateRequest() (request *MembercardPrivilegeGoodsUpdateRequest) {
	request = &MembercardPrivilegeGoodsUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "membercard/privilege/goods/update", "apigw")
	request.Method = api.POST
	return
}

func CreateMembercardPrivilegeGoodsUpdateResponse() (response *api.BaseResponse[MembercardPrivilegeGoodsUpdateResponse]) {
	response = api.CreateResponse[MembercardPrivilegeGoodsUpdateResponse](&MembercardPrivilegeGoodsUpdateResponse{})
	return
}

type MembercardPrivilegeGoodsUpdateRequestGoodsList struct {
	SkuList []MembercardPrivilegeGoodsUpdateRequestSkuList `json:"skuList,omitempty"`
	GoodsId int64                                          `json:"goodsId,omitempty"`
}

type MembercardPrivilegeGoodsUpdateRequestSkuList struct {
	SkuId     int64  `json:"skuId,omitempty"`
	SalePrice string `json:"salePrice,omitempty"`
	Discount  string `json:"discount,omitempty"`
}

type MembercardPrivilegeGoodsUpdateResponseFailInfoList struct {
	GoodsId int64  `json:"goodsId,omitempty"`
	FailMsg string `json:"failMsg,omitempty"`
}
