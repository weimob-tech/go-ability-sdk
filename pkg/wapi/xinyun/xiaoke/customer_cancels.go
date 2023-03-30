package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerCancels
// @id 1715
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除客户)
func (client *Client) CustomerCancels(request *CustomerCancelsRequest) (response *CustomerCancelsResponse, err error) {
	rpcResponse := CreateCustomerCancelsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerCancelsRequest struct {
	*api.BaseRequest

	Keys []string `json:"keys,omitempty"`
	Wid  int      `json:"wid,omitempty"`
}

type CustomerCancelsResponse struct {
	Message string `json:"message,omitempty"`
}

func CreateCustomerCancelsRequest() (request *CustomerCancelsRequest) {
	request = &CustomerCancelsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "customer/cancels", "api")
	request.Method = api.POST
	return
}

func CreateCustomerCancelsResponse() (response *api.BaseResponse[CustomerCancelsResponse]) {
	response = api.CreateResponse[CustomerCancelsResponse](&CustomerCancelsResponse{})
	return
}
