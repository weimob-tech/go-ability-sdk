package weimob_live

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RoomInfoGetList
// @id 2324
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2324?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询直播场次列表)
func (client *Client) RoomInfoGetList(request *RoomInfoGetListRequest) (response *RoomInfoGetListResponse, err error) {
	rpcResponse := CreateRoomInfoGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RoomInfoGetListRequest struct {
	*api.BaseRequest

	RoomTitle  string `json:"roomTitle,omitempty"`
	RoomId     string `json:"roomId,omitempty"`
	LiveRoomId int64  `json:"liveRoomId,omitempty"`
	PageSize   int64  `json:"pageSize,omitempty"`
	PageNum    int64  `json:"pageNum,omitempty"`
}

type RoomInfoGetListResponse struct {
	Data      []RoomInfoGetListResponseData `json:"data,omitempty"`
	TotalSize int64                         `json:"totalSize,omitempty"`
	PageSize  int64                         `json:"pageSize,omitempty"`
	PageNum   int64                         `json:"pageNum,omitempty"`
}

func CreateRoomInfoGetListRequest() (request *RoomInfoGetListRequest) {
	request = &RoomInfoGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_live", "v2.0", "room/info/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateRoomInfoGetListResponse() (response *api.BaseResponse[RoomInfoGetListResponse]) {
	response = api.CreateResponse[RoomInfoGetListResponse](&RoomInfoGetListResponse{})
	return
}

type RoomInfoGetListResponseData struct {
	RoomStatus int64  `json:"roomStatus,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
	PlanTime   string `json:"planTime,omitempty"`
	LiveRoomId int64  `json:"liveRoomId,omitempty"`
	StartTime  string `json:"startTime,omitempty"`
	RoomTitle  string `json:"roomTitle,omitempty"`
	RoomId     int64  `json:"roomId,omitempty"`
}
