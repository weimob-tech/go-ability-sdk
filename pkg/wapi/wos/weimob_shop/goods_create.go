package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsCreate
// @id 1643
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1643?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=添加在售商品)
func (client *Client) GoodsCreate(request *GoodsCreateRequest) (response *GoodsCreateResponse, err error) {
	rpcResponse := CreateGoodsCreateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsCreateRequest struct {
	*api.BaseRequest

	BasicInfo                    GoodsCreateRequestBasicInfo                      `json:"basicInfo,omitempty"`
	CategoryPropertyList         []GoodsCreateRequestCategoryPropertyList         `json:"categoryPropertyList,omitempty"`
	GoodsPropertyList            []GoodsCreateRequestGoodsPropertyList            `json:"goodsPropertyList,omitempty"`
	GoodsRightsList              []GoodsCreateRequestGoodsRightsList              `json:"goodsRightsList,omitempty"`
	InnerPropertyList            []GoodsCreateRequestInnerPropertyList            `json:"innerPropertyList,omitempty"`
	LimitInfo                    GoodsCreateRequestLimitInfo                      `json:"limitInfo,omitempty"`
	PerformanceWay               GoodsCreateRequestPerformanceWay                 `json:"performanceWay,omitempty"`
	SpecInfoList                 []GoodsCreateRequestSpecInfoList                 `json:"specInfoList,omitempty"`
	SkuList                      []GoodsCreateRequestSkuList                      `json:"skuList,omitempty"`
	UnifiedCyclePurchaseInfoList []GoodsCreateRequestUnifiedCyclePurchaseInfoList `json:"unifiedCyclePurchaseInfoList,omitempty"`
	SizeTableInfo                GoodsCreateRequestSizeTableInfo                  `json:"sizeTableInfo,omitempty"`
	GoodsMsgInfoList             []GoodsCreateRequestGoodsMsgInfoList             `json:"goodsMsgInfoList,omitempty"`
	PointDeductRule              GoodsCreateRequestPointDeductRule                `json:"pointDeductRule,omitempty"`
	BrandId                      int64                                            `json:"brandId,omitempty"`
	CategoryId                   int64                                            `json:"categoryId,omitempty"`
	DeductStockType              int64                                            `json:"deductStockType,omitempty"`
	DefaultImageUrl              string                                           `json:"defaultImageUrl,omitempty"`
	GoodsClassifyIdList          []int64                                          `json:"goodsClassifyIdList,omitempty"`
	GoodsDesc                    string                                           `json:"goodsDesc,omitempty"`
	GoodsImageUrl                []int64                                          `json:"goodsImageUrl,omitempty"`
	GoodsTemplateId              int64                                            `json:"goodsTemplateId,omitempty"`
	GoodsType                    int64                                            `json:"goodsType,omitempty"`
	GoodsVideoImageUrl           string                                           `json:"goodsVideoImageUrl,omitempty"`
	GoodsVideoUrl                string                                           `json:"goodsVideoUrl,omitempty"`
	InitSales                    int64                                            `json:"initSales,omitempty"`
	InvoiceTitle                 string                                           `json:"invoiceTitle,omitempty"`
	InvoiceTitleType             int64                                            `json:"invoiceTitleType,omitempty"`
	IsCanSell                    bool                                             `json:"isCanSell,omitempty"`
	IsMultiSku                   bool                                             `json:"isMultiSku,omitempty"`
	IsOnline                     bool                                             `json:"isOnline,omitempty"`
	LimitSwitch                  bool                                             `json:"limitSwitch,omitempty"`
	OuterGoodsCode               string                                           `json:"outerGoodsCode,omitempty"`
	SellUnitId                   int64                                            `json:"sellUnitId,omitempty"`
	SubGoodsType                 int64                                            `json:"subGoodsType,omitempty"`
	SubTitle                     string                                           `json:"subTitle,omitempty"`
	TagId                        int64                                            `json:"tagId,omitempty"`
	Title                        string                                           `json:"title,omitempty"`
	Wid                          int64                                            `json:"wid,omitempty"`
	GoodsDeliveryMode            int64                                            `json:"goodsDeliveryMode,omitempty"`
	SkuCycleSetting              int64                                            `json:"skuCycleSetting,omitempty"`
	CycleDeliveryNum             int64                                            `json:"cycleDeliveryNum,omitempty"`
	SaleChannelType              int64                                            `json:"saleChannelType,omitempty"`
	CategoryType                 int64                                            `json:"categoryType,omitempty"`
}

type GoodsCreateResponse struct {
	SkuList []GoodsCreateResponseSkuList `json:"skuList,omitempty"`
	GoodsId int64                        `json:"goodsId,omitempty"`
}

func CreateGoodsCreateRequest() (request *GoodsCreateRequest) {
	request = &GoodsCreateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/create", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsCreateResponse() (response *api.BaseResponse[GoodsCreateResponse]) {
	response = api.CreateResponse[GoodsCreateResponse](&GoodsCreateResponse{})
	return
}

type GoodsCreateRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type GoodsCreateRequestCategoryPropertyList struct {
	CategoryId  int64 `json:"categoryId,omitempty"`
	PropId      int64 `json:"propId,omitempty"`
	PropValueId int64 `json:"propValueId,omitempty"`
}

type GoodsCreateRequestGoodsPropertyList struct {
	PropId      int64 `json:"propId,omitempty"`
	PropValueId int64 `json:"propValueId,omitempty"`
}

type GoodsCreateRequestGoodsRightsList struct {
	PropId      int64 `json:"propId,omitempty"`
	PropValueId int64 `json:"propValueId,omitempty"`
}

type GoodsCreateRequestInnerPropertyList struct {
	PropId      int64 `json:"propId,omitempty"`
	PropValueId int64 `json:"propValueId,omitempty"`
}

type GoodsCreateRequestLimitInfo struct {
	GoodsLimitNum   int64 `json:"goodsLimitNum,omitempty"`
	GoodsLimitCycle int64 `json:"goodsLimitCycle,omitempty"`
}

type GoodsCreateRequestPerformanceWay struct {
	DeliveryList         []GoodsCreateRequestDeliveryList       `json:"deliveryList,omitempty"`
	MallCycleGoodsConfig GoodsCreateRequestMallCycleGoodsConfig `json:"mallCycleGoodsConfig,omitempty"`
}

type GoodsCreateRequestDeliveryList struct {
	DeliveryId         int64 `json:"deliveryId,omitempty"`
	DeliveryNodeShipId int64 `json:"deliveryNodeShipId,omitempty"`
	DeliveryType       int64 `json:"deliveryType,omitempty"`
	TemplateId         int64 `json:"templateId,omitempty"`
}

type GoodsCreateRequestMallCycleGoodsConfig struct {
	FulfillCycleConfig GoodsCreateRequestFulfillCycleConfig `json:"fulfillCycleConfig,omitempty"`
	FulfillCycleInt    int64                                `json:"fulfillCycleInt,omitempty"`
	DelayFulfill       int64                                `json:"delayFulfill,omitempty"`
	DelayCurrFulfill   int64                                `json:"delayCurrFulfill,omitempty"`
	DelayDay           int64                                `json:"delayDay,omitempty"`
	ClosingTime        int64                                `json:"closingTime,omitempty"`
	ClosingTimeDetail  string                               `json:"closingTimeDetail,omitempty"`
	AlterAddress       int64                                `json:"alterAddress,omitempty"`
}

type GoodsCreateRequestFulfillCycleConfig struct {
	CycleUnit          int64   `json:"cycleUnit,omitempty"`
	CycleValues        []int64 `json:"cycleValues,omitempty"`
	CycleInterval      int64   `json:"cycleInterval,omitempty"`
	CycleTimes         int64   `json:"cycleTimes,omitempty"`
	RemoveWeekend      int64   `json:"removeWeekend,omitempty"`
	FulfillTimeOptions []int64 `json:"fulfillTimeOptions,omitempty"`
}

type GoodsCreateRequestSpecInfoList struct {
	SkuSpecValueList    []GoodsCreateRequestSkuSpecValueList `json:"skuSpecValueList,omitempty"`
	SpecId              int64                                `json:"specId,omitempty"`
	SpecName            string                               `json:"specName,omitempty"`
	SpecImgEnable       bool                                 `json:"specImgEnable,omitempty"`
	IsOpenSizeRecommend bool                                 `json:"isOpenSizeRecommend,omitempty"`
}

type GoodsCreateRequestSkuSpecValueList struct {
	SpecValueName string `json:"specValueName,omitempty"`
	ImageUrl      string `json:"imageUrl,omitempty"`
	SpecValueId   int64  `json:"specValueId,omitempty"`
	Sort          int64  `json:"sort,omitempty"`
}

type GoodsCreateRequestSkuList struct {
	CycleSkuList     []GoodsCreateRequestCycleSkuList      `json:"cycleSkuList,omitempty"`
	SkuSpecValueList []GoodsCreateRequestSkuSpecValueList2 `json:"skuSpecValueList,omitempty"`
	ImageUrl         string                                `json:"imageUrl,omitempty"`
	CostPrice        int64                                 `json:"costPrice,omitempty"`
	SkuStockNum      int64                                 `json:"skuStockNum,omitempty"`
	MarketPrice      int64                                 `json:"marketPrice,omitempty"`
	OuterSkuCode     string                                `json:"outerSkuCode,omitempty"`
	SalePrice        int64                                 `json:"salePrice,omitempty"`
	Volume           int64                                 `json:"volume,omitempty"`
	Weight           int64                                 `json:"weight,omitempty"`
	SkuBarCode       string                                `json:"skuBarCode,omitempty"`
}

type GoodsCreateRequestCycleSkuList struct {
	CycleNum          int64  `json:"cycleNum	,omitempty"`
	CycleName         string `json:"cycleName,omitempty"`
	CycleMarkingPrice int64  `json:"cycleMarkingPrice,omitempty"`
	CyclePrice        int64  `json:"cyclePrice,omitempty"`
	CycleStockNum     int64  `json:"cycleStockNum,omitempty"`
}

type GoodsCreateRequestSkuSpecValueList2 struct {
	SpecId      int64 `json:"specId,omitempty"`
	SpecValueId int64 `json:"specValueId,omitempty"`
}

type GoodsCreateRequestUnifiedCyclePurchaseInfoList struct {
	CycleNum          int64  `json:"cycleNum,omitempty"`
	CycleName         string `json:"cycleName,omitempty"`
	CycleMarkingPrice int64  `json:"cycleMarkingPrice,omitempty"`
	CyclePrice        int64  `json:"cyclePrice,omitempty"`
	CycleStockNum     int64  `json:"cycleStockNum,omitempty"`
	IsValid           bool   `json:"isValid,omitempty"`
}

type GoodsCreateRequestSizeTableInfo struct {
	ParamList           []GoodsCreateRequestParamList           `json:"paramList,omitempty"`
	SizeTableDetailList []GoodsCreateRequestSizeTableDetailList `json:"sizeTableDetailList,omitempty"`
	SpecId              int64                                   `json:"specId,omitempty"`
	SpecName            string                                  `json:"specName,omitempty"`
}

type GoodsCreateRequestParamList struct {
	ParamName         string `json:"paramName,omitempty"`
	IsSystemParam     bool   `json:"isSystemParam,omitempty"`
	IsDefaultSelected bool   `json:"isDefaultSelected,omitempty"`
}

type GoodsCreateRequestSizeTableDetailList struct {
	ParamValueList []GoodsCreateRequestParamValueList `json:"paramValueList,omitempty"`
	SpecValueId    int64                              `json:"specValueId,omitempty"`
	SpecValue      string                             `json:"specValue,omitempty"`
}

type GoodsCreateRequestParamValueList struct {
	ParamName  string `json:"paramName,omitempty"`
	ParamValue int64  `json:"paramValue,omitempty"`
}

type GoodsCreateRequestGoodsMsgInfoList struct {
	Name             string  `json:"name,omitempty"`
	Type             int64   `json:"type,omitempty"`
	Required         bool    `json:"required,omitempty"`
	Desc             string  `json:"desc,omitempty"`
	Sort             int64   `json:"sort,omitempty"`
	CustomOptionList []int64 `json:"customOptionList,omitempty"`
}

type GoodsCreateRequestPointDeductRule struct {
	OpenPointDeduct int64 `json:"openPointDeduct,omitempty"`
	UseCondition    int64 `json:"useCondition,omitempty"`
	UseDiscount     int64 `json:"useDiscount,omitempty"`
	DeductAmount    int64 `json:"deductAmount,omitempty"`
	DeductDiscount  int64 `json:"deductDiscount,omitempty"`
	DeductionType   int64 `json:"deductionType,omitempty"`
}

type GoodsCreateResponseSkuList struct {
	OuterSkuCode string `json:"outerSkuCode,omitempty"`
	SkuId        int64  `json:"skuId,omitempty"`
}
