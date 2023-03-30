package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AiUpdateTaskStatus
// @id 3737
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=AI外呼更新任务状态)
func (client *Client) AiUpdateTaskStatus(request *AiUpdateTaskStatusRequest) (response *AiUpdateTaskStatusResponse, err error) {
	rpcResponse := CreateAiUpdateTaskStatusResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AiUpdateTaskStatusRequest struct {
	*api.BaseRequest

	TaskId       int64 `json:"taskId,omitempty"`
	LoginUserWid int64 `json:"loginUserWid,omitempty"`
	Status       int64 `json:"status,omitempty"`
}

type AiUpdateTaskStatusResponse struct {
}

func CreateAiUpdateTaskStatusRequest() (request *AiUpdateTaskStatusRequest) {
	request = &AiUpdateTaskStatusRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "ai/updateTaskStatus", "api")
	request.Method = api.POST
	return
}

func CreateAiUpdateTaskStatusResponse() (response *api.BaseResponse[AiUpdateTaskStatusResponse]) {
	response = api.CreateResponse[AiUpdateTaskStatusResponse](&AiUpdateTaskStatusResponse{})
	return
}
