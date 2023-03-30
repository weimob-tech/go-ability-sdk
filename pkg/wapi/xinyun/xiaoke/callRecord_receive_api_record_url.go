package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CallRecordReceiveApiRecordUrl
// @id 3742
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=API线路录音上传接口)
func (client *Client) CallRecordReceiveApiRecordUrl(request *CallRecordReceiveApiRecordUrlRequest) (response *CallRecordReceiveApiRecordUrlResponse, err error) {
	rpcResponse := CreateCallRecordReceiveApiRecordUrlResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CallRecordReceiveApiRecordUrlRequest struct {
	*api.BaseRequest

	CallId    string `json:"callId,omitempty"`
	RecordUrl string `json:"recordUrl,omitempty"`
}

type CallRecordReceiveApiRecordUrlResponse struct {
}

func CreateCallRecordReceiveApiRecordUrlRequest() (request *CallRecordReceiveApiRecordUrlRequest) {
	request = &CallRecordReceiveApiRecordUrlRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "callRecord/receiveApiRecordUrl", "api")
	request.Method = api.POST
	return
}

func CreateCallRecordReceiveApiRecordUrlResponse() (response *api.BaseResponse[CallRecordReceiveApiRecordUrlResponse]) {
	response = api.CreateResponse[CallRecordReceiveApiRecordUrlResponse](&CallRecordReceiveApiRecordUrlResponse{})
	return
}
