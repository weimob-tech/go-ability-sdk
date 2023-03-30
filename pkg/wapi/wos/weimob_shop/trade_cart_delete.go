package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TradeCartDelete
// @id 2206
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2206?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除购物车中商品)
func (client *Client) TradeCartDelete(request *TradeCartDeleteRequest) (response *TradeCartDeleteResponse, err error) {
	rpcResponse := CreateTradeCartDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TradeCartDeleteRequest struct {
	*api.BaseRequest

	GoodsList  []TradeCartDeleteRequestGoodsList `json:"goodsList,omitempty"`
	ExtendInfo TradeCartDeleteRequestExtendInfo  `json:"extendInfo,omitempty"`
	BasicInfo  TradeCartDeleteRequestBasicInfo   `json:"basicInfo,omitempty"`
	BizWid     int64                             `json:"bizWid,omitempty"`
}

type TradeCartDeleteResponse struct {
	ValidBizResult TradeCartDeleteResponseValidBizResult `json:"validBizResult,omitempty"`
	Status         bool                                  `json:"status,omitempty"`
}

func CreateTradeCartDeleteRequest() (request *TradeCartDeleteRequest) {
	request = &TradeCartDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "trade/cart/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateTradeCartDeleteResponse() (response *api.BaseResponse[TradeCartDeleteResponse]) {
	response = api.CreateResponse[TradeCartDeleteResponse](&TradeCartDeleteResponse{})
	return
}

type TradeCartDeleteRequestGoodsList struct {
	Vid    int64  `json:"vid,omitempty"`
	ItemId string `json:"itemId,omitempty"`
}

type TradeCartDeleteRequestExtendInfo struct {
	Refer          string `json:"refer,omitempty"`
	TemplateId     int64  `json:"templateId,omitempty"`
	BusinessSource int64  `json:"businessSource,omitempty"`
	Source         int64  `json:"source,omitempty"`
}

type TradeCartDeleteRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type TradeCartDeleteResponseValidBizResult struct {
	ValidBizInfo TradeCartDeleteResponseValidBizInfo `json:"validBizInfo,omitempty"`
	ValidSuccess bool                                `json:"validSuccess,omitempty"`
	ValidBizType string                              `json:"validBizType,omitempty"`
}

type TradeCartDeleteResponseValidBizInfo struct {
}
