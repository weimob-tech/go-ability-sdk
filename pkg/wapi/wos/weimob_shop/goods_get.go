package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsGet
// @id 1646
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1646?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询在售商品详情)
func (client *Client) GoodsGet(request *GoodsGetRequest) (response *GoodsGetResponse, err error) {
	rpcResponse := CreateGoodsGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsGetRequest struct {
	*api.BaseRequest

	BasicInfo GoodsGetRequestBasicInfo `json:"basicInfo,omitempty"`
	GoodsId   int64                    `json:"goodsId,omitempty"`
}

type GoodsGetResponse struct {
	PerformanceWay       GoodsGetResponsePerformanceWay         `json:"performanceWay,omitempty"`
	CategoryList         []GoodsGetResponseCategoryList         `json:"categoryList,omitempty"`
	BrandInfo            GoodsGetResponseBrandInfo              `json:"brandInfo,omitempty"`
	InnerPropertyList    []GoodsGetResponseInnerPropertyList    `json:"innerPropertyList,omitempty"`
	GoodsPropertyList    []GoodsGetResponseGoodsPropertyList    `json:"goodsPropertyList,omitempty"`
	GoodsRightsList      []GoodsGetResponseGoodsRightsList      `json:"goodsRightsList,omitempty"`
	CategoryPropertyList []GoodsGetResponseCategoryPropertyList `json:"categoryPropertyList,omitempty"`
	TagInfo              GoodsGetResponseTagInfo                `json:"tagInfo,omitempty"`
	SpecInfoList         []GoodsGetResponseSpecInfoList         `json:"specInfoList,omitempty"`
	PointDeductRule      GoodsGetResponsePointDeductRule        `json:"pointDeductRule,omitempty"`
	GoodsClassifyList    []GoodsGetResponseGoodsClassifyList    `json:"goodsClassifyList,omitempty"`
	SkuList              []GoodsGetResponseSkuList              `json:"skuList,omitempty"`
	SizeTableInfo        GoodsGetResponseSizeTableInfo          `json:"sizeTableInfo,omitempty"`
	IsMultiSku           bool                                   `json:"isMultiSku,omitempty"`
	SoldType             int64                                  `json:"soldType,omitempty"`
	LimitSwitch          bool                                   `json:"limitSwitch,omitempty"`
	GoodsTemplateId      int64                                  `json:"goodsTemplateId,omitempty"`
	GoodsId              int64                                  `json:"goodsId,omitempty"`
	IsOnline             bool                                   `json:"isOnline,omitempty"`
	Title                string                                 `json:"title,omitempty"`
	IsCanSell            bool                                   `json:"isCanSell,omitempty"`
	Vid                  int64                                  `json:"vid,omitempty"`
	SubGoodsType         int64                                  `json:"subGoodsType,omitempty"`
	SubTitle             string                                 `json:"subTitle,omitempty"`
	SellUnitTitle        string                                 `json:"sellUnitTitle,omitempty"`
	SellUnitId           int64                                  `json:"sellUnitId,omitempty"`
	CreateVid            int64                                  `json:"createVid,omitempty"`
	GoodsDesc            string                                 `json:"goodsDesc,omitempty"`
	IsCombine            bool                                   `json:"isCombine,omitempty"`
	Sort                 int64                                  `json:"sort,omitempty"`
	OuterGoodsCode       string                                 `json:"outerGoodsCode,omitempty"`
	GoodsVideoUrl        string                                 `json:"goodsVideoUrl,omitempty"`
	GoodsVideoImageUrl   string                                 `json:"goodsVideoImageUrl,omitempty"`
	DefaultImageUrl      string                                 `json:"defaultImageUrl,omitempty"`
	SaleChannelType      int64                                  `json:"saleChannelType,omitempty"`
	GoodsType            int64                                  `json:"goodsType,omitempty"`
	DeductStockType      int64                                  `json:"deductStockType,omitempty"`
	GoodsImageUrl        []int64                                `json:"goodsImageUrl,omitempty"`
	LatestFulfillDate    int64                                  `json:"latestFulfillDate,omitempty"`
}

func CreateGoodsGetRequest() (request *GoodsGetRequest) {
	request = &GoodsGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/get", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsGetResponse() (response *api.BaseResponse[GoodsGetResponse]) {
	response = api.CreateResponse[GoodsGetResponse](&GoodsGetResponse{})
	return
}

type GoodsGetRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type GoodsGetResponsePerformanceWay struct {
	DeliveryList         []GoodsGetResponseDeliveryList       `json:"deliveryList,omitempty"`
	MallCycleGoodsConfig GoodsGetResponseMallCycleGoodsConfig `json:"mallCycleGoodsConfig,omitempty"`
}

type GoodsGetResponseDeliveryList struct {
	DeliveryId         int64 `json:"deliveryId,omitempty"`
	DeliveryNodeShipId int64 `json:"deliveryNodeShipId,omitempty"`
	DeliveryType       int64 `json:"deliveryType,omitempty"`
	TemplateId         int64 `json:"templateId,omitempty"`
}

type GoodsGetResponseMallCycleGoodsConfig struct {
	FulfillCycleConfig GoodsGetResponseFulfillCycleConfig `json:"fulfillCycleConfig,omitempty"`
	FulfillCycleInt    int64                              `json:"fulfillCycleInt,omitempty"`
	DelayFulfill       int64                              `json:"delayFulfill,omitempty"`
	DelayCurrFulfill   int64                              `json:"delayCurrFulfill,omitempty"`
	DelayDay           int64                              `json:"delayDay,omitempty"`
	ClosingTime        int64                              `json:"closingTime,omitempty"`
	ClosingTimeDetail  string                             `json:"closingTimeDetail,omitempty"`
	AlterAddress       int64                              `json:"alterAddress,omitempty"`
}

type GoodsGetResponseFulfillCycleConfig struct {
	CycleUnit          int64   `json:"cycleUnit,omitempty"`
	CycleValues        []int64 `json:"cycleValues,omitempty"`
	CycleInterval      int64   `json:"cycleInterval,omitempty"`
	CycleTimes         int64   `json:"cycleTimes,omitempty"`
	RemoveWeekend      int64   `json:"removeWeekend,omitempty"`
	FulfillTimeOptions []int64 `json:"fulfillTimeOptions,omitempty"`
}

type GoodsGetResponseCategoryList struct {
	CategoryId       int64  `json:"categoryId,omitempty"`
	CategoryName     string `json:"categoryName,omitempty"`
	ParentCategoryId int64  `json:"parentCategoryId,omitempty"`
	CategoryLevel    int64  `json:"categoryLevel,omitempty"`
	CategoryType     int64  `json:"categoryType,omitempty"`
}

type GoodsGetResponseBrandInfo struct {
	BrandId int64  `json:"brandId,omitempty"`
	Log     string `json:"log,omitempty"`
	Name    string `json:"name,omitempty"`
	Type    int64  `json:"type,omitempty"`
}

type GoodsGetResponseInnerPropertyList struct {
	PropId        int64  `json:"propId,omitempty"`
	PropName      string `json:"propName,omitempty"`
	PropValueName string `json:"propValueName,omitempty"`
	PropValueId   int64  `json:"propValueId,omitempty"`
}

type GoodsGetResponseGoodsPropertyList struct {
	PropId        int64  `json:"propId,omitempty"`
	PropName      string `json:"propName,omitempty"`
	PropValueName string `json:"propValueName,omitempty"`
	PropValueId   int64  `json:"propValueId,omitempty"`
}

type GoodsGetResponseGoodsRightsList struct {
	PropId        int64  `json:"propId,omitempty"`
	PropName      string `json:"propName,omitempty"`
	PropValueName string `json:"propValueName,omitempty"`
	PropValueId   int64  `json:"propValueId,omitempty"`
}

type GoodsGetResponseCategoryPropertyList struct {
	CategoryId    int64  `json:"categoryId,omitempty"`
	PropId        int64  `json:"propId,omitempty"`
	PropName      string `json:"propName,omitempty"`
	PropValueName string `json:"propValueName,omitempty"`
	PropValueId   int64  `json:"propValueId,omitempty"`
}

type GoodsGetResponseTagInfo struct {
	TagId int64  `json:"tagId,omitempty"`
	Name  string `json:"name,omitempty"`
}

type GoodsGetResponseSpecInfoList struct {
	SkuSpecValueList    []GoodsGetResponseSkuSpecValueList `json:"skuSpecValueList,omitempty"`
	SpecId              int64                              `json:"specId,omitempty"`
	SpecName            string                             `json:"specName,omitempty"`
	SpecImgEnable       bool                               `json:"specImgEnable,omitempty"`
	IsOpenSizeRecommend bool                               `json:"isOpenSizeRecommend,omitempty"`
}

type GoodsGetResponseSkuSpecValueList struct {
	SpecValueName string `json:"specValueName,omitempty"`
	ImageUrl      string `json:"imageUrl,omitempty"`
	SpecValueId   int64  `json:"specValueId,omitempty"`
	Sort          int64  `json:"sort,omitempty"`
}

type GoodsGetResponsePointDeductRule struct {
	OpenPointDeduct int64 `json:"openPointDeduct,omitempty"`
	UseCondition    int64 `json:"useCondition,omitempty"`
	UseDiscount     int64 `json:"useDiscount,omitempty"`
	DeductAmount    int64 `json:"deductAmount,omitempty"`
	DeductDiscount  int64 `json:"deductDiscount,omitempty"`
	DeductionType   int64 `json:"deductionType,omitempty"`
}

type GoodsGetResponseGoodsClassifyList struct {
	ClassifyId    int64  `json:"classifyId,omitempty"`
	Name          string `json:"name,omitempty"`
	ParentId      int64  `json:"parentId,omitempty"`
	ClassifyLevel int64  `json:"classifyLevel,omitempty"`
}

type GoodsGetResponseSkuList struct {
	SkuSpecValueList []GoodsGetResponseSkuSpecValueList2 `json:"skuSpecValueList,omitempty"`
	Volume           int64                               `json:"volume,omitempty"`
	MarketPrice      int64                               `json:"marketPrice,omitempty"`
	SkuStockNum      int64                               `json:"skuStockNum,omitempty"`
	SalePrice        int64                               `json:"salePrice,omitempty"`
	CostPrice        int64                               `json:"costPrice,omitempty"`
	OuterSkuCode     string                              `json:"outerSkuCode,omitempty"`
	Weight           int64                               `json:"weight,omitempty"`
	SkuBarCode       string                              `json:"skuBarCode,omitempty"`
	IsDisabled       bool                                `json:"isDisabled,omitempty"`
	SkuId            int64                               `json:"skuId,omitempty"`
}

type GoodsGetResponseSkuSpecValueList2 struct {
	SpecId        int64  `json:"specId,omitempty"`
	SpecName      string `json:"specName,omitempty"`
	SpecValueId   int64  `json:"specValueId,omitempty"`
	SpecValueName string `json:"specValueName,omitempty"`
}

type GoodsGetResponseSizeTableInfo struct {
	ParamList           []GoodsGetResponseParamList           `json:"paramList,omitempty"`
	SizeTableDetailList []GoodsGetResponseSizeTableDetailList `json:"sizeTableDetailList,omitempty"`
	SpecId              int64                                 `json:"specId,omitempty"`
	SpecName            string                                `json:"specName,omitempty"`
}

type GoodsGetResponseParamList struct {
	IsDefaultSelected bool   `json:"isDefaultSelected,omitempty"`
	ParamName         string `json:"paramName,omitempty"`
	IsSystemParam     bool   `json:"isSystemParam,omitempty"`
}

type GoodsGetResponseSizeTableDetailList struct {
	ParamValueList []GoodsGetResponseParamValueList `json:"paramValueList,omitempty"`
	SpecValue      string                           `json:"specValue,omitempty"`
	SpecValueId    int64                            `json:"specValueId,omitempty"`
}

type GoodsGetResponseParamValueList struct {
	ParamName  string `json:"paramName,omitempty"`
	ParamValue int64  `json:"paramValue,omitempty"`
}
