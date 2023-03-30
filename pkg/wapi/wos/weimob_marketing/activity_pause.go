package weimob_marketing

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ActivityPause
// @id 1764
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1764?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=暂停活动)
func (client *Client) ActivityPause(request *ActivityPauseRequest) (response *ActivityPauseResponse, err error) {
	rpcResponse := CreateActivityPauseResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ActivityPauseRequest struct {
	*api.BaseRequest

	BosId             int64 `json:"bosId,omitempty"`
	ProductId         int64 `json:"productId,omitempty"`
	ProductInstanceId int64 `json:"productInstanceId,omitempty"`
	Id                int64 `json:"id,omitempty"`
	Vid               int64 `json:"vid,omitempty"`
}

type ActivityPauseResponse struct {
}

func CreateActivityPauseRequest() (request *ActivityPauseRequest) {
	request = &ActivityPauseRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_marketing", "v2.0", "activity/pause", "apigw")
	request.Method = api.POST
	return
}

func CreateActivityPauseResponse() (response *api.BaseResponse[ActivityPauseResponse]) {
	response = api.CreateResponse[ActivityPauseResponse](&ActivityPauseResponse{})
	return
}
