package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RoomstatusBatchQueryRoomTypeInfoInPeriod
// @id 1162
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取一段时间可用房型)
func (client *Client) RoomstatusBatchQueryRoomTypeInfoInPeriod(request *RoomstatusBatchQueryRoomTypeInfoInPeriodRequest) (response *RoomstatusBatchQueryRoomTypeInfoInPeriodResponse, err error) {
	rpcResponse := CreateRoomstatusBatchQueryRoomTypeInfoInPeriodResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RoomstatusBatchQueryRoomTypeInfoInPeriodRequest struct {
	*api.BaseRequest

	StoreId      int64   `json:"storeId,omitempty"`
	RoomTypeIds  []int64 `json:"roomTypeIds,omitempty"`
	CheckInTime  int     `json:"checkInTime,omitempty"`
	CheckOutTime int     `json:"checkOutTime,omitempty"`
}

type RoomstatusBatchQueryRoomTypeInfoInPeriodResponse struct {
}

func CreateRoomstatusBatchQueryRoomTypeInfoInPeriodRequest() (request *RoomstatusBatchQueryRoomTypeInfoInPeriodRequest) {
	request = &RoomstatusBatchQueryRoomTypeInfoInPeriodRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "roomstatus/batchQueryRoomTypeInfoInPeriod", "api")
	request.Method = api.POST
	return
}

func CreateRoomstatusBatchQueryRoomTypeInfoInPeriodResponse() (response *api.BaseResponse[RoomstatusBatchQueryRoomTypeInfoInPeriodResponse]) {
	response = api.CreateResponse[RoomstatusBatchQueryRoomTypeInfoInPeriodResponse](&RoomstatusBatchQueryRoomTypeInfoInPeriodResponse{})
	return
}
