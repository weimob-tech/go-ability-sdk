package weimob_shop

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobShopPointPaasCalcPointService struct {
	handler spi.WosInvocationHandler[PaasWeimobShopPointPaasCalcPointRequest, PaasWeimobShopPointPaasCalcPointResponse]
}

func (s PaasWeimobShopPointPaasCalcPointService) NewRequest() spi.InvocationRequest[PaasWeimobShopPointPaasCalcPointRequest] {
	return &spi.WosInvocationRequest[PaasWeimobShopPointPaasCalcPointRequest]{
		Params: &PaasWeimobShopPointPaasCalcPointRequest{},
	}
}

func (s PaasWeimobShopPointPaasCalcPointService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobShopPointPaasCalcPointRequest],
) (
	spi.InvocationResponse[PaasWeimobShopPointPaasCalcPointResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobShopPointPaasCalcPointRequest]))
}

type PaasWeimobShopPointPaasCalcPointRequest struct {
	BasicInfo             PaasWeimobShopPointPaasCalcPointRequestBasicInfo               `json:"basicInfo,omitempty"`
	ExtendInfo            PaasWeimobShopPointPaasCalcPointRequestExtendInfo              `json:"extendInfo,omitempty"`
	GoodsTradeMetas       []PaasWeimobShopPointPaasCalcPointRequestGoodsTradeMetas       `json:"goodsTradeMetas,omitempty"`
	InputUserDiscountList []PaasWeimobShopPointPaasCalcPointRequestInputUserDiscountList `json:"inputUserDiscountList,omitempty"`
	Tcode                 string                                                         `json:"tcode,omitempty"`
	Wid                   int64                                                          `json:"wid,omitempty"`
	ProductId             int64                                                          `json:"productId,omitempty"`
	ProductInstanceId     int64                                                          `json:"productInstanceId,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointRequestBasicInfo struct {
	Vid               string `json:"vid,omitempty"`
	ProductId         int64  `json:"productId,omitempty"`
	BosId             int64  `json:"bosId,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	VidType           int64  `json:"vidType,omitempty"`
	ProductVersionId  int64  `json:"productVersionId,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointRequestExtendInfo struct {
	GlobalTicket string `json:"globalTicket,omitempty"`
	TradeTrackId string `json:"tradeTrackId,omitempty"`
	Source       int64  `json:"source,omitempty"`
	SalesChannel int64  `json:"salesChannel,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointRequestGoodsTradeMetas struct {
	InputItemDiscountList []PaasWeimobShopPointPaasCalcPointRequestInputItemDiscountList `json:"inputItemDiscountList,omitempty"`
	TradeGoodsInfo        PaasWeimobShopPointPaasCalcPointRequestTradeGoodsInfo          `json:"tradeGoodsInfo,omitempty"`
	DiscountInfoList      []PaasWeimobShopPointPaasCalcPointRequestDiscountInfoList      `json:"discountInfoList,omitempty"`
	Mutex                 PaasWeimobShopPointPaasCalcPointRequestMutex                   `json:"mutex,omitempty"`
	BasicGoodsInfo        PaasWeimobShopPointPaasCalcPointRequestBasicGoodsInfo          `json:"basicGoodsInfo,omitempty"`
	FeeTypeDiscountInfos  []PaasWeimobShopPointPaasCalcPointRequestFeeTypeDiscountInfos  `json:"feeTypeDiscountInfos,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointRequestInputItemDiscountList struct {
	DiscountLevel  int64  `json:"discountLevel,omitempty"`
	DiscountStatus int64  `json:"discountStatus,omitempty"`
	DeductionType  int64  `json:"deductionType,omitempty"`
	CostAmount     string `json:"costAmount,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	DiscountId     string `json:"discountId,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointRequestTradeGoodsInfo struct {
	ItemId         int64  `json:"itemId,omitempty"`
	TradeGoodsType int64  `json:"tradeGoodsType,omitempty"`
	TradePriceType int64  `json:"tradePriceType,omitempty"`
	TradePrice     string `json:"tradePrice,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointRequestDiscountInfoList struct {
	DiscountLevel  int64  `json:"discountLevel,omitempty"`
	CheckAmount    string `json:"checkAmount,omitempty"`
	DeductionType  int64  `json:"deductionType,omitempty"`
	DiscountAmount string `json:"discountAmount,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	DiscountId     string `json:"discountId,omitempty"`
	Order          int64  `json:"order,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointRequestMutex struct {
	CustomDiscountTypeList []int64 `json:"customDiscountTypeList,omitempty"`
	CustomDiscountType     int64   `json:"customDiscountType,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointRequestBasicGoodsInfo struct {
	PriceInfo     []PaasWeimobShopPointPaasCalcPointRequestPriceInfo `json:"priceInfo,omitempty"`
	PreSellInfo   PaasWeimobShopPointPaasCalcPointRequestPreSellInfo `json:"preSellInfo,omitempty"`
	Vid           int64                                              `json:"vid,omitempty"`
	SellModelType int64                                              `json:"sellModelType,omitempty"`
	GoodsId       int64                                              `json:"goodsId,omitempty"`
	Num           int64                                              `json:"num,omitempty"`
	GoodsCode     string                                             `json:"goodsCode,omitempty"`
	SkuId         int64                                              `json:"skuId,omitempty"`
	VidType       int64                                              `json:"vidType,omitempty"`
	SkuCode       string                                             `json:"skuCode,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointRequestPriceInfo struct {
	Price string `json:"price,omitempty"`
	Type  int64  `json:"type,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointRequestPreSellInfo struct {
	Type int64 `json:"type,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointRequestFeeTypeDiscountInfos struct {
	DiscountInfoList []PaasWeimobShopPointPaasCalcPointRequestDiscountInfoList2 `json:"discountInfoList,omitempty"`
	Mutex            PaasWeimobShopPointPaasCalcPointRequestMutex2              `json:"mutex,omitempty"`
	FeeType          int64                                                      `json:"feeType,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointRequestDiscountInfoList2 struct {
	CheckAmount    string `json:"checkAmount,omitempty"`
	DiscountAmount string `json:"discountAmount,omitempty"`
	CostAmount     string `json:"costAmount,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	DiscountId     string `json:"discountId,omitempty"`
	Order          int64  `json:"order,omitempty"`
	DiscountLevel  int64  `json:"discountLevel,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointRequestMutex2 struct {
	CustomDiscountTypeList []int64 `json:"customDiscountTypeList,omitempty"`
	CustomDiscountType     int64   `json:"customDiscountType,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointRequestInputUserDiscountList struct {
	DiscountId     string `json:"discountId,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	DeductionType  int64  `json:"deductionType,omitempty"`
	DiscountStatus int64  `json:"discountStatus,omitempty"`
	CostAmount     string `json:"costAmount,omitempty"`
	Order          int64  `json:"order,omitempty"`
	DiscountLevel  int64  `json:"discountLevel,omitempty"`
}

type PaasWeimobShopPointPaasCalcPointResponse struct {
	CalcRuleList    []PaasWeimobShopPointPaasCalcPointResponseCalcRuleList    `json:"calcRuleList,omitempty"`
	GoodsTradeMetas []PaasWeimobShopPointPaasCalcPointResponseGoodsTradeMetas `json:"goodsTradeMetas,omitempty"`
	GlobalTicket    string                                                    `json:"globalTicket,omitempty"`
	OutTicket       string                                                    `json:"outTicket,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointResponseCalcRuleList struct {
	DiscountUseInfo PaasWeimobShopPointPaasCalcPointResponseDiscountUseInfo `json:"discountUseInfo,omitempty"`
	Mutex           PaasWeimobShopPointPaasCalcPointResponseMutex           `json:"mutex,omitempty"`
	DiscountType    int64                                                   `json:"discountType,omitempty"`
	DiscountId      string                                                  `json:"discountId,omitempty"`
	TotalCostAmount string                                                  `json:"totalCostAmount,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointResponseDiscountUseInfo struct {
	UsedDiscountDetailList     []PaasWeimobShopPointPaasCalcPointResponseUsedDiscountDetailList `json:"usedDiscountDetailList,omitempty"`
	MaxDiscountDetailList      []PaasWeimobShopPointPaasCalcPointResponseMaxDiscountDetailList  `json:"maxDiscountDetailList,omitempty"`
	MaxDiscountTotalCostAmount string                                                           `json:"maxDiscountTotalCostAmount,omitempty"`
	MaxDiscountTotalAmount     string                                                           `json:"maxDiscountTotalAmount,omitempty"`
	UseDiscountTotalAmount     string                                                           `json:"useDiscountTotalAmount,omitempty"`
	UseDiscountTotalCostAmount string                                                           `json:"useDiscountTotalCostAmount,omitempty"`
	UseStatus                  bool                                                             `json:"useStatus,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointResponseUsedDiscountDetailList struct {
	DeductionType  int64  `json:"deductionType,omitempty"`
	DiscountAmount string `json:"discountAmount,omitempty"`
	CostAmount     string `json:"costAmount,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointResponseMaxDiscountDetailList struct {
	DeductionType  int64  `json:"deductionType,omitempty"`
	DiscountAmount string `json:"discountAmount,omitempty"`
	CostAmount     string `json:"costAmount,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointResponseMutex struct {
	CustomDiscountTypeList string `json:"customDiscountTypeList,omitempty"`
	CustomDiscountType     int64  `json:"customDiscountType,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointResponseGoodsTradeMetas struct {
	TradeGoodsInfo       PaasWeimobShopPointPaasCalcPointResponseTradeGoodsInfo         `json:"tradeGoodsInfo,omitempty"`
	DiscountInfoList     []PaasWeimobShopPointPaasCalcPointResponseDiscountInfoList     `json:"discountInfoList,omitempty"`
	FeeTypeDiscountInfos []PaasWeimobShopPointPaasCalcPointResponseFeeTypeDiscountInfos `json:"feeTypeDiscountInfos,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointResponseTradeGoodsInfo struct {
	ItemId         int64  `json:"itemId,omitempty"`
	TradeGoodsType int64  `json:"tradeGoodsType,omitempty"`
	TradePriceType int64  `json:"tradePriceType,omitempty"`
	TradePrice     string `json:"tradePrice,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointResponseDiscountInfoList struct {
	CheckAmount    string `json:"checkAmount,omitempty"`
	DiscountAmount string `json:"discountAmount,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	DiscountId     int64  `json:"discountId,omitempty"`
	Order          int64  `json:"order,omitempty"`
	CostAmount     string `json:"costAmount,omitempty"`
	DiscountLevel  int64  `json:"discountLevel,omitempty"`
	DiscountStatus int64  `json:"discountStatus,omitempty"`
	DeductionType  int64  `json:"deductionType,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointResponseFeeTypeDiscountInfos struct {
	DiscountInfoList []PaasWeimobShopPointPaasCalcPointResponseDiscountInfoList2 `json:"discountInfoList,omitempty"`
	Mutex            PaasWeimobShopPointPaasCalcPointResponseMutex2              `json:"mutex,omitempty"`
	FeeType          int64                                                       `json:"feeType,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointResponseDiscountInfoList2 struct {
	CheckAmount    string `json:"checkAmount,omitempty"`
	DiscountAmount string `json:"discountAmount,omitempty"`
	CostAmount     string `json:"costAmount,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	DiscountId     string `json:"discountId,omitempty"`
	Order          int64  `json:"order,omitempty"`
	DiscountLevel  int64  `json:"discountLevel,omitempty"`
	DiscountStatus int64  `json:"discountStatus,omitempty"`
	DeductionType  int64  `json:"deductionType,omitempty"`
}
type PaasWeimobShopPointPaasCalcPointResponseMutex2 struct {
	CustomDiscountTypeList []int64 `json:"customDiscountTypeList,omitempty"`
	CustomDiscountType     int64   `json:"customDiscountType,omitempty"`
}

func CreatePaasWeimobShopPointPaasCalcPointResponse() *PaasWeimobShopPointPaasCalcPointResponse {
	return &PaasWeimobShopPointPaasCalcPointResponse{}
}

// OnPaasWeimobShopPointPaasCalcPointServiceInvocation
// @id 1168
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1168?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=积分批价能力)
func (s *Service) OnPaasWeimobShopPointPaasCalcPointServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobShopPointPaasCalcPointRequest, PaasWeimobShopPointPaasCalcPointResponse],
) (service *Service) {
	var invocation = &PaasWeimobShopPointPaasCalcPointService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobShopPointPaasCalcPointRequest, PaasWeimobShopPointPaasCalcPointResponse](invocation),
		"PaasWeimobShopPointPaasCalcPointService",
		"weimobShop.point.paasCalcPoint",
	)
	return s
}
