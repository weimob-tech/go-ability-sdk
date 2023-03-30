package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobShopPayPassPaycheckService struct {
	handler spi.WosInvocationHandler[PaasWeimobShopPayPassPaycheckRequest, PaasWeimobShopPayPassPaycheckResponse]
}

func (s PaasWeimobShopPayPassPaycheckService) NewRequest() spi.InvocationRequest[PaasWeimobShopPayPassPaycheckRequest] {
	return &spi.WosInvocationRequest[PaasWeimobShopPayPassPaycheckRequest]{
		Params: &PaasWeimobShopPayPassPaycheckRequest{},
	}
}

func (s PaasWeimobShopPayPassPaycheckService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobShopPayPassPaycheckRequest],
) (
	spi.InvocationResponse[PaasWeimobShopPayPassPaycheckResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobShopPayPassPaycheckRequest]))
}

type PaasWeimobShopPayPassPaycheckRequest struct {
	Orders                  []PaasWeimobShopPayPassPaycheckRequestOrders `json:"orders,omitempty"`
	ParentOrderNo           int64                                        `json:"parentOrderNo,omitempty"`
	MerchantId              int64                                        `json:"merchantId,omitempty"`
	BosId                   int64                                        `json:"bosId,omitempty"`
	ProductId               int64                                        `json:"productId,omitempty"`
	ProductInstanceId       int64                                        `json:"productInstanceId,omitempty"`
	TargetProductId         int64                                        `json:"targetProductId,omitempty"`
	TargetProductInstanceId int64                                        `json:"targetProductInstanceId,omitempty"`
	Tcode                   string                                       `json:"tcode,omitempty"`
}
type PaasWeimobShopPayPassPaycheckRequestOrders struct {
	OrderBase PaasWeimobShopPayPassPaycheckRequestOrderBase `json:"orderBase,omitempty"`
	Items     []PaasWeimobShopPayPassPaycheckRequestItems   `json:"items,omitempty"`
	Pay       PaasWeimobShopPayPassPaycheckRequestPay       `json:"pay,omitempty"`
	Merchant  PaasWeimobShopPayPassPaycheckRequestMerchant  `json:"merchant,omitempty"`
	Buyer     PaasWeimobShopPayPassPaycheckRequestBuyer     `json:"buyer,omitempty"`
}
type PaasWeimobShopPayPassPaycheckRequestOrderBase struct {
	OrderNo     int64 `json:"orderNo,omitempty"`
	PaymentType int64 `json:"paymentType,omitempty"`
	OrderType   int64 `json:"orderType,omitempty"`
	CreateTime  int64 `json:"createTime,omitempty"`
}
type PaasWeimobShopPayPassPaycheckRequestItems struct {
	Goods    PaasWeimobShopPayPassPaycheckRequestGoods    `json:"goods,omitempty"`
	PayInfo  PaasWeimobShopPayPassPaycheckRequestPayInfo  `json:"payInfo,omitempty"`
	BizInfos PaasWeimobShopPayPassPaycheckRequestBizInfos `json:"bizInfos,omitempty"`
	ItemExt  PaasWeimobShopPayPassPaycheckRequestItemExt  `json:"itemExt,omitempty"`
	ItemId   int64                                        `json:"itemId,omitempty"`
}
type PaasWeimobShopPayPassPaycheckRequestGoods struct {
	PriceInfos      []PaasWeimobShopPayPassPaycheckRequestPriceInfos `json:"priceInfos,omitempty"`
	GoodsExt        PaasWeimobShopPayPassPaycheckRequestGoodsExt     `json:"goodsExt,omitempty"`
	GoodsId         int64                                            `json:"goodsId,omitempty"`
	SkuNum          int64                                            `json:"skuNum,omitempty"`
	GoodsType       int64                                            `json:"goodsType,omitempty"`
	SubGoodsType    int64                                            `json:"subGoodsType,omitempty"`
	SalePrice       string                                           `json:"salePrice,omitempty"`
	GoodsCode       string                                           `json:"goodsCode,omitempty"`
	DeductStockType int64                                            `json:"deductStockType,omitempty"`
	SkuCode         string                                           `json:"skuCode,omitempty"`
	SkuId           int64                                            `json:"skuId,omitempty"`
	SkuBarCode      string                                           `json:"skuBarCode,omitempty"`
}
type PaasWeimobShopPayPassPaycheckRequestPriceInfos struct {
	Type   int64  `json:"type,omitempty"`
	Amount string `json:"amount,omitempty"`
}
type PaasWeimobShopPayPassPaycheckRequestGoodsExt struct {
}
type PaasWeimobShopPayPassPaycheckRequestPayInfo struct {
}
type PaasWeimobShopPayPassPaycheckRequestBizInfos struct {
}
type PaasWeimobShopPayPassPaycheckRequestItemExt struct {
}
type PaasWeimobShopPayPassPaycheckRequestPay struct {
	AmountInfos         []PaasWeimobShopPayPassPaycheckRequestAmountInfos `json:"amountInfos,omitempty"`
	PayAmount           string                                            `json:"payAmount,omitempty"`
	TotalAmount         string                                            `json:"totalAmount,omitempty"`
	TotalDiscountAmount string                                            `json:"totalDiscountAmount,omitempty"`
	ShouldPayAmount     string                                            `json:"shouldPayAmount,omitempty"`
}
type PaasWeimobShopPayPassPaycheckRequestAmountInfos struct {
	Type            int64  `json:"type,omitempty"`
	Amount          string `json:"amount,omitempty"`
	PayAmount       string `json:"payAmount,omitempty"`
	ShouldPayAmount string `json:"shouldPayAmount,omitempty"`
}
type PaasWeimobShopPayPassPaycheckRequestMerchant struct {
	BosId          int64 `json:"bosId,omitempty"`
	Vid            int64 `json:"vid,omitempty"`
	VidType        int64 `json:"vidType,omitempty"`
	ProcessVid     int64 `json:"processVid,omitempty"`
	ProcessVidType int64 `json:"processVidType,omitempty"`
}
type PaasWeimobShopPayPassPaycheckRequestBuyer struct {
	Wid int64 `json:"wid,omitempty"`
}

type PaasWeimobShopPayPassPaycheckResponse struct {
	CheckPayOrders  []PaasWeimobShopPayPassPaycheckResponseCheckPayOrders `json:"checkPayOrders,omitempty"`
	EnablePay       bool                                                  `json:"enablePay,omitempty"`
	DisableMessages []int64                                               `json:"disableMessages,omitempty"`
}
type PaasWeimobShopPayPassPaycheckResponseCheckPayOrders struct {
	CheckPayItems   []PaasWeimobShopPayPassPaycheckResponseCheckPayItems `json:"checkPayItems,omitempty"`
	OrderNo         string                                               `json:"orderNo,omitempty"`
	EnablePay       bool                                                 `json:"enablePay,omitempty"`
	DisableMessages []int64                                              `json:"disableMessages,omitempty"`
}
type PaasWeimobShopPayPassPaycheckResponseCheckPayItems struct {
	ItemId          int64   `json:"itemId,omitempty"`
	EnablePay       bool    `json:"enablePay,omitempty"`
	DisableMessages []int64 `json:"disableMessages,omitempty"`
}

func CreatePaasWeimobShopPayPassPaycheckResponse() *PaasWeimobShopPayPassPaycheckResponse {
	return &PaasWeimobShopPayPassPaycheckResponse{}
}

// OnPaasWeimobShopPayPassPaycheckServiceInvocation
// @id 899
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/899?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=支付校验)
func (s *Service) OnPaasWeimobShopPayPassPaycheckServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobShopPayPassPaycheckRequest, PaasWeimobShopPayPassPaycheckResponse],
) (service *Service) {
	var invocation = &PaasWeimobShopPayPassPaycheckService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobShopPayPassPaycheckRequest, PaasWeimobShopPayPassPaycheckResponse](invocation),
		"PaasWeimobShopPayPassPaycheckService",
		"weimobShop.pay.passPaycheck",
	)
	return s
}
