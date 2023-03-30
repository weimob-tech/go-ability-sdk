package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionQueryNynjDetail
// @id 1842
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询N元N件活动详情)
func (client *Client) PromotionQueryNynjDetail(request *PromotionQueryNynjDetailRequest) (response *PromotionQueryNynjDetailResponse, err error) {
	rpcResponse := CreatePromotionQueryNynjDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionQueryNynjDetailRequest struct {
	*api.BaseRequest

	ActivityId int `json:"activityId,omitempty"`
}

type PromotionQueryNynjDetailResponse struct {
}

func CreatePromotionQueryNynjDetailRequest() (request *PromotionQueryNynjDetailRequest) {
	request = &PromotionQueryNynjDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "promotion/queryNynjDetail", "api")
	request.Method = api.POST
	return
}

func CreatePromotionQueryNynjDetailResponse() (response *api.BaseResponse[PromotionQueryNynjDetailResponse]) {
	response = api.CreateResponse[PromotionQueryNynjDetailResponse](&PromotionQueryNynjDetailResponse{})
	return
}
