package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FieldsClueDetail
// @id 1647
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询线索字段验证规则)
func (client *Client) FieldsClueDetail(request *FieldsClueDetailRequest) (response *FieldsClueDetailResponse, err error) {
	rpcResponse := CreateFieldsClueDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FieldsClueDetailRequest struct {
	*api.BaseRequest

	Wid         int64  `json:"wid,omitempty"`
	PropertyKey string `json:"propertyKey,omitempty"`
}

type FieldsClueDetailResponse struct {
}

func CreateFieldsClueDetailRequest() (request *FieldsClueDetailRequest) {
	request = &FieldsClueDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "fields/clueDetail", "api")
	request.Method = api.POST
	return
}

func CreateFieldsClueDetailResponse() (response *api.BaseResponse[FieldsClueDetailResponse]) {
	response = api.CreateResponse[FieldsClueDetailResponse](&FieldsClueDetailResponse{})
	return
}
