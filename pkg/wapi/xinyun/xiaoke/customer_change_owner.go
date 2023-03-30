package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerChangeOwner
// @id 1720
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=变更客户所属人)
func (client *Client) CustomerChangeOwner(request *CustomerChangeOwnerRequest) (response *CustomerChangeOwnerResponse, err error) {
	rpcResponse := CreateCustomerChangeOwnerResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerChangeOwnerRequest struct {
	*api.BaseRequest

	Owner int64  `json:"owner,omitempty"`
	Key   string `json:"key,omitempty"`
	Wid   int64  `json:"wid,omitempty"`
}

type CustomerChangeOwnerResponse struct {
}

func CreateCustomerChangeOwnerRequest() (request *CustomerChangeOwnerRequest) {
	request = &CustomerChangeOwnerRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "customer/changeOwner", "api")
	request.Method = api.POST
	return
}

func CreateCustomerChangeOwnerResponse() (response *api.BaseResponse[CustomerChangeOwnerResponse]) {
	response = api.CreateResponse[CustomerChangeOwnerResponse](&CustomerChangeOwnerResponse{})
	return
}
