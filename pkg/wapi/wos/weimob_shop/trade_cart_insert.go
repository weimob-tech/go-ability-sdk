package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TradeCartInsert
// @id 1689
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1689?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=加入购物车)
func (client *Client) TradeCartInsert(request *TradeCartInsertRequest) (response *TradeCartInsertResponse, err error) {
	rpcResponse := CreateTradeCartInsertResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TradeCartInsertRequest struct {
	*api.BaseRequest

	GoodsList  []TradeCartInsertRequestGoodsList `json:"goodsList,omitempty"`
	ExtendInfo TradeCartInsertRequestExtendInfo  `json:"extendInfo,omitempty"`
	BasicInfo  TradeCartInsertRequestBasicInfo   `json:"basicInfo,omitempty"`
	BizWid     int64                             `json:"bizWid,omitempty"`
}

type TradeCartInsertResponse struct {
	ValidBizResult TradeCartInsertResponseValidBizResult `json:"validBizResult,omitempty"`
	GoodsNum       int64                                 `json:"goodsNum,omitempty"`
	Status         bool                                  `json:"status,omitempty"`
}

func CreateTradeCartInsertRequest() (request *TradeCartInsertRequest) {
	request = &TradeCartInsertRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "trade/cart/insert", "apigw")
	request.Method = api.POST
	return
}

func CreateTradeCartInsertResponse() (response *api.BaseResponse[TradeCartInsertResponse]) {
	response = api.CreateResponse[TradeCartInsertResponse](&TradeCartInsertResponse{})
	return
}

type TradeCartInsertRequestGoodsList struct {
	LabelList      []TradeCartInsertRequestLabelList  `json:"labelList,omitempty"`
	GoodsAbility   TradeCartInsertRequestGoodsAbility `json:"goodsAbility,omitempty"`
	Activities     []TradeCartInsertRequestActivities `json:"activities,omitempty"`
	Vid            int64                              `json:"vid,omitempty"`
	Num            int64                              `json:"num,omitempty"`
	GoodsId        int64                              `json:"goodsId,omitempty"`
	TradeGoodsType int64                              `json:"tradeGoodsType,omitempty"`
	SkuId          int64                              `json:"skuId,omitempty"`
	IsSelected     int64                              `json:"isSelected,omitempty"`
	CloudCustom    string                             `json:"cloudCustom,omitempty"`
}

type TradeCartInsertRequestLabelList struct {
	Attachment TradeCartInsertRequestAttachment `json:"attachment,omitempty"`
	LabelType  string                           `json:"labelType,omitempty"`
	AttachId   string                           `json:"attachId,omitempty"`
	CreateTime int64                            `json:"createTime,omitempty"`
}

type TradeCartInsertRequestAttachment struct {
}

type TradeCartInsertRequestGoodsAbility struct {
	CycleGoods TradeCartInsertRequestCycleGoods `json:"cycleGoods,omitempty"`
}

type TradeCartInsertRequestCycleGoods struct {
	BizId                    int64   `json:"bizId,omitempty"`
	CycleUnit                int64   `json:"cycleUnit,omitempty"`
	SelectCycleValues        []int64 `json:"selectCycleValues,omitempty"`
	SelectFulfillTimeOptions string  `json:"selectFulfillTimeOptions,omitempty"`
}

type TradeCartInsertRequestActivities struct {
	ActivityId      int64 `json:"activityId,omitempty"`
	ActivityType    int64 `json:"activityType,omitempty"`
	SubActivityType int64 `json:"subActivityType,omitempty"`
	UseStatus       int64 `json:"useStatus,omitempty"`
	LevelId         int64 `json:"levelId,omitempty"`
}

type TradeCartInsertRequestExtendInfo struct {
	Refer          string `json:"refer,omitempty"`
	TemplateId     int64  `json:"templateId,omitempty"`
	Source         int64  `json:"source,omitempty"`
	BusinessSource int64  `json:"businessSource,omitempty"`
}

type TradeCartInsertRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type TradeCartInsertResponseValidBizResult struct {
	ValidBizInfo TradeCartInsertResponseValidBizInfo `json:"validBizInfo,omitempty"`
	ValidBizType int64                               `json:"validBizType,omitempty"`
	ValidSuccess bool                                `json:"validSuccess,omitempty"`
}

type TradeCartInsertResponseValidBizInfo struct {
}
