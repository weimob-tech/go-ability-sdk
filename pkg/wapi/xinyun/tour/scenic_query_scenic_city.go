package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ScenicQueryScenicCity
// @id 1101
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询景区产品城市)
func (client *Client) ScenicQueryScenicCity(request *ScenicQueryScenicCityRequest) (response *ScenicQueryScenicCityResponse, err error) {
	rpcResponse := CreateScenicQueryScenicCityResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ScenicQueryScenicCityRequest struct {
	*api.BaseRequest
}

type ScenicQueryScenicCityResponse struct {
}

func CreateScenicQueryScenicCityRequest() (request *ScenicQueryScenicCityRequest) {
	request = &ScenicQueryScenicCityRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "scenic/queryScenicCity", "api")
	request.Method = api.POST
	return
}

func CreateScenicQueryScenicCityResponse() (response *api.BaseResponse[ScenicQueryScenicCityResponse]) {
	response = api.CreateResponse[ScenicQueryScenicCityResponse](&ScenicQueryScenicCityResponse{})
	return
}
