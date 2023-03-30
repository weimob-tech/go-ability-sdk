package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ClueDetail
// @id 1656
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询线索详情)
func (client *Client) ClueDetail(request *ClueDetailRequest) (response *ClueDetailResponse, err error) {
	rpcResponse := CreateClueDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ClueDetailRequest struct {
	*api.BaseRequest

	Wid int    `json:"wid,omitempty"`
	Key string `json:"key,omitempty"`
}

type ClueDetailResponse struct {
}

func CreateClueDetailRequest() (request *ClueDetailRequest) {
	request = &ClueDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "clue/detail", "api")
	request.Method = api.POST
	return
}

func CreateClueDetailResponse() (response *api.BaseResponse[ClueDetailResponse]) {
	response = api.CreateResponse[ClueDetailResponse](&ClueDetailResponse{})
	return
}
