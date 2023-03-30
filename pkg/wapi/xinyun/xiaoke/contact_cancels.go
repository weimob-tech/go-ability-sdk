package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ContactCancels
// @id 1728
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除联系人)
func (client *Client) ContactCancels(request *ContactCancelsRequest) (response *ContactCancelsResponse, err error) {
	rpcResponse := CreateContactCancelsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ContactCancelsRequest struct {
	*api.BaseRequest

	Keys []string `json:"keys,omitempty"`
	Wid  int      `json:"wid,omitempty"`
}

type ContactCancelsResponse struct {
	Message string `json:"message,omitempty"`
}

func CreateContactCancelsRequest() (request *ContactCancelsRequest) {
	request = &ContactCancelsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "contact/cancels", "api")
	request.Method = api.POST
	return
}

func CreateContactCancelsResponse() (response *api.BaseResponse[ContactCancelsResponse]) {
	response = api.CreateResponse[ContactCancelsResponse](&ContactCancelsResponse{})
	return
}
