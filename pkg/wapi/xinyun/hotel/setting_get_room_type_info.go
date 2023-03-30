package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SettingGetRoomTypeInfo
// @id 527
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=房型信息)
func (client *Client) SettingGetRoomTypeInfo(request *SettingGetRoomTypeInfoRequest) (response *SettingGetRoomTypeInfoResponse, err error) {
	rpcResponse := CreateSettingGetRoomTypeInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SettingGetRoomTypeInfoRequest struct {
	*api.BaseRequest

	StoreId    int64 `json:"storeId,omitempty"`
	RoomTypeId int64 `json:"roomTypeId,omitempty"`
}

type SettingGetRoomTypeInfoResponse struct {
}

func CreateSettingGetRoomTypeInfoRequest() (request *SettingGetRoomTypeInfoRequest) {
	request = &SettingGetRoomTypeInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "setting/getRoomTypeInfo", "api")
	request.Method = api.POST
	return
}

func CreateSettingGetRoomTypeInfoResponse() (response *api.BaseResponse[SettingGetRoomTypeInfoResponse]) {
	response = api.CreateResponse[SettingGetRoomTypeInfoResponse](&SettingGetRoomTypeInfoResponse{})
	return
}
