package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightTicketQueryRightTicketList
// @id 1087
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询门票维权列表)
func (client *Client) RightTicketQueryRightTicketList(request *RightTicketQueryRightTicketListRequest) (response *RightTicketQueryRightTicketListResponse, err error) {
	rpcResponse := CreateRightTicketQueryRightTicketListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightTicketQueryRightTicketListRequest struct {
	*api.BaseRequest

	RightNo     string `json:"rightNo,omitempty"`
	RightStatus int    `json:"rightStatus,omitempty"`
	RightApply  int    `json:"rightApply,omitempty"`
	BeginTime   int    `json:"beginTime,omitempty"`
	EndTime     int    `json:"endTime,omitempty"`
	PageIndex   int    `json:"pageIndex,omitempty"`
	PageSize    int    `json:"pageSize,omitempty"`
}

type RightTicketQueryRightTicketListResponse struct {
}

func CreateRightTicketQueryRightTicketListRequest() (request *RightTicketQueryRightTicketListRequest) {
	request = &RightTicketQueryRightTicketListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "rightTicket/queryRightTicketList", "api")
	request.Method = api.POST
	return
}

func CreateRightTicketQueryRightTicketListResponse() (response *api.BaseResponse[RightTicketQueryRightTicketListResponse]) {
	response = api.CreateResponse[RightTicketQueryRightTicketListResponse](&RightTicketQueryRightTicketListResponse{})
	return
}
