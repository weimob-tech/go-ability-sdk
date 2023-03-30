package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembershipBatchInsertUserTag
// @id 1408
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量添加用户标签关系)
func (client *Client) MembershipBatchInsertUserTag(request *MembershipBatchInsertUserTagRequest) (response *MembershipBatchInsertUserTagResponse, err error) {
	rpcResponse := CreateMembershipBatchInsertUserTagResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembershipBatchInsertUserTagRequest struct {
	*api.BaseRequest

	TagIds []int `json:"tagIds,omitempty"`
	Wids   []int `json:"wids,omitempty"`
}

type MembershipBatchInsertUserTagResponse struct {
}

func CreateMembershipBatchInsertUserTagRequest() (request *MembershipBatchInsertUserTagRequest) {
	request = &MembershipBatchInsertUserTagRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "membership/batchInsertUserTag", "api")
	request.Method = api.POST
	return
}

func CreateMembershipBatchInsertUserTagResponse() (response *api.BaseResponse[MembershipBatchInsertUserTagResponse]) {
	response = api.CreateResponse[MembershipBatchInsertUserTagResponse](&MembershipBatchInsertUserTagResponse{})
	return
}
