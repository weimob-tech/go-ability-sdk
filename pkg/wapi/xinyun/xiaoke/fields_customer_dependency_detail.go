package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FieldsCustomerDependencyDetail
// @id 1741
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询客户级联字段的依赖关系)
func (client *Client) FieldsCustomerDependencyDetail(request *FieldsCustomerDependencyDetailRequest) (response *FieldsCustomerDependencyDetailResponse, err error) {
	rpcResponse := CreateFieldsCustomerDependencyDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FieldsCustomerDependencyDetailRequest struct {
	*api.BaseRequest

	Wid          int64 `json:"wid,omitempty"`
	DependencyId int   `json:"dependencyId,omitempty"`
}

type FieldsCustomerDependencyDetailResponse struct {
}

func CreateFieldsCustomerDependencyDetailRequest() (request *FieldsCustomerDependencyDetailRequest) {
	request = &FieldsCustomerDependencyDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "fields/customerDependencyDetail", "api")
	request.Method = api.POST
	return
}

func CreateFieldsCustomerDependencyDetailResponse() (response *api.BaseResponse[FieldsCustomerDependencyDetailResponse]) {
	response = api.CreateResponse[FieldsCustomerDependencyDetailResponse](&FieldsCustomerDependencyDetailResponse{})
	return
}
