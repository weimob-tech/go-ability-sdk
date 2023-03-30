package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SettingQueryWayCityList
// @id 1015
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询线路产品城市)
func (client *Client) SettingQueryWayCityList(request *SettingQueryWayCityListRequest) (response *SettingQueryWayCityListResponse, err error) {
	rpcResponse := CreateSettingQueryWayCityListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SettingQueryWayCityListRequest struct {
	*api.BaseRequest
}

type SettingQueryWayCityListResponse struct {
}

func CreateSettingQueryWayCityListRequest() (request *SettingQueryWayCityListRequest) {
	request = &SettingQueryWayCityListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "setting/queryWayCityList", "api")
	request.Method = api.POST
	return
}

func CreateSettingQueryWayCityListResponse() (response *api.BaseResponse[SettingQueryWayCityListResponse]) {
	response = api.CreateResponse[SettingQueryWayCityListResponse](&SettingQueryWayCityListResponse{})
	return
}
