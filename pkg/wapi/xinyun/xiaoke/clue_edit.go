package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ClueEdit
// @id 1650
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=编辑线索)
func (client *Client) ClueEdit(request *ClueEditRequest) (response *ClueEditResponse, err error) {
	rpcResponse := CreateClueEditResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ClueEditRequest struct {
	*api.BaseRequest

	ClueInfo ClueEditRequestClueInfo `json:"clueInfo,omitempty"`
	Key      string                  `json:"key,omitempty"`
	Wid      int64                   `json:"wid,omitempty"`
}

type ClueEditResponse struct {
}

func CreateClueEditRequest() (request *ClueEditRequest) {
	request = &ClueEditRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "clue/edit", "api")
	request.Method = api.POST
	return
}

func CreateClueEditResponse() (response *api.BaseResponse[ClueEditResponse]) {
	response = api.CreateResponse[ClueEditResponse](&ClueEditResponse{})
	return
}

type ClueEditRequestClueInfo struct {
}
