package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CluePublicseaObtain
// @id 1721
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=领取/分配线索公海数据)
func (client *Client) CluePublicseaObtain(request *CluePublicseaObtainRequest) (response *CluePublicseaObtainResponse, err error) {
	rpcResponse := CreateCluePublicseaObtainResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CluePublicseaObtainRequest struct {
	*api.BaseRequest

	Owner      int64  `json:"owner,omitempty"`
	Keys       string `json:"keys,omitempty"`
	Wid        int64  `json:"wid,omitempty"`
	OptionType int    `json:"optionType,omitempty"`
}

type CluePublicseaObtainResponse struct {
}

func CreateCluePublicseaObtainRequest() (request *CluePublicseaObtainRequest) {
	request = &CluePublicseaObtainRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "clue/publicseaObtain", "api")
	request.Method = api.POST
	return
}

func CreateCluePublicseaObtainResponse() (response *api.BaseResponse[CluePublicseaObtainResponse]) {
	response = api.CreateResponse[CluePublicseaObtainResponse](&CluePublicseaObtainResponse{})
	return
}
