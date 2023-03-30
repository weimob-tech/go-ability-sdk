package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SettingGetRoomTypeList
// @id 526
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=房型列表)
func (client *Client) SettingGetRoomTypeList(request *SettingGetRoomTypeListRequest) (response *SettingGetRoomTypeListResponse, err error) {
	rpcResponse := CreateSettingGetRoomTypeListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SettingGetRoomTypeListRequest struct {
	*api.BaseRequest

	StoreId      int64  `json:"storeId,omitempty"`
	RoomTypeName string `json:"roomTypeName,omitempty"`
	PageIndex    int    `json:"pageIndex,omitempty"`
	PageSize     int    `json:"pageSize,omitempty"`
}

type SettingGetRoomTypeListResponse struct {
}

func CreateSettingGetRoomTypeListRequest() (request *SettingGetRoomTypeListRequest) {
	request = &SettingGetRoomTypeListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "setting/getRoomTypeList", "api")
	request.Method = api.POST
	return
}

func CreateSettingGetRoomTypeListResponse() (response *api.BaseResponse[SettingGetRoomTypeListResponse]) {
	response = api.CreateResponse[SettingGetRoomTypeListResponse](&SettingGetRoomTypeListResponse{})
	return
}
