package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TicketVerifyVerify
// @id 1090
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=门票核销)
func (client *Client) TicketVerifyVerify(request *TicketVerifyVerifyRequest) (response *TicketVerifyVerifyResponse, err error) {
	rpcResponse := CreateTicketVerifyVerifyResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TicketVerifyVerifyRequest struct {
	*api.BaseRequest

	Code string `json:"code,omitempty"`
}

type TicketVerifyVerifyResponse struct {
}

func CreateTicketVerifyVerifyRequest() (request *TicketVerifyVerifyRequest) {
	request = &TicketVerifyVerifyRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "ticketVerify/verify", "api")
	request.Method = api.POST
	return
}

func CreateTicketVerifyVerifyResponse() (response *api.BaseResponse[TicketVerifyVerifyResponse]) {
	response = api.CreateResponse[TicketVerifyVerifyResponse](&TicketVerifyVerifyResponse{})
	return
}
