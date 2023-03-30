package kld

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// KLDMemberCardRechargeMember
// @id 330
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=会员线下充值)
func (client *Client) KLDMemberCardRechargeMember(request *KLDMemberCardRechargeMemberRequest) (response *KLDMemberCardRechargeMemberResponse, err error) {
	rpcResponse := CreateKLDMemberCardRechargeMemberResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type KLDMemberCardRechargeMemberRequest struct {
	*api.BaseRequest

	MemberCardNo string `json:"memberCardNo,omitempty"`
	Amount       string `json:"amount,omitempty"`
	StoreId      int64  `json:"storeId,omitempty"`
	Password     string `json:"password,omitempty"`
	Operator     string `json:"operator,omitempty"`
	Comment      string `json:"comment,omitempty"`
}

type KLDMemberCardRechargeMemberResponse struct {
}

func CreateKLDMemberCardRechargeMemberRequest() (request *KLDMemberCardRechargeMemberRequest) {
	request = &KLDMemberCardRechargeMemberRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("kld", "1_0", "KLDMemberCard/RechargeMember", "api")
	request.Method = api.POST
	return
}

func CreateKLDMemberCardRechargeMemberResponse() (response *api.BaseResponse[KLDMemberCardRechargeMemberResponse]) {
	response = api.CreateResponse[KLDMemberCardRechargeMemberResponse](&KLDMemberCardRechargeMemberResponse{})
	return
}
