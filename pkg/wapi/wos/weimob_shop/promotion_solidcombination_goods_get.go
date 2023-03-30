package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionSolidcombinationGoodsGet
// @id 1920
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1920?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询固定套装商品详情)
func (client *Client) PromotionSolidcombinationGoodsGet(request *PromotionSolidcombinationGoodsGetRequest) (response *PromotionSolidcombinationGoodsGetResponse, err error) {
	rpcResponse := CreatePromotionSolidcombinationGoodsGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionSolidcombinationGoodsGetRequest struct {
	*api.BaseRequest

	BasicInfo    PromotionSolidcombinationGoodsGetRequestBasicInfo `json:"basicInfo,omitempty"`
	ActivityType int64                                             `json:"activityType,omitempty"`
	GroupId      int64                                             `json:"groupId,omitempty"`
	GoodsIdList  []int64                                           `json:"goodsIdList,omitempty"`
	ActivityId   int64                                             `json:"activityId,omitempty"`
}

type PromotionSolidcombinationGoodsGetResponse struct {
	PageList []PromotionSolidcombinationGoodsGetResponsePageList `json:"pageList,omitempty"`
}

func CreatePromotionSolidcombinationGoodsGetRequest() (request *PromotionSolidcombinationGoodsGetRequest) {
	request = &PromotionSolidcombinationGoodsGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "promotion/solidcombination/goods/get", "apigw")
	request.Method = api.POST
	return
}

func CreatePromotionSolidcombinationGoodsGetResponse() (response *api.BaseResponse[PromotionSolidcombinationGoodsGetResponse]) {
	response = api.CreateResponse[PromotionSolidcombinationGoodsGetResponse](&PromotionSolidcombinationGoodsGetResponse{})
	return
}

type PromotionSolidcombinationGoodsGetRequestBasicInfo struct {
	BosId int64 `json:"bosId,omitempty"`
	Vid   int64 `json:"vid,omitempty"`
}

type PromotionSolidcombinationGoodsGetResponsePageList struct {
	SkuList         []PromotionSolidcombinationGoodsGetResponseSkuList `json:"skuList,omitempty"`
	GoodsId         int64                                              `json:"goodsId,omitempty"`
	Title           string                                             `json:"title,omitempty"`
	EachLimitNum    int64                                              `json:"eachLimitNum,omitempty"`
	ShowGoodsDetail int64                                              `json:"showGoodsDetail,omitempty"`
	GroupId         int64                                              `json:"groupId,omitempty"`
	BuyNum          int64                                              `json:"buyNum,omitempty"`
}

type PromotionSolidcombinationGoodsGetResponseSkuList struct {
	SkuId         int64 `json:"skuId,omitempty"`
	CanSaleBuyNum int64 `json:"canSaleBuyNum,omitempty"`
	ActivityPrice int64 `json:"activityPrice,omitempty"`
}
