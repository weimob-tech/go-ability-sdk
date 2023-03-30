package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ScenicGetScenicByCode
// @id 1100
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询景区产品详情)
func (client *Client) ScenicGetScenicByCode(request *ScenicGetScenicByCodeRequest) (response *ScenicGetScenicByCodeResponse, err error) {
	rpcResponse := CreateScenicGetScenicByCodeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ScenicGetScenicByCodeRequest struct {
	*api.BaseRequest

	ScenicCode string `json:"scenicCode,omitempty"`
}

type ScenicGetScenicByCodeResponse struct {
}

func CreateScenicGetScenicByCodeRequest() (request *ScenicGetScenicByCodeRequest) {
	request = &ScenicGetScenicByCodeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "scenic/getScenicByCode", "api")
	request.Method = api.POST
	return
}

func CreateScenicGetScenicByCodeResponse() (response *api.BaseResponse[ScenicGetScenicByCodeResponse]) {
	response = api.CreateResponse[ScenicGetScenicByCodeResponse](&ScenicGetScenicByCodeResponse{})
	return
}
