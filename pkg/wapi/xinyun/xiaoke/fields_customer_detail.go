package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FieldsCustomerDetail
// @id 1739
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询客户字段校验规则)
func (client *Client) FieldsCustomerDetail(request *FieldsCustomerDetailRequest) (response *FieldsCustomerDetailResponse, err error) {
	rpcResponse := CreateFieldsCustomerDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FieldsCustomerDetailRequest struct {
	*api.BaseRequest

	Wid         int64  `json:"wid,omitempty"`
	PropertyKey string `json:"propertyKey,omitempty"`
}

type FieldsCustomerDetailResponse struct {
}

func CreateFieldsCustomerDetailRequest() (request *FieldsCustomerDetailRequest) {
	request = &FieldsCustomerDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "fields/customerDetail", "api")
	request.Method = api.POST
	return
}

func CreateFieldsCustomerDetailResponse() (response *api.BaseResponse[FieldsCustomerDetailResponse]) {
	response = api.CreateResponse[FieldsCustomerDetailResponse](&FieldsCustomerDetailResponse{})
	return
}
