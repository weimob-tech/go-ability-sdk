package weimob_marketing

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ActivityBehaviorActionAdd
// @id 1761
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1761?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=上报活动转化行为)
func (client *Client) ActivityBehaviorActionAdd(request *ActivityBehaviorActionAddRequest) (response *ActivityBehaviorActionAddResponse, err error) {
	rpcResponse := CreateActivityBehaviorActionAddResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ActivityBehaviorActionAddRequest struct {
	*api.BaseRequest

	Action     string `json:"action,omitempty"`
	ActivityId int64  `json:"activityId,omitempty"`
	BosId      int64  `json:"bosId,omitempty"`
	ExtraInfo  string `json:"extraInfo,omitempty"`
	Pid        int64  `json:"pid,omitempty"`
	Piid       int64  `json:"piid,omitempty"`
	Vid        int64  `json:"vid,omitempty"`
	Wid        int64  `json:"wid,omitempty"`
}

type ActivityBehaviorActionAddResponse struct {
}

func CreateActivityBehaviorActionAddRequest() (request *ActivityBehaviorActionAddRequest) {
	request = &ActivityBehaviorActionAddRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_marketing", "v2.0", "activity/behavior/action/add", "apigw")
	request.Method = api.POST
	return
}

func CreateActivityBehaviorActionAddResponse() (response *api.BaseResponse[ActivityBehaviorActionAddResponse]) {
	response = api.CreateResponse[ActivityBehaviorActionAddResponse](&ActivityBehaviorActionAddResponse{})
	return
}
