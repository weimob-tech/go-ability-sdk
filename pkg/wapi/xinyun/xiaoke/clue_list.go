package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ClueList
// @id 1653
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询线索列表)
func (client *Client) ClueList(request *ClueListRequest) (response *ClueListResponse, err error) {
	rpcResponse := CreateClueListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ClueListRequest struct {
	*api.BaseRequest

	Wid            int    `json:"wid,omitempty"`
	PageNum        int    `json:"pageNum,omitempty"`
	PageSize       int    `json:"pageSize,omitempty"`
	Type           int    `json:"type,omitempty"`
	ResourcesScope int    `json:"resourcesScope,omitempty"`
	OrderByRule    int    `json:"orderByRule,omitempty"`
	PublicSeaId    string `json:"publicSeaId,omitempty"`
}

type ClueListResponse struct {
	TotalPage  int   `json:"totalPage,omitempty"`
	PageSize   int   `json:"pageSize,omitempty"`
	TotalCount int64 `json:"totalCount,omitempty"`
	PageNum    int64 `json:"pageNum,omitempty"`
}

func CreateClueListRequest() (request *ClueListRequest) {
	request = &ClueListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "clue/list", "api")
	request.Method = api.POST
	return
}

func CreateClueListResponse() (response *api.BaseResponse[ClueListResponse]) {
	response = api.CreateResponse[ClueListResponse](&ClueListResponse{})
	return
}
