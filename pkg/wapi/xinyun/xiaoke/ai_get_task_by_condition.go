package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AiGetTaskByCondition
// @id 3726
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取任务详情)
func (client *Client) AiGetTaskByCondition(request *AiGetTaskByConditionRequest) (response *AiGetTaskByConditionResponse, err error) {
	rpcResponse := CreateAiGetTaskByConditionResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AiGetTaskByConditionRequest struct {
	*api.BaseRequest

	TaskId int64 `json:"taskId,omitempty"`
}

type AiGetTaskByConditionResponse struct {
	CallTimeParam         AiGetTaskByConditionResponseCallTimeParam         `json:"callTimeParam,omitempty"`
	TaskRecallConfigParam AiGetTaskByConditionResponseTaskRecallConfigParam `json:"taskRecallConfigParam,omitempty"`
	CallOutTimeRangeList  AiGetTaskByConditionResponseCallOutTimeRangeList  `json:"callOutTimeRangeList,omitempty"`
	LoginUserWid          int64                                             `json:"loginUserWid,omitempty"`
	SalesTalkId           int64                                             `json:"salesTalkId,omitempty"`
	ConcurrentNum         int64                                             `json:"concurrentNum,omitempty"`
	UserWids              []int64                                           `json:"userWids,omitempty"`
	CreatorName           string                                            `json:"creatorName,omitempty"`
	CreatePhone           string                                            `json:"createPhone,omitempty"`
	TaskName              string                                            `json:"taskName,omitempty"`
	CallLineId            int64                                             `json:"callLineId,omitempty"`
	TaskId                int64                                             `json:"taskId,omitempty"`
}

func CreateAiGetTaskByConditionRequest() (request *AiGetTaskByConditionRequest) {
	request = &AiGetTaskByConditionRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "ai/getTaskByCondition", "api")
	request.Method = api.POST
	return
}

func CreateAiGetTaskByConditionResponse() (response *api.BaseResponse[AiGetTaskByConditionResponse]) {
	response = api.CreateResponse[AiGetTaskByConditionResponse](&AiGetTaskByConditionResponse{})
	return
}

type AiGetTaskByConditionResponseCallTimeParam struct {
	CallTimeFlag  bool    `json:"callTimeFlag,omitempty"`
	Weeks         []int64 `json:"weeks,omitempty"`
	CallTimeModel int64   `json:"callTimeModel,omitempty"`
}

type AiGetTaskByConditionResponseTaskRecallConfigParam struct {
	IntervalTimeType int64   `json:"intervalTimeType,omitempty"`
	RecallTimes      int64   `json:"recallTimes,omitempty"`
	RecallInterval   int64   `json:"recallInterval,omitempty"`
	RecallMode       int64   `json:"recallMode,omitempty"`
	RecallScenesList []int64 `json:"recallScenesList,omitempty"`
	NeedRecall       int64   `json:"needRecall,omitempty"`
}

type AiGetTaskByConditionResponseCallOutTimeRangeList struct {
	StartTime        int64   `json:"startTime,omitempty"`
	RecallScenesList []int64 `json:"recallScenesList,omitempty"`
}
