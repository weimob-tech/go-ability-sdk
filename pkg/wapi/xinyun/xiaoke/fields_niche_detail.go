package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FieldsNicheDetail
// @id 1743
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询商机字段校验规则)
func (client *Client) FieldsNicheDetail(request *FieldsNicheDetailRequest) (response *FieldsNicheDetailResponse, err error) {
	rpcResponse := CreateFieldsNicheDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FieldsNicheDetailRequest struct {
	*api.BaseRequest

	Wid         int64  `json:"wid,omitempty"`
	PropertyKey string `json:"propertyKey,omitempty"`
}

type FieldsNicheDetailResponse struct {
}

func CreateFieldsNicheDetailRequest() (request *FieldsNicheDetailRequest) {
	request = &FieldsNicheDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "fields/nicheDetail", "api")
	request.Method = api.POST
	return
}

func CreateFieldsNicheDetailResponse() (response *api.BaseResponse[FieldsNicheDetailResponse]) {
	response = api.CreateResponse[FieldsNicheDetailResponse](&FieldsNicheDetailResponse{})
	return
}
