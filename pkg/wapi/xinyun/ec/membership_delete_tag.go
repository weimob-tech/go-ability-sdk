package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembershipDeleteTag
// @id 1303
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除手动标签)
func (client *Client) MembershipDeleteTag(request *MembershipDeleteTagRequest) (response *MembershipDeleteTagResponse, err error) {
	rpcResponse := CreateMembershipDeleteTagResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembershipDeleteTagRequest struct {
	*api.BaseRequest

	TagId int64 `json:"tagId,omitempty"`
}

type MembershipDeleteTagResponse struct {
}

func CreateMembershipDeleteTagRequest() (request *MembershipDeleteTagRequest) {
	request = &MembershipDeleteTagRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "membership/deleteTag", "api")
	request.Method = api.POST
	return
}

func CreateMembershipDeleteTagResponse() (response *api.BaseResponse[MembershipDeleteTagResponse]) {
	response = api.CreateResponse[MembershipDeleteTagResponse](&MembershipDeleteTagResponse{})
	return
}
