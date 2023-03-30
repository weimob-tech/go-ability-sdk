package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SettingDeleteWayCityByCityId
// @id 1017
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除线路产品城市)
func (client *Client) SettingDeleteWayCityByCityId(request *SettingDeleteWayCityByCityIdRequest) (response *SettingDeleteWayCityByCityIdResponse, err error) {
	rpcResponse := CreateSettingDeleteWayCityByCityIdResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SettingDeleteWayCityByCityIdRequest struct {
	*api.BaseRequest

	Id int64 `json:"id,omitempty"`
}

type SettingDeleteWayCityByCityIdResponse struct {
}

func CreateSettingDeleteWayCityByCityIdRequest() (request *SettingDeleteWayCityByCityIdRequest) {
	request = &SettingDeleteWayCityByCityIdRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "setting/deleteWayCityByCityId", "api")
	request.Method = api.POST
	return
}

func CreateSettingDeleteWayCityByCityIdResponse() (response *api.BaseResponse[SettingDeleteWayCityByCityIdResponse]) {
	response = api.CreateResponse[SettingDeleteWayCityByCityIdResponse](&SettingDeleteWayCityByCityIdResponse{})
	return
}
