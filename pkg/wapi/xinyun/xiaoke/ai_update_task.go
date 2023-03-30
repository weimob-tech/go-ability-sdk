package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AiUpdateTask
// @id 3724
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=编辑任务)
func (client *Client) AiUpdateTask(request *AiUpdateTaskRequest) (response *AiUpdateTaskResponse, err error) {
	rpcResponse := CreateAiUpdateTaskResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AiUpdateTaskRequest struct {
	*api.BaseRequest

	TaskRecallConfigParam AiUpdateTaskRequestTaskRecallConfigParam `json:"taskRecallConfigParam,omitempty"`
	CallTimeParam         AiUpdateTaskRequestCallTimeParam         `json:"callTimeParam,omitempty"`
	Rule                  AiUpdateTaskRequestRule                  `json:"rule,omitempty"`
	TaskId                int64                                    `json:"taskId,omitempty"`
	CallLineId            int64                                    `json:"callLineId,omitempty"`
	ConcurrentNum         int64                                    `json:"concurrentNum,omitempty"`
	LoginUserWid          int64                                    `json:"loginUserWid,omitempty"`
	UserWids              []int64                                  `json:"userWids,omitempty"`
	TaskName              string                                   `json:"taskName,omitempty"`
}

type AiUpdateTaskResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateAiUpdateTaskRequest() (request *AiUpdateTaskRequest) {
	request = &AiUpdateTaskRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "ai/updateTask", "api")
	request.Method = api.POST
	return
}

func CreateAiUpdateTaskResponse() (response *api.BaseResponse[AiUpdateTaskResponse]) {
	response = api.CreateResponse[AiUpdateTaskResponse](&AiUpdateTaskResponse{})
	return
}

type AiUpdateTaskRequestTaskRecallConfigParam struct {
	NeedRecall       int64   `json:"needRecall,omitempty"`
	RecallScenesList []int64 `json:"recallScenesList,omitempty"`
	RecallTimes      int64   `json:"recallTimes,omitempty"`
	RecallInterval   int64   `json:"recallInterval,omitempty"`
	IntervalTimeType int64   `json:"intervalTimeType,omitempty"`
	RecallMode       int64   `json:"recallMode,omitempty"`
}

type AiUpdateTaskRequestCallTimeParam struct {
	CallOutTimeRangeList AiUpdateTaskRequestCallOutTimeRangeList `json:"callOutTimeRangeList,omitempty"`
	CallTimeFlag         bool                                    `json:"callTimeFlag,omitempty"`
	CallTimeModel        int64                                   `json:"callTimeModel,omitempty"`
	Weeks                []int64                                 `json:"weeks,omitempty"`
	StartTime            string                                  `json:"startTime,omitempty"`
}

type AiUpdateTaskRequestCallOutTimeRangeList struct {
	BeginTime string `json:"beginTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
}

type AiUpdateTaskRequestRule struct {
	HangUpMsgs    []AiUpdateTaskRequestHangUpMsgs  `json:"hangUpMsgs,omitempty"`
	AutoAddWechat AiUpdateTaskRequestAutoAddWechat `json:"autoAddWechat,omitempty"`
}

type AiUpdateTaskRequestHangUpMsgs struct {
	HangUpMsgItem AiUpdateTaskRequestHangUpMsgItem `json:"hangUpMsgItem,omitempty"`
}

type AiUpdateTaskRequestHangUpMsgItem struct {
	Levels    []map[string]any `json:"levels,omitempty"`
	LevelId   int64            `json:"levelId,omitempty"`
	LevelName string           `json:"levelName,omitempty"`
	MsgName   string           `json:"msgName,omitempty"`
	MsgTempId int64            `json:"msgTempId,omitempty"`
}

type AiUpdateTaskRequestAutoAddWechat struct {
	Users             []AiUpdateTaskRequestUsers `json:"users,omitempty"`
	AutoAddWechatFlag int64                      `json:"autoAddWechatFlag,omitempty"`
}

type AiUpdateTaskRequestUsers struct {
	UserName string `json:"userName,omitempty"`
	UserWid  int64  `json:"userWid,omitempty"`
}
