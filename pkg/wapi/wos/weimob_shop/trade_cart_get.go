package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TradeCartGet
// @id 2205
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2205?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查看购物车)
func (client *Client) TradeCartGet(request *TradeCartGetRequest) (response *TradeCartGetResponse, err error) {
	rpcResponse := CreateTradeCartGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TradeCartGetRequest struct {
	*api.BaseRequest

	ExtendInfo     TradeCartGetRequestExtendInfo     `json:"extendInfo,omitempty"`
	QueryParameter TradeCartGetRequestQueryParameter `json:"queryParameter,omitempty"`
	BasicInfo      TradeCartGetRequestBasicInfo      `json:"basicInfo,omitempty"`
	BizWid         int64                             `json:"bizWid,omitempty"`
}

type TradeCartGetResponse struct {
	BuyInfo        TradeCartGetResponseBuyInfo        `json:"buyInfo,omitempty"`
	PaymentInfo    TradeCartGetResponsePaymentInfo    `json:"paymentInfo,omitempty"`
	ValidBizResult TradeCartGetResponseValidBizResult `json:"validBizResult,omitempty"`
	VidInfoList    []TradeCartGetResponseVidInfoList  `json:"vidInfoList,omitempty"`
	ItemList       []TradeCartGetResponseItemList     `json:"itemList,omitempty"`
}

func CreateTradeCartGetRequest() (request *TradeCartGetRequest) {
	request = &TradeCartGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "trade/cart/get", "apigw")
	request.Method = api.POST
	return
}

func CreateTradeCartGetResponse() (response *api.BaseResponse[TradeCartGetResponse]) {
	response = api.CreateResponse[TradeCartGetResponse](&TradeCartGetResponse{})
	return
}

type TradeCartGetRequestExtendInfo struct {
	Refer          string `json:"refer,omitempty"`
	Source         int64  `json:"source,omitempty"`
	TemplateId     int64  `json:"templateId,omitempty"`
	BusinessSource int64  `json:"businessSource,omitempty"`
}

type TradeCartGetRequestQueryParameter struct {
	Activities []TradeCartGetRequestActivities `json:"activities,omitempty"`
}

type TradeCartGetRequestActivities struct {
	ActivityId   int64 `json:"activityId,omitempty"`
	ActivityType int64 `json:"activityType,omitempty"`
}

type TradeCartGetRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type TradeCartGetResponseBuyInfo struct {
	Wid int64 `json:"wid,omitempty"`
}

type TradeCartGetResponsePaymentInfo struct {
	DiscountList        []TradeCartGetResponseDiscountList `json:"discountList,omitempty"`
	IsAllSelected       bool                               `json:"isAllSelected,omitempty"`
	PaymentAmount       string                             `json:"paymentAmount,omitempty"`
	SelectedGoodsCount  int64                              `json:"selectedGoodsCount,omitempty"`
	TotalAmount         string                             `json:"totalAmount,omitempty"`
	TotalDiscountAmount string                             `json:"totalDiscountAmount,omitempty"`
	TotalGoodsNum       int64                              `json:"totalGoodsNum,omitempty"`
}

type TradeCartGetResponseDiscountList struct {
	DiscountId     int64  `json:"discountId,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	DiscountAmount string `json:"discountAmount,omitempty"`
}

type TradeCartGetResponseValidBizResult struct {
	ValidBizInfo TradeCartGetResponseValidBizInfo `json:"validBizInfo,omitempty"`
	ValidSuccess bool                             `json:"validSuccess,omitempty"`
	ValidBizType string                           `json:"validBizType,omitempty"`
}

type TradeCartGetResponseValidBizInfo struct {
}

type TradeCartGetResponseVidInfoList struct {
	BosId   int64  `json:"bosId,omitempty"`
	Vid     int64  `json:"vid,omitempty"`
	VidType int64  `json:"vidType,omitempty"`
	VidName string `json:"vidName,omitempty"`
}

type TradeCartGetResponseItemList struct {
	UsedActivityList          []TradeCartGetResponseUsedActivityList `json:"usedActivityList,omitempty"`
	DiscountList              []TradeCartGetResponseDiscountList2    `json:"discountList,omitempty"`
	AddTime                   TradeCartGetResponseAddTime            `json:"addTime,omitempty"`
	LabelList                 []TradeCartGetResponseLabelList        `json:"labelList,omitempty"`
	ItemId                    int64                                  `json:"itemId,omitempty"`
	CartStatus                int64                                  `json:"cartStatus,omitempty"`
	StatusDesc                string                                 `json:"statusDesc,omitempty"`
	TradeGoodsType            int64                                  `json:"tradeGoodsType,omitempty"`
	Vid                       int64                                  `json:"vid,omitempty"`
	GoodsId                   int64                                  `json:"goodsId,omitempty"`
	GoodsTitle                string                                 `json:"goodsTitle,omitempty"`
	SkuId                     int64                                  `json:"skuId,omitempty"`
	SkuAttrInfo               string                                 `json:"skuAttrInfo,omitempty"`
	Num                       int64                                  `json:"num,omitempty"`
	SkuImage                  string                                 `json:"skuImage,omitempty"`
	SalePrice                 string                                 `json:"salePrice,omitempty"`
	AvailableStockNum         int64                                  `json:"availableStockNum,omitempty"`
	ActivityAvailableStockNum int64                                  `json:"activityAvailableStockNum,omitempty"`
	CanBuyNum                 int64                                  `json:"canBuyNum,omitempty"`
}

type TradeCartGetResponseUsedActivityList struct {
	ActivityId       int64 `json:"activityId,omitempty"`
	ActivityType     int64 `json:"activityType,omitempty"`
	ActivityCategory int64 `json:"activityCategory,omitempty"`
	UseStatus        int64 `json:"useStatus,omitempty"`
	IsEnjoyDiscount  bool  `json:"isEnjoyDiscount,omitempty"`
}

type TradeCartGetResponseDiscountList2 struct {
	DiscountId     int64  `json:"discountId,omitempty"`
	DiscountType   int64  `json:"discountType,omitempty"`
	DiscountAmount string `json:"discountAmount,omitempty"`
}

type TradeCartGetResponseAddTime struct {
}

type TradeCartGetResponseLabelList struct {
	Attachment TradeCartGetResponseAttachment `json:"attachment,omitempty"`
	LabelType  string                         `json:"labelType,omitempty"`
	AttachId   string                         `json:"attachId,omitempty"`
	CreateTime int64                          `json:"createTime,omitempty"`
}

type TradeCartGetResponseAttachment struct {
}
