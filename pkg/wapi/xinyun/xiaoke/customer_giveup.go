package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerGiveup
// @id 1719
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=放弃公海)
func (client *Client) CustomerGiveup(request *CustomerGiveupRequest) (response *CustomerGiveupResponse, err error) {
	rpcResponse := CreateCustomerGiveupResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerGiveupRequest struct {
	*api.BaseRequest

	Keys        []string `json:"keys,omitempty"`
	Wid         int      `json:"wid,omitempty"`
	PublicSeaId string   `json:"publicSeaId,omitempty"`
	Reason      string   `json:"reason,omitempty"`
}

type CustomerGiveupResponse struct {
	Message string `json:"message,omitempty"`
}

func CreateCustomerGiveupRequest() (request *CustomerGiveupRequest) {
	request = &CustomerGiveupRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "customer/giveup", "api")
	request.Method = api.POST
	return
}

func CreateCustomerGiveupResponse() (response *api.BaseResponse[CustomerGiveupResponse]) {
	response = api.CreateResponse[CustomerGiveupResponse](&CustomerGiveupResponse{})
	return
}
