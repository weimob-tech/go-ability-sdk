package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpUpdateMemberRightGoodsList
// @id 1784
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新商品折扣信息)
func (client *Client) MbpUpdateMemberRightGoodsList(request *MbpUpdateMemberRightGoodsListRequest) (response *MbpUpdateMemberRightGoodsListResponse, err error) {
	rpcResponse := CreateMbpUpdateMemberRightGoodsListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpUpdateMemberRightGoodsListRequest struct {
	*api.BaseRequest

	GoodsList        []MbpUpdateMemberRightGoodsListRequestGoodsList `json:"goodsList,omitempty"`
	RuleId           string                                          `json:"ruleId,omitempty"`
	DiscountType     int64                                           `json:"discountType,omitempty"`
	IgnoreChangeType int64                                           `json:"ignoreChangeType,omitempty"`
}

type MbpUpdateMemberRightGoodsListResponse struct {
	FailGoods    []int64 `json:"failGoods,omitempty"`
	SuccessGoods []int64 `json:"successGoods,omitempty"`
}

func CreateMbpUpdateMemberRightGoodsListRequest() (request *MbpUpdateMemberRightGoodsListRequest) {
	request = &MbpUpdateMemberRightGoodsListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "mbp/updateMemberRightGoodsList", "api")
	request.Method = api.POST
	return
}

func CreateMbpUpdateMemberRightGoodsListResponse() (response *api.BaseResponse[MbpUpdateMemberRightGoodsListResponse]) {
	response = api.CreateResponse[MbpUpdateMemberRightGoodsListResponse](&MbpUpdateMemberRightGoodsListResponse{})
	return
}

type MbpUpdateMemberRightGoodsListRequestGoodsList struct {
	SkuLIst []MbpUpdateMemberRightGoodsListRequestSkuLIst `json:"skuLIst,omitempty"`
	GoodsId int64                                         `json:"goodsId,omitempty"`
}

type MbpUpdateMemberRightGoodsListRequestSkuLIst struct {
	Discount  int64 `json:"discount,omitempty"`
	SalePrice int64 `json:"salePrice,omitempty"`
	SkuId     int64 `json:"skuId,omitempty"`
}
