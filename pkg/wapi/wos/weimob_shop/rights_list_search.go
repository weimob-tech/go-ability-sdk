package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsListSearch
// @id 1848
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1848?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询售后单列表)
func (client *Client) RightsListSearch(request *RightsListSearchRequest) (response *RightsListSearchResponse, err error) {
	rpcResponse := CreateRightsListSearchResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsListSearchRequest struct {
	*api.BaseRequest

	QueryParameter RightsListSearchRequestQueryParameter `json:"queryParameter,omitempty"`
	PageSize       int64                                 `json:"pageSize,omitempty"`
	PageNum        int64                                 `json:"pageNum,omitempty"`
}

type RightsListSearchResponse struct {
	PageList   []RightsListSearchResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                              `json:"pageNum,omitempty"`
	PageSize   int64                              `json:"pageSize,omitempty"`
	TotalCount int64                              `json:"totalCount,omitempty"`
}

func CreateRightsListSearchRequest() (request *RightsListSearchRequest) {
	request = &RightsListSearchRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "rights/list/search", "apigw")
	request.Method = api.POST
	return
}

func CreateRightsListSearchResponse() (response *api.BaseResponse[RightsListSearchResponse]) {
	response = api.CreateResponse[RightsListSearchResponse](&RightsListSearchResponse{})
	return
}

type RightsListSearchRequestQueryParameter struct {
	QueryTime   RightsListSearchRequestQueryTime `json:"queryTime,omitempty"`
	RightsTypes []int64                          `json:"rightsTypes,omitempty"`
	ProcessVids []int64                          `json:"processVids,omitempty"`
	OrderNo     int64                            `json:"orderNo	,omitempty"`
	Statuses    []int64                          `json:"statuses,omitempty"`
	UserWid     int64                            `json:"userWid,omitempty"`
	SearchType  int64                            `json:"searchType,omitempty"`
	Keyword     string                           `json:"keyword,omitempty"`
}

type RightsListSearchRequestQueryTime struct {
	StartTime int64 `json:"startTime,omitempty"`
	EndTime   int64 `json:"endTime,omitempty"`
	Type      int64 `json:"type,omitempty"`
}

type RightsListSearchResponsePageList struct {
	RightsOrder      RightsListSearchResponseRightsOrder        `json:"rightsOrder,omitempty"`
	MerchantInfo     RightsListSearchResponseMerchantInfo       `json:"merchantInfo,omitempty"`
	BuyerInfo        RightsListSearchResponseBuyerInfo          `json:"buyerInfo,omitempty"`
	RightsItems      []RightsListSearchResponseRightsItems      `json:"rightsItems,omitempty"`
	RightsReason     RightsListSearchResponseRightsReason       `json:"rightsReason,omitempty"`
	RefundAccount    RightsListSearchResponseRefundAccount      `json:"refundAccount,omitempty"`
	FlagInfo         RightsListSearchResponseFlagInfo           `json:"flagInfo,omitempty"`
	RefundDetail     RightsListSearchResponseRefundDetail2      `json:"refundDetail,omitempty"`
	OriginOrder      RightsListSearchResponseOriginOrder        `json:"originOrder,omitempty"`
	ReturnOrder      RightsListSearchResponseReturnOrder        `json:"returnOrder,omitempty"`
	ExchangeOrder    RightsListSearchResponseExchangeOrder      `json:"exchangeOrder,omitempty"`
	OutRightsInfo    RightsListSearchResponseOutRightsInfo      `json:"outRightsInfo,omitempty"`
	RightsAssets     []RightsListSearchResponseRightsAssets2    `json:"rightsAssets,omitempty"`
	RightsStatusLogs []RightsListSearchResponseRightsStatusLogs `json:"rightsStatusLogs,omitempty"`
	RefuseInfo       RightsListSearchResponseRefuseInfo         `json:"refuseInfo,omitempty"`
}

type RightsListSearchResponseRightsOrder struct {
	RightsId        int64  `json:"rightsId,omitempty"`
	CreateTime      int64  `json:"createTime,omitempty"`
	UpdateTime      int64  `json:"updateTime,omitempty"`
	RightsStatus    int64  `json:"rightsStatus,omitempty"`
	RightsType      int64  `json:"rightsType,omitempty"`
	RightsSource    int64  `json:"rightsSource,omitempty"`
	RightsCauseType int64  `json:"rightsCauseType,omitempty"`
	RefundType      int64  `json:"refundType,omitempty"`
	Currency        string `json:"currency,omitempty"`
}

type RightsListSearchResponseMerchantInfo struct {
	MerchantId     int64  `json:"merchantId,omitempty"`
	BosId          int64  `json:"bosId,omitempty"`
	BosName        string `json:"bosName,omitempty"`
	Vid            int64  `json:"vid,omitempty"`
	VidName        string `json:"vidName,omitempty"`
	ProcessVid     int64  `json:"processVid,omitempty"`
	ProcessVidName string `json:"processVidName,omitempty"`
}

type RightsListSearchResponseBuyerInfo struct {
	MemberBenefits []RightsListSearchResponseMemberBenefits `json:"memberBenefits,omitempty"`
	Wid            int64                                    `json:"wid,omitempty"`
	UserNickName   string                                   `json:"userNickName,omitempty"`
}

type RightsListSearchResponseMemberBenefits struct {
	BenefitType int64 `json:"benefitType,omitempty"`
}

type RightsListSearchResponseRightsItems struct {
	GoodsInfo        RightsListSearchResponseGoodsInfo      `json:"goodsInfo,omitempty"`
	RefundDetail     RightsListSearchResponseRefundDetail   `json:"refundDetail,omitempty"`
	RightsAssets     []RightsListSearchResponseRightsAssets `json:"rightsAssets,omitempty"`
	RightsItemId     int64                                  `json:"rightsItemId,omitempty"`
	OrderItemId      int64                                  `json:"orderItemId,omitempty"`
	ApplyNum         string                                 `json:"applyNum,omitempty"`
	ReturnNum        string                                 `json:"returnNum,omitempty"`
	GoodsReceiveType int64                                  `json:"goodsReceiveType,omitempty"`
}

type RightsListSearchResponseGoodsInfo struct {
	LabelInfos        []RightsListSearchResponseLabelInfos     `json:"labelInfos,omitempty"`
	BizInfos          []RightsListSearchResponseBizInfos       `json:"bizInfos,omitempty"`
	GoodsAbilities    []RightsListSearchResponseGoodsAbilities `json:"goodsAbilities,omitempty"`
	GoodsType         int64                                    `json:"goodsType,omitempty"`
	SubGoodsType      int64                                    `json:"subGoodsType,omitempty"`
	GoodsId           int64                                    `json:"goodsId,omitempty"`
	GoodsCode         string                                   `json:"goodsCode,omitempty"`
	GoodsTitle        string                                   `json:"goodsTitle,omitempty"`
	ImageUrl          string                                   `json:"imageUrl,omitempty"`
	Price             string                                   `json:"price,omitempty"`
	SkuNum            string                                   `json:"skuNum,omitempty"`
	SkuId             int64                                    `json:"skuId,omitempty"`
	SkuCode           string                                   `json:"skuCode,omitempty"`
	SkuBarCode        string                                   `json:"skuBarCode,omitempty"`
	SkuAttrInfo       string                                   `json:"skuAttrInfo,omitempty"`
	RightsServiceType int64                                    `json:"rightsServiceType,omitempty"`
}

type RightsListSearchResponseLabelInfos struct {
	LabelType string `json:"labelType,omitempty"`
}

type RightsListSearchResponseBizInfos struct {
	BizId      int64  `json:"bizId,omitempty"`
	BizOrderId string `json:"bizOrderId,omitempty"`
	BizType    int64  `json:"bizType,omitempty"`
	SubBizType int64  `json:"subBizType,omitempty"`
}

type RightsListSearchResponseGoodsAbilities struct {
	AbilityCode string `json:"abilityCode,omitempty"`
}

type RightsListSearchResponseRefundDetail struct {
	ApplyAmountInfos  []RightsListSearchResponseApplyAmountInfos  `json:"applyAmountInfos,omitempty"`
	RefundAmountInfos []RightsListSearchResponseRefundAmountInfos `json:"refundAmountInfos,omitempty"`
	ApplyAmount       string                                      `json:"applyAmount,omitempty"`
	RefundAmount      string                                      `json:"refundAmount,omitempty"`
}

type RightsListSearchResponseApplyAmountInfos struct {
	Type   int64  `json:"type,omitempty"`
	Amount string `json:"amount,omitempty"`
}

type RightsListSearchResponseRefundAmountInfos struct {
	Type   int64  `json:"type,omitempty"`
	Amount string `json:"amount,omitempty"`
}

type RightsListSearchResponseRightsAssets struct {
	AssetsType   int64  `json:"assetsType,omitempty"`
	AssetsTarget int64  `json:"assetsTarget,omitempty"`
	AssetsAmount string `json:"assetsAmount,omitempty"`
	AssetsNum    string `json:"assetsNum,omitempty"`
}

type RightsListSearchResponseRightsReason struct {
	RightsReason    string  `json:"rightsReason,omitempty"`
	RightsRemark    string  `json:"rightsRemark,omitempty"`
	ReasonImageUrls []int64 `json:"reasonImageUrls,omitempty"`
}

type RightsListSearchResponseRefundAccount struct {
	AccountNum        string `json:"accountNum,omitempty"`
	OpeningBank       string `json:"openingBank,omitempty"`
	RefundPayMethodId int64  `json:"refundPayMethodId,omitempty"`
	UserName          string `json:"userName,omitempty"`
}

type RightsListSearchResponseFlagInfo struct {
	FlagRank    int64  `json:"flagRank,omitempty"`
	FlagContent string `json:"flagContent,omitempty"`
}

type RightsListSearchResponseRefundDetail2 struct {
	ApplyAmountInfos  []RightsListSearchResponseApplyAmountInfos2  `json:"applyAmountInfos	,omitempty"`
	RefundAmountInfos []RightsListSearchResponseRefundAmountInfos2 `json:"refundAmountInfos,omitempty"`
	ApplyAmount       string                                       `json:"applyAmount,omitempty"`
	RefundAmount      string                                       `json:"refundAmount,omitempty"`
}

type RightsListSearchResponseApplyAmountInfos2 struct {
	Type   int64  `json:"type,omitempty"`
	Amount string `json:"amount,omitempty"`
}

type RightsListSearchResponseRefundAmountInfos2 struct {
	Type   int64  `json:"type,omitempty"`
	Amount string `json:"amount,omitempty"`
}

type RightsListSearchResponseOriginOrder struct {
	OrderNo          int64  `json:"orderNo,omitempty"`
	OutOrderNo       string `json:"outOrderNo,omitempty"`
	ApplyOrderStatus int64  `json:"applyOrderStatus,omitempty"`
	OrderType        int64  `json:"orderType,omitempty"`
	OrderSource      int64  `json:"orderSource,omitempty"`
	ChannelType      int64  `json:"channelType,omitempty"`
	BizSourceType    int64  `json:"bizSourceType,omitempty"`
	PayType          int64  `json:"payType,omitempty"`
}

type RightsListSearchResponseReturnOrder struct {
	DeliveryStatus      int64  `json:"deliveryStatus,omitempty"`
	DeliveryType        int64  `json:"deliveryType,omitempty"`
	DeliveryCompanyCode string `json:"deliveryCompanyCode,omitempty"`
	DeliveryCompany     string `json:"deliveryCompany,omitempty"`
	DeliveryNo          string `json:"deliveryNo,omitempty"`
	ReceiverName        string `json:"receiverName,omitempty"`
	ReceiverPhone       string `json:"receiverPhone,omitempty"`
	ReceiverAddress     string `json:"receiverAddress,omitempty"`
	DeliveryMethod      int64  `json:"deliveryMethod,omitempty"`
}

type RightsListSearchResponseExchangeOrder struct {
	OrderItems []RightsListSearchResponseOrderItems `json:"orderItems,omitempty"`
	OrderNo    int64                                `json:"orderNo,omitempty"`
}

type RightsListSearchResponseOrderItems struct {
	GoodsInfo   RightsListSearchResponseGoodsInfo2 `json:"goodsInfo,omitempty"`
	OrderItemId int64                              `json:"orderItemId,omitempty"`
}

type RightsListSearchResponseGoodsInfo2 struct {
	BizInfos          []RightsListSearchResponseBizInfos2       `json:"bizInfos,omitempty"`
	LabelInfos        []RightsListSearchResponseLabelInfos2     `json:"labelInfos,omitempty"`
	GoodsAbilities    []RightsListSearchResponseGoodsAbilities2 `json:"goodsAbilities,omitempty"`
	GoodsCode         string                                    `json:"goodsCode,omitempty"`
	GoodsId           int64                                     `json:"goodsId,omitempty"`
	GoodsTitle        string                                    `json:"goodsTitle,omitempty"`
	GoodsType         int64                                     `json:"goodsType,omitempty"`
	ImageUrl          string                                    `json:"imageUrl,omitempty"`
	Price             string                                    `json:"price,omitempty"`
	RightsServiceType int64                                     `json:"rightsServiceType,omitempty"`
	SkuAttrInfo       string                                    `json:"skuAttrInfo,omitempty"`
	SkuBarCode        string                                    `json:"skuBarCode,omitempty"`
	SkuCode           string                                    `json:"skuCode,omitempty"`
	SkuId             int64                                     `json:"skuId,omitempty"`
	SkuNum            string                                    `json:"skuNum,omitempty"`
	SubGoodsType      int64                                     `json:"subGoodsType,omitempty"`
}

type RightsListSearchResponseBizInfos2 struct {
	BizId      int64  `json:"bizId,omitempty"`
	BizOrderId string `json:"bizOrderId,omitempty"`
	BizType    int64  `json:"bizType,omitempty"`
	SubBizType int64  `json:"subBizType,omitempty"`
}

type RightsListSearchResponseLabelInfos2 struct {
	LabelType string `json:"labelType,omitempty"`
}

type RightsListSearchResponseGoodsAbilities2 struct {
	AbilityCode string `json:"abilityCode,omitempty"`
	AbilityType int64  `json:"abilityType,omitempty"`
	BizId       int64  `json:"bizId,omitempty"`
}

type RightsListSearchResponseOutRightsInfo struct {
	OutRightsNo string `json:"outRightsNo,omitempty"`
}

type RightsListSearchResponseRightsAssets2 struct {
	AssetsType   int64  `json:"assetsType,omitempty"`
	AssetsTarget int64  `json:"assetsTarget,omitempty"`
	AssetsAmount string `json:"assetsAmount,omitempty"`
	AssetsNum    string `json:"assetsNum,omitempty"`
}

type RightsListSearchResponseRightsStatusLogs struct {
	DateTime int64  `json:"dateTime,omitempty"`
	Type     string `json:"type,omitempty"`
}

type RightsListSearchResponseRefuseInfo struct {
	RefusedReason string `json:"refusedReason,omitempty"`
}
