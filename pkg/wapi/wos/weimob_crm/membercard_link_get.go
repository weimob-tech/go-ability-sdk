package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembercardLinkGet
// @id 1511
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1511?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取会员卡页面链接)
func (client *Client) MembercardLinkGet(request *MembercardLinkGetRequest) (response *MembercardLinkGetResponse, err error) {
	rpcResponse := CreateMembercardLinkGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembercardLinkGetRequest struct {
	*api.BaseRequest

	CardType              int64 `json:"cardType,omitempty"`
	MembershipPlanId      int64 `json:"membershipPlanId,omitempty"`
	MembershipCardScene   int64 `json:"membershipCardScene,omitempty"`
	MembershipCardSceneId int64 `json:"membershipCardSceneId,omitempty"`
	Vid                   int64 `json:"vid,omitempty"`
}

type MembercardLinkGetResponse struct {
	Url      string `json:"url,omitempty"`
	ShortUrl string `json:"shortUrl,omitempty"`
}

func CreateMembercardLinkGetRequest() (request *MembercardLinkGetRequest) {
	request = &MembercardLinkGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "membercard/link/get", "apigw")
	request.Method = api.POST
	return
}

func CreateMembercardLinkGetResponse() (response *api.BaseResponse[MembercardLinkGetResponse]) {
	response = api.CreateResponse[MembercardLinkGetResponse](&MembercardLinkGetResponse{})
	return
}
