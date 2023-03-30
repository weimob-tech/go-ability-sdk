package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpQueryMemberGoodsDiscount
// @id 2344
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询会员权益商品折扣)
func (client *Client) MbpQueryMemberGoodsDiscount(request *MbpQueryMemberGoodsDiscountRequest) (response *MbpQueryMemberGoodsDiscountResponse, err error) {
	rpcResponse := CreateMbpQueryMemberGoodsDiscountResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpQueryMemberGoodsDiscountRequest struct {
	*api.BaseRequest

	RuleId      int     `json:"ruleId,omitempty"`
	GoodsIdList []int64 `json:"goodsIdList,omitempty"`
}

type MbpQueryMemberGoodsDiscountResponse struct {
}

func CreateMbpQueryMemberGoodsDiscountRequest() (request *MbpQueryMemberGoodsDiscountRequest) {
	request = &MbpQueryMemberGoodsDiscountRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "mbp/queryMemberGoodsDiscount", "api")
	request.Method = api.POST
	return
}

func CreateMbpQueryMemberGoodsDiscountResponse() (response *api.BaseResponse[MbpQueryMemberGoodsDiscountResponse]) {
	response = api.CreateResponse[MbpQueryMemberGoodsDiscountResponse](&MbpQueryMemberGoodsDiscountResponse{})
	return
}
