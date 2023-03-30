package kld

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// KLDMemberCardResetMemberExpireDateInfo
// @id 220
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=重置会员有效期)
func (client *Client) KLDMemberCardResetMemberExpireDateInfo(request *KLDMemberCardResetMemberExpireDateInfoRequest) (response *KLDMemberCardResetMemberExpireDateInfoResponse, err error) {
	rpcResponse := CreateKLDMemberCardResetMemberExpireDateInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type KLDMemberCardResetMemberExpireDateInfoRequest struct {
	*api.BaseRequest

	MemberCardNo string `json:"memberCardNo,omitempty"`
}

type KLDMemberCardResetMemberExpireDateInfoResponse struct {
}

func CreateKLDMemberCardResetMemberExpireDateInfoRequest() (request *KLDMemberCardResetMemberExpireDateInfoRequest) {
	request = &KLDMemberCardResetMemberExpireDateInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("kld", "1_0", "KLDMemberCard/ResetMemberExpireDateInfo", "api")
	request.Method = api.POST
	return
}

func CreateKLDMemberCardResetMemberExpireDateInfoResponse() (response *api.BaseResponse[KLDMemberCardResetMemberExpireDateInfoResponse]) {
	response = api.CreateResponse[KLDMemberCardResetMemberExpireDateInfoResponse](&KLDMemberCardResetMemberExpireDateInfoResponse{})
	return
}
