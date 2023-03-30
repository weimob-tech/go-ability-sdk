package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembershipBatchRemoveTag
// @id 1305
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量移除标签)
func (client *Client) MembershipBatchRemoveTag(request *MembershipBatchRemoveTagRequest) (response *MembershipBatchRemoveTagResponse, err error) {
	rpcResponse := CreateMembershipBatchRemoveTagResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembershipBatchRemoveTagRequest struct {
	*api.BaseRequest

	TagIds []int `json:"tagIds,omitempty"`
	Wids   []int `json:"wids,omitempty"`
}

type MembershipBatchRemoveTagResponse struct {
}

func CreateMembershipBatchRemoveTagRequest() (request *MembershipBatchRemoveTagRequest) {
	request = &MembershipBatchRemoveTagRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "membership/batchRemoveTag", "api")
	request.Method = api.POST
	return
}

func CreateMembershipBatchRemoveTagResponse() (response *api.BaseResponse[MembershipBatchRemoveTagResponse]) {
	response = api.CreateResponse[MembershipBatchRemoveTagResponse](&MembershipBatchRemoveTagResponse{})
	return
}
