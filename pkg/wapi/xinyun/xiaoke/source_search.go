package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SourceSearch
// @id 2136
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询客资录入方式列表)
func (client *Client) SourceSearch(request *SourceSearchRequest) (response *SourceSearchResponse, err error) {
	rpcResponse := CreateSourceSearchResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SourceSearchRequest struct {
	*api.BaseRequest

	Wid   int64  `json:"wid,omitempty"`
	Stage string `json:"stage,omitempty"`
}

type SourceSearchResponse map[string]any

func CreateSourceSearchRequest() (request *SourceSearchRequest) {
	request = &SourceSearchRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "source/search", "api")
	request.Method = api.POST
	return
}

func CreateSourceSearchResponse() (response *api.BaseResponse[SourceSearchResponse]) {
	response = api.CreateResponse[SourceSearchResponse](&SourceSearchResponse{})
	return
}
