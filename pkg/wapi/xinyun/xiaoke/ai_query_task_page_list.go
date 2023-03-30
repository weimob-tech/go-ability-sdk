package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AiQueryTaskPageList
// @id 3736
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=AI外呼查询任务列表)
func (client *Client) AiQueryTaskPageList(request *AiQueryTaskPageListRequest) (response *AiQueryTaskPageListResponse, err error) {
	rpcResponse := CreateAiQueryTaskPageListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AiQueryTaskPageListRequest struct {
	*api.BaseRequest

	QueryTime    AiQueryTaskPageListRequestQueryTime `json:"queryTime,omitempty"`
	LoginUserWid int64                               `json:"loginUserWid,omitempty"`
	Creator      string                              `json:"creator,omitempty"`
	PageNum      int64                               `json:"pageNum,omitempty"`
	PageSize     int64                               `json:"pageSize,omitempty"`
	TaskStatus   []int64                             `json:"taskStatus,omitempty"`
	TimeTypeKey  string                              `json:"timeTypeKey,omitempty"`
	TaskNameStr  string                              `json:"taskNameStr,omitempty"`
}

type AiQueryTaskPageListResponse struct {
	List       AiQueryTaskPageListResponselist `json:"list,omitempty"`
	StartIndex int64                           `json:"startIndex,omitempty"`
	TotalPage  int64                           `json:"totalPage,omitempty"`
	PageSize   int64                           `json:"pageSize,omitempty"`
	TotalCount int64                           `json:"totalCount,omitempty"`
	PageNum    int64                           `json:"pageNum,omitempty"`
	NextFlag   bool                            `json:"nextFlag,omitempty"`
}

func CreateAiQueryTaskPageListRequest() (request *AiQueryTaskPageListRequest) {
	request = &AiQueryTaskPageListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "ai/queryTaskPageList", "api")
	request.Method = api.POST
	return
}

func CreateAiQueryTaskPageListResponse() (response *api.BaseResponse[AiQueryTaskPageListResponse]) {
	response = api.CreateResponse[AiQueryTaskPageListResponse](&AiQueryTaskPageListResponse{})
	return
}

type AiQueryTaskPageListRequestQueryTime struct {
	QueryStartTime string `json:"queryStartTime,omitempty"`
	QueryEndTime   string `json:"queryEndTime,omitempty"`
}

type AiQueryTaskPageListResponselist struct {
	CreateTime     int64  `json:"createTime,omitempty"`
	TotalNumberCnt int64  `json:"totalNumberCnt,omitempty"`
	CreatorName    string `json:"creatorName,omitempty"`
	Name           string `json:"name,omitempty"`
	ProgressCnt    int64  `json:"progressCnt,omitempty"`
	Status         int64  `json:"status,omitempty"`
	TaskId         int64  `json:"taskId,omitempty"`
}
