package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembershipCreateTag
// @id 1301
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=创建手动标签)
func (client *Client) MembershipCreateTag(request *MembershipCreateTagRequest) (response *MembershipCreateTagResponse, err error) {
	rpcResponse := CreateMembershipCreateTagResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembershipCreateTagRequest struct {
	*api.BaseRequest

	TagName string `json:"tagName,omitempty"`
}

type MembershipCreateTagResponse struct {
}

func CreateMembershipCreateTagRequest() (request *MembershipCreateTagRequest) {
	request = &MembershipCreateTagRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "membership/createTag", "api")
	request.Method = api.POST
	return
}

func CreateMembershipCreateTagResponse() (response *api.BaseResponse[MembershipCreateTagResponse]) {
	response = api.CreateResponse[MembershipCreateTagResponse](&MembershipCreateTagResponse{})
	return
}
