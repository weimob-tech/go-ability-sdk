package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ClueDrop
// @id 1800
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=线索掉保)
func (client *Client) ClueDrop(request *ClueDropRequest) (response *ClueDropResponse, err error) {
	rpcResponse := CreateClueDropResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ClueDropRequest struct {
	*api.BaseRequest

	Keys []string `json:"keys,omitempty"`
	Wid  int      `json:"wid,omitempty"`
}

type ClueDropResponse struct {
}

func CreateClueDropRequest() (request *ClueDropRequest) {
	request = &ClueDropRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "clue/drop", "api")
	request.Method = api.POST
	return
}

func CreateClueDropResponse() (response *api.BaseResponse[ClueDropResponse]) {
	response = api.CreateResponse[ClueDropResponse](&ClueDropResponse{})
	return
}
