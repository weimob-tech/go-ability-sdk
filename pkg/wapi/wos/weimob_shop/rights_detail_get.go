package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsDetailGet
// @id 1704
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1704?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询售后单详情)
func (client *Client) RightsDetailGet(request *RightsDetailGetRequest) (response *RightsDetailGetResponse, err error) {
	rpcResponse := CreateRightsDetailGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsDetailGetRequest struct {
	*api.BaseRequest

	RightsId int64 `json:"rightsId	,omitempty"`
}

type RightsDetailGetResponse struct {
	RightsInfo RightsDetailGetResponseRightsInfo `json:"rightsInfo,omitempty"`
}

func CreateRightsDetailGetRequest() (request *RightsDetailGetRequest) {
	request = &RightsDetailGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "rights/detail/get", "apigw")
	request.Method = api.POST
	return
}

func CreateRightsDetailGetResponse() (response *api.BaseResponse[RightsDetailGetResponse]) {
	response = api.CreateResponse[RightsDetailGetResponse](&RightsDetailGetResponse{})
	return
}

type RightsDetailGetResponseRightsInfo struct {
	BuyerInfo        RightsDetailGetResponseBuyerInfo          `json:"buyerInfo,omitempty"`
	ExchangeOrder    RightsDetailGetResponseExchangeOrder      `json:"exchangeOrder,omitempty"`
	FlagInfo         RightsDetailGetResponseFlagInfo           `json:"flagInfo,omitempty"`
	MerchantInfo     RightsDetailGetResponseMerchantInfo       `json:"merchantInfo,omitempty"`
	OriginOrder      RightsDetailGetResponseOriginOrder        `json:"originOrder,omitempty"`
	OutRightsInfo    RightsDetailGetResponseOutRightsInfo      `json:"outRightsInfo,omitempty"`
	RefundAccount    RightsDetailGetResponseRefundAccount      `json:"refundAccount,omitempty"`
	ReturnOrder      RightsDetailGetResponseReturnOrder        `json:"returnOrder,omitempty"`
	RightsAssets     []RightsDetailGetResponseRightsAssets     `json:"rightsAssets,omitempty"`
	RightsItems      []RightsDetailGetResponseRightsItems      `json:"rightsItems,omitempty"`
	RightsOrder      RightsDetailGetResponseRightsOrder        `json:"rightsOrder,omitempty"`
	RightsReason     RightsDetailGetResponseRightsReason       `json:"rightsReason,omitempty"`
	RefundDetail     RightsDetailGetResponseRefundDetail2      `json:"refundDetail,omitempty"`
	RightsStatusLogs []RightsDetailGetResponseRightsStatusLogs `json:"rightsStatusLogs,omitempty"`
	RefuseInfo       RightsDetailGetResponseRefuseInfo         `json:"refuseInfo,omitempty"`
}

type RightsDetailGetResponseBuyerInfo struct {
	MemberBenefits []RightsDetailGetResponseMemberBenefits `json:"memberBenefits,omitempty"`
	UserNickName   string                                  `json:"userNickName,omitempty"`
	Wid            int64                                   `json:"wid,omitempty"`
}

type RightsDetailGetResponseMemberBenefits struct {
	BenefitType int64 `json:"benefitType,omitempty"`
}

type RightsDetailGetResponseExchangeOrder struct {
	OrderItems []RightsDetailGetResponseOrderItems `json:"orderItems,omitempty"`
	OrderNo    int64                               `json:"orderNo,omitempty"`
}

type RightsDetailGetResponseOrderItems struct {
	GoodsInfo   RightsDetailGetResponseGoodsInfo `json:"goodsInfo,omitempty"`
	OrderItemId int64                            `json:"orderItemId,omitempty"`
}

type RightsDetailGetResponseGoodsInfo struct {
	BizInfos          []RightsDetailGetResponseBizInfos       `json:"bizInfos,omitempty"`
	LabelInfos        []RightsDetailGetResponseLabelInfos     `json:"labelInfos,omitempty"`
	GoodsAbilities    []RightsDetailGetResponseGoodsAbilities `json:"goodsAbilities,omitempty"`
	ProductInfos      []RightsDetailGetResponseProductInfos   `json:"productInfos,omitempty"`
	GoodsCode         string                                  `json:"goodsCode,omitempty"`
	GoodsId           int64                                   `json:"goodsId,omitempty"`
	GoodsTitle        string                                  `json:"goodsTitle,omitempty"`
	GoodsType         int64                                   `json:"goodsType,omitempty"`
	ImageUrl          string                                  `json:"imageUrl,omitempty"`
	Price             string                                  `json:"price,omitempty"`
	RightsServiceType int64                                   `json:"rightsServiceType,omitempty"`
	SkuAttrInfo       string                                  `json:"skuAttrInfo,omitempty"`
	SkuBarCode        string                                  `json:"skuBarCode,omitempty"`
	SkuCode           string                                  `json:"skuCode,omitempty"`
	SkuId             int64                                   `json:"skuId,omitempty"`
	SkuNum            string                                  `json:"skuNum,omitempty"`
	SubGoodsType      int64                                   `json:"subGoodsType,omitempty"`
}

type RightsDetailGetResponseBizInfos struct {
	BizId      int64  `json:"bizId,omitempty"`
	BizOrderId string `json:"bizOrderId,omitempty"`
	BizType    int64  `json:"bizType,omitempty"`
	SubBizType int64  `json:"subBizType,omitempty"`
}

type RightsDetailGetResponseLabelInfos struct {
	LabelType string `json:"labelType,omitempty"`
}

type RightsDetailGetResponseGoodsAbilities struct {
	AbilityCode string `json:"abilityCode,omitempty"`
	AbilityType int64  `json:"abilityType,omitempty"`
	BizId       int64  `json:"bizId,omitempty"`
}

type RightsDetailGetResponseProductInfos struct {
	WarehouseInfos  []RightsDetailGetResponseWarehouseInfos `json:"warehouseInfos,omitempty"`
	CombSkuId       int64                                   `json:"combSkuId,omitempty"`
	CombTitle       string                                  `json:"combTitle,omitempty"`
	Title           string                                  `json:"title,omitempty"`
	ProductType     int64                                   `json:"productType,omitempty"`
	ItemSkuQuantity int64                                   `json:"itemSkuQuantity,omitempty"`
	Price           int64                                   `json:"price,omitempty"`
}

type RightsDetailGetResponseWarehouseInfos struct {
	WarehouseType string `json:"warehouseType,omitempty"`
	WarehouseId   int64  `json:"warehouseId,omitempty"`
	WarehouseName string `json:"warehouseName,omitempty"`
	Quantity      int64  `json:"quantity,omitempty"`
}

type RightsDetailGetResponseFlagInfo struct {
	FlagContent string `json:"flagContent,omitempty"`
	FlagRank    int64  `json:"flagRank,omitempty"`
}

type RightsDetailGetResponseMerchantInfo struct {
	MerchantId     int64  `json:"merchantId,omitempty"`
	BosId          int64  `json:"bosId,omitempty"`
	BosName        string `json:"bosName,omitempty"`
	Vid            int64  `json:"vid,omitempty"`
	VidName        string `json:"vidName,omitempty"`
	ProcessVid     int64  `json:"processVid,omitempty"`
	ProcessVidName string `json:"processVidName,omitempty"`
}

type RightsDetailGetResponseOriginOrder struct {
	OrderNo          int64  `json:"orderNo,omitempty"`
	OutOrderNo       string `json:"outOrderNo,omitempty"`
	ApplyOrderStatus int64  `json:"applyOrderStatus,omitempty"`
	OrderType        int64  `json:"orderType,omitempty"`
	OrderSource      int64  `json:"orderSource,omitempty"`
	ChannelType      int64  `json:"channelType,omitempty"`
	BizSourceType    int64  `json:"bizSourceType,omitempty"`
	PayType          int64  `json:"payType,omitempty"`
}

type RightsDetailGetResponseOutRightsInfo struct {
	OutRightsNo string `json:"outRightsNo,omitempty"`
}

type RightsDetailGetResponseRefundAccount struct {
	AccountNum     string `json:"accountNum,omitempty"`
	OpeningBank    string `json:"openingBank,omitempty"`
	RefundMethodId int64  `json:"refundMethodId,omitempty"`
	UserName       string `json:"userName,omitempty"`
}

type RightsDetailGetResponseReturnOrder struct {
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

type RightsDetailGetResponseRightsAssets struct {
	AssetsAmount string `json:"assetsAmount,omitempty"`
	AssetsNum    string `json:"assetsNum,omitempty"`
	AssetsTarget int64  `json:"assetsTarget,omitempty"`
	AssetsType   int64  `json:"assetsType,omitempty"`
}

type RightsDetailGetResponseRightsItems struct {
	GoodsInfo        RightsDetailGetResponseGoodsInfo2      `json:"goodsInfo,omitempty"`
	RightsAssets     []RightsDetailGetResponseRightsAssets2 `json:"rightsAssets,omitempty"`
	RefundDetail     RightsDetailGetResponseRefundDetail    `json:"refundDetail,omitempty"`
	ApplyNum         string                                 `json:"applyNum,omitempty"`
	GoodsReceiveType int64                                  `json:"goodsReceiveType,omitempty"`
	OrderItemId      int64                                  `json:"orderItemId,omitempty"`
	ReturnNum        string                                 `json:"returnNum,omitempty"`
	RightsItemId     int64                                  `json:"rightsItemId,omitempty"`
}

type RightsDetailGetResponseGoodsInfo2 struct {
	BizInfos          []RightsDetailGetResponseBizInfos2       `json:"bizInfos,omitempty"`
	LabelInfos        []RightsDetailGetResponseLabelInfos2     `json:"labelInfos,omitempty"`
	GoodsAbilities    []RightsDetailGetResponseGoodsAbilities2 `json:"goodsAbilities,omitempty"`
	ProductInfos      []RightsDetailGetResponseProductInfos2   `json:"productInfos,omitempty"`
	GoodsCode         string                                   `json:"goodsCode,omitempty"`
	GoodsId           int64                                    `json:"goodsId,omitempty"`
	GoodsTitle        string                                   `json:"goodsTitle,omitempty"`
	GoodsType         int64                                    `json:"goodsType,omitempty"`
	ImageUrl          string                                   `json:"imageUrl,omitempty"`
	Price             string                                   `json:"price,omitempty"`
	RightsServiceType int64                                    `json:"rightsServiceType,omitempty"`
	SkuAttrInfo       string                                   `json:"skuAttrInfo,omitempty"`
	SkuBarCode        string                                   `json:"skuBarCode,omitempty"`
	SkuCode           string                                   `json:"skuCode,omitempty"`
	SkuId             int64                                    `json:"skuId,omitempty"`
	SkuNum            string                                   `json:"skuNum,omitempty"`
	SubGoodsType      int64                                    `json:"subGoodsType,omitempty"`
}

type RightsDetailGetResponseBizInfos2 struct {
	BizId      int64  `json:"bizId,omitempty"`
	BizOrderId string `json:"bizOrderId,omitempty"`
	BizType    int64  `json:"bizType,omitempty"`
	SubBizType int64  `json:"subBizType,omitempty"`
}

type RightsDetailGetResponseLabelInfos2 struct {
	LabelType string `json:"labelType,omitempty"`
}

type RightsDetailGetResponseGoodsAbilities2 struct {
	AbilityCode string `json:"abilityCode,omitempty"`
	AbilityType int64  `json:"abilityType,omitempty"`
	BizId       int64  `json:"bizId,omitempty"`
}

type RightsDetailGetResponseProductInfos2 struct {
	WarehouseInfos  []RightsDetailGetResponseWarehouseInfos2 `json:"warehouseInfos,omitempty"`
	CombSkuId       int64                                    `json:"combSkuId,omitempty"`
	CombTitle       string                                   `json:"combTitle,omitempty"`
	Title           string                                   `json:"title,omitempty"`
	ProductType     int64                                    `json:"productType,omitempty"`
	ItemSkuQuantity int64                                    `json:"itemSkuQuantity,omitempty"`
	Price           int64                                    `json:"price,omitempty"`
}

type RightsDetailGetResponseWarehouseInfos2 struct {
	WarehouseType string `json:"warehouseType,omitempty"`
	WarehouseId   int64  `json:"warehouseId,omitempty"`
	WarehouseName string `json:"warehouseName,omitempty"`
	Quantity      int64  `json:"quantity,omitempty"`
}

type RightsDetailGetResponseRightsAssets2 struct {
	AssetsAmount string `json:"assetsAmount,omitempty"`
	AssetsNum    string `json:"assetsNum,omitempty"`
	AssetsTarget int64  `json:"assetsTarget,omitempty"`
	AssetsType   int64  `json:"assetsType,omitempty"`
}

type RightsDetailGetResponseRefundDetail struct {
	ApplyAmountInfos  []RightsDetailGetResponseApplyAmountInfos  `json:"applyAmountInfos,omitempty"`
	RefundAmountInfos []RightsDetailGetResponseRefundAmountInfos `json:"refundAmountInfos,omitempty"`
	ApplyAmount       string                                     `json:"applyAmount,omitempty"`
	RefundAmount      string                                     `json:"refundAmount,omitempty"`
}

type RightsDetailGetResponseApplyAmountInfos struct {
	Type   int64  `json:"type,omitempty"`
	Amount string `json:"amount,omitempty"`
}

type RightsDetailGetResponseRefundAmountInfos struct {
	Type   int64  `json:"type,omitempty"`
	Amount string `json:"amount,omitempty"`
}

type RightsDetailGetResponseRightsOrder struct {
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

type RightsDetailGetResponseRightsReason struct {
	ReasonImageUrls []int64 `json:"reasonImageUrls,omitempty"`
	RightsReason    string  `json:"rightsReason,omitempty"`
	RightsRemark    string  `json:"rightsRemark,omitempty"`
}

type RightsDetailGetResponseRefundDetail2 struct {
	ApplyAmountInfos  []RightsDetailGetResponseApplyAmountInfos2  `json:"applyAmountInfos,omitempty"`
	RefundAmountInfos []RightsDetailGetResponseRefundAmountInfos2 `json:"refundAmountInfos,omitempty"`
	ApplyAmount       string                                      `json:"applyAmount,omitempty"`
	RefundAmount      string                                      `json:"refundAmount,omitempty"`
}

type RightsDetailGetResponseApplyAmountInfos2 struct {
	Type   int64  `json:"type,omitempty"`
	Amount string `json:"amount,omitempty"`
}

type RightsDetailGetResponseRefundAmountInfos2 struct {
	Type   int64  `json:"type,omitempty"`
	Amount string `json:"amount,omitempty"`
}

type RightsDetailGetResponseRightsStatusLogs struct {
	DateTime int64  `json:"dateTime,omitempty"`
	Type     string `json:"type,omitempty"`
}

type RightsDetailGetResponseRefuseInfo struct {
	RefusedReason string `json:"refusedReason,omitempty"`
}
