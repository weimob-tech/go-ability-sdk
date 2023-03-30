package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// LogsClue
// @id 1643
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询线索跟进记录)
func (client *Client) LogsClue(request *LogsClueRequest) (response *LogsClueResponse, err error) {
	rpcResponse := CreateLogsClueResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type LogsClueRequest struct {
	*api.BaseRequest

	Key      string `json:"key,omitempty"`
	Wid      int64  `json:"wid,omitempty"`
	PageNum  int    `json:"pageNum,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
}

type LogsClueResponse struct {
}

func CreateLogsClueRequest() (request *LogsClueRequest) {
	request = &LogsClueRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "logs/clue", "api")
	request.Method = api.POST
	return
}

func CreateLogsClueResponse() (response *api.BaseResponse[LogsClueResponse]) {
	response = api.CreateResponse[LogsClueResponse](&LogsClueResponse{})
	return
}
