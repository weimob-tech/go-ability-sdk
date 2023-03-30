package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CallRecordStatus
// @id 1846
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=通话状态回调)
func (client *Client) CallRecordStatus(request *CallRecordStatusRequest) (response *CallRecordStatusResponse, err error) {
	rpcResponse := CreateCallRecordStatusResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CallRecordStatusRequest struct {
	*api.BaseRequest

	Platform string `json:"platform,omitempty"`
	Status   string `json:"status,omitempty"`
	CallId   string `json:"callId,omitempty"`
}

type CallRecordStatusResponse struct {
}

func CreateCallRecordStatusRequest() (request *CallRecordStatusRequest) {
	request = &CallRecordStatusRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "callRecord/status", "api")
	request.Method = api.POST
	return
}

func CreateCallRecordStatusResponse() (response *api.BaseResponse[CallRecordStatusResponse]) {
	response = api.CreateResponse[CallRecordStatusResponse](&CallRecordStatusResponse{})
	return
}
