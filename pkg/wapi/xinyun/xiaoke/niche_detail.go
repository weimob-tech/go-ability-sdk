package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NicheDetail
// @id 1737
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询商机详情)
func (client *Client) NicheDetail(request *NicheDetailRequest) (response *NicheDetailResponse, err error) {
	rpcResponse := CreateNicheDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NicheDetailRequest struct {
	*api.BaseRequest

	Wid int    `json:"wid,omitempty"`
	Key string `json:"key,omitempty"`
}

type NicheDetailResponse struct {
}

func CreateNicheDetailRequest() (request *NicheDetailRequest) {
	request = &NicheDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "niche/detail", "api")
	request.Method = api.POST
	return
}

func CreateNicheDetailResponse() (response *api.BaseResponse[NicheDetailResponse]) {
	response = api.CreateResponse[NicheDetailResponse](&NicheDetailResponse{})
	return
}
