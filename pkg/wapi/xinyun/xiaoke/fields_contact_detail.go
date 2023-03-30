package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FieldsContactDetail
// @id 1747
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询联系人字段校验规则)
func (client *Client) FieldsContactDetail(request *FieldsContactDetailRequest) (response *FieldsContactDetailResponse, err error) {
	rpcResponse := CreateFieldsContactDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FieldsContactDetailRequest struct {
	*api.BaseRequest

	Wid         int64  `json:"wid,omitempty"`
	PropertyKey string `json:"propertyKey,omitempty"`
}

type FieldsContactDetailResponse struct {
}

func CreateFieldsContactDetailRequest() (request *FieldsContactDetailRequest) {
	request = &FieldsContactDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "fields/contactDetail", "api")
	request.Method = api.POST
	return
}

func CreateFieldsContactDetailResponse() (response *api.BaseResponse[FieldsContactDetailResponse]) {
	response = api.CreateResponse[FieldsContactDetailResponse](&FieldsContactDetailResponse{})
	return
}
