package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NewmembercardResetMemberExpireDateInfo
// @id 112
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=重置会员有效期)
func (client *Client) NewmembercardResetMemberExpireDateInfo(request *NewmembercardResetMemberExpireDateInfoRequest) (response *NewmembercardResetMemberExpireDateInfoResponse, err error) {
	rpcResponse := CreateNewmembercardResetMemberExpireDateInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NewmembercardResetMemberExpireDateInfoRequest struct {
	*api.BaseRequest

	MemberCardNo string `json:"memberCardNo,omitempty"`
}

type NewmembercardResetMemberExpireDateInfoResponse struct {
}

func CreateNewmembercardResetMemberExpireDateInfoRequest() (request *NewmembercardResetMemberExpireDateInfoRequest) {
	request = &NewmembercardResetMemberExpireDateInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "newmembercard/ResetMemberExpireDateInfo", "api")
	request.Method = api.POST
	return
}

func CreateNewmembercardResetMemberExpireDateInfoResponse() (response *api.BaseResponse[NewmembercardResetMemberExpireDateInfoResponse]) {
	response = api.CreateResponse[NewmembercardResetMemberExpireDateInfoResponse](&NewmembercardResetMemberExpireDateInfoResponse{})
	return
}
