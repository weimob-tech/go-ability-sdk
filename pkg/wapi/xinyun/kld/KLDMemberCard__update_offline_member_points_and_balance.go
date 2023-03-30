package kld

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// KLDMemberCardUpdateOfflineMemberPointsAndBalance
// @id 335
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新实体店会员积分余额)
func (client *Client) KLDMemberCardUpdateOfflineMemberPointsAndBalance(request *KLDMemberCardUpdateOfflineMemberPointsAndBalanceRequest) (response *KLDMemberCardUpdateOfflineMemberPointsAndBalanceResponse, err error) {
	rpcResponse := CreateKLDMemberCardUpdateOfflineMemberPointsAndBalanceResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type KLDMemberCardUpdateOfflineMemberPointsAndBalanceRequest struct {
	*api.BaseRequest

	Name   string  `json:"name,omitempty"`
	Phone  string  `json:"phone,omitempty"`
	Points int     `json:"points,omitempty"`
	Amount float64 `json:"amount,omitempty"`
}

type KLDMemberCardUpdateOfflineMemberPointsAndBalanceResponse struct {
}

func CreateKLDMemberCardUpdateOfflineMemberPointsAndBalanceRequest() (request *KLDMemberCardUpdateOfflineMemberPointsAndBalanceRequest) {
	request = &KLDMemberCardUpdateOfflineMemberPointsAndBalanceRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("kld", "1_0", "KLDMemberCard/UpdateOfflineMemberPointsAndBalance", "api")
	request.Method = api.POST
	return
}

func CreateKLDMemberCardUpdateOfflineMemberPointsAndBalanceResponse() (response *api.BaseResponse[KLDMemberCardUpdateOfflineMemberPointsAndBalanceResponse]) {
	response = api.CreateResponse[KLDMemberCardUpdateOfflineMemberPointsAndBalanceResponse](&KLDMemberCardUpdateOfflineMemberPointsAndBalanceResponse{})
	return
}
