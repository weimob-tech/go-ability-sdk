package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionQueryDiscountDetail
// @id 1787
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询限时折扣活动详情)
func (client *Client) PromotionQueryDiscountDetail(request *PromotionQueryDiscountDetailRequest) (response *PromotionQueryDiscountDetailResponse, err error) {
	rpcResponse := CreatePromotionQueryDiscountDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionQueryDiscountDetailRequest struct {
	*api.BaseRequest

	DiscountId int64 `json:"discountId,omitempty"`
}

type PromotionQueryDiscountDetailResponse struct {
}

func CreatePromotionQueryDiscountDetailRequest() (request *PromotionQueryDiscountDetailRequest) {
	request = &PromotionQueryDiscountDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "promotion/queryDiscountDetail", "api")
	request.Method = api.POST
	return
}

func CreatePromotionQueryDiscountDetailResponse() (response *api.BaseResponse[PromotionQueryDiscountDetailResponse]) {
	response = api.CreateResponse[PromotionQueryDiscountDetailResponse](&PromotionQueryDiscountDetailResponse{})
	return
}
