package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CallRecordDetail
// @id 1847
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=通话详情回调)
func (client *Client) CallRecordDetail(request *CallRecordDetailRequest) (response *CallRecordDetailResponse, err error) {
	rpcResponse := CreateCallRecordDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CallRecordDetailRequest struct {
	*api.BaseRequest

	Platform    string `json:"platform,omitempty"`
	CallId      string `json:"callId,omitempty"`
	Caller      string `json:"caller,omitempty"`
	Callee      string `json:"callee,omitempty"`
	CallFlag    int    `json:"callFlag,omitempty"`
	Duration    int    `json:"duration,omitempty"`
	RecordUrl   string `json:"recordUrl,omitempty"`
	CallEndTime int    `json:"callEndTime,omitempty"`
}

type CallRecordDetailResponse struct {
}

func CreateCallRecordDetailRequest() (request *CallRecordDetailRequest) {
	request = &CallRecordDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "callRecord/detail", "api")
	request.Method = api.POST
	return
}

func CreateCallRecordDetailResponse() (response *api.BaseResponse[CallRecordDetailResponse]) {
	response = api.CreateResponse[CallRecordDetailResponse](&CallRecordDetailResponse{})
	return
}
