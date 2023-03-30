package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionCollocationGoodsGet
// @id 1926
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1926?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询搭配套装商品详情)
func (client *Client) PromotionCollocationGoodsGet(request *PromotionCollocationGoodsGetRequest) (response *PromotionCollocationGoodsGetResponse, err error) {
	rpcResponse := CreatePromotionCollocationGoodsGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionCollocationGoodsGetRequest struct {
	*api.BaseRequest

	BasicInfo    PromotionCollocationGoodsGetRequestBasicInfo `json:"basicInfo,omitempty"`
	ActivityType int64                                        `json:"activityType,omitempty"`
	GroupId      int64                                        `json:"groupId,omitempty"`
	GoodsIdList  []int64                                      `json:"goodsIdList,omitempty"`
	ActivityId   int64                                        `json:"activityId,omitempty"`
}

type PromotionCollocationGoodsGetResponse struct {
	PageList []PromotionCollocationGoodsGetResponsePageList `json:"pageList,omitempty"`
}

func CreatePromotionCollocationGoodsGetRequest() (request *PromotionCollocationGoodsGetRequest) {
	request = &PromotionCollocationGoodsGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "promotion/collocation/goods/get", "apigw")
	request.Method = api.POST
	return
}

func CreatePromotionCollocationGoodsGetResponse() (response *api.BaseResponse[PromotionCollocationGoodsGetResponse]) {
	response = api.CreateResponse[PromotionCollocationGoodsGetResponse](&PromotionCollocationGoodsGetResponse{})
	return
}

type PromotionCollocationGoodsGetRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type PromotionCollocationGoodsGetResponsePageList struct {
	SkuList         []PromotionCollocationGoodsGetResponseSkuList `json:"skuList,omitempty"`
	GoodsId         int64                                         `json:"goodsId,omitempty"`
	Title           string                                        `json:"title,omitempty"`
	EachLimitNum    int64                                         `json:"eachLimitNum,omitempty"`
	ShowGoodsDetail int64                                         `json:"showGoodsDetail,omitempty"`
	GroupId         int64                                         `json:"groupId,omitempty"`
}

type PromotionCollocationGoodsGetResponseSkuList struct {
	SkuId         int64 `json:"skuId,omitempty"`
	CanSaleBuyNum int64 `json:"canSaleBuyNum,omitempty"`
	ActivityPrice int64 `json:"activityPrice,omitempty"`
}
