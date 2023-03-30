package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromocodeTemplateDetailGet
// @id 1598
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1598?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询码模板详情)
func (client *Client) PromocodeTemplateDetailGet(request *PromocodeTemplateDetailGetRequest) (response *PromocodeTemplateDetailGetResponse, err error) {
	rpcResponse := CreatePromocodeTemplateDetailGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromocodeTemplateDetailGetRequest struct {
	*api.BaseRequest

	CodeTemplateId int64 `json:"codeTemplateId,omitempty"`
}

type PromocodeTemplateDetailGetResponse struct {
	BaseInfo   PromocodeTemplateDetailGetResponseBaseInfo   `json:"baseInfo,omitempty"`
	ReduceRule PromocodeTemplateDetailGetResponseReduceRule `json:"reduceRule,omitempty"`
	UseRule    PromocodeTemplateDetailGetResponseUseRule    `json:"useRule,omitempty"`
	CreateInfo PromocodeTemplateDetailGetResponseCreateInfo `json:"createInfo,omitempty"`
}

func CreatePromocodeTemplateDetailGetRequest() (request *PromocodeTemplateDetailGetRequest) {
	request = &PromocodeTemplateDetailGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "promocode/template/detail/get", "apigw")
	request.Method = api.POST
	return
}

func CreatePromocodeTemplateDetailGetResponse() (response *api.BaseResponse[PromocodeTemplateDetailGetResponse]) {
	response = api.CreateResponse[PromocodeTemplateDetailGetResponse](&PromocodeTemplateDetailGetResponse{})
	return
}

type PromocodeTemplateDetailGetResponseBaseInfo struct {
	CodeSetting    PromocodeTemplateDetailGetResponseCodeSetting `json:"codeSetting,omitempty"`
	CodeTemplateId int64                                         `json:"codeTemplateId,omitempty"`
	Name           string                                        `json:"name,omitempty"`
	Desc           string                                        `json:"desc,omitempty"`
	CouponType     int64                                         `json:"couponType,omitempty"`
	CodeType       int64                                         `json:"codeType,omitempty"`
	GenerateType   int64                                         `json:"generateType,omitempty"`
	Stock          int64                                         `json:"stock,omitempty"`
	ImportFileUrl  string                                        `json:"importFileUrl,omitempty"`
	ValidStartTime string                                        `json:"validStartTime,omitempty"`
	ValidEndTime   string                                        `json:"validEndTime,omitempty"`
	UserLimitNum   int64                                         `json:"userLimitNum,omitempty"`
	LogoUrl        string                                        `json:"logoUrl,omitempty"`
	Code           string                                        `json:"code,omitempty"`
	BelongVid      int64                                         `json:"belongVid,omitempty"`
	BelongVidName  string                                        `json:"belongVidName,omitempty"`
}

type PromocodeTemplateDetailGetResponseCodeSetting struct {
	FixLeft    string `json:"fixLeft,omitempty"`
	FixRight   string `json:"fixRight,omitempty"`
	CodeLength int64  `json:"codeLength,omitempty"`
	IsLetter   int64  `json:"isLetter,omitempty"`
}

type PromocodeTemplateDetailGetResponseReduceRule struct {
	CouponAmount       int64 `json:"couponAmount,omitempty"`
	CouponCostType     int64 `json:"couponCostType,omitempty"`
	CouponCost         int64 `json:"couponCost,omitempty"`
	CostOfOrganization int64 `json:"costOfOrganization,omitempty"`
	CostOfStore        int64 `json:"costOfStore,omitempty"`
	PricingType        int64 `json:"pricingType,omitempty"`
	DiscountLimitType  int64 `json:"discountLimitType,omitempty"`
	DiscountLimitValue int64 `json:"discountLimitValue,omitempty"`
}

type PromocodeTemplateDetailGetResponseUseRule struct {
	AcceptTypeDTO       PromocodeTemplateDetailGetResponseAcceptTypeDTO  `json:"acceptTypeDTO,omitempty"`
	CycleUseRule        []PromocodeTemplateDetailGetResponseCycleUseRule `json:"cycleUseRule,omitempty"`
	CanUseDiscount      PromocodeTemplateDetailGetResponseCanUseDiscount `json:"canUseDiscount,omitempty"`
	AllSceneDTO         PromocodeTemplateDetailGetResponseAllSceneDTO    `json:"allSceneDTO,omitempty"`
	UserTakeLimit       int64                                            `json:"userTakeLimit,omitempty"`
	HasUseThreshold     bool                                             `json:"hasUseThreshold,omitempty"`
	CanCashTicket       bool                                             `json:"canCashTicket,omitempty"`
	GoodsCashTicket     int64                                            `json:"goodsCashTicket,omitempty"`
	CanGoodsNumber      bool                                             `json:"canGoodsNumber,omitempty"`
	MinGoodsNumber      int64                                            `json:"minGoodsNumber,omitempty"`
	MaxGoodsNumber      int64                                            `json:"maxGoodsNumber,omitempty"`
	IsAllTimeUse        bool                                             `json:"isAllTimeUse,omitempty"`
	IsAllUseScene       bool                                             `json:"isAllUseScene,omitempty"`
	OrderDeductList     []int64                                          `json:"orderDeductList,omitempty"`
	ImportGoodsFlag     bool                                             `json:"importGoodsFlag,omitempty"`
	ImportGoodsFileName string                                           `json:"importGoodsFileName,omitempty"`
}

type PromocodeTemplateDetailGetResponseAcceptTypeDTO struct {
	AcceptStoreType   int64   `json:"acceptStoreType,omitempty"`
	AcceptStoreIds    []int64 `json:"acceptStoreIds,omitempty"`
	AcceptGoodsType   int64   `json:"acceptGoodsType,omitempty"`
	AcceptGoodsIds    []int64 `json:"acceptGoodsIds,omitempty"`
	ExistExcludeGoods bool    `json:"existExcludeGoods,omitempty"`
	ExcludeGoodsIds   []int64 `json:"excludeGoodsIds,omitempty"`
}

type PromocodeTemplateDetailGetResponseCycleUseRule struct {
	TimeSegment []PromocodeTemplateDetailGetResponseTimeSegment `json:"timeSegment,omitempty"`
	WeekDay     []int64                                         `json:"weekDay,omitempty"`
}

type PromocodeTemplateDetailGetResponseTimeSegment struct {
	BeginHour   int64 `json:"beginHour,omitempty"`
	EndHour     int64 `json:"endHour,omitempty"`
	BeginMinute int64 `json:"beginMinute,omitempty"`
	EndMinute   int64 `json:"endMinute,omitempty"`
}

type PromocodeTemplateDetailGetResponseCanUseDiscount struct {
	CanUseWithOtherDiscount bool    `json:"canUseWithOtherDiscount,omitempty"`
	ShoppingMallDiscount    []int64 `json:"shoppingMallDiscount,omitempty"`
	ForHereDiscount         []int64 `json:"forHereDiscount,omitempty"`
	TakeOutDiscount         []int64 `json:"takeOutDiscount,omitempty"`
	WineTourDiscount        []int64 `json:"wineTourDiscount,omitempty"`
}

type PromocodeTemplateDetailGetResponseAllSceneDTO struct {
	ShoppingMallSceneList []int64 `json:"shoppingMallSceneList,omitempty"`
}

type PromocodeTemplateDetailGetResponseCreateInfo struct {
	CreateTime string `json:"createTime,omitempty"`
	OperatorId int64  `json:"operatorId,omitempty"`
	Source     string `json:"source,omitempty"`
}
