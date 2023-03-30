package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsImport
// @id 2285
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2285?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=导入售后单)
func (client *Client) RightsImport(request *RightsImportRequest) (response *RightsImportResponse, err error) {
	rpcResponse := CreateRightsImportResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsImportRequest struct {
	*api.BaseRequest

	BuyerInfo        RightsImportRequestBuyerInfo          `json:"buyerInfo,omitempty"`
	OriginOrder      RightsImportRequestOriginOrder        `json:"originOrder,omitempty"`
	RefundDetail     RightsImportRequestRefundDetail       `json:"refundDetail,omitempty"`
	RightsAssets     []RightsImportRequestRightsAssets     `json:"rightsAssets,omitempty"`
	MerchantInfo     RightsImportRequestMerchantInfo       `json:"merchantInfo,omitempty"`
	RightsItems      []RightsImportRequestRightsItems      `json:"rightsItems,omitempty"`
	RefundAccount    RightsImportRequestRefundAccount      `json:"refundAccount,omitempty"`
	RightsOrder      RightsImportRequestRightsOrder        `json:"rightsOrder,omitempty"`
	GuiderInfo       RightsImportRequestGuiderInfo         `json:"guiderInfo,omitempty"`
	AssociateBizList []RightsImportRequestAssociateBizList `json:"associateBizList,omitempty"`
}

type RightsImportResponse struct {
	RightsOrderNo int64 `json:"rightsOrderNo,omitempty"`
}

func CreateRightsImportRequest() (request *RightsImportRequest) {
	request = &RightsImportRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "rights/import", "apigw")
	request.Method = api.POST
	return
}

func CreateRightsImportResponse() (response *api.BaseResponse[RightsImportResponse]) {
	response = api.CreateResponse[RightsImportResponse](&RightsImportResponse{})
	return
}

type RightsImportRequestBuyerInfo struct {
	Wid        int64  `json:"wid,omitempty"`
	BuyerPhone string `json:"buyerPhone,omitempty"`
	OutCardNo  string `json:"outCardNo,omitempty"`
}

type RightsImportRequestOriginOrder struct {
	OutOriginalOrderNo string `json:"outOriginalOrderNo,omitempty"`
	OutOrderNo         string `json:"outOrderNo,omitempty"`
}

type RightsImportRequestRefundDetail struct {
	ApplyAmount  string `json:"applyAmount,omitempty"`
	RefundAmount string `json:"refundAmount,omitempty"`
}

type RightsImportRequestRightsAssets struct {
	AssetAmount           string `json:"assetAmount,omitempty"`
	AssetsTarget          int64  `json:"assetsTarget,omitempty"`
	CouponType            int64  `json:"couponType,omitempty"`
	AssetNum              string `json:"assetNum,omitempty"`
	CouponCode            string `json:"couponCode,omitempty"`
	CouponDeductionAmount string `json:"couponDeductionAmount,omitempty"`
	AssetType             int64  `json:"assetType,omitempty"`
	CouponTitle           string `json:"couponTitle,omitempty"`
}

type RightsImportRequestMerchantInfo struct {
	Vid             int64  `json:"vid,omitempty"`
	OuterMerchantId string `json:"outerMerchantId,omitempty"`
}

type RightsImportRequestRightsItems struct {
	RefundDetail    RightsImportRequestRefundDetail2   `json:"refundDetail,omitempty"`
	RightsAssets    []RightsImportRequestRightsAssets2 `json:"rightsAssets,omitempty"`
	OutOrderItemId  string                             `json:"outOrderItemId,omitempty"`
	MarketPrice     string                             `json:"marketPrice,omitempty"`
	GuiderName      string                             `json:"guiderName,omitempty"`
	OutGuiderNo     string                             `json:"outGuiderNo,omitempty"`
	SalePrice       string                             `json:"salePrice,omitempty"`
	OrderItemId     int64                              `json:"orderItemId,omitempty"`
	GoodsId         int64                              `json:"goodsId,omitempty"`
	GuiderPhone     string                             `json:"guiderPhone,omitempty"`
	ApplyNum        string                             `json:"applyNum,omitempty"`
	SkuName         string                             `json:"skuName,omitempty"`
	PayAmount       string                             `json:"payAmount,omitempty"`
	OutGoodsId      string                             `json:"outGoodsId,omitempty"`
	GuiderWid       int64                              `json:"guiderWid,omitempty"`
	OutRightsItemId string                             `json:"outRightsItemId,omitempty"`
	GoodsTitle      string                             `json:"goodsTitle,omitempty"`
	OutSkuId        string                             `json:"outSkuId,omitempty"`
	SkuId           int64                              `json:"skuId,omitempty"`
	GuiderNo        string                             `json:"guiderNo,omitempty"`
}

type RightsImportRequestRefundDetail2 struct {
	ApplyAmount  string `json:"applyAmount,omitempty"`
	RefundAmount string `json:"refundAmount,omitempty"`
}

type RightsImportRequestRightsAssets2 struct {
	AssetAmount  string `json:"assetAmount,omitempty"`
	AssetsTarget int64  `json:"assetsTarget,omitempty"`
	AssetNum     string `json:"assetNum,omitempty"`
	AssetType    int64  `json:"assetType,omitempty"`
}

type RightsImportRequestRefundAccount struct {
	RefundMethodId  int64 `json:"refundMethodId,omitempty"`
	PayTradeOrderId int64 `json:"payTradeOrderId,omitempty"`
}

type RightsImportRequestRightsOrder struct {
	Reason          string `json:"reason,omitempty"`
	FinishTime      int64  `json:"finishTime,omitempty"`
	RefundMethod    int64  `json:"refundMethod,omitempty"`
	RightsStatus    int64  `json:"rightsStatus,omitempty"`
	RightsType      int64  `json:"rightsType,omitempty"`
	AgreeTime       int64  `json:"agreeTime,omitempty"`
	PlatformType    int64  `json:"platformType,omitempty"`
	ChannelType     int64  `json:"channelType,omitempty"`
	Remark          string `json:"remark,omitempty"`
	RefundType      int64  `json:"refundType,omitempty"`
	CreateTime      int64  `json:"createTime,omitempty"`
	OutRightsNo     string `json:"outRightsNo,omitempty"`
	RightsCauseType int64  `json:"rightsCauseType,omitempty"`
}

type RightsImportRequestGuiderInfo struct {
	GuiderNo        string `json:"guiderNo,omitempty"`
	OutGuiderNo     string `json:"outGuiderNo,omitempty"`
	PrivateGuiderNo string `json:"privateGuiderNo,omitempty"`
}

type RightsImportRequestAssociateBizList struct {
	BizValue int64 `json:"bizValue,omitempty"`
}
