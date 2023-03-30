package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FieldsContact
// @id 1746
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询联系人字段列表)
func (client *Client) FieldsContact(request *FieldsContactRequest) (response *FieldsContactResponse, err error) {
	rpcResponse := CreateFieldsContactResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FieldsContactRequest struct {
	*api.BaseRequest

	Wid int64 `json:"wid,omitempty"`
}

type FieldsContactResponse struct {
}

func CreateFieldsContactRequest() (request *FieldsContactRequest) {
	request = &FieldsContactRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "fields/contact", "api")
	request.Method = api.POST
	return
}

func CreateFieldsContactResponse() (response *api.BaseResponse[FieldsContactResponse]) {
	response = api.CreateResponse[FieldsContactResponse](&FieldsContactResponse{})
	return
}
