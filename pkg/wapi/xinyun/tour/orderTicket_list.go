package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderTicketList
// @id 1110
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询门票订单列表)
func (client *Client) OrderTicketList(request *OrderTicketListRequest) (response *OrderTicketListResponse, err error) {
	rpcResponse := CreateOrderTicketListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderTicketListRequest struct {
	*api.BaseRequest

	OrderStatus int    `json:"orderStatus,omitempty"`
	TimeType    int    `json:"timeType,omitempty"`
	StartTime   int    `json:"startTime,omitempty"`
	EndTime     int    `json:"endTime,omitempty"`
	Keyword     int    `json:"keyword,omitempty"`
	Value       string `json:"value,omitempty"`
	PageIndex   int    `json:"pageIndex,omitempty"`
	PageSize    int    `json:"pageSize,omitempty"`
}

type OrderTicketListResponse struct {
}

func CreateOrderTicketListRequest() (request *OrderTicketListRequest) {
	request = &OrderTicketListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "orderTicket/list", "api")
	request.Method = api.POST
	return
}

func CreateOrderTicketListResponse() (response *api.BaseResponse[OrderTicketListResponse]) {
	response = api.CreateResponse[OrderTicketListResponse](&OrderTicketListResponse{})
	return
}
