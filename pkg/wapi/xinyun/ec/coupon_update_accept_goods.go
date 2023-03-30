package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponUpdateAcceptGoods
// @id 1697
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新优惠券模板适用商品)
func (client *Client) CouponUpdateAcceptGoods(request *CouponUpdateAcceptGoodsRequest) (response *CouponUpdateAcceptGoodsResponse, err error) {
	rpcResponse := CreateCouponUpdateAcceptGoodsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponUpdateAcceptGoodsRequest struct {
	*api.BaseRequest

	AcceptGoodsIds  []string `json:"acceptGoodsIds,omitempty"`
	StoreId         int64    `json:"storeId,omitempty"`
	CardTemplateId  int64    `json:"cardTemplateId,omitempty"`
	AcceptGoodsType int      `json:"acceptGoodsType,omitempty"`
}

type CouponUpdateAcceptGoodsResponse struct {
}

func CreateCouponUpdateAcceptGoodsRequest() (request *CouponUpdateAcceptGoodsRequest) {
	request = &CouponUpdateAcceptGoodsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "coupon/updateAcceptGoods", "api")
	request.Method = api.POST
	return
}

func CreateCouponUpdateAcceptGoodsResponse() (response *api.BaseResponse[CouponUpdateAcceptGoodsResponse]) {
	response = api.CreateResponse[CouponUpdateAcceptGoodsResponse](&CouponUpdateAcceptGoodsResponse{})
	return
}
