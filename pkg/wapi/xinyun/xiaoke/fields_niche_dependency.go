package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FieldsNicheDependency
// @id 1744
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询设置级联关系的字段)
func (client *Client) FieldsNicheDependency(request *FieldsNicheDependencyRequest) (response *FieldsNicheDependencyResponse, err error) {
	rpcResponse := CreateFieldsNicheDependencyResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FieldsNicheDependencyRequest struct {
	*api.BaseRequest

	Wid int64 `json:"wid,omitempty"`
}

type FieldsNicheDependencyResponse struct {
}

func CreateFieldsNicheDependencyRequest() (request *FieldsNicheDependencyRequest) {
	request = &FieldsNicheDependencyRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "fields/nicheDependency", "api")
	request.Method = api.POST
	return
}

func CreateFieldsNicheDependencyResponse() (response *api.BaseResponse[FieldsNicheDependencyResponse]) {
	response = api.CreateResponse[FieldsNicheDependencyResponse](&FieldsNicheDependencyResponse{})
	return
}
