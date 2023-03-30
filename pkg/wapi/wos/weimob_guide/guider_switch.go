package weimob_guide

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuiderSwitch
// @id 1758
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1758?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=导购切换门店)
func (client *Client) GuiderSwitch(request *GuiderSwitchRequest) (response *GuiderSwitchResponse, err error) {
	rpcResponse := CreateGuiderSwitchResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuiderSwitchRequest struct {
	*api.BaseRequest

	GuiderWid     int64 `json:"guiderWid,omitempty"`
	TargetVid     int64 `json:"targetVid,omitempty"`
	KeepCustomers int64 `json:"keepCustomers,omitempty"`
}

type GuiderSwitchResponse struct {
	Result bool   `json:"result,omitempty"`
	TaskId string `json:"taskId,omitempty"`
}

func CreateGuiderSwitchRequest() (request *GuiderSwitchRequest) {
	request = &GuiderSwitchRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_guide", "v2.0", "guider/switch", "apigw")
	request.Method = api.POST
	return
}

func CreateGuiderSwitchResponse() (response *api.BaseResponse[GuiderSwitchResponse]) {
	response = api.CreateResponse[GuiderSwitchResponse](&GuiderSwitchResponse{})
	return
}
