package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AiCreateTask
// @id 3723
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=创建任务)
func (client *Client) AiCreateTask(request *AiCreateTaskRequest) (response *AiCreateTaskResponse, err error) {
	rpcResponse := CreateAiCreateTaskResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AiCreateTaskRequest struct {
	*api.BaseRequest

	TaskRecallConfigParam AiCreateTaskRequestTaskRecallConfigParam `json:"taskRecallConfigParam,omitempty"`
	CallTimeParam         AiCreateTaskRequestCallTimeParam         `json:"callTimeParam,omitempty"`
	CallOutTimeRangeList  AiCreateTaskRequestCallOutTimeRangeList  `json:"callOutTimeRangeList,omitempty"`
	Rule                  AiCreateTaskRequestRule                  `json:"rule,omitempty"`
	TaskName              string                                   `json:"taskName,omitempty"`
	SalesTalkId           int64                                    `json:"salesTalkId,omitempty"`
	LoginUserWid          int64                                    `json:"loginUserWid,omitempty"`
	CallLineId            int64                                    `json:"callLineId,omitempty"`
	UserWids              []int64                                  `json:"userWids,omitempty"`
	StartTime             string                                   `json:"startTime,omitempty"`
	CreatePhone           string                                   `json:"createPhone,omitempty"`
	ConcurrentNum         int64                                    `json:"concurrentNum,omitempty"`
}

type AiCreateTaskResponse struct {
	TaskId int64 `json:"taskId,omitempty"`
}

func CreateAiCreateTaskRequest() (request *AiCreateTaskRequest) {
	request = &AiCreateTaskRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "ai/createTask", "api")
	request.Method = api.POST
	return
}

func CreateAiCreateTaskResponse() (response *api.BaseResponse[AiCreateTaskResponse]) {
	response = api.CreateResponse[AiCreateTaskResponse](&AiCreateTaskResponse{})
	return
}

type AiCreateTaskRequestTaskRecallConfigParam struct {
	NeedRecall       int64   `json:"needRecall,omitempty"`
	RecallScenesList []int64 `json:"recallScenesList,omitempty"`
	RecallTimes      int64   `json:"recallTimes,omitempty"`
	RecallInterval   int64   `json:"recallInterval,omitempty"`
	IntervalTimeType int64   `json:"intervalTimeType,omitempty"`
	RecallMode       int64   `json:"recallMode,omitempty"`
}

type AiCreateTaskRequestCallTimeParam struct {
	CallTimeFlag  bool    `json:"callTimeFlag,omitempty"`
	CallTimeModel int64   `json:"callTimeModel,omitempty"`
	Weeks         []int64 `json:"weeks,omitempty"`
}

type AiCreateTaskRequestCallOutTimeRangeList struct {
	BeginTime string `json:"beginTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
}

type AiCreateTaskRequestRule struct {
	AutoAddWechat AiCreateTaskRequestAutoAddWechat `json:"autoAddWechat,omitempty"`
	HangUpMsgs    []AiCreateTaskRequestHangUpMsgs  `json:"hangUpMsgs,omitempty"`
}

type AiCreateTaskRequestAutoAddWechat struct {
	Users             []AiCreateTaskRequestUsers `json:"users,omitempty"`
	AutoAddWechatFlag int64                      `json:"autoAddWechatFlag,omitempty"`
}

type AiCreateTaskRequestUsers struct {
	UserName string `json:"userName,omitempty"`
	UserWid  int64  `json:"userWid,omitempty"`
}

type AiCreateTaskRequestHangUpMsgs struct {
	HangUpMsgItem AiCreateTaskRequestHangUpMsgItem `json:"hangUpMsgItem,omitempty"`
}

type AiCreateTaskRequestHangUpMsgItem struct {
	Levels    []map[string]any `json:"levels,omitempty"`
	LevelId   int64            `json:"levelId,omitempty"`
	LevelName string           `json:"levelName,omitempty"`
	MsgName   string           `json:"msgName,omitempty"`
	MsgTempId int64            `json:"msgTempId,omitempty"`
}
