package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FieldsContactDependencyDetail
// @id 1749
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询俩新人级联字段的依赖关系)
func (client *Client) FieldsContactDependencyDetail(request *FieldsContactDependencyDetailRequest) (response *FieldsContactDependencyDetailResponse, err error) {
	rpcResponse := CreateFieldsContactDependencyDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FieldsContactDependencyDetailRequest struct {
	*api.BaseRequest

	Wid          int64 `json:"wid,omitempty"`
	DependencyId int   `json:"dependencyId,omitempty"`
}

type FieldsContactDependencyDetailResponse struct {
}

func CreateFieldsContactDependencyDetailRequest() (request *FieldsContactDependencyDetailRequest) {
	request = &FieldsContactDependencyDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "fields/contactDependencyDetail", "api")
	request.Method = api.POST
	return
}

func CreateFieldsContactDependencyDetailResponse() (response *api.BaseResponse[FieldsContactDependencyDetailResponse]) {
	response = api.CreateResponse[FieldsContactDependencyDetailResponse](&FieldsContactDependencyDetailResponse{})
	return
}
