package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RoomstatusBatchGetRoomStatus
// @id 530
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量获取房态信息)
func (client *Client) RoomstatusBatchGetRoomStatus(request *RoomstatusBatchGetRoomStatusRequest) (response *RoomstatusBatchGetRoomStatusResponse, err error) {
	rpcResponse := CreateRoomstatusBatchGetRoomStatusResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RoomstatusBatchGetRoomStatusRequest struct {
	*api.BaseRequest

	StoreId int64   `json:"storeId,omitempty"`
	RoomIds []int64 `json:"roomIds,omitempty"`
	From    int     `json:"from,omitempty"`
	To      int     `json:"to,omitempty"`
}

type RoomstatusBatchGetRoomStatusResponse struct {
}

func CreateRoomstatusBatchGetRoomStatusRequest() (request *RoomstatusBatchGetRoomStatusRequest) {
	request = &RoomstatusBatchGetRoomStatusRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "roomstatus/batchGetRoomStatus", "api")
	request.Method = api.POST
	return
}

func CreateRoomstatusBatchGetRoomStatusResponse() (response *api.BaseResponse[RoomstatusBatchGetRoomStatusResponse]) {
	response = api.CreateResponse[RoomstatusBatchGetRoomStatusResponse](&RoomstatusBatchGetRoomStatusResponse{})
	return
}
