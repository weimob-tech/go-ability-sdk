package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RoomstatusEditRoomStatus
// @id 531
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新房态)
func (client *Client) RoomstatusEditRoomStatus(request *RoomstatusEditRoomStatusRequest) (response *RoomstatusEditRoomStatusResponse, err error) {
	rpcResponse := CreateRoomstatusEditRoomStatusResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RoomstatusEditRoomStatusRequest struct {
	*api.BaseRequest

	StoreId       int64  `json:"storeId,omitempty"`
	RoomId        int64  `json:"roomId,omitempty"`
	RoomStartDate int    `json:"roomStartDate,omitempty"`
	RoomEndDate   int    `json:"roomEndDate,omitempty"`
	RoomTypeId    int64  `json:"roomTypeId,omitempty"`
	RoomStatus    int    `json:"roomStatus,omitempty"`
	RemarkVo      string `json:"remarkVo,omitempty"`
}

type RoomstatusEditRoomStatusResponse struct {
}

func CreateRoomstatusEditRoomStatusRequest() (request *RoomstatusEditRoomStatusRequest) {
	request = &RoomstatusEditRoomStatusRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "roomstatus/editRoomStatus", "api")
	request.Method = api.POST
	return
}

func CreateRoomstatusEditRoomStatusResponse() (response *api.BaseResponse[RoomstatusEditRoomStatusResponse]) {
	response = api.CreateResponse[RoomstatusEditRoomStatusResponse](&RoomstatusEditRoomStatusResponse{})
	return
}
