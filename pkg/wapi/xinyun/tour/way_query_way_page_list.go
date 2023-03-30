package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// WayQueryWayPageList
// @id 997
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=线路产品列表)
func (client *Client) WayQueryWayPageList(request *WayQueryWayPageListRequest) (response *WayQueryWayPageListResponse, err error) {
	rpcResponse := CreateWayQueryWayPageListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WayQueryWayPageListRequest struct {
	*api.BaseRequest

	WayNature       []map[string]any `json:"wayNature,omitempty"`
	WayType         []map[string]any `json:"wayType,omitempty"`
	DepartureCity   []int            `json:"departureCity,omitempty"`
	DestinationCity []int            `json:"destinationCity,omitempty"`
	PageIndex       int              `json:"pageIndex,omitempty"`
	PageSize        int              `json:"pageSize,omitempty"`
	GoodsCode       string           `json:"goodsCode,omitempty"`
	Status          int              `json:"status,omitempty"`
	GoodsName       string           `json:"goodsName,omitempty"`
	StartTime       int              `json:"startTime,omitempty"`
	EndTime         int              `json:"endTime,omitempty"`
	ContainsDel     bool             `json:"containsDel,omitempty"`
}

type WayQueryWayPageListResponse struct {
}

func CreateWayQueryWayPageListRequest() (request *WayQueryWayPageListRequest) {
	request = &WayQueryWayPageListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "way/queryWayPageList", "api")
	request.Method = api.POST
	return
}

func CreateWayQueryWayPageListResponse() (response *api.BaseResponse[WayQueryWayPageListResponse]) {
	response = api.CreateResponse[WayQueryWayPageListResponse](&WayQueryWayPageListResponse{})
	return
}
