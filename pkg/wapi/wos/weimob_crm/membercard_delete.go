package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembercardDelete
// @id 2331
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2331?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除会员卡)
func (client *Client) MembercardDelete(request *MembercardDeleteRequest) (response *MembercardDeleteResponse, err error) {
	rpcResponse := CreateMembercardDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembercardDeleteRequest struct {
	*api.BaseRequest

	Reason           string `json:"reason,omitempty"`
	Wid              int64  `json:"wid,omitempty"`
	OperatorWid      int64  `json:"operatorWid,omitempty"`
	MembershipPlanId int64  `json:"membershipPlanId,omitempty"`
}

type MembercardDeleteResponse struct {
	Status bool `json:"status,omitempty"`
}

func CreateMembercardDeleteRequest() (request *MembercardDeleteRequest) {
	request = &MembercardDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "membercard/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateMembercardDeleteResponse() (response *api.BaseResponse[MembercardDeleteResponse]) {
	response = api.CreateResponse[MembercardDeleteResponse](&MembercardDeleteResponse{})
	return
}
