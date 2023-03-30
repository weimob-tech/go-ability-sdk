package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// LogsAddContactLog
// @id 1855
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新建联系人跟进日志)
func (client *Client) LogsAddContactLog(request *LogsAddContactLogRequest) (response *LogsAddContactLogResponse, err error) {
	rpcResponse := CreateLogsAddContactLogResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type LogsAddContactLogRequest struct {
	*api.BaseRequest

	Key     string `json:"key,omitempty"`
	Wid     int64  `json:"wid,omitempty"`
	Content string `json:"content,omitempty"`
}

type LogsAddContactLogResponse struct {
}

func CreateLogsAddContactLogRequest() (request *LogsAddContactLogRequest) {
	request = &LogsAddContactLogRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "logs/addContactLog", "api")
	request.Method = api.POST
	return
}

func CreateLogsAddContactLogResponse() (response *api.BaseResponse[LogsAddContactLogResponse]) {
	response = api.CreateResponse[LogsAddContactLogResponse](&LogsAddContactLogResponse{})
	return
}
