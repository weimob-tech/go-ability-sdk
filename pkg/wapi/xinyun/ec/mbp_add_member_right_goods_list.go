package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpAddMemberRightGoodsList
// @id 1783
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=添加商品折扣信息)
func (client *Client) MbpAddMemberRightGoodsList(request *MbpAddMemberRightGoodsListRequest) (response *MbpAddMemberRightGoodsListResponse, err error) {
	rpcResponse := CreateMbpAddMemberRightGoodsListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpAddMemberRightGoodsListRequest struct {
	*api.BaseRequest

	GoodsList        []MbpAddMemberRightGoodsListRequestGoodsList `json:"goodsList,omitempty"`
	RuleId           string                                       `json:"ruleId,omitempty"`
	DiscountType     int64                                        `json:"discountType,omitempty"`
	IgnoreChangeType int64                                        `json:"ignoreChangeType,omitempty"`
}

type MbpAddMemberRightGoodsListResponse struct {
	FailGoods    []int64 `json:"failGoods,omitempty"`
	SuccessGoods []int64 `json:"successGoods,omitempty"`
}

func CreateMbpAddMemberRightGoodsListRequest() (request *MbpAddMemberRightGoodsListRequest) {
	request = &MbpAddMemberRightGoodsListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "mbp/addMemberRightGoodsList", "api")
	request.Method = api.POST
	return
}

func CreateMbpAddMemberRightGoodsListResponse() (response *api.BaseResponse[MbpAddMemberRightGoodsListResponse]) {
	response = api.CreateResponse[MbpAddMemberRightGoodsListResponse](&MbpAddMemberRightGoodsListResponse{})
	return
}

type MbpAddMemberRightGoodsListRequestGoodsList struct {
	SkuLIst []MbpAddMemberRightGoodsListRequestSkuLIst `json:"skuLIst,omitempty"`
	GoodsId int64                                      `json:"goodsId,omitempty"`
}

type MbpAddMemberRightGoodsListRequestSkuLIst struct {
	Discount  int64 `json:"discount,omitempty"`
	SalePrice int64 `json:"salePrice,omitempty"`
	SkuId     int64 `json:"skuId,omitempty"`
}
