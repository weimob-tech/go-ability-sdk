package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionQueryRedemptionDetail
// @id 1792
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询加价购活动详情)
func (client *Client) PromotionQueryRedemptionDetail(request *PromotionQueryRedemptionDetailRequest) (response *PromotionQueryRedemptionDetailResponse, err error) {
	rpcResponse := CreatePromotionQueryRedemptionDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionQueryRedemptionDetailRequest struct {
	*api.BaseRequest

	ActivityId int64 `json:"activityId,omitempty"`
}

type PromotionQueryRedemptionDetailResponse struct {
}

func CreatePromotionQueryRedemptionDetailRequest() (request *PromotionQueryRedemptionDetailRequest) {
	request = &PromotionQueryRedemptionDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "promotion/queryRedemptionDetail", "api")
	request.Method = api.POST
	return
}

func CreatePromotionQueryRedemptionDetailResponse() (response *api.BaseResponse[PromotionQueryRedemptionDetailResponse]) {
	response = api.CreateResponse[PromotionQueryRedemptionDetailResponse](&PromotionQueryRedemptionDetailResponse{})
	return
}
