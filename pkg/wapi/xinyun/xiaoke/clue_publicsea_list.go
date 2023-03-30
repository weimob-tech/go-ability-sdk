package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CluePublicseaList
// @id 1642
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询线索公海列表)
func (client *Client) CluePublicseaList(request *CluePublicseaListRequest) (response *CluePublicseaListResponse, err error) {
	rpcResponse := CreateCluePublicseaListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CluePublicseaListRequest struct {
	*api.BaseRequest

	Wid int `json:"wid,omitempty"`
}

type CluePublicseaListResponse struct {
}

func CreateCluePublicseaListRequest() (request *CluePublicseaListRequest) {
	request = &CluePublicseaListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "clue/publicseaList", "api")
	request.Method = api.POST
	return
}

func CreateCluePublicseaListResponse() (response *api.BaseResponse[CluePublicseaListResponse]) {
	response = api.CreateResponse[CluePublicseaListResponse](&CluePublicseaListResponse{})
	return
}
