package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembercardMembercodeGet
// @id 1743
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1743?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询会员码信息)
func (client *Client) MembercardMembercodeGet(request *MembercardMembercodeGetRequest) (response *MembercardMembercodeGetResponse, err error) {
	rpcResponse := CreateMembercardMembercodeGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembercardMembercodeGetRequest struct {
	*api.BaseRequest

	CodeSource  int64  `json:"codeSource,omitempty"`
	Wid         int64  `json:"wid,omitempty"`
	DynamicCode string `json:"dynamicCode,omitempty"`
}

type MembercardMembercodeGetResponse struct {
	Wid                 int64  `json:"wid,omitempty"`
	MembershipPlanId    int64  `json:"membershipPlanId,omitempty"`
	Phone               int64  `json:"phone,omitempty"`
	DynamicExpireAtTime int64  `json:"dynamicExpireAtTime,omitempty"`
	ExpireTime          int64  `json:"expireTime,omitempty"`
	WxCode              string `json:"wxCode,omitempty"`
	CustomCardNo        string `json:"customCardNo,omitempty"`
}

func CreateMembercardMembercodeGetRequest() (request *MembercardMembercodeGetRequest) {
	request = &MembercardMembercodeGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "membercard/membercode/get", "apigw")
	request.Method = api.POST
	return
}

func CreateMembercardMembercodeGetResponse() (response *api.BaseResponse[MembercardMembercodeGetResponse]) {
	response = api.CreateResponse[MembercardMembercodeGetResponse](&MembercardMembercodeGetResponse{})
	return
}
