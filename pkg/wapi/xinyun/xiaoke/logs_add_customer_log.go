package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// LogsAddCustomerLog
// @id 1726
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新建客户跟进记录)
func (client *Client) LogsAddCustomerLog(request *LogsAddCustomerLogRequest) (response *LogsAddCustomerLogResponse, err error) {
	rpcResponse := CreateLogsAddCustomerLogResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type LogsAddCustomerLogRequest struct {
	*api.BaseRequest

	Key     string `json:"key,omitempty"`
	Wid     int64  `json:"wid,omitempty"`
	Content string `json:"content,omitempty"`
}

type LogsAddCustomerLogResponse struct {
}

func CreateLogsAddCustomerLogRequest() (request *LogsAddCustomerLogRequest) {
	request = &LogsAddCustomerLogRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "logs/addCustomerLog", "api")
	request.Method = api.POST
	return
}

func CreateLogsAddCustomerLogResponse() (response *api.BaseResponse[LogsAddCustomerLogResponse]) {
	response = api.CreateResponse[LogsAddCustomerLogResponse](&LogsAddCustomerLogResponse{})
	return
}
