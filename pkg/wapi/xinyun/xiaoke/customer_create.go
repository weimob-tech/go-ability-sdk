package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerCreate
// @id 1713
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新建客户)
func (client *Client) CustomerCreate(request *CustomerCreateRequest) (response *CustomerCreateResponse, err error) {
	rpcResponse := CreateCustomerCreateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerCreateRequest struct {
	*api.BaseRequest

	CustomerInfo CustomerCreateRequestCustomerInfo `json:"customerInfo,omitempty"`
	PublicSeaId  string                            `json:"publicSeaId,omitempty"`
	Wid          int64                             `json:"wid,omitempty"`
}

type CustomerCreateResponse struct {
}

func CreateCustomerCreateRequest() (request *CustomerCreateRequest) {
	request = &CustomerCreateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "customer/create", "api")
	request.Method = api.POST
	return
}

func CreateCustomerCreateResponse() (response *api.BaseResponse[CustomerCreateResponse]) {
	response = api.CreateResponse[CustomerCreateResponse](&CustomerCreateResponse{})
	return
}

type CustomerCreateRequestCustomerInfo struct {
}
