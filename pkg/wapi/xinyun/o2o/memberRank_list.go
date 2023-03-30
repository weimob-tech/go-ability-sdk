package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberRankList
// @id 89
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询该商户下面所有的会员等级类型)
func (client *Client) MemberRankList(request *MemberRankListRequest) (response *MemberRankListResponse, err error) {
	rpcResponse := CreateMemberRankListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberRankListRequest struct {
	*api.BaseRequest

	MerchantId int64 `json:"merchantId,omitempty"`
}

type MemberRankListResponse struct {
}

func CreateMemberRankListRequest() (request *MemberRankListRequest) {
	request = &MemberRankListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "memberRank/list", "api")
	request.Method = api.POST
	return
}

func CreateMemberRankListResponse() (response *api.BaseResponse[MemberRankListResponse]) {
	response = api.CreateResponse[MemberRankListResponse](&MemberRankListResponse{})
	return
}
