package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberRankUpdate
// @id 86
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=对已经设置的会员等级进行更新)
func (client *Client) MemberRankUpdate(request *MemberRankUpdateRequest) (response *MemberRankUpdateResponse, err error) {
	rpcResponse := CreateMemberRankUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberRankUpdateRequest struct {
	*api.BaseRequest

	MerchantId      int64  `json:"merchantId,omitempty"`
	RankId          int64  `json:"rankId,omitempty"`
	ExpType         int    `json:"expType,omitempty"`
	RankDiscount    int    `json:"rankDiscount,omitempty"`
	RankName        string `json:"rankName,omitempty"`
	GrowthDownLimit int64  `json:"growthDownLimit,omitempty"`
	GrowthUpLimit   int64  `json:"growthUpLimit,omitempty"`
	ExpNum          int    `json:"expNum,omitempty"`
	SubtractPoints  int64  `json:"subtractPoints,omitempty"`
}

type MemberRankUpdateResponse struct {
}

func CreateMemberRankUpdateRequest() (request *MemberRankUpdateRequest) {
	request = &MemberRankUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "memberRank/update", "api")
	request.Method = api.POST
	return
}

func CreateMemberRankUpdateResponse() (response *api.BaseResponse[MemberRankUpdateResponse]) {
	response = api.CreateResponse[MemberRankUpdateResponse](&MemberRankUpdateResponse{})
	return
}
