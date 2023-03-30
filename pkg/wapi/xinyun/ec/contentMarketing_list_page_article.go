package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ContentMarketingListPageArticle
// @id 1281
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=专题文章列表)
func (client *Client) ContentMarketingListPageArticle(request *ContentMarketingListPageArticleRequest) (response *ContentMarketingListPageArticleResponse, err error) {
	rpcResponse := CreateContentMarketingListPageArticleResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ContentMarketingListPageArticleRequest struct {
	*api.BaseRequest

	QueryParameter ContentMarketingListPageArticleRequestQueryParameter `json:"queryParameter,omitempty"`
	StoreId        int64                                                `json:"storeId,omitempty"`
	PageSize       int                                                  `json:"pageSize,omitempty"`
	StartIndex     int64                                                `json:"startIndex,omitempty"`
}

type ContentMarketingListPageArticleResponse struct {
}

func CreateContentMarketingListPageArticleRequest() (request *ContentMarketingListPageArticleRequest) {
	request = &ContentMarketingListPageArticleRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "contentMarketing/listPageArticle", "api")
	request.Method = api.POST
	return
}

func CreateContentMarketingListPageArticleResponse() (response *api.BaseResponse[ContentMarketingListPageArticleResponse]) {
	response = api.CreateResponse[ContentMarketingListPageArticleResponse](&ContentMarketingListPageArticleResponse{})
	return
}

type ContentMarketingListPageArticleRequestQueryParameter struct {
}
