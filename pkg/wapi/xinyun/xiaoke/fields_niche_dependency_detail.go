package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FieldsNicheDependencyDetail
// @id 1745
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询商机级联字段的依赖关系)
func (client *Client) FieldsNicheDependencyDetail(request *FieldsNicheDependencyDetailRequest) (response *FieldsNicheDependencyDetailResponse, err error) {
	rpcResponse := CreateFieldsNicheDependencyDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FieldsNicheDependencyDetailRequest struct {
	*api.BaseRequest

	Wid          int64 `json:"wid,omitempty"`
	DependencyId int   `json:"dependencyId,omitempty"`
}

type FieldsNicheDependencyDetailResponse struct {
}

func CreateFieldsNicheDependencyDetailRequest() (request *FieldsNicheDependencyDetailRequest) {
	request = &FieldsNicheDependencyDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "fields/nicheDependencyDetail", "api")
	request.Method = api.POST
	return
}

func CreateFieldsNicheDependencyDetailResponse() (response *api.BaseResponse[FieldsNicheDependencyDetailResponse]) {
	response = api.CreateResponse[FieldsNicheDependencyDetailResponse](&FieldsNicheDependencyDetailResponse{})
	return
}
