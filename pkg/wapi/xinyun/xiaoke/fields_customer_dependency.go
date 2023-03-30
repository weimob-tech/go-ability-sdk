package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FieldsCustomerDependency
// @id 1740
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询客户设置级联关系的字段)
func (client *Client) FieldsCustomerDependency(request *FieldsCustomerDependencyRequest) (response *FieldsCustomerDependencyResponse, err error) {
	rpcResponse := CreateFieldsCustomerDependencyResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FieldsCustomerDependencyRequest struct {
	*api.BaseRequest

	Wid int64 `json:"wid,omitempty"`
}

type FieldsCustomerDependencyResponse struct {
}

func CreateFieldsCustomerDependencyRequest() (request *FieldsCustomerDependencyRequest) {
	request = &FieldsCustomerDependencyRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "fields/customerDependency", "api")
	request.Method = api.POST
	return
}

func CreateFieldsCustomerDependencyResponse() (response *api.BaseResponse[FieldsCustomerDependencyResponse]) {
	response = api.CreateResponse[FieldsCustomerDependencyResponse](&FieldsCustomerDependencyResponse{})
	return
}
