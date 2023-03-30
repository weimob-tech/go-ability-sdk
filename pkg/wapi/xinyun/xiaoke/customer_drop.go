package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerDrop
// @id 1799
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=客户掉保)
func (client *Client) CustomerDrop(request *CustomerDropRequest) (response *CustomerDropResponse, err error) {
	rpcResponse := CreateCustomerDropResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerDropRequest struct {
	*api.BaseRequest

	Keys []string `json:"keys,omitempty"`
	Wid  int      `json:"wid,omitempty"`
}

type CustomerDropResponse struct {
}

func CreateCustomerDropRequest() (request *CustomerDropRequest) {
	request = &CustomerDropRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "customer/drop", "api")
	request.Method = api.POST
	return
}

func CreateCustomerDropResponse() (response *api.BaseResponse[CustomerDropResponse]) {
	response = api.CreateResponse[CustomerDropResponse](&CustomerDropResponse{})
	return
}
