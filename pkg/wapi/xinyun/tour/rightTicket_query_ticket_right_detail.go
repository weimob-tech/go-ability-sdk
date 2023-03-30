package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightTicketQueryTicketRightDetail
// @id 1088
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询门票维权详情)
func (client *Client) RightTicketQueryTicketRightDetail(request *RightTicketQueryTicketRightDetailRequest) (response *RightTicketQueryTicketRightDetailResponse, err error) {
	rpcResponse := CreateRightTicketQueryTicketRightDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightTicketQueryTicketRightDetailRequest struct {
	*api.BaseRequest

	RightNo string `json:"rightNo,omitempty"`
	OrderNo string `json:"orderNo,omitempty"`
}

type RightTicketQueryTicketRightDetailResponse struct {
}

func CreateRightTicketQueryTicketRightDetailRequest() (request *RightTicketQueryTicketRightDetailRequest) {
	request = &RightTicketQueryTicketRightDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "rightTicket/queryTicketRightDetail", "api")
	request.Method = api.POST
	return
}

func CreateRightTicketQueryTicketRightDetailResponse() (response *api.BaseResponse[RightTicketQueryTicketRightDetailResponse]) {
	response = api.CreateResponse[RightTicketQueryTicketRightDetailResponse](&RightTicketQueryTicketRightDetailResponse{})
	return
}
