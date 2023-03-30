package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ContactDetail
// @id 1731
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询联系人详情)
func (client *Client) ContactDetail(request *ContactDetailRequest) (response *ContactDetailResponse, err error) {
	rpcResponse := CreateContactDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ContactDetailRequest struct {
	*api.BaseRequest

	Wid int    `json:"wid,omitempty"`
	Key string `json:"key,omitempty"`
}

type ContactDetailResponse struct {
}

func CreateContactDetailRequest() (request *ContactDetailRequest) {
	request = &ContactDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "contact/detail", "api")
	request.Method = api.POST
	return
}

func CreateContactDetailResponse() (response *api.BaseResponse[ContactDetailResponse]) {
	response = api.CreateResponse[ContactDetailResponse](&ContactDetailResponse{})
	return
}
