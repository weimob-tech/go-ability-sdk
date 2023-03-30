package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ClueCreate
// @id 1651
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新建线索)
func (client *Client) ClueCreate(request *ClueCreateRequest) (response *ClueCreateResponse, err error) {
	rpcResponse := CreateClueCreateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ClueCreateRequest struct {
	*api.BaseRequest

	ClueInfo    ClueCreateRequestClueInfo `json:"clueInfo,omitempty"`
	PublicSeaId string                    `json:"publicSeaId,omitempty"`
	Wid         int64                     `json:"wid,omitempty"`
}

type ClueCreateResponse struct {
}

func CreateClueCreateRequest() (request *ClueCreateRequest) {
	request = &ClueCreateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "clue/create", "api")
	request.Method = api.POST
	return
}

func CreateClueCreateResponse() (response *api.BaseResponse[ClueCreateResponse]) {
	response = api.CreateResponse[ClueCreateResponse](&ClueCreateResponse{})
	return
}

type ClueCreateRequestClueInfo struct {
}
