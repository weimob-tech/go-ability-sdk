package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AiQueryUserInfos
// @id 3727
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询企业用户列表)
func (client *Client) AiQueryUserInfos(request *AiQueryUserInfosRequest) (response *AiQueryUserInfosResponse, err error) {
	rpcResponse := CreateAiQueryUserInfosResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AiQueryUserInfosRequest struct {
	*api.BaseRequest

	PageNum  int64 `json:"pageNum,omitempty"`
	PageSize int64 `json:"pageSize,omitempty"`
}

type AiQueryUserInfosResponse struct {
	UserInfos  AiQueryUserInfosResponseUserInfos `json:"userInfos,omitempty"`
	TotalPage  int64                             `json:"totalPage,omitempty"`
	TotalCount int64                             `json:"totalCount,omitempty"`
}

func CreateAiQueryUserInfosRequest() (request *AiQueryUserInfosRequest) {
	request = &AiQueryUserInfosRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "ai/queryUserInfos", "api")
	request.Method = api.POST
	return
}

func CreateAiQueryUserInfosResponse() (response *api.BaseResponse[AiQueryUserInfosResponse]) {
	response = api.CreateResponse[AiQueryUserInfosResponse](&AiQueryUserInfosResponse{})
	return
}

type AiQueryUserInfosResponseUserInfos struct {
	AccountId int64  `json:"accountId,omitempty"`
	Name      string `json:"name,omitempty"`
	UserWid   int64  `json:"userWid,omitempty"`
}
