package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TradeSettlementCalcDiscount
// @id 2317
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2317?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=计算最优价接口)
func (client *Client) TradeSettlementCalcDiscount(request *TradeSettlementCalcDiscountRequest) (response *TradeSettlementCalcDiscountResponse, err error) {
	rpcResponse := CreateTradeSettlementCalcDiscountResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TradeSettlementCalcDiscountRequest struct {
	*api.BaseRequest

	GoodsList  []TradeSettlementCalcDiscountRequestGoodsList `json:"goodsList,omitempty"`
	ExtendInfo TradeSettlementCalcDiscountRequestExtendInfo  `json:"extendInfo,omitempty"`
	BasicInfo  TradeSettlementCalcDiscountRequestBasicInfo   `json:"basicInfo,omitempty"`
	BizWid     int64                                         `json:"bizWid,omitempty"`
	ChannelId  int64                                         `json:"channelId,omitempty"`
	BosId      int64                                         `json:"bosId,omitempty"`
}

type TradeSettlementCalcDiscountResponse struct {
	GoodsList      []TradeSettlementCalcDiscountResponseGoodsList    `json:"goodsList,omitempty"`
	OrderFundInfo  TradeSettlementCalcDiscountResponseOrderFundInfo  `json:"orderFundInfo,omitempty"`
	ValidBizResult TradeSettlementCalcDiscountResponseValidBizResult `json:"validBizResult,omitempty"`
}

func CreateTradeSettlementCalcDiscountRequest() (request *TradeSettlementCalcDiscountRequest) {
	request = &TradeSettlementCalcDiscountRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "trade/settlement/calcDiscount", "apigw")
	request.Method = api.POST
	return
}

func CreateTradeSettlementCalcDiscountResponse() (response *api.BaseResponse[TradeSettlementCalcDiscountResponse]) {
	response = api.CreateResponse[TradeSettlementCalcDiscountResponse](&TradeSettlementCalcDiscountResponse{})
	return
}

type TradeSettlementCalcDiscountRequestGoodsList struct {
	Activities     []TradeSettlementCalcDiscountRequestActivities `json:"activities,omitempty"`
	Vid            int64                                          `json:"vid,omitempty"`
	ItemId         int64                                          `json:"itemId,omitempty"`
	BuyNum         int64                                          `json:"buyNum,omitempty"`
	TradeGoodsType int64                                          `json:"tradeGoodsType,omitempty"`
	GoodsId        int64                                          `json:"goodsId,omitempty"`
	SkuId          int64                                          `json:"skuId,omitempty"`
}

type TradeSettlementCalcDiscountRequestActivities struct {
	ActivityId      string `json:"activityId,omitempty"`
	DiscountTicket  string `json:"discountTicket,omitempty"`
	Level           int64  `json:"level,omitempty"`
	SubDiscountType int64  `json:"subDiscountType,omitempty"`
	Sort            int64  `json:"sort,omitempty"`
	ActivityType    int64  `json:"activityType,omitempty"`
	DiscountStep    string `json:"discountStep,omitempty"`
	UseType         int64  `json:"useType,omitempty"`
	GroupKey        string `json:"groupKey,omitempty"`
}

type TradeSettlementCalcDiscountRequestExtendInfo struct {
	DeviceInfo TradeSettlementCalcDiscountRequestDeviceInfo `json:"deviceInfo,omitempty"`
	OcdAppId   int64                                        `json:"ocdAppId,omitempty"`
	Refer      string                                       `json:"refer,omitempty"`
	UserIp     string                                       `json:"userIp,omitempty"`
	Source     int64                                        `json:"source,omitempty"`
	Uuid       string                                       `json:"uuid,omitempty"`
}

type TradeSettlementCalcDiscountRequestDeviceInfo struct {
	DeviceType string `json:"deviceType,omitempty"`
	OsVersion  string `json:"osVersion,omitempty"`
	OsType     int64  `json:"osType,omitempty"`
	DeviceNo   string `json:"deviceNo,omitempty"`
}

type TradeSettlementCalcDiscountRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type TradeSettlementCalcDiscountResponseGoodsList struct {
	UsedActivityList []TradeSettlementCalcDiscountResponseUsedActivityList `json:"usedActivityList,omitempty"`
	ItemId           int64                                                 `json:"itemId,omitempty"`
	GoodsId          int64                                                 `json:"goodsId,omitempty"`
	SkuId            int64                                                 `json:"skuId,omitempty"`
	BuyNum           int64                                                 `json:"buyNum,omitempty"`
	TradeGoodsType   int64                                                 `json:"tradeGoodsType,omitempty"`
	SkuMarketPrice   string                                                `json:"skuMarketPrice,omitempty"`
	SkuSalePrice     string                                                `json:"skuSalePrice,omitempty"`
}

type TradeSettlementCalcDiscountResponseUsedActivityList struct {
	ActivityId      string `json:"activityId,omitempty"`
	ActivityName    string `json:"activityName,omitempty"`
	ActivityType    int64  `json:"activityType,omitempty"`
	ActivitySubType int64  `json:"activitySubType,omitempty"`
	DiscountAmount  string `json:"discountAmount,omitempty"`
}

type TradeSettlementCalcDiscountResponseOrderFundInfo struct {
	UsedActivityList    []TradeSettlementCalcDiscountResponseUsedActivityList2 `json:"usedActivityList,omitempty"`
	OrderAmount         string                                                 `json:"orderAmount,omitempty"`
	OrderDiscountAmount string                                                 `json:"orderDiscountAmount,omitempty"`
}

type TradeSettlementCalcDiscountResponseUsedActivityList2 struct {
	ActivityId      string `json:"activityId,omitempty"`
	ActivityName    string `json:"activityName,omitempty"`
	ActivityType    int64  `json:"activityType,omitempty"`
	ActivitySubType int64  `json:"activitySubType,omitempty"`
	DiscountAmount  string `json:"discountAmount,omitempty"`
}

type TradeSettlementCalcDiscountResponseValidBizResult struct {
	ValidBizType int64  `json:"validBizType,omitempty"`
	ValidBizInfo string `json:"validBizInfo,omitempty"`
	ValidSuccess bool   `json:"validSuccess,omitempty"`
}
