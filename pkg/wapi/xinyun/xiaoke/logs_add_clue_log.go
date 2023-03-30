package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// LogsAddClueLog
// @id 1725
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新建线索跟进记录)
func (client *Client) LogsAddClueLog(request *LogsAddClueLogRequest) (response *LogsAddClueLogResponse, err error) {
	rpcResponse := CreateLogsAddClueLogResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type LogsAddClueLogRequest struct {
	*api.BaseRequest

	Key     string `json:"key,omitempty"`
	Wid     int64  `json:"wid,omitempty"`
	Content string `json:"content,omitempty"`
}

type LogsAddClueLogResponse struct {
}

func CreateLogsAddClueLogRequest() (request *LogsAddClueLogRequest) {
	request = &LogsAddClueLogRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "logs/addClueLog", "api")
	request.Method = api.POST
	return
}

func CreateLogsAddClueLogResponse() (response *api.BaseResponse[LogsAddClueLogResponse]) {
	response = api.CreateResponse[LogsAddClueLogResponse](&LogsAddClueLogResponse{})
	return
}
