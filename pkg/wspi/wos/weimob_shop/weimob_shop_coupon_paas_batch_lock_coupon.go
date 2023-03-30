package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobShopCouponPaasBatchLockCouponService struct {
	handler spi.WosInvocationHandler[PaasWeimobShopCouponPaasBatchLockCouponRequest, PaasWeimobShopCouponPaasBatchLockCouponResponse]
}

func (s PaasWeimobShopCouponPaasBatchLockCouponService) NewRequest() spi.InvocationRequest[PaasWeimobShopCouponPaasBatchLockCouponRequest] {
	return &spi.WosInvocationRequest[PaasWeimobShopCouponPaasBatchLockCouponRequest]{
		Params: &PaasWeimobShopCouponPaasBatchLockCouponRequest{},
	}
}

func (s PaasWeimobShopCouponPaasBatchLockCouponService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobShopCouponPaasBatchLockCouponRequest],
) (
	spi.InvocationResponse[PaasWeimobShopCouponPaasBatchLockCouponResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobShopCouponPaasBatchLockCouponRequest]))
}

type PaasWeimobShopCouponPaasBatchLockCouponRequest struct {
	CouponList         []PaasWeimobShopCouponPaasBatchLockCouponRequestCouponList         `json:"couponList,omitempty"`
	StoreGoodsInfoDTOS []PaasWeimobShopCouponPaasBatchLockCouponRequestStoreGoodsInfoDTOS `json:"storeGoodsInfoDTOS,omitempty"`
	Wid                int64                                                              `json:"wid,omitempty"`
	OrderAmount        int64                                                              `json:"orderAmount,omitempty"`
	IsIgnoreGoodsInfo  bool                                                               `json:"isIgnoreGoodsInfo,omitempty"`
	UseScene           int64                                                              `json:"useScene,omitempty"`
}
type PaasWeimobShopCouponPaasBatchLockCouponRequestCouponList struct {
	BosId            int64  `json:"bosId,omitempty"`
	Vid              int64  `json:"vid,omitempty"`
	Wid              int64  `json:"wid,omitempty"`
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	OrderId          string `json:"orderId,omitempty"`
	OrderNo          string `json:"orderNo,omitempty"`
	Code             string `json:"code,omitempty"`
	Amount           int64  `json:"amount,omitempty"`
	UseScene         int64  `json:"useScene,omitempty"`
}
type PaasWeimobShopCouponPaasBatchLockCouponRequestStoreGoodsInfoDTOS struct {
	VidNodeList PaasWeimobShopCouponPaasBatchLockCouponRequestVidNodeList  `json:"vidNodeList,omitempty"`
	GoodsInfos  []PaasWeimobShopCouponPaasBatchLockCouponRequestGoodsInfos `json:"goodsInfos,omitempty"`
}
type PaasWeimobShopCouponPaasBatchLockCouponRequestVidNodeList struct {
	ParentVids []PaasWeimobShopCouponPaasBatchLockCouponRequestParentVids `json:"parentVids,omitempty"`
}
type PaasWeimobShopCouponPaasBatchLockCouponRequestParentVids struct {
	MerchantId int64  `json:"merchantId,omitempty"`
	Vid        int64  `json:"vid,omitempty"`
	VidName    string `json:"vidName,omitempty"`
	VidType    int64  `json:"vidType,omitempty"`
}
type PaasWeimobShopCouponPaasBatchLockCouponRequestGoodsInfos struct {
	Skus        []PaasWeimobShopCouponPaasBatchLockCouponRequestSkus `json:"skus,omitempty"`
	Id          int64                                                `json:"id,omitempty"`
	CategoryIds []int64                                              `json:"categoryIds,omitempty"`
	Tags        []int64                                              `json:"tags,omitempty"`
	GroupIds    []int64                                              `json:"groupIds,omitempty"`
}
type PaasWeimobShopCouponPaasBatchLockCouponRequestSkus struct {
}

type PaasWeimobShopCouponPaasBatchLockCouponResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreatePaasWeimobShopCouponPaasBatchLockCouponResponse() *PaasWeimobShopCouponPaasBatchLockCouponResponse {
	return &PaasWeimobShopCouponPaasBatchLockCouponResponse{}
}

// OnPaasWeimobShopCouponPaasBatchLockCouponServiceInvocation
// @id 580
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/580?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=锁定用户优惠券)
func (s *Service) OnPaasWeimobShopCouponPaasBatchLockCouponServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobShopCouponPaasBatchLockCouponRequest, PaasWeimobShopCouponPaasBatchLockCouponResponse],
) (service *Service) {
	var invocation = &PaasWeimobShopCouponPaasBatchLockCouponService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobShopCouponPaasBatchLockCouponRequest, PaasWeimobShopCouponPaasBatchLockCouponResponse](invocation),
		"PaasWeimobShopCouponPaasBatchLockCouponService",
		"weimobShop.coupon.paasBatchLockCoupon",
	)
	return s
}
