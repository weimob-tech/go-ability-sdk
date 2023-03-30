package weimob_live

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RoomUserInfoGetList
// @id 2325
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2325?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=直播场次用户列表)
func (client *Client) RoomUserInfoGetList(request *RoomUserInfoGetListRequest) (response *RoomUserInfoGetListResponse, err error) {
	rpcResponse := CreateRoomUserInfoGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RoomUserInfoGetListRequest struct {
	*api.BaseRequest

	RoomId      int64  `json:"roomId,omitempty"`
	ScrollIndex string `json:"scrollIndex,omitempty"`
	PageSize    int64  `json:"pageSize,omitempty"`
	PageNum     int64  `json:"pageNum,omitempty"`
}

type RoomUserInfoGetListResponse struct {
	Data      []RoomUserInfoGetListResponseData `json:"data,omitempty"`
	PageSize  int64                             `json:"pageSize,omitempty"`
	PageNum   int64                             `json:"pageNum,omitempty"`
	TotalSize int64                             `json:"totalSize,omitempty"`
}

func CreateRoomUserInfoGetListRequest() (request *RoomUserInfoGetListRequest) {
	request = &RoomUserInfoGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_live", "v2.0", "room/user/info/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateRoomUserInfoGetListResponse() (response *api.BaseResponse[RoomUserInfoGetListResponse]) {
	response = api.CreateResponse[RoomUserInfoGetListResponse](&RoomUserInfoGetListResponse{})
	return
}

type RoomUserInfoGetListResponseData struct {
	Id          int64  `json:"id,omitempty"`
	BosId       int64  `json:"bosId,omitempty"`
	RoomId      int64  `json:"roomId,omitempty"`
	PlatformWid int64  `json:"platformWid,omitempty"`
	Wid         int64  `json:"wid,omitempty"`
	NickName    string `json:"nickName,omitempty"`
	FansRank    string `json:"fansRank,omitempty"`
	PrizeInfo   int64  `json:"prizeInfo,omitempty"`
	Appointed   int64  `json:"appointed,omitempty"`
	Watched     int64  `json:"watched,omitempty"`
	Commented   int64  `json:"commented,omitempty"`
	Subscribed  int64  `json:"subscribed,omitempty"`
}
