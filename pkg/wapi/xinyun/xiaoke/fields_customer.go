package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FieldsCustomer
// @id 1738
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询客户字段列表)
func (client *Client) FieldsCustomer(request *FieldsCustomerRequest) (response *FieldsCustomerResponse, err error) {
	rpcResponse := CreateFieldsCustomerResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FieldsCustomerRequest struct {
	*api.BaseRequest

	Wid int64 `json:"wid,omitempty"`
}

type FieldsCustomerResponse struct {
}

func CreateFieldsCustomerRequest() (request *FieldsCustomerRequest) {
	request = &FieldsCustomerRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "fields/customer", "api")
	request.Method = api.POST
	return
}

func CreateFieldsCustomerResponse() (response *api.BaseResponse[FieldsCustomerResponse]) {
	response = api.CreateResponse[FieldsCustomerResponse](&FieldsCustomerResponse{})
	return
}
