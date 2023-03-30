package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ScenicDeleteScenicByCode
// @id 1098
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除景区)
func (client *Client) ScenicDeleteScenicByCode(request *ScenicDeleteScenicByCodeRequest) (response *ScenicDeleteScenicByCodeResponse, err error) {
	rpcResponse := CreateScenicDeleteScenicByCodeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ScenicDeleteScenicByCodeRequest struct {
	*api.BaseRequest

	ScenicCode string `json:"scenicCode,omitempty"`
}

type ScenicDeleteScenicByCodeResponse struct {
}

func CreateScenicDeleteScenicByCodeRequest() (request *ScenicDeleteScenicByCodeRequest) {
	request = &ScenicDeleteScenicByCodeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "scenic/deleteScenicByCode", "api")
	request.Method = api.POST
	return
}

func CreateScenicDeleteScenicByCodeResponse() (response *api.BaseResponse[ScenicDeleteScenicByCodeResponse]) {
	response = api.CreateResponse[ScenicDeleteScenicByCodeResponse](&ScenicDeleteScenicByCodeResponse{})
	return
}
