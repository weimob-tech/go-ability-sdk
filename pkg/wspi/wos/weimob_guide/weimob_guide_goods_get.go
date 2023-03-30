package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideGoodsGetService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideGoodsGetRequest, PaasWeimobGuideGoodsGetResponse]
}

func (s PaasWeimobGuideGoodsGetService) NewRequest() spi.InvocationRequest[PaasWeimobGuideGoodsGetRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideGoodsGetRequest]{
		Params: &PaasWeimobGuideGoodsGetRequest{},
	}
}

func (s PaasWeimobGuideGoodsGetService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideGoodsGetRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideGoodsGetResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideGoodsGetRequest]))
}

type PaasWeimobGuideGoodsGetRequest struct {
	FilterOption PaasWeimobGuideGoodsGetRequestFilterOption `json:"filterOption,omitempty"`
	GoodsIdList  []int64                                    `json:"goodsIdList,omitempty"`
	ActivityId   int64                                      `json:"activityId,omitempty"`
	ActivityType int64                                      `json:"activityType,omitempty"`
}
type PaasWeimobGuideGoodsGetRequestFilterOption struct {
	IsQueryTagInfo        bool `json:"isQueryTagInfo,omitempty"`
	IsQueryClassifyInfo   bool `json:"isQueryClassifyInfo,omitempty"`
	IsQuerySaleInfo       bool `json:"isQuerySaleInfo,omitempty"`
	IsQueryPropertyInfo   bool `json:"isQueryPropertyInfo,omitempty"`
	IsQueryCategoryInfo   bool `json:"isQueryCategoryInfo,omitempty"`
	IsQueryBrandInfo      bool `json:"isQueryBrandInfo,omitempty"`
	IsFilterDeleted       bool `json:"isFilterDeleted,omitempty"`
	IsQueryGoodsStockInfo bool `json:"isQueryGoodsStockInfo,omitempty"`
	IsQueryGoodsPriceInfo bool `json:"isQueryGoodsPriceInfo,omitempty"`
	IsQuerySkuInfo        bool `json:"isQuerySkuInfo,omitempty"`
	IsQueryActivityInfo   bool `json:"isQueryActivityInfo,omitempty"`
	IsQueryActivity       bool `json:"isQueryActivity,omitempty"`
	IsQueryActivityPrice  bool `json:"isQueryActivityPrice,omitempty"`
	GenerateLinkFlag      bool `json:"generateLinkFlag,omitempty"`
	IsQueryItemSkuInfo    bool `json:"isQueryItemSkuInfo,omitempty"`
	IsQueryMember         bool `json:"isQueryMember,omitempty"`
}

type PaasWeimobGuideGoodsGetResponse struct {
	GoodsStock         PaasWeimobGuideGoodsGetResponseGoodsStock    `json:"goodsStock,omitempty"`
	GoodsLimit         PaasWeimobGuideGoodsGetResponseGoodsLimit    `json:"goodsLimit,omitempty"`
	PreSellGoods       PaasWeimobGuideGoodsGetResponsePreSellGoods  `json:"preSellGoods,omitempty"`
	SkuSpecList        []PaasWeimobGuideGoodsGetResponseSkuSpecList `json:"skuSpecList,omitempty"`
	GoodsTag           PaasWeimobGuideGoodsGetResponseGoodsTag      `json:"goodsTag,omitempty"`
	SkuList            []PaasWeimobGuideGoodsGetResponseSkuList     `json:"skuList,omitempty"`
	Vid                int64                                        `json:"vid,omitempty"`
	CreateVid          int64                                        `json:"createVid,omitempty"`
	SubGoodsType       int64                                        `json:"subGoodsType,omitempty"`
	OuterGoodsCode     string                                       `json:"outerGoodsCode,omitempty"`
	GoodsId            int64                                        `json:"goodsId,omitempty"`
	Title              string                                       `json:"title,omitempty"`
	SubTitle           string                                       `json:"subTitle,omitempty"`
	SoldType           int64                                        `json:"soldType,omitempty"`
	RealSaleNum        int64                                        `json:"realSaleNum,omitempty"`
	DefaultImageUrl    string                                       `json:"defaultImageUrl,omitempty"`
	SaleChannelType    int64                                        `json:"saleChannelType,omitempty"`
	IsCanSell          bool                                         `json:"isCanSell,omitempty"`
	IsOnline           bool                                         `json:"isOnline,omitempty"`
	Sort               int64                                        `json:"sort,omitempty"`
	IsDeleted          bool                                         `json:"isDeleted,omitempty"`
	Source             int64                                        `json:"source,omitempty"`
	GoodsImageUrl      []int64                                      `json:"goodsImageUrl,omitempty"`
	GoodsVideoUrl      string                                       `json:"goodsVideoUrl,omitempty"`
	GoodsVideoImageUrl string                                       `json:"goodsVideoImageUrl,omitempty"`
	GoodsDesc          string                                       `json:"goodsDesc,omitempty"`
	OuterGoodsId       int64                                        `json:"outerGoodsId,omitempty"`
	DeductStockType    int64                                        `json:"deductStockType,omitempty"`
	CategoryId         int64                                        `json:"categoryId,omitempty"`
	LimitSwitch        bool                                         `json:"limitSwitch,omitempty"`
	GoodsShowType      int64                                        `json:"goodsShowType,omitempty"`
}
type PaasWeimobGuideGoodsGetResponseGoodsStock struct {
	Stock            []PaasWeimobGuideGoodsGetResponseStock `json:"stock,omitempty"`
	ExistNullStock   bool                                   `json:"existNullStock,omitempty"`
	IsAllStockEmpty  bool                                   `json:"isAllStockEmpty,omitempty"`
	GoodsStockNum    int64                                  `json:"goodsStockNum,omitempty"`
	PreGoodsStockNum int64                                  `json:"preGoodsStockNum,omitempty"`
}
type PaasWeimobGuideGoodsGetResponseStock struct {
	ChannelType            int64 `json:"channelType,omitempty"`
	ExistNullStock         bool  `json:"existNullStock,omitempty"`
	IsAllStockEmpty        bool  `json:"isAllStockEmpty,omitempty"`
	GoodsStockNum          int64 `json:"goodsStockNum,omitempty"`
	FrozenQuantity         int64 `json:"frozenQuantity,omitempty"`
	AllowOversoldQuantity  int64 `json:"allowOversoldQuantity,omitempty"`
	AlreadySellQuantity    int64 `json:"alreadySellQuantity,omitempty"`
	PreGoodsStockNum       int64 `json:"preGoodsStockNum,omitempty"`
	PreFrozenQuantity      int64 `json:"preFrozenQuantity,omitempty"`
	PreAlreadySellQuantity int64 `json:"preAlreadySellQuantity,omitempty"`
}
type PaasWeimobGuideGoodsGetResponseGoodsLimit struct {
	GoodsLimitNum   int64 `json:"goodsLimitNum,omitempty"`
	GoodsLimitCycle int64 `json:"goodsLimitCycle,omitempty"`
}
type PaasWeimobGuideGoodsGetResponsePreSellGoods struct {
	FinalPaymentInfo    PaasWeimobGuideGoodsGetResponseFinalPaymentInfo `json:"finalPaymentInfo,omitempty"`
	DeliveryInfo        PaasWeimobGuideGoodsGetResponseDeliveryInfo     `json:"deliveryInfo,omitempty"`
	GoodsId             int64                                           `json:"goodsId,omitempty"`
	DepositPayType      int64                                           `json:"depositPayType,omitempty"`
	Deposit             int64                                           `json:"deposit,omitempty"`
	IgnoreChangeType    int64                                           `json:"ignoreChangeType,omitempty"`
	PresellTimeType     int64                                           `json:"presellTimeType,omitempty"`
	StartPresellTime    int64                                           `json:"startPresellTime,omitempty"`
	EndPresellTime      int64                                           `json:"endPresellTime,omitempty"`
	ExpireProcessType   int64                                           `json:"expireProcessType,omitempty"`
	SelectWarehouseType int64                                           `json:"selectWarehouseType,omitempty"`
	WarehouseId         int64                                           `json:"warehouseId,omitempty"`
	WarehouseName       string                                          `json:"warehouseName,omitempty"`
	GoodsLimitRule      int64                                           `json:"goodsLimitRule,omitempty"`
}
type PaasWeimobGuideGoodsGetResponseFinalPaymentInfo struct {
	FinalPaymentType          int64 `json:"finalPaymentType,omitempty"`
	StartFinalPayTime         int64 `json:"startFinalPayTime,omitempty"`
	EndFinalPayTime           int64 `json:"endFinalPayTime,omitempty"`
	AfterDepositStartFinalDay int64 `json:"afterDepositStartFinalDay,omitempty"`
	AfterDepositEndFinalDay   int64 `json:"afterDepositEndFinalDay,omitempty"`
	AfterDepositEndFinalHour  int64 `json:"afterDepositEndFinalHour,omitempty"`
}
type PaasWeimobGuideGoodsGetResponseDeliveryInfo struct {
	DeliveryType              int64 `json:"deliveryType,omitempty"`
	AfterFinalPayDeliveryDay  int64 `json:"afterFinalPayDeliveryDay,omitempty"`
	AfterFinalPayDeliveryHour int64 `json:"afterFinalPayDeliveryHour,omitempty"`
	ScheduleDeliveryTime      int64 `json:"scheduleDeliveryTime,omitempty"`
}
type PaasWeimobGuideGoodsGetResponseSkuSpecList struct {
	SpecId        int64  `json:"specId,omitempty"`
	SpecName      string `json:"specName,omitempty"`
	SpecImgEnable bool   `json:"specImgEnable,omitempty"`
}
type PaasWeimobGuideGoodsGetResponseGoodsTag struct {
	TagId int64  `json:"tagId,omitempty"`
	Name  string `json:"name,omitempty"`
}
type PaasWeimobGuideGoodsGetResponseSkuList struct {
	SkuSpecValueList []PaasWeimobGuideGoodsGetResponseSkuSpecValueList `json:"skuSpecValueList,omitempty"`
	GoodsId          int64                                             `json:"goodsId,omitempty"`
	SkuId            int64                                             `json:"skuId,omitempty"`
	CostPrice        int64                                             `json:"costPrice,omitempty"`
	SalePrice        int64                                             `json:"salePrice,omitempty"`
}
type PaasWeimobGuideGoodsGetResponseSkuSpecValueList struct {
	SpecId        int64  `json:"specId,omitempty"`
	SpecValueName string `json:"specValueName,omitempty"`
	SpecName      string `json:"specName,omitempty"`
	SpecValueId   int64  `json:"specValueId,omitempty"`
}

func CreatePaasWeimobGuideGoodsGetResponse() *PaasWeimobGuideGoodsGetResponse {
	return &PaasWeimobGuideGoodsGetResponse{}
}

// OnPaasWeimobGuideGoodsGetServiceInvocation
// @id 740
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/740?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询商品信息)
func (s *Service) OnPaasWeimobGuideGoodsGetServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideGoodsGetRequest, PaasWeimobGuideGoodsGetResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideGoodsGetService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideGoodsGetRequest, PaasWeimobGuideGoodsGetResponse](invocation),
		"PaasWeimobGuideGoodsGetService",
		"weimobGuide.goods.get",
	)
	return s
}
