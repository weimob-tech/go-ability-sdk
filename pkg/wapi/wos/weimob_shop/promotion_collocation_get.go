package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionCollocationGet
// @id 1925
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1925?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询搭配套装活动（活动基础规则+活动商品规则）)
func (client *Client) PromotionCollocationGet(request *PromotionCollocationGetRequest) (response *PromotionCollocationGetResponse, err error) {
	rpcResponse := CreatePromotionCollocationGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionCollocationGetRequest struct {
	*api.BaseRequest

	BasicInfo  PromotionCollocationGetRequestBasicInfo `json:"basicInfo,omitempty"`
	ActivityId int64                                   `json:"activityId,omitempty"`
}

type PromotionCollocationGetResponse struct {
	GroupInfoList              []PromotionCollocationGetResponseGroupInfoList `json:"groupInfoList,omitempty"`
	TotalLimitNum              int64                                          `json:"totalLimitNum,omitempty"`
	OverLimitCanOriginPriceBuy int64                                          `json:"overLimitCanOriginPriceBuy,omitempty"`
	EachLimitNum               int64                                          `json:"eachLimitNum,omitempty"`
	ActivityId                 int64                                          `json:"activityId,omitempty"`
	ActivityType               int64                                          `json:"activityType,omitempty"`
	PromotionName              string                                         `json:"promotionName,omitempty"`
	StartTime                  int64                                          `json:"startTime,omitempty"`
	EndTime                    int64                                          `json:"endTime,omitempty"`
	MarkingType                int64                                          `json:"markingType,omitempty"`
	TagRuleId                  int64                                          `json:"tagRuleId,omitempty"`
	SelectPeopleType           int64                                          `json:"selectPeopleType,omitempty"`
	SelectPeopleRuleId         int64                                          `json:"selectPeopleRuleId,omitempty"`
	ApplicationGoodsType       int64                                          `json:"applicationGoodsType,omitempty"`
	CheckType                  int64                                          `json:"checkType,omitempty"`
	SubIds                     []int64                                        `json:"subIds,omitempty"`
	SelectStoreType            int64                                          `json:"selectStoreType,omitempty"`
	SelectStoreIds             []int64                                        `json:"selectStoreIds,omitempty"`
	TradeScene                 []int64                                        `json:"tradeScene,omitempty"`
	AvailableDeduct            []int64                                        `json:"availableDeduct,omitempty"`
	AvailablePromotion         []int64                                        `json:"availablePromotion,omitempty"`
}

func CreatePromotionCollocationGetRequest() (request *PromotionCollocationGetRequest) {
	request = &PromotionCollocationGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "promotion/collocation/get", "apigw")
	request.Method = api.POST
	return
}

func CreatePromotionCollocationGetResponse() (response *api.BaseResponse[PromotionCollocationGetResponse]) {
	response = api.CreateResponse[PromotionCollocationGetResponse](&PromotionCollocationGetResponse{})
	return
}

type PromotionCollocationGetRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type PromotionCollocationGetResponseGroupInfoList struct {
	Sort        int64   `json:"sort,omitempty"`
	GroupId     int64   `json:"groupId,omitempty"`
	GroupName   string  `json:"groupName,omitempty"`
	IsMustBuy   int64   `json:"isMustBuy,omitempty"`
	MustBuyNum  int64   `json:"mustBuyNum,omitempty"`
	MaxBuyNum   int64   `json:"maxBuyNum,omitempty"`
	GoodsIdList []int64 `json:"goodsIdList,omitempty"`
}
