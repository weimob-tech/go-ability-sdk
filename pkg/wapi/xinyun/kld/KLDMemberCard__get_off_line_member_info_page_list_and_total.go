package kld

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// KLDMemberCardGetOffLineMemberInfoPageListAndTotal
// @id 222
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取实体店会员列表)
func (client *Client) KLDMemberCardGetOffLineMemberInfoPageListAndTotal(request *KLDMemberCardGetOffLineMemberInfoPageListAndTotalRequest) (response *KLDMemberCardGetOffLineMemberInfoPageListAndTotalResponse, err error) {
	rpcResponse := CreateKLDMemberCardGetOffLineMemberInfoPageListAndTotalResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type KLDMemberCardGetOffLineMemberInfoPageListAndTotalRequest struct {
	*api.BaseRequest

	Name      string `json:"name,omitempty"`
	Phone     string `json:"phone,omitempty"`
	PageIndex int    `json:"pageIndex,omitempty"`
	PageSize  int    `json:"pageSize,omitempty"`
}

type KLDMemberCardGetOffLineMemberInfoPageListAndTotalResponse struct {
}

func CreateKLDMemberCardGetOffLineMemberInfoPageListAndTotalRequest() (request *KLDMemberCardGetOffLineMemberInfoPageListAndTotalRequest) {
	request = &KLDMemberCardGetOffLineMemberInfoPageListAndTotalRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("kld", "1_0", "KLDMemberCard/GetOffLineMemberInfoPageListAndTotal", "api")
	request.Method = api.POST
	return
}

func CreateKLDMemberCardGetOffLineMemberInfoPageListAndTotalResponse() (response *api.BaseResponse[KLDMemberCardGetOffLineMemberInfoPageListAndTotalResponse]) {
	response = api.CreateResponse[KLDMemberCardGetOffLineMemberInfoPageListAndTotalResponse](&KLDMemberCardGetOffLineMemberInfoPageListAndTotalResponse{})
	return
}
