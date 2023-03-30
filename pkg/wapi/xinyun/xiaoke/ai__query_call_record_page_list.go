package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AiQueryCallRecordPageList
// @id 3741
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=AI外呼查询通话记录)
func (client *Client) AiQueryCallRecordPageList(request *AiQueryCallRecordPageListRequest) (response *AiQueryCallRecordPageListResponse, err error) {
	rpcResponse := CreateAiQueryCallRecordPageListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AiQueryCallRecordPageListRequest struct {
	*api.BaseRequest

	Duration     AiQueryCallRecordPageListRequestDuration  `json:"duration,omitempty"`
	QueryTime    AiQueryCallRecordPageListRequestQueryTime `json:"queryTime,omitempty"`
	LoginUserWid int64                                     `json:"loginUserWid,omitempty"`
	PageNum      int64                                     `json:"pageNum,omitempty"`
	PageSize     int64                                     `json:"pageSize,omitempty"`
	TimeTypeKey  string                                    `json:"timeTypeKey,omitempty"`
	TaskId       int64                                     `json:"taskId,omitempty"`
	LevelNames   string                                    `json:"levelNames,omitempty"`
	UserWids     string                                    `json:"userWids,omitempty"`
	CallType     int64                                     `json:"callType,omitempty"`
}

type AiQueryCallRecordPageListResponse struct {
	List       AiQueryCallRecordPageListResponselist `json:"list,omitempty"`
	StartIndex int64                                 `json:"startIndex,omitempty"`
	TotalPage  int64                                 `json:"totalPage,omitempty"`
	PageSize   int64                                 `json:"pageSize,omitempty"`
	TotalCount int64                                 `json:"totalCount,omitempty"`
	PageNum    int64                                 `json:"pageNum,omitempty"`
}

func CreateAiQueryCallRecordPageListRequest() (request *AiQueryCallRecordPageListRequest) {
	request = &AiQueryCallRecordPageListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "ai/QueryCallRecordPageList", "api")
	request.Method = api.POST
	return
}

func CreateAiQueryCallRecordPageListResponse() (response *api.BaseResponse[AiQueryCallRecordPageListResponse]) {
	response = api.CreateResponse[AiQueryCallRecordPageListResponse](&AiQueryCallRecordPageListResponse{})
	return
}

type AiQueryCallRecordPageListRequestDuration struct {
	DurationType int64 `json:"durationType,omitempty"`
	MinDuration  int64 `json:"minDuration,omitempty"`
	MaxDuration  int64 `json:"maxDuration,omitempty"`
}

type AiQueryCallRecordPageListRequestQueryTime struct {
	QueryStartTime string `json:"queryStartTime,omitempty"`
	QueryEndTime   string `json:"queryEndTime,omitempty"`
}

type AiQueryCallRecordPageListResponselist struct {
}
