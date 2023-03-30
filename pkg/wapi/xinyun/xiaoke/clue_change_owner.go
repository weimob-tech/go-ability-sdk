package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ClueChangeOwner
// @id 1652
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=变更线索所属人)
func (client *Client) ClueChangeOwner(request *ClueChangeOwnerRequest) (response *ClueChangeOwnerResponse, err error) {
	rpcResponse := CreateClueChangeOwnerResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ClueChangeOwnerRequest struct {
	*api.BaseRequest

	Owner int64  `json:"owner,omitempty"`
	Key   string `json:"key,omitempty"`
	Wid   int64  `json:"wid,omitempty"`
}

type ClueChangeOwnerResponse struct {
}

func CreateClueChangeOwnerRequest() (request *ClueChangeOwnerRequest) {
	request = &ClueChangeOwnerRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "clue/changeOwner", "api")
	request.Method = api.POST
	return
}

func CreateClueChangeOwnerResponse() (response *api.BaseResponse[ClueChangeOwnerResponse]) {
	response = api.CreateResponse[ClueChangeOwnerResponse](&ClueChangeOwnerResponse{})
	return
}
