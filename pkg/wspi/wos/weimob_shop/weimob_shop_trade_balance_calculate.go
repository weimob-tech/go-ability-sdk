package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobShopTradeBalanceCalculateService struct {
	handler spi.WosInvocationHandler[PaasWeimobShopTradeBalanceCalculateRequest, PaasWeimobShopTradeBalanceCalculateResponse]
}

func (s PaasWeimobShopTradeBalanceCalculateService) NewRequest() spi.InvocationRequest[PaasWeimobShopTradeBalanceCalculateRequest] {
	return &spi.WosInvocationRequest[PaasWeimobShopTradeBalanceCalculateRequest]{
		Params: &PaasWeimobShopTradeBalanceCalculateRequest{},
	}
}

func (s PaasWeimobShopTradeBalanceCalculateService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobShopTradeBalanceCalculateRequest],
) (
	spi.InvocationResponse[PaasWeimobShopTradeBalanceCalculateResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobShopTradeBalanceCalculateRequest]))
}

type PaasWeimobShopTradeBalanceCalculateRequest struct {
	BasicInfo             PaasWeimobShopTradeBalanceCalculateRequestBasicInfo               `json:"basicInfo,omitempty"`
	ExtendInfo            PaasWeimobShopTradeBalanceCalculateRequestExtendInfo              `json:"extendInfo,omitempty"`
	GoodsTradeMetas       []PaasWeimobShopTradeBalanceCalculateRequestGoodsTradeMetas       `json:"goodsTradeMetas,omitempty"`
	InputUserDiscountList []PaasWeimobShopTradeBalanceCalculateRequestInputUserDiscountList `json:"inputUserDiscountList,omitempty"`
	Wid                   int64                                                             `json:"wid,omitempty"`
}
type PaasWeimobShopTradeBalanceCalculateRequestBasicInfo struct {
	BosId             int64  `json:"bosId,omitempty"`
	ProductId         int64  `json:"productId,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	ProductVersionId  string `json:"productVersionId,omitempty"`
	Tcode             string `json:"tcode,omitempty"`
	Vid               int64  `json:"vid,omitempty"`
	VidType           int64  `json:"vidType,omitempty"`
}
type PaasWeimobShopTradeBalanceCalculateRequestExtendInfo struct {
	GlobalTicket string `json:"globalTicket,omitempty"`
	SalesChannel int64  `json:"salesChannel,omitempty"`
	Source       int64  `json:"source,omitempty"`
	TradeScene   int64  `json:"tradeScene,omitempty"`
	TradeTrackId string `json:"tradeTrackId,omitempty"`
}
type PaasWeimobShopTradeBalanceCalculateRequestGoodsTradeMetas struct {
	BasicGoodsInfo        PaasWeimobShopTradeBalanceCalculateRequestBasicGoodsInfo          `json:"basicGoodsInfo,omitempty"`
	DiscountInfoList      []PaasWeimobShopTradeBalanceCalculateRequestDiscountInfoList      `json:"discountInfoList,omitempty"`
	InputItemDiscountList []PaasWeimobShopTradeBalanceCalculateRequestInputItemDiscountList `json:"inputItemDiscountList,omitempty"`
	Mutex                 PaasWeimobShopTradeBalanceCalculateRequestMutex                   `json:"mutex,omitempty"`
	PostageDiscountInfo   PaasWeimobShopTradeBalanceCalculateRequestPostageDiscountInfo     `json:"postageDiscountInfo,omitempty"`
	TradeGoodsInfo        PaasWeimobShopTradeBalanceCalculateRequestTradeGoodsInfo          `json:"tradeGoodsInfo,omitempty"`
}
type PaasWeimobShopTradeBalanceCalculateRequestBasicGoodsInfo struct {
	PreSellInfo   PaasWeimobShopTradeBalanceCalculateRequestPreSellInfo `json:"preSellInfo,omitempty"`
	PriceInfo     []PaasWeimobShopTradeBalanceCalculateRequestPriceInfo `json:"priceInfo,omitempty"`
	GoodsCode     string                                                `json:"goodsCode,omitempty"`
	GoodsId       int64                                                 `json:"goodsId,omitempty"`
	Num           int64                                                 `json:"num,omitempty"`
	SellModelType int64                                                 `json:"sellModelType,omitempty"`
	SkuId         int64                                                 `json:"skuId,omitempty"`
	Vid           int64                                                 `json:"vid,omitempty"`
	VidType       int64                                                 `json:"vidType,omitempty"`
	SkuCode       string                                                `json:"skuCode,omitempty"`
}
type PaasWeimobShopTradeBalanceCalculateRequestPreSellInfo struct {
	Type int64 `json:"type,omitempty"`
}
type PaasWeimobShopTradeBalanceCalculateRequestPriceInfo struct {
	Price string `json:"price,omitempty"`
	Type  int64  `json:"type,omitempty"`
}
type PaasWeimobShopTradeBalanceCalculateRequestDiscountInfoList struct {
	CheckAmount    string `json:"checkAmount,omitempty"`
	DiscountAmount string `json:"discountAmount,omitempty"`
	DiscountId     int64  `json:"discountId,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	DiscountLevel  int64  `json:"discountLevel,omitempty"`
	DiscountStatus int64  `json:"discountStatus,omitempty"`
	Order          int64  `json:"order,omitempty"`
	CostAmount     string `json:"costAmount,omitempty"`
}
type PaasWeimobShopTradeBalanceCalculateRequestInputItemDiscountList struct {
	CostAmount     string `json:"costAmount,omitempty"`
	DiscountId     string `json:"discountId,omitempty"`
	DiscountStatus int64  `json:"discountStatus,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	Order          int64  `json:"order,omitempty"`
	DiscountLevel  int64  `json:"discountLevel,omitempty"`
	DeductionType  int64  `json:"deductionType,omitempty"`
}
type PaasWeimobShopTradeBalanceCalculateRequestMutex struct {
	CustomDiscountType     int64   `json:"customDiscountType,omitempty"`
	CustomDiscountTypeList []int64 `json:"customDiscountTypeList,omitempty"`
}
type PaasWeimobShopTradeBalanceCalculateRequestPostageDiscountInfo struct {
	DiscountInfoList []PaasWeimobShopTradeBalanceCalculateRequestDiscountInfoList2 `json:"discountInfoList,omitempty"`
	Mutex            PaasWeimobShopTradeBalanceCalculateRequestMutex2              `json:"mutex,omitempty"`
}
type PaasWeimobShopTradeBalanceCalculateRequestDiscountInfoList2 struct {
}
type PaasWeimobShopTradeBalanceCalculateRequestMutex2 struct {
	CustomDiscountType     int64   `json:"customDiscountType,omitempty"`
	CustomDiscountTypeList []int64 `json:"customDiscountTypeList,omitempty"`
}
type PaasWeimobShopTradeBalanceCalculateRequestTradeGoodsInfo struct {
	ItemId         int64  `json:"itemId,omitempty"`
	TradeGoodsType int64  `json:"tradeGoodsType,omitempty"`
	TradePrice     string `json:"tradePrice,omitempty"`
	TradePriceType int64  `json:"tradePriceType,omitempty"`
}
type PaasWeimobShopTradeBalanceCalculateRequestInputUserDiscountList struct {
	DiscountId     string `json:"discountId,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	DiscountLevel  int64  `json:"discountLevel,omitempty"`
	DiscountStatus int64  `json:"discountStatus,omitempty"`
	Order          int64  `json:"order,omitempty"`
	CostAmount     string `json:"costAmount,omitempty"`
	DeductionType  int64  `json:"deductionType,omitempty"`
}

type PaasWeimobShopTradeBalanceCalculateResponse struct {
	GoodsTradeMetas    []PaasWeimobShopTradeBalanceCalculateResponseGoodsTradeMetas `json:"goodsTradeMetas,omitempty"`
	CalcRuleList       []PaasWeimobShopTradeBalanceCalculateResponseCalcRuleList    `json:"calcRuleList,omitempty"`
	GlobalTicket       string                                                       `json:"globalTicket,omitempty"`
	OutTicket          string                                                       `json:"outTicket,omitempty"`
	EnableInputBalance bool                                                         `json:"enableInputBalance,omitempty"`
}
type PaasWeimobShopTradeBalanceCalculateResponseGoodsTradeMetas struct {
	TradeGoodsInfo      PaasWeimobShopTradeBalanceCalculateResponseTradeGoodsInfo      `json:"tradeGoodsInfo,omitempty"`
	DiscountInfoList    []PaasWeimobShopTradeBalanceCalculateResponseDiscountInfoList  `json:"discountInfoList,omitempty"`
	PostageDiscountInfo PaasWeimobShopTradeBalanceCalculateResponsePostageDiscountInfo `json:"postageDiscountInfo,omitempty"`
}
type PaasWeimobShopTradeBalanceCalculateResponseTradeGoodsInfo struct {
	ItemId         int64  `json:"itemId,omitempty"`
	TradeGoodsType int64  `json:"tradeGoodsType,omitempty"`
	TradePrice     string `json:"tradePrice,omitempty"`
	TradePriceType int64  `json:"tradePriceType,omitempty"`
}
type PaasWeimobShopTradeBalanceCalculateResponseDiscountInfoList struct {
	CheckAmount    string `json:"checkAmount,omitempty"`
	DiscountAmount string `json:"discountAmount,omitempty"`
	DiscountId     string `json:"discountId,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	DiscountStatus int64  `json:"discountStatus,omitempty"`
	Order          int64  `json:"order,omitempty"`
	CostAmount     string `json:"costAmount,omitempty"`
}
type PaasWeimobShopTradeBalanceCalculateResponsePostageDiscountInfo struct {
	Mutex            PaasWeimobShopTradeBalanceCalculateResponseMutex               `json:"mutex,omitempty"`
	DiscountInfoList []PaasWeimobShopTradeBalanceCalculateResponseDiscountInfoList2 `json:"discountInfoList,omitempty"`
}
type PaasWeimobShopTradeBalanceCalculateResponseMutex struct {
	CustomDiscountType     int64   `json:"customDiscountType,omitempty"`
	CustomDiscountTypeList []int64 `json:"customDiscountTypeList,omitempty"`
}
type PaasWeimobShopTradeBalanceCalculateResponseDiscountInfoList2 struct {
	CheckAmount    string `json:"checkAmount,omitempty"`
	DiscountAmount string `json:"discountAmount,omitempty"`
	DiscountId     string `json:"discountId,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	DiscountStatus int64  `json:"discountStatus,omitempty"`
	Order          int64  `json:"order,omitempty"`
	CostAmount     string `json:"costAmount,omitempty"`
}
type PaasWeimobShopTradeBalanceCalculateResponseCalcRuleList struct {
	Mutex           PaasWeimobShopTradeBalanceCalculateResponseMutex2          `json:"mutex,omitempty"`
	DiscountUseInfo PaasWeimobShopTradeBalanceCalculateResponseDiscountUseInfo `json:"discountUseInfo,omitempty"`
	DiscountId      string                                                     `json:"discountId,omitempty"`
	DiscountType    string                                                     `json:"discountType,omitempty"`
	TotalCostAmount string                                                     `json:"totalCostAmount,omitempty"`
}
type PaasWeimobShopTradeBalanceCalculateResponseMutex2 struct {
	CustomDiscountType     int64   `json:"customDiscountType,omitempty"`
	CustomDiscountTypeList []int64 `json:"customDiscountTypeList,omitempty"`
}
type PaasWeimobShopTradeBalanceCalculateResponseDiscountUseInfo struct {
	UsedDiscountDetailList     []PaasWeimobShopTradeBalanceCalculateResponseUsedDiscountDetailList `json:"usedDiscountDetailList,omitempty"`
	MaxDiscountDetailList      []PaasWeimobShopTradeBalanceCalculateResponseMaxDiscountDetailList  `json:"maxDiscountDetailList,omitempty"`
	MaxDiscountTotalAmount     string                                                              `json:"maxDiscountTotalAmount,omitempty"`
	MaxDiscountTotalCostAmount string                                                              `json:"maxDiscountTotalCostAmount,omitempty"`
	UseDiscountTotalAmount     string                                                              `json:"useDiscountTotalAmount,omitempty"`
	UseDiscountTotalCostAmount string                                                              `json:"useDiscountTotalCostAmount,omitempty"`
	UseStatus                  bool                                                                `json:"useStatus,omitempty"`
	Toast                      string                                                              `json:"toast,omitempty"`
}
type PaasWeimobShopTradeBalanceCalculateResponseUsedDiscountDetailList struct {
	DiscountAmount string `json:"discountAmount,omitempty"`
	CostAmount     string `json:"costAmount,omitempty"`
	DeductionType  string `json:"deductionType,omitempty"`
}
type PaasWeimobShopTradeBalanceCalculateResponseMaxDiscountDetailList struct {
	DiscountAmount string `json:"discountAmount,omitempty"`
	CostAmount     string `json:"costAmount,omitempty"`
	DeductionType  int64  `json:"deductionType,omitempty"`
}

func CreatePaasWeimobShopTradeBalanceCalculateResponse() *PaasWeimobShopTradeBalanceCalculateResponse {
	return &PaasWeimobShopTradeBalanceCalculateResponse{}
}

// OnPaasWeimobShopTradeBalanceCalculateServiceInvocation
// @id 1433
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1433?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=余额批价)
func (s *Service) OnPaasWeimobShopTradeBalanceCalculateServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobShopTradeBalanceCalculateRequest, PaasWeimobShopTradeBalanceCalculateResponse],
) (service *Service) {
	var invocation = &PaasWeimobShopTradeBalanceCalculateService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobShopTradeBalanceCalculateRequest, PaasWeimobShopTradeBalanceCalculateResponse](invocation),
		"PaasWeimobShopTradeBalanceCalculateService",
		"weimobShop.trade.balance.calculate",
	)
	return s
}
