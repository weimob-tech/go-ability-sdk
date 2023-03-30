package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightTicketAllowTicketRight
// @id 1089
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=同意门票维权)
func (client *Client) RightTicketAllowTicketRight(request *RightTicketAllowTicketRightRequest) (response *RightTicketAllowTicketRightResponse, err error) {
	rpcResponse := CreateRightTicketAllowTicketRightResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightTicketAllowTicketRightRequest struct {
	*api.BaseRequest

	OrderNo string `json:"orderNo,omitempty"`
	RightNo string `json:"rightNo,omitempty"`
}

type RightTicketAllowTicketRightResponse struct {
}

func CreateRightTicketAllowTicketRightRequest() (request *RightTicketAllowTicketRightRequest) {
	request = &RightTicketAllowTicketRightRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "rightTicket/allowTicketRight", "api")
	request.Method = api.POST
	return
}

func CreateRightTicketAllowTicketRightResponse() (response *api.BaseResponse[RightTicketAllowTicketRightResponse]) {
	response = api.CreateResponse[RightTicketAllowTicketRightResponse](&RightTicketAllowTicketRightResponse{})
	return
}
