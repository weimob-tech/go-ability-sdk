package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TradeCartUpdate
// @id 2207
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2207?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=修改购物车)
func (client *Client) TradeCartUpdate(request *TradeCartUpdateRequest) (response *TradeCartUpdateResponse, err error) {
	rpcResponse := CreateTradeCartUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TradeCartUpdateRequest struct {
	*api.BaseRequest

	GoodsList   []TradeCartUpdateRequestGoodsList `json:"goodsList,omitempty"`
	ExtendInfo  TradeCartUpdateRequestExtendInfo  `json:"extendInfo,omitempty"`
	BasicInfo   TradeCartUpdateRequestBasicInfo   `json:"basicInfo,omitempty"`
	OperateType int64                             `json:"operateType,omitempty"`
	BizWid      int64                             `json:"bizWid,omitempty"`
}

type TradeCartUpdateResponse struct {
	ValidBizResult TradeCartUpdateResponseValidBizResult `json:"validBizResult,omitempty"`
	Status         bool                                  `json:"status,omitempty"`
}

func CreateTradeCartUpdateRequest() (request *TradeCartUpdateRequest) {
	request = &TradeCartUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "trade/cart/update", "apigw")
	request.Method = api.POST
	return
}

func CreateTradeCartUpdateResponse() (response *api.BaseResponse[TradeCartUpdateResponse]) {
	response = api.CreateResponse[TradeCartUpdateResponse](&TradeCartUpdateResponse{})
	return
}

type TradeCartUpdateRequestGoodsList struct {
	Vid            int64  `json:"vid,omitempty"`
	ItemId         string `json:"itemId,omitempty"`
	TradeGoodsType int64  `json:"tradeGoodsType,omitempty"`
	GoodsId        int64  `json:"goodsId,omitempty"`
	Num            int64  `json:"num,omitempty"`
	IsSelected     int64  `json:"isSelected,omitempty"`
	SkuId          int64  `json:"skuId,omitempty"`
}

type TradeCartUpdateRequestExtendInfo struct {
	Refer          string `json:"refer,omitempty"`
	TemplateId     int64  `json:"templateId,omitempty"`
	BusinessSource int64  `json:"businessSource,omitempty"`
	Source         int64  `json:"source,omitempty"`
}

type TradeCartUpdateRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type TradeCartUpdateResponseValidBizResult struct {
	ValidBizInfo TradeCartUpdateResponseValidBizInfo `json:"validBizInfo,omitempty"`
	ValidBizType string                              `json:"validBizType,omitempty"`
	ValidSuccess bool                                `json:"validSuccess,omitempty"`
}

type TradeCartUpdateResponseValidBizInfo struct {
}
