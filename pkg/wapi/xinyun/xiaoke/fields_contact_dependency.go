package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FieldsContactDependency
// @id 1748
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询联系人设置级联关系的字段)
func (client *Client) FieldsContactDependency(request *FieldsContactDependencyRequest) (response *FieldsContactDependencyResponse, err error) {
	rpcResponse := CreateFieldsContactDependencyResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FieldsContactDependencyRequest struct {
	*api.BaseRequest

	Wid int64 `json:"wid,omitempty"`
}

type FieldsContactDependencyResponse struct {
}

func CreateFieldsContactDependencyRequest() (request *FieldsContactDependencyRequest) {
	request = &FieldsContactDependencyRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "fields/contactDependency", "api")
	request.Method = api.POST
	return
}

func CreateFieldsContactDependencyResponse() (response *api.BaseResponse[FieldsContactDependencyResponse]) {
	response = api.CreateResponse[FieldsContactDependencyResponse](&FieldsContactDependencyResponse{})
	return
}
