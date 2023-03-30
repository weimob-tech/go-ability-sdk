package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionQueryXjxzDetail
// @id 1793
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询第X件X折活动详情)
func (client *Client) PromotionQueryXjxzDetail(request *PromotionQueryXjxzDetailRequest) (response *PromotionQueryXjxzDetailResponse, err error) {
	rpcResponse := CreatePromotionQueryXjxzDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionQueryXjxzDetailRequest struct {
	*api.BaseRequest

	ActivityId int64 `json:"activityId,omitempty"`
}

type PromotionQueryXjxzDetailResponse struct {
}

func CreatePromotionQueryXjxzDetailRequest() (request *PromotionQueryXjxzDetailRequest) {
	request = &PromotionQueryXjxzDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "promotion/queryXjxzDetail", "api")
	request.Method = api.POST
	return
}

func CreatePromotionQueryXjxzDetailResponse() (response *api.BaseResponse[PromotionQueryXjxzDetailResponse]) {
	response = api.CreateResponse[PromotionQueryXjxzDetailResponse](&PromotionQueryXjxzDetailResponse{})
	return
}
