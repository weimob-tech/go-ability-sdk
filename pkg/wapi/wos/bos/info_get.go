package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// InfoGet
// @id 1366
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1366?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询bos基本信息)
func (client *Client) InfoGet(request *InfoGetRequest) (response *InfoGetResponse, err error) {
	rpcResponse := CreateInfoGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type InfoGetRequest struct {
	*api.BaseRequest
}

type InfoGetResponse struct {
	BosName string `json:"bosName,omitempty"`
	BosLogo string `json:"bosLogo,omitempty"`
	Desc    string `json:"desc,omitempty"`
	BosId   int64  `json:"bosId,omitempty"`
}

func CreateInfoGetRequest() (request *InfoGetRequest) {
	request = &InfoGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "info/get", "apigw")
	request.Method = api.POST
	return
}

func CreateInfoGetResponse() (response *api.BaseResponse[InfoGetResponse]) {
	response = api.CreateResponse[InfoGetResponse](&InfoGetResponse{})
	return
}
