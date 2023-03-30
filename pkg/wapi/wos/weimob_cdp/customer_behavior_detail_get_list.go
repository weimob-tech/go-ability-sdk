package weimob_cdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerBehaviorDetailGetList
// @id 1910
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1910?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询客户行为明细列表)
func (client *Client) CustomerBehaviorDetailGetList(request *CustomerBehaviorDetailGetListRequest) (response *CustomerBehaviorDetailGetListResponse, err error) {
	rpcResponse := CreateCustomerBehaviorDetailGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerBehaviorDetailGetListRequest struct {
	*api.BaseRequest

	FilterList              []CustomerBehaviorDetailGetListRequestFilterList `json:"filterList,omitempty"`
	BosId                   int64                                            `json:"bosId,omitempty"`
	TargetProductInstanceId int64                                            `json:"targetProductInstanceId,omitempty"`
	EndTime                 int64                                            `json:"endTime,omitempty"`
	GuiderId                string                                           `json:"guiderId,omitempty"`
	GuiderIdType            int64                                            `json:"guiderIdType,omitempty"`
	GuiderWid               string                                           `json:"guiderWid,omitempty"`
	PageNum                 int64                                            `json:"pageNum,omitempty"`
	PageSize                int64                                            `json:"pageSize,omitempty"`
	SessionId               string                                           `json:"sessionId,omitempty"`
	StartTime               int64                                            `json:"startTime,omitempty"`
	UkeyList                []int64                                          `json:"ukeyList,omitempty"`
}

type CustomerBehaviorDetailGetListResponse struct {
	BehaviorRecordList CustomerBehaviorDetailGetListResponseBehaviorRecordList `json:"behaviorRecordList,omitempty"`
}

func CreateCustomerBehaviorDetailGetListRequest() (request *CustomerBehaviorDetailGetListRequest) {
	request = &CustomerBehaviorDetailGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_cdp", "v2.0", "customer/behavior/detail/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerBehaviorDetailGetListResponse() (response *api.BaseResponse[CustomerBehaviorDetailGetListResponse]) {
	response = api.CreateResponse[CustomerBehaviorDetailGetListResponse](&CustomerBehaviorDetailGetListResponse{})
	return
}

type CustomerBehaviorDetailGetListRequestFilterList struct {
	En       string `json:"en,omitempty"`
	PageName string `json:"pageName,omitempty"`
}

type CustomerBehaviorDetailGetListResponseBehaviorRecordList struct {
	List     []CustomerBehaviorDetailGetListResponselist `json:"list,omitempty"`
	PageSize int64                                       `json:"pageSize,omitempty"`
	PageNum  int64                                       `json:"pageNum,omitempty"`
	Total    int64                                       `json:"total,omitempty"`
}

type CustomerBehaviorDetailGetListResponselist struct {
	GoodsTitleList        []CustomerBehaviorDetailGetListResponseGoodsTitleList `json:"goodsTitleList,omitempty"`
	SkuNameList           []CustomerBehaviorDetailGetListResponseSkuNameList    `json:"skuNameList,omitempty"`
	SkuNumList            []CustomerBehaviorDetailGetListResponseSkuNumList     `json:"skuNumList,omitempty"`
	PageTitle             string                                                `json:"pageTitle,omitempty"`
	GoodsId               string                                                `json:"goodsId,omitempty"`
	GoodsName             string                                                `json:"goodsName,omitempty"`
	ActivityTypeName      string                                                `json:"activityTypeName,omitempty"`
	ActivityName          string                                                `json:"activityName,omitempty"`
	Keyword               string                                                `json:"keyword,omitempty"`
	MatName               string                                                `json:"matName,omitempty"`
	SkuName               string                                                `json:"skuName,omitempty"`
	SkuPrice              int64                                                 `json:"skuPrice,omitempty"`
	CartSkuNum            int64                                                 `json:"cartSkuNum,omitempty"`
	GlobalStoreName       string                                                `json:"globalStoreName,omitempty"`
	RoundName             string                                                `json:"roundName,omitempty"`
	RoomName              string                                                `json:"roomName,omitempty"`
	PageState             string                                                `json:"pageState,omitempty"`
	PageStateName         string                                                `json:"pageStateName,omitempty"`
	PrizeName             string                                                `json:"prizeName,omitempty"`
	PrizeType             string                                                `json:"prizeType,omitempty"`
	AppId                 string                                                `json:"appId,omitempty"`
	AppName               string                                                `json:"appName,omitempty"`
	OrderId               string                                                `json:"orderId,omitempty"`
	GoodsFirstCategory    string                                                `json:"goodsFirstCategory,omitempty"`
	GoodsSecondCategory   string                                                `json:"goodsSecondCategory,omitempty"`
	ButtonStateName       string                                                `json:"buttonStateName,omitempty"`
	CouponName            string                                                `json:"couponName,omitempty"`
	PageSize              int64                                                 `json:"pageSize,omitempty"`
	PayAmt                int64                                                 `json:"payAmt,omitempty"`
	GoodsPrice            int64                                                 `json:"goodsPrice,omitempty"`
	VidName               string                                                `json:"vidName,omitempty"`
	CardName              string                                                `json:"cardName,omitempty"`
	BosId                 string                                                `json:"bosId,omitempty"`
	ChannelType           string                                                `json:"channelType,omitempty"`
	ChannelId             string                                                `json:"channelId,omitempty"`
	SchannelType          string                                                `json:"schannelType,omitempty"`
	SchannelId            string                                                `json:"schannelId,omitempty"`
	Ukey                  string                                                `json:"ukey,omitempty"`
	UkeyType              string                                                `json:"ukeyType,omitempty"`
	UpdateTime            int64                                                 `json:"updateTime,omitempty"`
	SessionId             string                                                `json:"sessionId,omitempty"`
	PageName              string                                                `json:"pageName,omitempty"`
	PageNameReal          string                                                `json:"pageNameReal,omitempty"`
	Vid                   string                                                `json:"vid,omitempty"`
	VidType               string                                                `json:"vidType,omitempty"`
	En                    string                                                `json:"en,omitempty"`
	EventType             string                                                `json:"eventType,omitempty"`
	EventName             string                                                `json:"eventName,omitempty"`
	ElementId             string                                                `json:"elementId,omitempty"`
	OriginInfo            string                                                `json:"originInfo,omitempty"`
	BindGuiderId          string                                                `json:"bindGuiderId,omitempty"`
	ShareGuiderId         string                                                `json:"shareGuiderId,omitempty"`
	DataJson              string                                                `json:"dataJson,omitempty"`
	GlobalStoreId         string                                                `json:"globalStoreId,omitempty"`
	SdpTicket             string                                                `json:"sdpTicket,omitempty"`
	SourceEnd             string                                                `json:"sourceEnd,omitempty"`
	AtyName               string                                                `json:"atyName,omitempty"`
	ActivityId            string                                                `json:"activityId,omitempty"`
	ActivityType          string                                                `json:"activityType,omitempty"`
	SkuId                 string                                                `json:"skuId,omitempty"`
	CouponCode            string                                                `json:"couponCode,omitempty"`
	Marketing             string                                                `json:"marketing,omitempty"`
	CouponDesc            string                                                `json:"couponDesc,omitempty"`
	CouponAmountType      string                                                `json:"couponAmountType,omitempty"`
	CouponUseAmount       string                                                `json:"couponUseAmount,omitempty"`
	GoodsMainImageUrlList string                                                `json:"goodsMainImageUrlList,omitempty"`
	UseStartTime          string                                                `json:"useStartTime,omitempty"`
	UseEndTime            string                                                `json:"useEndTime,omitempty"`
	CouponType            string                                                `json:"couponType,omitempty"`
	StoreName             string                                                `json:"storeName,omitempty"`
	GroupName             string                                                `json:"groupName,omitempty"`
	VisitDuration         string                                                `json:"visitDuration,omitempty"`
	UseThreshold          string                                                `json:"useThreshold,omitempty"`
	TuikeId               string                                                `json:"tuikeId,omitempty"`
	Scene                 string                                                `json:"scene,omitempty"`
	TjIndex               string                                                `json:"tjIndex,omitempty"`
	TjType                string                                                `json:"tjType,omitempty"`
}

type CustomerBehaviorDetailGetListResponseGoodsTitleList struct {
}

type CustomerBehaviorDetailGetListResponseSkuNameList struct {
}

type CustomerBehaviorDetailGetListResponseSkuNumList struct {
}
