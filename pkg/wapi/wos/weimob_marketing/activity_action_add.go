package weimob_marketing

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ActivityActionAdd
// @id 1911
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1911?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=上报活动行为)
func (client *Client) ActivityActionAdd(request *ActivityActionAddRequest) (response *ActivityActionAddResponse, err error) {
	rpcResponse := CreateActivityActionAddResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ActivityActionAddRequest struct {
	*api.BaseRequest

	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	BosId             int64  `json:"bosId,omitempty"`
	Vid               int64  `json:"vid,omitempty"`
	ActivityId        int64  `json:"activityId,omitempty"`
	Wid               int64  `json:"wid,omitempty"`
	Action            string `json:"action,omitempty"`
	ExtraInfo         string `json:"extraInfo,omitempty"`
	ProductId         int64  `json:"productId,omitempty"`
	OccurTime         int64  `json:"occurTime,omitempty"`
}

type ActivityActionAddResponse struct {
}

func CreateActivityActionAddRequest() (request *ActivityActionAddRequest) {
	request = &ActivityActionAddRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_marketing", "v2.0", "activity/action/add", "apigw")
	request.Method = api.POST
	return
}

func CreateActivityActionAddResponse() (response *api.BaseResponse[ActivityActionAddResponse]) {
	response = api.CreateResponse[ActivityActionAddResponse](&ActivityActionAddResponse{})
	return
}
