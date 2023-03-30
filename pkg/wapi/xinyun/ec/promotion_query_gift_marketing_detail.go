package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionQueryGiftMarketingDetail
// @id 1791
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询满赠活动详情)
func (client *Client) PromotionQueryGiftMarketingDetail(request *PromotionQueryGiftMarketingDetailRequest) (response *PromotionQueryGiftMarketingDetailResponse, err error) {
	rpcResponse := CreatePromotionQueryGiftMarketingDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionQueryGiftMarketingDetailRequest struct {
	*api.BaseRequest

	ActivityId int64 `json:"activityId,omitempty"`
}

type PromotionQueryGiftMarketingDetailResponse struct {
}

func CreatePromotionQueryGiftMarketingDetailRequest() (request *PromotionQueryGiftMarketingDetailRequest) {
	request = &PromotionQueryGiftMarketingDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "promotion/queryGiftMarketingDetail", "api")
	request.Method = api.POST
	return
}

func CreatePromotionQueryGiftMarketingDetailResponse() (response *api.BaseResponse[PromotionQueryGiftMarketingDetailResponse]) {
	response = api.CreateResponse[PromotionQueryGiftMarketingDetailResponse](&PromotionQueryGiftMarketingDetailResponse{})
	return
}
