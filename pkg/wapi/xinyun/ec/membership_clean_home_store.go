package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembershipCleanHomeStore
// @id 1457
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=解绑用户归属门店)
func (client *Client) MembershipCleanHomeStore(request *MembershipCleanHomeStoreRequest) (response *MembershipCleanHomeStoreResponse, err error) {
	rpcResponse := CreateMembershipCleanHomeStoreResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembershipCleanHomeStoreRequest struct {
	*api.BaseRequest

	Wid    int64  `json:"wid,omitempty"`
	Reason string `json:"reason,omitempty"`
}

type MembershipCleanHomeStoreResponse struct {
}

func CreateMembershipCleanHomeStoreRequest() (request *MembershipCleanHomeStoreRequest) {
	request = &MembershipCleanHomeStoreRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "membership/cleanHomeStore", "api")
	request.Method = api.POST
	return
}

func CreateMembershipCleanHomeStoreResponse() (response *api.BaseResponse[MembershipCleanHomeStoreResponse]) {
	response = api.CreateResponse[MembershipCleanHomeStoreResponse](&MembershipCleanHomeStoreResponse{})
	return
}
