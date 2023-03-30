package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MarketactivityGetList
// @id 1922
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1922?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询促销活动列表)
func (client *Client) MarketactivityGetList(request *MarketactivityGetListRequest) (response *MarketactivityGetListResponse, err error) {
	rpcResponse := CreateMarketactivityGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MarketactivityGetListRequest struct {
	*api.BaseRequest

	QueryParameter MarketactivityGetListRequestQueryParameter `json:"queryParameter,omitempty"`
	BasicInfo      MarketactivityGetListRequestBasicInfo      `json:"basicInfo,omitempty"`
	PageNum        int64                                      `json:"pageNum,omitempty"`
	PageSize       int64                                      `json:"pageSize,omitempty"`
}

type MarketactivityGetListResponse struct {
	PageList   []MarketactivityGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                   `json:"pageNum,omitempty"`
	PageSize   int64                                   `json:"pageSize,omitempty"`
	TotalCount int64                                   `json:"totalCount,omitempty"`
}

func CreateMarketactivityGetListRequest() (request *MarketactivityGetListRequest) {
	request = &MarketactivityGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "marketactivity/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateMarketactivityGetListResponse() (response *api.BaseResponse[MarketactivityGetListResponse]) {
	response = api.CreateResponse[MarketactivityGetListResponse](&MarketactivityGetListResponse{})
	return
}

type MarketactivityGetListRequestQueryParameter struct {
	ActivityType    int64   `json:"activityType,omitempty"`
	ActivitySubType int64   `json:"activitySubType,omitempty"`
	Status          []int64 `json:"status,omitempty"`
	StartTime       int64   `json:"startTime,omitempty"`
	EndTime         int64   `json:"endTime,omitempty"`
	PromotionName   string  `json:"promotionName,omitempty"`
}

type MarketactivityGetListRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type MarketactivityGetListResponsePageList struct {
	CycleTimeList        []MarketactivityGetListResponseCycleTimeList `json:"cycleTimeList,omitempty"`
	ActivityId           int64                                        `json:"activityId,omitempty"`
	BelongVid            int64                                        `json:"belongVid,omitempty"`
	BelongVidType        int64                                        `json:"belongVidType,omitempty"`
	VidName              string                                       `json:"vidName,omitempty"`
	PromotionName        string                                       `json:"promotionName,omitempty"`
	TimeType             int64                                        `json:"timeType,omitempty"`
	StartTime            int64                                        `json:"startTime,omitempty"`
	ReceivableMoney      int64                                        `json:"receivableMoney,omitempty"`
	PayNum               int64                                        `json:"payNum,omitempty"`
	OrderAmount          int64                                        `json:"orderAmount,omitempty"`
	OrderAveragePrice    int64                                        `json:"orderAveragePrice,omitempty"`
	PageView             int64                                        `json:"pageView,omitempty"`
	UniqueView           int64                                        `json:"uniqueView,omitempty"`
	PageClick            int64                                        `json:"pageClick,omitempty"`
	UniqueClick          int64                                        `json:"uniqueClick,omitempty"`
	Day                  string                                       `json:"day,omitempty"`
	EndTime              int64                                        `json:"endTime,omitempty"`
	PromotionStatus      int64                                        `json:"promotionStatus,omitempty"`
	ActivityType         int64                                        `json:"activityType,omitempty"`
	ActivitySubType      int64                                        `json:"activitySubType,omitempty"`
	PromotionImage       string                                       `json:"promotionImage,omitempty"`
	IsPromotionLongTime  int64                                        `json:"isPromotionLongTime,omitempty"`
	ShareCustomerNumbers int64                                        `json:"shareCustomerNumbers,omitempty"`
	CreateTime           string                                       `json:"createTime,omitempty"`
}

type MarketactivityGetListResponseCycleTimeList struct {
	RepeatType          int64  `json:"repeatType,omitempty"`
	RepeatDay           string `json:"repeatDay,omitempty"`
	RepeatStartInterval string `json:"repeatStartInterval,omitempty"`
	RepeatEndInterval   string `json:"repeatEndInterval,omitempty"`
}
