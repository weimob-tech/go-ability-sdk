package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ClueDelete
// @id 1801
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=强制删除线索)
func (client *Client) ClueDelete(request *ClueDeleteRequest) (response *ClueDeleteResponse, err error) {
	rpcResponse := CreateClueDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ClueDeleteRequest struct {
	*api.BaseRequest

	Keys []string `json:"keys,omitempty"`
	Wid  int      `json:"wid,omitempty"`
}

type ClueDeleteResponse struct {
	Message string `json:"message,omitempty"`
}

func CreateClueDeleteRequest() (request *ClueDeleteRequest) {
	request = &ClueDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "clue/delete", "api")
	request.Method = api.POST
	return
}

func CreateClueDeleteResponse() (response *api.BaseResponse[ClueDeleteResponse]) {
	response = api.CreateResponse[ClueDeleteResponse](&ClueDeleteResponse{})
	return
}
