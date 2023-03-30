package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsDetailGetByOrderNos
// @id 1839
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1839?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=根据批量订单号查询售后单详情)
func (client *Client) RightsDetailGetByOrderNos(request *RightsDetailGetByOrderNosRequest) (response *RightsDetailGetByOrderNosResponse, err error) {
	rpcResponse := CreateRightsDetailGetByOrderNosResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsDetailGetByOrderNosRequest struct {
	*api.BaseRequest

	OrderNos []int64 `json:"orderNos,omitempty"`
}

type RightsDetailGetByOrderNosResponse struct {
	RightsInfos []RightsDetailGetByOrderNosResponseRightsInfos `json:"rightsInfos,omitempty"`
}

func CreateRightsDetailGetByOrderNosRequest() (request *RightsDetailGetByOrderNosRequest) {
	request = &RightsDetailGetByOrderNosRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "rights/detail/getByOrderNos", "apigw")
	request.Method = api.POST
	return
}

func CreateRightsDetailGetByOrderNosResponse() (response *api.BaseResponse[RightsDetailGetByOrderNosResponse]) {
	response = api.CreateResponse[RightsDetailGetByOrderNosResponse](&RightsDetailGetByOrderNosResponse{})
	return
}

type RightsDetailGetByOrderNosResponseRightsInfos struct {
	BuyerInfo        RightsDetailGetByOrderNosResponseBuyerInfo          `json:"buyerInfo,omitempty"`
	ExchangeOrder    RightsDetailGetByOrderNosResponseExchangeOrder      `json:"exchangeOrder,omitempty"`
	FlagInfo         RightsDetailGetByOrderNosResponseFlagInfo           `json:"flagInfo,omitempty"`
	MerchantInfo     RightsDetailGetByOrderNosResponseMerchantInfo       `json:"merchantInfo,omitempty"`
	OriginOrder      RightsDetailGetByOrderNosResponseOriginOrder        `json:"originOrder,omitempty"`
	OutRightsInfo    RightsDetailGetByOrderNosResponseOutRightsInfo      `json:"outRightsInfo,omitempty"`
	RefundAccount    RightsDetailGetByOrderNosResponseRefundAccount      `json:"refundAccount,omitempty"`
	ReturnOrder      RightsDetailGetByOrderNosResponseReturnOrder        `json:"returnOrder,omitempty"`
	RightsAssets     []RightsDetailGetByOrderNosResponseRightsAssets     `json:"rightsAssets,omitempty"`
	RightsItems      []RightsDetailGetByOrderNosResponseRightsItems      `json:"rightsItems,omitempty"`
	RightsOrder      RightsDetailGetByOrderNosResponseRightsOrder        `json:"rightsOrder,omitempty"`
	RightsReason     RightsDetailGetByOrderNosResponseRightsReason       `json:"rightsReason,omitempty"`
	RefundDetail     RightsDetailGetByOrderNosResponseRefundDetail2      `json:"refundDetail,omitempty"`
	RightsStatusLogs []RightsDetailGetByOrderNosResponseRightsStatusLogs `json:"rightsStatusLogs,omitempty"`
	RefuseInfo       RightsDetailGetByOrderNosResponseRefuseInfo         `json:"refuseInfo,omitempty"`
}

type RightsDetailGetByOrderNosResponseBuyerInfo struct {
	MemberBenefits []RightsDetailGetByOrderNosResponseMemberBenefits `json:"memberBenefits,omitempty"`
	UserNickName   string                                            `json:"userNickName,omitempty"`
	Wid            int64                                             `json:"wid,omitempty"`
}

type RightsDetailGetByOrderNosResponseMemberBenefits struct {
	BenefitType int64 `json:"benefitType,omitempty"`
}

type RightsDetailGetByOrderNosResponseExchangeOrder struct {
	OrderItems []RightsDetailGetByOrderNosResponseOrderItems `json:"orderItems,omitempty"`
	OrderNo    int64                                         `json:"orderNo,omitempty"`
}

type RightsDetailGetByOrderNosResponseOrderItems struct {
	GoodsInfo   RightsDetailGetByOrderNosResponseGoodsInfo `json:"goodsInfo,omitempty"`
	OrderItemId string                                     `json:"orderItemId,omitempty"`
}

type RightsDetailGetByOrderNosResponseGoodsInfo struct {
	BizInfos          []RightsDetailGetByOrderNosResponseBizInfos       `json:"bizInfos,omitempty"`
	LabelInfos        []RightsDetailGetByOrderNosResponseLabelInfos     `json:"labelInfos,omitempty"`
	GoodsAbilities    []RightsDetailGetByOrderNosResponseGoodsAbilities `json:"goodsAbilities,omitempty"`
	GoodsCode         string                                            `json:"goodsCode,omitempty"`
	GoodsId           string                                            `json:"goodsId,omitempty"`
	GoodsTitle        string                                            `json:"goodsTitle,omitempty"`
	GoodsType         string                                            `json:"goodsType,omitempty"`
	ImageUrl          string                                            `json:"imageUrl,omitempty"`
	Price             string                                            `json:"price,omitempty"`
	RightsServiceType int64                                             `json:"rightsServiceType,omitempty"`
	SkuAttrInfo       string                                            `json:"skuAttrInfo,omitempty"`
	SkuBarCode        string                                            `json:"skuBarCode,omitempty"`
	SkuCode           string                                            `json:"skuCode,omitempty"`
	SkuId             int64                                             `json:"skuId,omitempty"`
	SkuNum            string                                            `json:"skuNum,omitempty"`
	SubGoodsType      int64                                             `json:"subGoodsType,omitempty"`
}

type RightsDetailGetByOrderNosResponseBizInfos struct {
	BizId      int64  `json:"bizId,omitempty"`
	BizOrderId string `json:"bizOrderId,omitempty"`
	BizType    int64  `json:"bizType,omitempty"`
	SubBizType int64  `json:"subBizType,omitempty"`
}

type RightsDetailGetByOrderNosResponseLabelInfos struct {
	LabelType string `json:"labelType,omitempty"`
}

type RightsDetailGetByOrderNosResponseGoodsAbilities struct {
	AbilityCode string `json:"abilityCode,omitempty"`
	BizId       int64  `json:"bizId,omitempty"`
}

type RightsDetailGetByOrderNosResponseFlagInfo struct {
	FlagContent string `json:"flagContent,omitempty"`
	FlagRank    int64  `json:"flagRank,omitempty"`
}

type RightsDetailGetByOrderNosResponseMerchantInfo struct {
	MerchantId     int64  `json:"merchantId,omitempty"`
	BosId          int64  `json:"bosId,omitempty"`
	BosName        string `json:"bosName,omitempty"`
	Vid            int64  `json:"vid,omitempty"`
	VidName        string `json:"vidName,omitempty"`
	ProcessVid     int64  `json:"processVid,omitempty"`
	ProcessVidName string `json:"processVidName,omitempty"`
}

type RightsDetailGetByOrderNosResponseOriginOrder struct {
	OrderNo          int64  `json:"orderNo,omitempty"`
	OutOrderNo       string `json:"outOrderNo,omitempty"`
	ApplyOrderStatus int64  `json:"applyOrderStatus,omitempty"`
	OrderType        int64  `json:"orderType,omitempty"`
	OrderSource      int64  `json:"orderSource,omitempty"`
	ChannelType      int64  `json:"channelType,omitempty"`
	BizSourceType    int64  `json:"bizSourceType,omitempty"`
	PayType          int64  `json:"payType,omitempty"`
}

type RightsDetailGetByOrderNosResponseOutRightsInfo struct {
	OutRightsNo string `json:"outRightsNo,omitempty"`
}

type RightsDetailGetByOrderNosResponseRefundAccount struct {
	AccountNum     string `json:"accountNum,omitempty"`
	OpeningBank    string `json:"openingBank,omitempty"`
	RefundMethodId int64  `json:"refundMethodId,omitempty"`
	UserName       string `json:"userName,omitempty"`
}

type RightsDetailGetByOrderNosResponseReturnOrder struct {
	DeliveryCompany     string `json:"deliveryCompany,omitempty"`
	DeliveryCompanyCode string `json:"deliveryCompanyCode,omitempty"`
	DeliveryNo          string `json:"deliveryNo,omitempty"`
	DeliveryStatus      int64  `json:"deliveryStatus,omitempty"`
	DeliveryType        int64  `json:"deliveryType,omitempty"`
	ReceiverAddress     string `json:"receiverAddress,omitempty"`
	ReceiverName        string `json:"receiverName,omitempty"`
	ReceiverPhone       string `json:"receiverPhone,omitempty"`
	DeliveryMethod      int64  `json:"deliveryMethod,omitempty"`
}

type RightsDetailGetByOrderNosResponseRightsAssets struct {
	AssetsAmount string `json:"assetsAmount,omitempty"`
	AssetsNum    string `json:"assetsNum,omitempty"`
	AssetsTarget int64  `json:"assetsTarget,omitempty"`
	AssetsType   int64  `json:"assetsType,omitempty"`
}

type RightsDetailGetByOrderNosResponseRightsItems struct {
	GoodsInfo        RightsDetailGetByOrderNosResponseGoodsInfo2      `json:"goodsInfo,omitempty"`
	RightsAssets     []RightsDetailGetByOrderNosResponseRightsAssets2 `json:"rightsAssets,omitempty"`
	RefundDetail     RightsDetailGetByOrderNosResponseRefundDetail    `json:"refundDetail,omitempty"`
	ApplyNum         string                                           `json:"applyNum,omitempty"`
	GoodsReceiveType int64                                            `json:"goodsReceiveType,omitempty"`
	OrderItemId      int64                                            `json:"orderItemId,omitempty"`
	ReturnNum        string                                           `json:"returnNum,omitempty"`
	RightsItemId     int64                                            `json:"rightsItemId,omitempty"`
}

type RightsDetailGetByOrderNosResponseGoodsInfo2 struct {
	BizInfos          []RightsDetailGetByOrderNosResponseBizInfos2       `json:"bizInfos,omitempty"`
	LabelInfos        []RightsDetailGetByOrderNosResponseLabelInfos2     `json:"labelInfos,omitempty"`
	GoodsAbilities    []RightsDetailGetByOrderNosResponseGoodsAbilities2 `json:"goodsAbilities,omitempty"`
	GoodsCode         string                                             `json:"goodsCode,omitempty"`
	GoodsId           int64                                              `json:"goodsId,omitempty"`
	GoodsTitle        string                                             `json:"goodsTitle,omitempty"`
	GoodsType         int64                                              `json:"goodsType,omitempty"`
	ImageUrl          string                                             `json:"imageUrl,omitempty"`
	Price             string                                             `json:"price,omitempty"`
	RightsServiceType int64                                              `json:"rightsServiceType,omitempty"`
	SkuAttrInfo       string                                             `json:"skuAttrInfo,omitempty"`
	SkuBarCode        string                                             `json:"skuBarCode,omitempty"`
	SkuCode           string                                             `json:"skuCode,omitempty"`
	SkuId             int64                                              `json:"skuId,omitempty"`
	SkuNum            string                                             `json:"skuNum,omitempty"`
	SubGoodsType      int64                                              `json:"subGoodsType,omitempty"`
}

type RightsDetailGetByOrderNosResponseBizInfos2 struct {
	BizId      int64  `json:"bizId,omitempty"`
	BizOrderId string `json:"bizOrderId,omitempty"`
	BizType    int64  `json:"bizType,omitempty"`
	SubBizType int64  `json:"subBizType,omitempty"`
}

type RightsDetailGetByOrderNosResponseLabelInfos2 struct {
	LabelType string `json:"labelType,omitempty"`
}

type RightsDetailGetByOrderNosResponseGoodsAbilities2 struct {
	AbilityCode string `json:"abilityCode,omitempty"`
}

type RightsDetailGetByOrderNosResponseRightsAssets2 struct {
	AssetsAmount string `json:"assetsAmount,omitempty"`
	AssetsNum    string `json:"assetsNum,omitempty"`
	AssetsTarget int64  `json:"assetsTarget,omitempty"`
	AssetsType   int64  `json:"assetsType,omitempty"`
}

type RightsDetailGetByOrderNosResponseRefundDetail struct {
	ApplyAmountInfos  []RightsDetailGetByOrderNosResponseApplyAmountInfos  `json:"applyAmountInfos,omitempty"`
	RefundAmountInfos []RightsDetailGetByOrderNosResponseRefundAmountInfos `json:"refundAmountInfos,omitempty"`
	ApplyAmount       string                                               `json:"applyAmount,omitempty"`
	RefundAmount      string                                               `json:"refundAmount,omitempty"`
}

type RightsDetailGetByOrderNosResponseApplyAmountInfos struct {
	Type   int64  `json:"type,omitempty"`
	Amount string `json:"amount,omitempty"`
}

type RightsDetailGetByOrderNosResponseRefundAmountInfos struct {
	Type   int64  `json:"type,omitempty"`
	Amount string `json:"amount,omitempty"`
}

type RightsDetailGetByOrderNosResponseRightsOrder struct {
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

type RightsDetailGetByOrderNosResponseRightsReason struct {
	ReasonImageUrls []int64 `json:"reasonImageUrls,omitempty"`
	RightsReason    string  `json:"rightsReason,omitempty"`
	RightsRemark    string  `json:"rightsRemark,omitempty"`
}

type RightsDetailGetByOrderNosResponseRefundDetail2 struct {
	ApplyAmountInfos  []RightsDetailGetByOrderNosResponseApplyAmountInfos2  `json:"applyAmountInfos,omitempty"`
	RefundAmountInfos []RightsDetailGetByOrderNosResponseRefundAmountInfos2 `json:"refundAmountInfos,omitempty"`
	ApplyAmount       string                                                `json:"applyAmount,omitempty"`
	RefundAmount      string                                                `json:"refundAmount,omitempty"`
}

type RightsDetailGetByOrderNosResponseApplyAmountInfos2 struct {
	Type   int64  `json:"type,omitempty"`
	Amount string `json:"amount,omitempty"`
}

type RightsDetailGetByOrderNosResponseRefundAmountInfos2 struct {
	Type   int64  `json:"type,omitempty"`
	Amount string `json:"amount,omitempty"`
}

type RightsDetailGetByOrderNosResponseRightsStatusLogs struct {
	DateTime int64  `json:"dateTime,omitempty"`
	Type     string `json:"type,omitempty"`
}

type RightsDetailGetByOrderNosResponseRefuseInfo struct {
	RefusedReason string `json:"refusedReason,omitempty"`
}
