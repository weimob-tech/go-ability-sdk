package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembershipUpdateCardRank
// @id 1962
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新会员卡等级)
func (client *Client) MembershipUpdateCardRank(request *MembershipUpdateCardRankRequest) (response *MembershipUpdateCardRankResponse, err error) {
	rpcResponse := CreateMembershipUpdateCardRankResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembershipUpdateCardRankRequest struct {
	*api.BaseRequest

	RequestVo MembershipUpdateCardRankRequestRequestVo `json:"requestVo,omitempty"`
	Params    MembershipUpdateCardRankRequestParams    `json:"params,omitempty"`
	Sign      string                                   `json:"sign,omitempty"`
	Timestamp string                                   `json:"timestamp,omitempty"`
}

type MembershipUpdateCardRankResponse struct {
}

func CreateMembershipUpdateCardRankRequest() (request *MembershipUpdateCardRankRequest) {
	request = &MembershipUpdateCardRankRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "membership/updateCardRank", "api")
	request.Method = api.POST
	return
}

func CreateMembershipUpdateCardRankResponse() (response *api.BaseResponse[MembershipUpdateCardRankResponse]) {
	response = api.CreateResponse[MembershipUpdateCardRankResponse](&MembershipUpdateCardRankResponse{})
	return
}

type MembershipUpdateCardRankRequestRequestVo struct {
	MerchantId        int64 `json:"merchantId,omitempty"`
	BosId             int64 `json:"bosId,omitempty"`
	ProductInstanceId int64 `json:"productInstanceId,omitempty"`
	ProductId         int64 `json:"productId,omitempty"`
}

type MembershipUpdateCardRankRequestParams struct {
}
