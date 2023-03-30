package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ClueGiveup
// @id 1655
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=放弃公海)
func (client *Client) ClueGiveup(request *ClueGiveupRequest) (response *ClueGiveupResponse, err error) {
	rpcResponse := CreateClueGiveupResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ClueGiveupRequest struct {
	*api.BaseRequest

	Keys        []string `json:"keys,omitempty"`
	Wid         int      `json:"wid,omitempty"`
	PublicSeaId string   `json:"publicSeaId,omitempty"`
	Reason      string   `json:"reason,omitempty"`
}

type ClueGiveupResponse struct {
	Message string `json:"message,omitempty"`
}

func CreateClueGiveupRequest() (request *ClueGiveupRequest) {
	request = &ClueGiveupRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "clue/giveup", "api")
	request.Method = api.POST
	return
}

func CreateClueGiveupResponse() (response *api.BaseResponse[ClueGiveupResponse]) {
	response = api.CreateResponse[ClueGiveupResponse](&ClueGiveupResponse{})
	return
}
