package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AiQueryLineList
// @id 3740
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=AI外呼查询线路列表)
func (client *Client) AiQueryLineList(request *AiQueryLineListRequest) (response *AiQueryLineListResponse, err error) {
	rpcResponse := CreateAiQueryLineListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AiQueryLineListRequest struct {
	*api.BaseRequest
}

type AiQueryLineListResponse struct {
	LineDtoList AiQueryLineListResponseLineDtoList `json:"lineDtoList,omitempty"`
}

func CreateAiQueryLineListRequest() (request *AiQueryLineListRequest) {
	request = &AiQueryLineListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "ai/queryLineList", "api")
	request.Method = api.POST
	return
}

func CreateAiQueryLineListResponse() (response *api.BaseResponse[AiQueryLineListResponse]) {
	response = api.CreateResponse[AiQueryLineListResponse](&AiQueryLineListResponse{})
	return
}

type AiQueryLineListResponseLineDtoList struct {
	ConcurrentLimit int64  `json:"concurrentLimit,omitempty"`
	ExtId           string `json:"extId,omitempty"`
	AvailableNum    int64  `json:"availableNum,omitempty"`
	Source          int64  `json:"source,omitempty"`
	Id              int64  `json:"id,omitempty"`
	ExtName         int64  `json:"extName,omitempty"`
}
