package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FieldsClueDependencyDetail
// @id 1648
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询级联字段依赖关系)
func (client *Client) FieldsClueDependencyDetail(request *FieldsClueDependencyDetailRequest) (response *FieldsClueDependencyDetailResponse, err error) {
	rpcResponse := CreateFieldsClueDependencyDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FieldsClueDependencyDetailRequest struct {
	*api.BaseRequest

	Wid          int64 `json:"wid,omitempty"`
	DependencyId int   `json:"dependencyId,omitempty"`
}

type FieldsClueDependencyDetailResponse struct {
}

func CreateFieldsClueDependencyDetailRequest() (request *FieldsClueDependencyDetailRequest) {
	request = &FieldsClueDependencyDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "fields/clueDependencyDetail", "api")
	request.Method = api.POST
	return
}

func CreateFieldsClueDependencyDetailResponse() (response *api.BaseResponse[FieldsClueDependencyDetailResponse]) {
	response = api.CreateResponse[FieldsClueDependencyDetailResponse](&FieldsClueDependencyDetailResponse{})
	return
}
