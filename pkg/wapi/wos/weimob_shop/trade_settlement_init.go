package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TradeSettlementInit
// @id 1690
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1690?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=确认订单)
func (client *Client) TradeSettlementInit(request *TradeSettlementInitRequest) (response *TradeSettlementInitResponse, err error) {
	rpcResponse := CreateTradeSettlementInitResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TradeSettlementInitRequest struct {
	*api.BaseRequest

	GoodsList             []TradeSettlementInitRequestGoodsList  `json:"goodsList,omitempty"`
	ExtendInfo            TradeSettlementInitRequestExtendInfo   `json:"extendInfo,omitempty"`
	BasicInfo             TradeSettlementInitRequestBasicInfo    `json:"basicInfo,omitempty"`
	CouponList            []TradeSettlementInitRequestCouponList `json:"couponList,omitempty"`
	Remark                TradeSettlementInitRequestRemark       `json:"remark,omitempty"`
	Delivery              TradeSettlementInitRequestDelivery     `json:"delivery,omitempty"`
	Openid                string                                 `json:"openid,omitempty"`
	OrderPageSource       int64                                  `json:"orderPageSource,omitempty"`
	BosId                 int64                                  `json:"bosId,omitempty"`
	BizWid                int64                                  `json:"bizWid,omitempty"`
	Appid                 string                                 `json:"appid,omitempty"`
	OrderNo               int64                                  `json:"orderNo,omitempty"`
	OperationSource       string                                 `json:"operationSource,omitempty"`
	TemplateId            int64                                  `json:"templateId,omitempty"`
	BizVid                int64                                  `json:"bizVid,omitempty"`
	BizVidType            int64                                  `json:"bizVidType,omitempty"`
	Channel               int64                                  `json:"channel,omitempty"`
	ChannelId             int64                                  `json:"channelId,omitempty"`
	TradeTrackId          string                                 `json:"tradeTrackId,omitempty"`
	ChannelScene          int64                                  `json:"channelScene,omitempty"`
	TradeComponentVersion string                                 `json:"tradeComponentVersion,omitempty"`
	OuterOrderNo          string                                 `json:"outerOrderNo,omitempty"`
	OuterBusinessType     string                                 `json:"outerBusinessType,omitempty"`
}

type TradeSettlementInitResponse struct {
	CommitOrderSwitches     TradeSettlementInitResponseCommitOrderSwitches `json:"commitOrderSwitches,omitempty"`
	ValidBizResult          TradeSettlementInitResponseValidBizResult      `json:"validBizResult,omitempty"`
	TargetProductInstanceId int64                                          `json:"targetProductInstanceId,omitempty"`
	TradeTrackId            string                                         `json:"tradeTrackId,omitempty"`
	ConfirmOrderKey         string                                         `json:"confirmOrderKey,omitempty"`
}

func CreateTradeSettlementInitRequest() (request *TradeSettlementInitRequest) {
	request = &TradeSettlementInitRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "trade/settlement/init", "apigw")
	request.Method = api.POST
	return
}

func CreateTradeSettlementInitResponse() (response *api.BaseResponse[TradeSettlementInitResponse]) {
	response = api.CreateResponse[TradeSettlementInitResponse](&TradeSettlementInitResponse{})
	return
}

type TradeSettlementInitRequestGoodsList struct {
	LabelList      []TradeSettlementInitRequestLabelList     `json:"labelList,omitempty"`
	Activities     []TradeSettlementInitRequestActivities    `json:"activities,omitempty"`
	GoodsAbility   TradeSettlementInitRequestGoodsAbility    `json:"goodsAbility,omitempty"`
	FeeInfoList    []TradeSettlementInitRequestFeeInfoList   `json:"feeInfoList,omitempty"`
	ItemMsgList    []TradeSettlementInitRequestItemMsgList   `json:"itemMsgList,omitempty"`
	WarehouseList  []TradeSettlementInitRequestWarehouseList `json:"warehouseList,omitempty"`
	Vid            int64                                     `json:"vid,omitempty"`
	BuyNum         int64                                     `json:"buyNum,omitempty"`
	GoodsId        int64                                     `json:"goodsId,omitempty"`
	SkuId          int64                                     `json:"skuId,omitempty"`
	TradeGoodsType int64                                     `json:"tradeGoodsType,omitempty"`
	ItemId         int64                                     `json:"itemId,omitempty"`
	CloudCustom    string                                    `json:"cloudCustom,omitempty"`
	CartItemId     int64                                     `json:"cartItemId,omitempty"`
}

type TradeSettlementInitRequestLabelList struct {
	LabelType  string `json:"labelType,omitempty"`
	AttachId   string `json:"attachId,omitempty"`
	Attachment string `json:"attachment,omitempty"`
	CreateTime int64  `json:"createTime,omitempty"`
}

type TradeSettlementInitRequestActivities struct {
	DiscountTicket  string `json:"discountTicket,omitempty"`
	ActivityType    int64  `json:"activityType,omitempty"`
	DiscountStep    string `json:"discountStep,omitempty"`
	ActivityId      string `json:"activityId,omitempty"`
	Sort            int64  `json:"sort,omitempty"`
	GroupKey        string `json:"groupKey,omitempty"`
	Level           int64  `json:"level,omitempty"`
	UseType         int64  `json:"useType,omitempty"`
	SubDiscountType int64  `json:"subDiscountType,omitempty"`
}

type TradeSettlementInitRequestGoodsAbility struct {
	CycleGoods TradeSettlementInitRequestCycleGoods `json:"cycleGoods,omitempty"`
	PointGoods TradeSettlementInitRequestPointGoods `json:"pointGoods,omitempty"`
	DrawGoods  TradeSettlementInitRequestDrawGoods  `json:"drawGoods,omitempty"`
}

type TradeSettlementInitRequestCycleGoods struct {
	AbilityCode              string  `json:"abilityCode,omitempty"`
	BizId                    string  `json:"bizId,omitempty"`
	CycleUnit                string  `json:"cycleUnit,omitempty"`
	SelectCycleValues        []int64 `json:"selectCycleValues,omitempty"`
	LatestFulfillDate        string  `json:"latestFulfillDate,omitempty"`
	SelectFulfillTimeOptions string  `json:"selectFulfillTimeOptions,omitempty"`
}

type TradeSettlementInitRequestPointGoods struct {
	AbilityCode string `json:"abilityCode,omitempty"`
	BizId       int64  `json:"bizId,omitempty"`
}

type TradeSettlementInitRequestDrawGoods struct {
	AbilityCode string `json:"abilityCode,omitempty"`
	BizId       int64  `json:"bizId,omitempty"`
}

type TradeSettlementInitRequestFeeInfoList struct {
	Type                   int64   `json:"type,omitempty"`
	Name                   string  `json:"name,omitempty"`
	Amount                 int64   `json:"amount,omitempty"`
	Remark                 string  `json:"remark,omitempty"`
	CustomDiscountTypeList []int64 `json:"customDiscountTypeList,omitempty"`
}

type TradeSettlementInitRequestItemMsgList struct {
	ItemMsgReq TradeSettlementInitRequestItemMsgReq `json:"ItemMsgReq,omitempty"`
}

type TradeSettlementInitRequestItemMsgReq struct {
	Key   string `json:"key,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type TradeSettlementInitRequestWarehouseList struct {
	WarehouseReq TradeSettlementInitRequestWarehouseReq `json:"WarehouseReq,omitempty"`
}

type TradeSettlementInitRequestWarehouseReq struct {
	WarehouseId   int64  `json:"warehouseId,omitempty"`
	WarehouseName string `json:"warehouseName,omitempty"`
	WarehouseType string `json:"warehouseType,omitempty"`
	Quantity      int64  `json:"quantity,omitempty"`
}

type TradeSettlementInitRequestExtendInfo struct {
	DeviceInfo TradeSettlementInitRequestDeviceInfo `json:"deviceInfo,omitempty"`
	Refer      string                               `json:"refer,omitempty"`
	Source     int64                                `json:"source,omitempty"`
	UserIp     string                               `json:"userIp,omitempty"`
	OcdAppId   int64                                `json:"ocdAppId,omitempty"`
	Uuid       string                               `json:"uuid,omitempty"`
}

type TradeSettlementInitRequestDeviceInfo struct {
	OsType     int64  `json:"osType,omitempty"`
	DeviceType string `json:"deviceType,omitempty"`
	OsVersion  string `json:"osVersion,omitempty"`
}

type TradeSettlementInitRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type TradeSettlementInitRequestCouponList struct {
	ResourceTemplateId int64  `json:"resourceTemplateId,omitempty"`
	HasReceived        string `json:"hasReceived,omitempty"`
	CouponCode         string `json:"couponCode,omitempty"`
}

type TradeSettlementInitRequestRemark struct {
	Remark string `json:"remark,omitempty"`
}

type TradeSettlementInitRequestDelivery struct {
	DeliveryType        TradeSettlementInitRequestDeliveryType          `json:"deliveryType,omitempty"`
	DeliveryDateTime    TradeSettlementInitRequestDeliveryDateTime      `json:"deliveryDateTime,omitempty"`
	DeliveryCustomField TradeSettlementInitRequestDeliveryCustomField   `json:"deliveryCustomField,omitempty"`
	ReceiveAddress      TradeSettlementInitRequestReceiveAddress        `json:"receiveAddress,omitempty"`
	CustomFieldDataList []TradeSettlementInitRequestCustomFieldDataList `json:"customFieldDataList,omitempty"`
}

type TradeSettlementInitRequestDeliveryType struct {
	DeliveryType     int64  `json:"deliveryType,omitempty"`
	DeliveryTypeName string `json:"deliveryTypeName,omitempty"`
	ProcessVid       int64  `json:"processVid,omitempty"`
}

type TradeSettlementInitRequestDeliveryDateTime struct {
	DeliveryDatetimeType int64   `json:"deliveryDatetimeType,omitempty"`
	SelectedDatetimeList []int64 `json:"selectedDatetimeList,omitempty"`
}

type TradeSettlementInitRequestDeliveryCustomField struct {
	SelfPickUpSiteId   string `json:"selfPickUpSiteId,omitempty"`
	SelfPickUpSiteName string `json:"selfPickUpSiteName,omitempty"`
	SelfPickupMobile   string `json:"selfPickupMobile,omitempty"`
	SelfPickupUserName string `json:"selfPickupUserName,omitempty"`
}

type TradeSettlementInitRequestReceiveAddress struct {
	ReceiveName   string `json:"receiveName,omitempty"`
	ReceiveTel    string `json:"receiveTel,omitempty"`
	ProvinceCode  string `json:"provinceCode,omitempty"`
	ProvinceName  string `json:"provinceName,omitempty"`
	CityCode      string `json:"cityCode,omitempty"`
	CityName      string `json:"cityName,omitempty"`
	CountyCode    string `json:"countyCode,omitempty"`
	CountryName   string `json:"countryName,omitempty"`
	AreaCode      string `json:"areaCode,omitempty"`
	AreaName      string `json:"areaName,omitempty"`
	DetailAddress string `json:"detailAddress,omitempty"`
	HouseNumber   string `json:"houseNumber,omitempty"`
	Longitude     string `json:"longitude,omitempty"`
	Latitude      string `json:"latitude,omitempty"`
	PostCode      string `json:"postCode,omitempty"`
	PhoneZone     string `json:"phoneZone,omitempty"`
}

type TradeSettlementInitRequestCustomFieldDataList struct {
	CustomFieldReq TradeSettlementInitRequestCustomFieldReq `json:"CustomFieldReq,omitempty"`
}

type TradeSettlementInitRequestCustomFieldReq struct {
	ComponentType string `json:"componentType,omitempty"`
	Key           string `json:"key,omitempty"`
	Name          string `json:"name,omitempty"`
	Value         string `json:"value,omitempty"`
}

type TradeSettlementInitResponseCommitOrderSwitches struct {
	UnableCommitFactorList []TradeSettlementInitResponseUnableCommitFactorList `json:"unableCommitFactorList,omitempty"`
	EnableCommit           bool                                                `json:"enableCommit,omitempty"`
}

type TradeSettlementInitResponseUnableCommitFactorList struct {
	UnableCommitType   int64  `json:"unableCommitType,omitempty"`
	UnableCommitReason string `json:"unableCommitReason,omitempty"`
}

type TradeSettlementInitResponseValidBizResult struct {
	ValidBizType int64  `json:"validBizType,omitempty"`
	ValidBizInfo string `json:"validBizInfo,omitempty"`
	ValidSuccess bool   `json:"validSuccess,omitempty"`
}
