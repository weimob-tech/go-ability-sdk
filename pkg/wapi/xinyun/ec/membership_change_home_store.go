package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembershipChangeHomeStore
// @id 1335
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=变更用户归属门店)
func (client *Client) MembershipChangeHomeStore(request *MembershipChangeHomeStoreRequest) (response *MembershipChangeHomeStoreResponse, err error) {
	rpcResponse := CreateMembershipChangeHomeStoreResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembershipChangeHomeStoreRequest struct {
	*api.BaseRequest

	HomeStoreId  int64   `json:"homeStoreId,omitempty"`
	ChangeReason string  `json:"changeReason,omitempty"`
	WidList      []int64 `json:"widList,omitempty"`
}

type MembershipChangeHomeStoreResponse struct {
}

func CreateMembershipChangeHomeStoreRequest() (request *MembershipChangeHomeStoreRequest) {
	request = &MembershipChangeHomeStoreRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "membership/changeHomeStore", "api")
	request.Method = api.POST
	return
}

func CreateMembershipChangeHomeStoreResponse() (response *api.BaseResponse[MembershipChangeHomeStoreResponse]) {
	response = api.CreateResponse[MembershipChangeHomeStoreResponse](&MembershipChangeHomeStoreResponse{})
	return
}
