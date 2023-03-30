package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AiQueryLevelList
// @id 3928
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取意向等级列表)
func (client *Client) AiQueryLevelList(request *AiQueryLevelListRequest) (response *AiQueryLevelListResponse, err error) {
	rpcResponse := CreateAiQueryLevelListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AiQueryLevelListRequest struct {
	*api.BaseRequest

	TaskId int64 `json:"taskId,omitempty"`
}

type AiQueryLevelListResponse struct {
	List []AiQueryLevelListResponselist `json:"list,omitempty"`
}

func CreateAiQueryLevelListRequest() (request *AiQueryLevelListRequest) {
	request = &AiQueryLevelListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "ai/queryLevelList", "api")
	request.Method = api.POST
	return
}

func CreateAiQueryLevelListResponse() (response *api.BaseResponse[AiQueryLevelListResponse]) {
	response = api.CreateResponse[AiQueryLevelListResponse](&AiQueryLevelListResponse{})
	return
}

type AiQueryLevelListResponselist struct {
	LevelAlias string `json:"levelAlias,omitempty"`
	LevelId    int64  `json:"levelId,omitempty"`
	LevelName  string `json:"levelName,omitempty"`
}
