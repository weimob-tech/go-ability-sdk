package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SettingSaveWayCity
// @id 1016
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增/修改线路产品城市)
func (client *Client) SettingSaveWayCity(request *SettingSaveWayCityRequest) (response *SettingSaveWayCityResponse, err error) {
	rpcResponse := CreateSettingSaveWayCityResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SettingSaveWayCityRequest struct {
	*api.BaseRequest

	Id      int64  `json:"id,omitempty"`
	Country string `json:"country,omitempty"`
	Name    string `json:"name,omitempty"`
	Type    int    `json:"type,omitempty"`
	IsHot   int    `json:"isHot,omitempty"`
}

type SettingSaveWayCityResponse struct {
}

func CreateSettingSaveWayCityRequest() (request *SettingSaveWayCityRequest) {
	request = &SettingSaveWayCityRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "setting/saveWayCity", "api")
	request.Method = api.POST
	return
}

func CreateSettingSaveWayCityResponse() (response *api.BaseResponse[SettingSaveWayCityResponse]) {
	response = api.CreateResponse[SettingSaveWayCityResponse](&SettingSaveWayCityResponse{})
	return
}
