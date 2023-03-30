package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponLock
// @id 1142
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1142?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=批量锁券)
func (client *Client) CouponLock(request *CouponLockRequest) (response *CouponLockResponse, err error) {
	rpcResponse := CreateCouponLockResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponLockRequest struct {
	*api.BaseRequest

	CouponList         []CouponLockRequestCouponList         `json:"couponList,omitempty"`
	StoreGoodsInfoDTOS []CouponLockRequestStoreGoodsInfoDTOS `json:"storeGoodsInfoDTOS,omitempty"`
	UseScene           int64                                 `json:"useScene,omitempty"`
	OrderAmount        string                                `json:"orderAmount,omitempty"`
	IsIgnoreGoodsInfo  bool                                  `json:"isIgnoreGoodsInfo,omitempty"`
	Wid                int64                                 `json:"wid,omitempty"`
	VidType            int64                                 `json:"vidType,omitempty"`
	Vid                int64                                 `json:"vid,omitempty"`
}

type CouponLockResponse struct {
}

func CreateCouponLockRequest() (request *CouponLockRequest) {
	request = &CouponLockRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/lock", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponLockResponse() (response *api.BaseResponse[CouponLockResponse]) {
	response = api.CreateResponse[CouponLockResponse](&CouponLockResponse{})
	return
}

type CouponLockRequestCouponList struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	OrderNo          string `json:"orderNo,omitempty"`
	Code             string `json:"code,omitempty"`
	Amount           int64  `json:"amount,omitempty"`
}

type CouponLockRequestStoreGoodsInfoDTOS struct {
	VidNodeList CouponLockRequestVidNodeList  `json:"vidNodeList,omitempty"`
	GoodsInfos  []CouponLockRequestGoodsInfos `json:"goodsInfos,omitempty"`
}

type CouponLockRequestVidNodeList struct {
	BosId   int64 `json:"bosId,omitempty"`
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type CouponLockRequestGoodsInfos struct {
	Skus        []CouponLockRequestSkus `json:"skus,omitempty"`
	Id          int64                   `json:"id,omitempty"`
	CategoryIds []int64                 `json:"categoryIds,omitempty"`
	Tags        []int64                 `json:"tags,omitempty"`
	GroupIds    []int64                 `json:"groupIds,omitempty"`
}

type CouponLockRequestSkus struct {
	Id       int64  `json:"id,omitempty"`
	Price    string `json:"price,omitempty"`
	Num      int64  `json:"num,omitempty"`
	Selected bool   `json:"selected,omitempty"`
}
