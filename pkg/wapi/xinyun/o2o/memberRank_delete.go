package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberRankDelete
// @id 87
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除会员等级)
func (client *Client) MemberRankDelete(request *MemberRankDeleteRequest) (response *MemberRankDeleteResponse, err error) {
	rpcResponse := CreateMemberRankDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberRankDeleteRequest struct {
	*api.BaseRequest

	MerchantId int64 `json:"merchantId,omitempty"`
	RankId     int64 `json:"rankId,omitempty"`
}

type MemberRankDeleteResponse struct {
}

func CreateMemberRankDeleteRequest() (request *MemberRankDeleteRequest) {
	request = &MemberRankDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "memberRank/delete", "api")
	request.Method = api.POST
	return
}

func CreateMemberRankDeleteResponse() (response *api.BaseResponse[MemberRankDeleteResponse]) {
	response = api.CreateResponse[MemberRankDeleteResponse](&MemberRankDeleteResponse{})
	return
}
