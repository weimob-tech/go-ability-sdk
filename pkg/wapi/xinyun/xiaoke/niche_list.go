package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NicheList
// @id 1732
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询商机列表)
func (client *Client) NicheList(request *NicheListRequest) (response *NicheListResponse, err error) {
	rpcResponse := CreateNicheListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NicheListRequest struct {
	*api.BaseRequest

	Wid            int `json:"wid,omitempty"`
	PageNum        int `json:"pageNum,omitempty"`
	PageSize       int `json:"pageSize,omitempty"`
	Type           int `json:"type,omitempty"`
	ResourcesScope int `json:"resourcesScope,omitempty"`
	OrderByRule    int `json:"orderByRule,omitempty"`
}

type NicheListResponse struct {
	TotalPage  int   `json:"totalPage,omitempty"`
	PageSize   int   `json:"pageSize,omitempty"`
	TotalCount int64 `json:"totalCount,omitempty"`
	PageNum    int64 `json:"pageNum,omitempty"`
}

func CreateNicheListRequest() (request *NicheListRequest) {
	request = &NicheListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "niche/list", "api")
	request.Method = api.POST
	return
}

func CreateNicheListResponse() (response *api.BaseResponse[NicheListResponse]) {
	response = api.CreateResponse[NicheListResponse](&NicheListResponse{})
	return
}
