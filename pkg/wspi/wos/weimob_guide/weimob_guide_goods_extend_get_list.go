package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideGoodsExtendGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideGoodsExtendGetListRequest, PaasWeimobGuideGoodsExtendGetListResponse]
}

func (s PaasWeimobGuideGoodsExtendGetListService) NewRequest() spi.InvocationRequest[PaasWeimobGuideGoodsExtendGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideGoodsExtendGetListRequest]{
		Params: &PaasWeimobGuideGoodsExtendGetListRequest{},
	}
}

func (s PaasWeimobGuideGoodsExtendGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideGoodsExtendGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideGoodsExtendGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideGoodsExtendGetListRequest]))
}

type PaasWeimobGuideGoodsExtendGetListRequest struct {
	QueryParameter PaasWeimobGuideGoodsExtendGetListRequestQueryParameter `json:"queryParameter,omitempty"`
	PageNum        int64                                                  `json:"pageNum,omitempty"`
	PageSize       int64                                                  `json:"pageSize,omitempty"`
}
type PaasWeimobGuideGoodsExtendGetListRequestQueryParameter struct {
	ClassifyId           int64   `json:"classifyId,omitempty"`
	CategoryId           int64   `json:"categoryId,omitempty"`
	InnerPropertyValueId int64   `json:"innerPropertyValueId,omitempty"`
	PropertyValueId      int64   `json:"propertyValueId,omitempty"`
	IsCanSell            int64   `json:"isCanSell,omitempty"`
	IsCombine            int64   `json:"isCombine,omitempty"`
	GoodsStatus          int64   `json:"goodsStatus,omitempty"`
	SaleChannelType      []int64 `json:"saleChannelType,omitempty"`
	Source               int64   `json:"source,omitempty"`
	GoodsBelongToType    int64   `json:"goodsBelongToType,omitempty"`
	Search               string  `json:"search,omitempty"`
	SearchType           int64   `json:"searchType,omitempty"`
	GoodsIdList          []int64 `json:"goodsIdList,omitempty"`
	GoodsTagId           int64   `json:"goodsTagId,omitempty"`
	RightsValueId        int64   `json:"rightsValueId,omitempty"`
	GoodsTypes           []int64 `json:"goodsTypes,omitempty"`
	SubGoodsTypes        []int64 `json:"subGoodsTypes,omitempty"`
	SoldType             []int64 `json:"soldType,omitempty"`
	Sort                 int64   `json:"sort,omitempty"`
	SortType             int64   `json:"sortType,omitempty"`
	MinSalePrice         int64   `json:"minSalePrice,omitempty"`
	MaxSalePrice         int64   `json:"maxSalePrice,omitempty"`
}

type PaasWeimobGuideGoodsExtendGetListResponse struct {
	PageList   []PaasWeimobGuideGoodsExtendGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                               `json:"pageNum,omitempty"`
	PageSize   int64                                               `json:"pageSize,omitempty"`
	TotalCount int64                                               `json:"totalCount,omitempty"`
}
type PaasWeimobGuideGoodsExtendGetListResponsePageList struct {
	GoodsPrice        PaasWeimobGuideGoodsExtendGetListResponseGoodsPrice `json:"goodsPrice,omitempty"`
	GoodsStock        PaasWeimobGuideGoodsExtendGetListResponseGoodsStock `json:"goodsStock,omitempty"`
	Vid               int64                                               `json:"vid,omitempty"`
	CreateVid         int64                                               `json:"createVid,omitempty"`
	GoodsType         int64                                               `json:"goodsType,omitempty"`
	SubGoodsType      int64                                               `json:"subGoodsType,omitempty"`
	GoodsId           int64                                               `json:"goodsId,omitempty"`
	Title             string                                              `json:"title,omitempty"`
	SoldType          int64                                               `json:"soldType,omitempty"`
	RealSaleNum       int64                                               `json:"realSaleNum,omitempty"`
	DefaultImageUrl   string                                              `json:"defaultImageUrl,omitempty"`
	SaleChannelType   int64                                               `json:"saleChannelType,omitempty"`
	IsCanSell         bool                                                `json:"isCanSell,omitempty"`
	IsOnline          bool                                                `json:"isOnline,omitempty"`
	IsMultiSku        bool                                                `json:"isMultiSku,omitempty"`
	IsAssigned        bool                                                `json:"isAssigned,omitempty"`
	Sort              int64                                               `json:"sort,omitempty"`
	Source            int64                                               `json:"source,omitempty"`
	ProductId         int64                                               `json:"productId,omitempty"`
	ProductInstanceId int64                                               `json:"productInstanceId,omitempty"`
}
type PaasWeimobGuideGoodsExtendGetListResponseGoodsPrice struct {
	MaxSalePrice   int64 `json:"maxSalePrice,omitempty"`
	MinSalePrice   int64 `json:"minSalePrice,omitempty"`
	MaxCostPrice   int64 `json:"maxCostPrice,omitempty"`
	MinCostPrice   int64 `json:"minCostPrice,omitempty"`
	MaxMarketPrice int64 `json:"maxMarketPrice,omitempty"`
	MinMarketPrice int64 `json:"minMarketPrice,omitempty"`
}
type PaasWeimobGuideGoodsExtendGetListResponseGoodsStock struct {
	Stock                  []PaasWeimobGuideGoodsExtendGetListResponseStock `json:"stock,omitempty"`
	ExistNullStock         bool                                             `json:"existNullStock,omitempty"`
	IsAllStockEmpty        bool                                             `json:"isAllStockEmpty,omitempty"`
	GoodsStockNum          int64                                            `json:"goodsStockNum,omitempty"`
	FrozenQuantity         int64                                            `json:"frozenQuantity,omitempty"`
	AllowOversoldQuantity  int64                                            `json:"allowOversoldQuantity,omitempty"`
	AlreadySellQuantity    int64                                            `json:"alreadySellQuantity,omitempty"`
	PreGoodsStockNum       int64                                            `json:"preGoodsStockNum,omitempty"`
	PreFrozenQuantity      int64                                            `json:"preFrozenQuantity,omitempty"`
	PreAlreadySellQuantity int64                                            `json:"preAlreadySellQuantity,omitempty"`
}
type PaasWeimobGuideGoodsExtendGetListResponseStock struct {
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

func CreatePaasWeimobGuideGoodsExtendGetListResponse() *PaasWeimobGuideGoodsExtendGetListResponse {
	return &PaasWeimobGuideGoodsExtendGetListResponse{}
}

// OnPaasWeimobGuideGoodsExtendGetListServiceInvocation
// @id 738
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/738?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询商品列表)
func (s *Service) OnPaasWeimobGuideGoodsExtendGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideGoodsExtendGetListRequest, PaasWeimobGuideGoodsExtendGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideGoodsExtendGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideGoodsExtendGetListRequest, PaasWeimobGuideGoodsExtendGetListResponse](invocation),
		"PaasWeimobGuideGoodsExtendGetListService",
		"weimobGuide.goods.extend.getList",
	)
	return s
}
