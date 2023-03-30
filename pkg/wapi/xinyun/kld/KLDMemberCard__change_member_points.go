package kld

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// KLDMemberCardChangeMemberPoints
// @id 230
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=增加/减少积分)
func (client *Client) KLDMemberCardChangeMemberPoints(request *KLDMemberCardChangeMemberPointsRequest) (response *KLDMemberCardChangeMemberPointsResponse, err error) {
	rpcResponse := CreateKLDMemberCardChangeMemberPointsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type KLDMemberCardChangeMemberPointsRequest struct {
	*api.BaseRequest

	MemberCardNo       string `json:"memberCardNo,omitempty"`
	Points             int    `json:"points,omitempty"`
	StoreId            int64  `json:"storeId,omitempty"`
	Title              string `json:"title,omitempty"`
	Remark             string `json:"remark,omitempty"`
	IsAboutGrowthValue bool   `json:"isAboutGrowthValue,omitempty"`
	Operator           string `json:"operator,omitempty"`
	PointsPayType      int    `json:"pointsPayType,omitempty"`
}

type KLDMemberCardChangeMemberPointsResponse struct {
}

func CreateKLDMemberCardChangeMemberPointsRequest() (request *KLDMemberCardChangeMemberPointsRequest) {
	request = &KLDMemberCardChangeMemberPointsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("kld", "1_0", "KLDMemberCard/ChangeMemberPoints", "api")
	request.Method = api.POST
	return
}

func CreateKLDMemberCardChangeMemberPointsResponse() (response *api.BaseResponse[KLDMemberCardChangeMemberPointsResponse]) {
	response = api.CreateResponse[KLDMemberCardChangeMemberPointsResponse](&KLDMemberCardChangeMemberPointsResponse{})
	return
}
