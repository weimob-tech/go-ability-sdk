package weimob_marketing

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ActivityExecute
// @id 1908
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1908?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=操作活动)
func (client *Client) ActivityExecute(request *ActivityExecuteRequest) (response *ActivityExecuteResponse, err error) {
	rpcResponse := CreateActivityExecuteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ActivityExecuteRequest struct {
	*api.BaseRequest

	Action            string `json:"action,omitempty"`
	Id                int64  `json:"id,omitempty"`
	BosId             int64  `json:"bosId,omitempty"`
	ProductId         int64  `json:"productId,omitempty"`
	ProductInstanceId int64  `json:"productInstanceId,omitempty"`
	Vid               int64  `json:"vid,omitempty"`
}

type ActivityExecuteResponse struct {
}

func CreateActivityExecuteRequest() (request *ActivityExecuteRequest) {
	request = &ActivityExecuteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_marketing", "v2.0", "activity/execute", "apigw")
	request.Method = api.POST
	return
}

func CreateActivityExecuteResponse() (response *api.BaseResponse[ActivityExecuteResponse]) {
	response = api.CreateResponse[ActivityExecuteResponse](&ActivityExecuteResponse{})
	return
}
