package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembershipQueryUserInfoOpen
// @id 1309
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询客户详情)
func (client *Client) MembershipQueryUserInfoOpen(request *MembershipQueryUserInfoOpenRequest) (response *MembershipQueryUserInfoOpenResponse, err error) {
	rpcResponse := CreateMembershipQueryUserInfoOpenResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembershipQueryUserInfoOpenRequest struct {
	*api.BaseRequest

	Wid     int64 `json:"wid,omitempty"`
	StoreId int64 `json:"storeId,omitempty"`
}

type MembershipQueryUserInfoOpenResponse struct {
}

func CreateMembershipQueryUserInfoOpenRequest() (request *MembershipQueryUserInfoOpenRequest) {
	request = &MembershipQueryUserInfoOpenRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "membership/queryUserInfoOpen", "api")
	request.Method = api.POST
	return
}

func CreateMembershipQueryUserInfoOpenResponse() (response *api.BaseResponse[MembershipQueryUserInfoOpenResponse]) {
	response = api.CreateResponse[MembershipQueryUserInfoOpenResponse](&MembershipQueryUserInfoOpenResponse{})
	return
}
