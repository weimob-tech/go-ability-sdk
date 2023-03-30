package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembershipEditTag
// @id 1302
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=编辑手动标签)
func (client *Client) MembershipEditTag(request *MembershipEditTagRequest) (response *MembershipEditTagResponse, err error) {
	rpcResponse := CreateMembershipEditTagResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembershipEditTagRequest struct {
	*api.BaseRequest

	TagId   int64  `json:"tagId,omitempty"`
	TagName string `json:"tagName,omitempty"`
}

type MembershipEditTagResponse struct {
}

func CreateMembershipEditTagRequest() (request *MembershipEditTagRequest) {
	request = &MembershipEditTagRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "membership/editTag", "api")
	request.Method = api.POST
	return
}

func CreateMembershipEditTagResponse() (response *api.BaseResponse[MembershipEditTagResponse]) {
	response = api.CreateResponse[MembershipEditTagResponse](&MembershipEditTagResponse{})
	return
}
