package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AiQueryMessageTemplateList
// @id 3927
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取短信模版列表)
func (client *Client) AiQueryMessageTemplateList(request *AiQueryMessageTemplateListRequest) (response *AiQueryMessageTemplateListResponse, err error) {
	rpcResponse := CreateAiQueryMessageTemplateListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AiQueryMessageTemplateListRequest struct {
	*api.BaseRequest

	PageNum  int64 `json:"pageNum,omitempty"`
	PageSize int64 `json:"pageSize,omitempty"`
}

type AiQueryMessageTemplateListResponse struct {
	List       []AiQueryMessageTemplateListResponselist `json:"list,omitempty"`
	StartIndex int64                                    `json:"startIndex,omitempty"`
	TotalPage  int64                                    `json:"totalPage,omitempty"`
	PageSize   int64                                    `json:"pageSize,omitempty"`
	PageNum    int64                                    `json:"pageNum,omitempty"`
}

func CreateAiQueryMessageTemplateListRequest() (request *AiQueryMessageTemplateListRequest) {
	request = &AiQueryMessageTemplateListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "ai/queryMessageTemplateList", "api")
	request.Method = api.POST
	return
}

func CreateAiQueryMessageTemplateListResponse() (response *api.BaseResponse[AiQueryMessageTemplateListResponse]) {
	response = api.CreateResponse[AiQueryMessageTemplateListResponse](&AiQueryMessageTemplateListResponse{})
	return
}

type AiQueryMessageTemplateListResponselist struct {
	TemplateText string `json:"templateText,omitempty"`
	CreateTime   int64  `json:"createTime,omitempty"`
	TemplateName string `json:"templateName,omitempty"`
	UpdateTime   int64  `json:"updateTime,omitempty"`
	Id           int64  `json:"id,omitempty"`
}
