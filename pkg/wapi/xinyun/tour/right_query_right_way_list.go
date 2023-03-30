package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightQueryRightWayList
// @id 1006
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询线路维权列表)
func (client *Client) RightQueryRightWayList(request *RightQueryRightWayListRequest) (response *RightQueryRightWayListResponse, err error) {
	rpcResponse := CreateRightQueryRightWayListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightQueryRightWayListRequest struct {
	*api.BaseRequest

	RightNo     string `json:"rightNo,omitempty"`
	RightStatus int    `json:"rightStatus,omitempty"`
	RightApply  int    `json:"rightApply,omitempty"`
	BeginTime   int    `json:"beginTime,omitempty"`
	EndTime     int    `json:"endTime,omitempty"`
	PageIndex   int    `json:"pageIndex,omitempty"`
	PageSize    int    `json:"pageSize,omitempty"`
}

type RightQueryRightWayListResponse struct {
}

func CreateRightQueryRightWayListRequest() (request *RightQueryRightWayListRequest) {
	request = &RightQueryRightWayListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "right/queryRightWayList", "api")
	request.Method = api.POST
	return
}

func CreateRightQueryRightWayListResponse() (response *api.BaseResponse[RightQueryRightWayListResponse]) {
	response = api.CreateResponse[RightQueryRightWayListResponse](&RightQueryRightWayListResponse{})
	return
}
