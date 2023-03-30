package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ClueCancels
// @id 1657
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除线索)
func (client *Client) ClueCancels(request *ClueCancelsRequest) (response *ClueCancelsResponse, err error) {
	rpcResponse := CreateClueCancelsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ClueCancelsRequest struct {
	*api.BaseRequest

	Keys []map[string]any `json:"keys,omitempty"`
	Wid  int              `json:"wid,omitempty"`
}

type ClueCancelsResponse struct {
	Message string `json:"message,omitempty"`
}

func CreateClueCancelsRequest() (request *ClueCancelsRequest) {
	request = &ClueCancelsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "clue/cancels", "api")
	request.Method = api.POST
	return
}

func CreateClueCancelsResponse() (response *api.BaseResponse[ClueCancelsResponse]) {
	response = api.CreateResponse[ClueCancelsResponse](&ClueCancelsResponse{})
	return
}
