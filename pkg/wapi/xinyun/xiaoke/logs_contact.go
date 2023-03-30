package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// LogsContact
// @id 1854
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询联系人跟进记录)
func (client *Client) LogsContact(request *LogsContactRequest) (response *LogsContactResponse, err error) {
	rpcResponse := CreateLogsContactResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type LogsContactRequest struct {
	*api.BaseRequest

	Key      string `json:"key,omitempty"`
	Wid      int64  `json:"wid,omitempty"`
	PageNum  int    `json:"pageNum,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
}

type LogsContactResponse struct {
}

func CreateLogsContactRequest() (request *LogsContactRequest) {
	request = &LogsContactRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "logs/contact", "api")
	request.Method = api.POST
	return
}

func CreateLogsContactResponse() (response *api.BaseResponse[LogsContactResponse]) {
	response = api.CreateResponse[LogsContactResponse](&LogsContactResponse{})
	return
}
