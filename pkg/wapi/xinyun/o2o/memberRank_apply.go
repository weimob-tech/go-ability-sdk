package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberRankApply
// @id 88
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=使增加或者修改的会员等级类型立即生效)
func (client *Client) MemberRankApply(request *MemberRankApplyRequest) (response *MemberRankApplyResponse, err error) {
	rpcResponse := CreateMemberRankApplyResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberRankApplyRequest struct {
	*api.BaseRequest

	MerchantId int64 `json:"merchantId,omitempty"`
}

type MemberRankApplyResponse struct {
}

func CreateMemberRankApplyRequest() (request *MemberRankApplyRequest) {
	request = &MemberRankApplyRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "memberRank/apply", "api")
	request.Method = api.POST
	return
}

func CreateMemberRankApplyResponse() (response *api.BaseResponse[MemberRankApplyResponse]) {
	response = api.CreateResponse[MemberRankApplyResponse](&MemberRankApplyResponse{})
	return
}
