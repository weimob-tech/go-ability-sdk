package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ScenicDeleteScenicCityByCode
// @id 1099
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除景区产品城市)
func (client *Client) ScenicDeleteScenicCityByCode(request *ScenicDeleteScenicCityByCodeRequest) (response *ScenicDeleteScenicCityByCodeResponse, err error) {
	rpcResponse := CreateScenicDeleteScenicCityByCodeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ScenicDeleteScenicCityByCodeRequest struct {
	*api.BaseRequest

	CityCode string `json:"cityCode,omitempty"`
}

type ScenicDeleteScenicCityByCodeResponse struct {
}

func CreateScenicDeleteScenicCityByCodeRequest() (request *ScenicDeleteScenicCityByCodeRequest) {
	request = &ScenicDeleteScenicCityByCodeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "scenic/deleteScenicCityByCode", "api")
	request.Method = api.POST
	return
}

func CreateScenicDeleteScenicCityByCodeResponse() (response *api.BaseResponse[ScenicDeleteScenicCityByCodeResponse]) {
	response = api.CreateResponse[ScenicDeleteScenicCityByCodeResponse](&ScenicDeleteScenicCityByCodeResponse{})
	return
}
