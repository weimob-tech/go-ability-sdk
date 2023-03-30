package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberRankAdd
// @id 85
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增会员等级)
func (client *Client) MemberRankAdd(request *MemberRankAddRequest) (response *MemberRankAddResponse, err error) {
	rpcResponse := CreateMemberRankAddResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberRankAddRequest struct {
	*api.BaseRequest

	MerchantId      int64  `json:"merchantId,omitempty"`
	ExpType         int    `json:"expType,omitempty"`
	RankDiscount    int    `json:"rankDiscount,omitempty"`
	RankName        string `json:"rankName,omitempty"`
	GrowthDownLimit int64  `json:"growthDownLimit,omitempty"`
	GrowthUpLimit   int64  `json:"growthUpLimit,omitempty"`
	ExpNum          int    `json:"expNum,omitempty"`
	SubtractPoints  int64  `json:"subtractPoints,omitempty"`
}

type MemberRankAddResponse struct {
}

func CreateMemberRankAddRequest() (request *MemberRankAddRequest) {
	request = &MemberRankAddRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "memberRank/add", "api")
	request.Method = api.POST
	return
}

func CreateMemberRankAddResponse() (response *api.BaseResponse[MemberRankAddResponse]) {
	response = api.CreateResponse[MemberRankAddResponse](&MemberRankAddResponse{})
	return
}
