package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsUpdate
// @id 1644
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1644?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新在售商品)
func (client *Client) GoodsUpdate(request *GoodsUpdateRequest) (response *GoodsUpdateResponse, err error) {
	rpcResponse := CreateGoodsUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsUpdateRequest struct {
	*api.BaseRequest

	BasicInfo                    GoodsUpdateRequestBasicInfo                      `json:"basicInfo,omitempty"`
	CategoryPropertyList         []GoodsUpdateRequestCategoryPropertyList         `json:"categoryPropertyList,omitempty"`
	GoodsPropertyList            []GoodsUpdateRequestGoodsPropertyList            `json:"goodsPropertyList,omitempty"`
	GoodsRightsList              []GoodsUpdateRequestGoodsRightsList              `json:"goodsRightsList,omitempty"`
	InnerPropertyList            []GoodsUpdateRequestInnerPropertyList            `json:"innerPropertyList,omitempty"`
	LimitInfo                    GoodsUpdateRequestLimitInfo                      `json:"limitInfo,omitempty"`
	PerformanceWay               GoodsUpdateRequestPerformanceWay                 `json:"performanceWay,omitempty"`
	SpecInfoList                 []GoodsUpdateRequestSpecInfoList                 `json:"specInfoList,omitempty"`
	SkuList                      []GoodsUpdateRequestSkuList                      `json:"skuList,omitempty"`
	UnifiedCyclePurchaseInfoList []GoodsUpdateRequestUnifiedCyclePurchaseInfoList `json:"unifiedCyclePurchaseInfoList,omitempty"`
	SizeTableInfo                GoodsUpdateRequestSizeTableInfo                  `json:"sizeTableInfo,omitempty"`
	GoodsMsgInfoList             []GoodsUpdateRequestGoodsMsgInfoList             `json:"goodsMsgInfoList,omitempty"`
	BrandId                      int64                                            `json:"brandId,omitempty"`
	CategoryId                   int64                                            `json:"categoryId,omitempty"`
	DeductStockType              int64                                            `json:"deductStockType,omitempty"`
	DefaultImageUrl              string                                           `json:"defaultImageUrl,omitempty"`
	GoodsClassifyIdList          []int64                                          `json:"goodsClassifyIdList,omitempty"`
	GoodsDesc                    string                                           `json:"goodsDesc,omitempty"`
	GoodsId                      int64                                            `json:"goodsId,omitempty"`
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

type GoodsUpdateResponse struct {
	SkuList []GoodsUpdateResponseSkuList `json:"skuList,omitempty"`
	GoodsId int64                        `json:"goodsId,omitempty"`
}

func CreateGoodsUpdateRequest() (request *GoodsUpdateRequest) {
	request = &GoodsUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/update", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsUpdateResponse() (response *api.BaseResponse[GoodsUpdateResponse]) {
	response = api.CreateResponse[GoodsUpdateResponse](&GoodsUpdateResponse{})
	return
}

type GoodsUpdateRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type GoodsUpdateRequestCategoryPropertyList struct {
	CategoryId  int64 `json:"categoryId,omitempty"`
	PropId      int64 `json:"propId,omitempty"`
	PropValueId int64 `json:"propValueId,omitempty"`
}

type GoodsUpdateRequestGoodsPropertyList struct {
	PropId      int64 `json:"propId,omitempty"`
	PropValueId int64 `json:"propValueId,omitempty"`
}

type GoodsUpdateRequestGoodsRightsList struct {
	PropId      int64 `json:"propId,omitempty"`
	PropValueId int64 `json:"propValueId,omitempty"`
}

type GoodsUpdateRequestInnerPropertyList struct {
	PropId      int64 `json:"propId,omitempty"`
	PropValueId int64 `json:"propValueId,omitempty"`
}

type GoodsUpdateRequestLimitInfo struct {
	GoodsLimitNum   int64 `json:"goodsLimitNum,omitempty"`
	GoodsLimitCycle int64 `json:"goodsLimitCycle,omitempty"`
}

type GoodsUpdateRequestPerformanceWay struct {
	DeliveryList         []GoodsUpdateRequestDeliveryList       `json:"deliveryList,omitempty"`
	MallCycleGoodsConfig GoodsUpdateRequestMallCycleGoodsConfig `json:"mallCycleGoodsConfig,omitempty"`
}

type GoodsUpdateRequestDeliveryList struct {
	DeliveryId         int64 `json:"deliveryId,omitempty"`
	DeliveryNodeShipId int64 `json:"deliveryNodeShipId,omitempty"`
	DeliveryType       int64 `json:"deliveryType,omitempty"`
	TemplateId         int64 `json:"templateId,omitempty"`
}

type GoodsUpdateRequestMallCycleGoodsConfig struct {
	FulfillCycleConfig GoodsUpdateRequestFulfillCycleConfig `json:"fulfillCycleConfig,omitempty"`
	FulfillCycleInt    int64                                `json:"fulfillCycleInt,omitempty"`
	DelayFulfill       int64                                `json:"delayFulfill,omitempty"`
	DelayCurrFulfill   int64                                `json:"delayCurrFulfill,omitempty"`
	DelayDay           int64                                `json:"delayDay,omitempty"`
	ClosingTime        int64                                `json:"closingTime,omitempty"`
	ClosingTimeDetail  string                               `json:"closingTimeDetail,omitempty"`
	AlterAddress       int64                                `json:"alterAddress,omitempty"`
}

type GoodsUpdateRequestFulfillCycleConfig struct {
	CycleUnit          int64   `json:"cycleUnit,omitempty"`
	CycleValues        []int64 `json:"cycleValues,omitempty"`
	CycleInterval      int64   `json:"cycleInterval,omitempty"`
	CycleTimes         int64   `json:"cycleTimes,omitempty"`
	RemoveWeekend      int64   `json:"removeWeekend	,omitempty"`
	FulfillTimeOptions []int64 `json:"fulfillTimeOptions,omitempty"`
}

type GoodsUpdateRequestSpecInfoList struct {
	SkuSpecValueList    []GoodsUpdateRequestSkuSpecValueList `json:"skuSpecValueList,omitempty"`
	SpecId              int64                                `json:"specId,omitempty"`
	SpecName            string                               `json:"specName,omitempty"`
	SpecImgEnable       bool                                 `json:"specImgEnable,omitempty"`
	IsOpenSizeRecommend bool                                 `json:"isOpenSizeRecommend,omitempty"`
}

type GoodsUpdateRequestSkuSpecValueList struct {
	SpecValueName string `json:"specValueName,omitempty"`
	ImageUrl      string `json:"imageUrl,omitempty"`
	SpecValueId   int64  `json:"specValueId,omitempty"`
	Sort          int64  `json:"sort,omitempty"`
}

type GoodsUpdateRequestSkuList struct {
	SkuSpecValueList []GoodsUpdateRequestSkuSpecValueList2 `json:"skuSpecValueList,omitempty"`
	CycleSkuList     []GoodsUpdateRequestCycleSkuList      `json:"cycleSkuList,omitempty"`
	SkuId            int64                                 `json:"skuId,omitempty"`
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

type GoodsUpdateRequestSkuSpecValueList2 struct {
	SpecId      int64 `json:"specId,omitempty"`
	SpecValueId int64 `json:"specValueId,omitempty"`
}

type GoodsUpdateRequestCycleSkuList struct {
	CycleNum          int64  `json:"cycleNum,omitempty"`
	CycleName         string `json:"cycleName,omitempty"`
	CycleMarkingPrice int64  `json:"cycleMarkingPrice,omitempty"`
	CyclePrice        int64  `json:"cyclePrice,omitempty"`
	CycleStockNum     int64  `json:"cycleStockNum,omitempty"`
}

type GoodsUpdateRequestUnifiedCyclePurchaseInfoList struct {
	CycleNum          int64  `json:"cycleNum,omitempty"`
	CycleName         string `json:"cycleName,omitempty"`
	CycleMarkingPrice int64  `json:"cycleMarkingPrice,omitempty"`
	CyclePrice        int64  `json:"cyclePrice,omitempty"`
	CycleStockNum     int64  `json:"cycleStockNum,omitempty"`
	IsValid           bool   `json:"isValid,omitempty"`
}

type GoodsUpdateRequestSizeTableInfo struct {
	ParamList           []GoodsUpdateRequestParamList           `json:"paramList,omitempty"`
	SizeTableDetailList []GoodsUpdateRequestSizeTableDetailList `json:"sizeTableDetailList,omitempty"`
	SpecId              int64                                   `json:"specId,omitempty"`
	SpecName            string                                  `json:"specName,omitempty"`
}

type GoodsUpdateRequestParamList struct {
	IsDefaultSelected bool   `json:"isDefaultSelected,omitempty"`
	ParamName         string `json:"paramName,omitempty"`
	IsSystemParam     bool   `json:"isSystemParam,omitempty"`
}

type GoodsUpdateRequestSizeTableDetailList struct {
	ParamValueList []GoodsUpdateRequestParamValueList `json:"paramValueList,omitempty"`
	SpecValue      int64                              `json:"specValue,omitempty"`
	SpecValueId    string                             `json:"specValueId,omitempty"`
}

type GoodsUpdateRequestParamValueList struct {
	ParamName  string `json:"paramName,omitempty"`
	ParamValue int64  `json:"paramValue,omitempty"`
}

type GoodsUpdateRequestGoodsMsgInfoList struct {
	Name             string  `json:"name,omitempty"`
	Type             int64   `json:"type,omitempty"`
	Required         bool    `json:"required,omitempty"`
	Desc             string  `json:"desc,omitempty"`
	Sort             int64   `json:"sort,omitempty"`
	CustomOptionList []int64 `json:"customOptionList,omitempty"`
}

type GoodsUpdateResponseSkuList struct {
	OuterSkuCode string `json:"outerSkuCode,omitempty"`
	SkuId        int64  `json:"skuId,omitempty"`
}
