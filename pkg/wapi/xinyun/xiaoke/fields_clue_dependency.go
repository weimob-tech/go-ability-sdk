package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FieldsClueDependency
// @id 1649
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询设置级联关系列表)
func (client *Client) FieldsClueDependency(request *FieldsClueDependencyRequest) (response *FieldsClueDependencyResponse, err error) {
	rpcResponse := CreateFieldsClueDependencyResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FieldsClueDependencyRequest struct {
	*api.BaseRequest

	Wid int64 `json:"wid,omitempty"`
}

type FieldsClueDependencyResponse struct {
}

func CreateFieldsClueDependencyRequest() (request *FieldsClueDependencyRequest) {
	request = &FieldsClueDependencyRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "fields/clueDependency", "api")
	request.Method = api.POST
	return
}

func CreateFieldsClueDependencyResponse() (response *api.BaseResponse[FieldsClueDependencyResponse]) {
	response = api.CreateResponse[FieldsClueDependencyResponse](&FieldsClueDependencyResponse{})
	return
}
