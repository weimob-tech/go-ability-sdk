package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromotionCreatePromotionTag
// @id 1607
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=创建活动短标题)
func (client *Client) PromotionCreatePromotionTag(request *PromotionCreatePromotionTagRequest) (response *PromotionCreatePromotionTagResponse, err error) {
	rpcResponse := CreatePromotionCreatePromotionTagResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromotionCreatePromotionTagRequest struct {
	*api.BaseRequest

	ActivityType int    `json:"activityType,omitempty"`
	TagName      string `json:"TagName,omitempty"`
}

type PromotionCreatePromotionTagResponse struct {
}

func CreatePromotionCreatePromotionTagRequest() (request *PromotionCreatePromotionTagRequest) {
	request = &PromotionCreatePromotionTagRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "promotion/createPromotionTag", "api")
	request.Method = api.POST
	return
}

func CreatePromotionCreatePromotionTagResponse() (response *api.BaseResponse[PromotionCreatePromotionTagResponse]) {
	response = api.CreateResponse[PromotionCreatePromotionTagResponse](&PromotionCreatePromotionTagResponse{})
	return
}
