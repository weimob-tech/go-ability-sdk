package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SettingDeleteWayTagByTagId
// @id 1014
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除线路标签)
func (client *Client) SettingDeleteWayTagByTagId(request *SettingDeleteWayTagByTagIdRequest) (response *SettingDeleteWayTagByTagIdResponse, err error) {
	rpcResponse := CreateSettingDeleteWayTagByTagIdResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SettingDeleteWayTagByTagIdRequest struct {
	*api.BaseRequest

	Id int64 `json:"id,omitempty"`
}

type SettingDeleteWayTagByTagIdResponse struct {
}

func CreateSettingDeleteWayTagByTagIdRequest() (request *SettingDeleteWayTagByTagIdRequest) {
	request = &SettingDeleteWayTagByTagIdRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "setting/deleteWayTagByTagId", "api")
	request.Method = api.POST
	return
}

func CreateSettingDeleteWayTagByTagIdResponse() (response *api.BaseResponse[SettingDeleteWayTagByTagIdResponse]) {
	response = api.CreateResponse[SettingDeleteWayTagByTagIdResponse](&SettingDeleteWayTagByTagIdResponse{})
	return
}
