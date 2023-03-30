package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberOpenPlatformServiceAddOfflineMember
// @id 200
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增线下会员)
func (client *Client) MemberOpenPlatformServiceAddOfflineMember(request *MemberOpenPlatformServiceAddOfflineMemberRequest) (response *MemberOpenPlatformServiceAddOfflineMemberResponse, err error) {
	rpcResponse := CreateMemberOpenPlatformServiceAddOfflineMemberResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberOpenPlatformServiceAddOfflineMemberRequest struct {
	*api.BaseRequest

	Name                   string  `json:"name,omitempty"`
	Phone                  string  `json:"phone,omitempty"`
	Province               string  `json:"province,omitempty"`
	City                   string  `json:"city,omitempty"`
	District               string  `json:"district,omitempty"`
	Address                string  `json:"address,omitempty"`
	Birthday               int64   `json:"birthday,omitempty"`
	Sex                    int     `json:"sex,omitempty"`
	Growthvalue            int     `json:"growthvalue,omitempty"`
	Points                 int     `json:"points,omitempty"`
	TotalPoints            int     `json:"total_points,omitempty"`
	Balance                float64 `json:"balance,omitempty"`
	TotalTransactionAmount float64 `json:"total_transaction_amount,omitempty"`
	Openid                 string  `json:"openid,omitempty"`
}

type MemberOpenPlatformServiceAddOfflineMemberResponse struct {
}

func CreateMemberOpenPlatformServiceAddOfflineMemberRequest() (request *MemberOpenPlatformServiceAddOfflineMemberRequest) {
	request = &MemberOpenPlatformServiceAddOfflineMemberRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "MemberOpenPlatformService/AddOfflineMember", "api")
	request.Method = api.POST
	return
}

func CreateMemberOpenPlatformServiceAddOfflineMemberResponse() (response *api.BaseResponse[MemberOpenPlatformServiceAddOfflineMemberResponse]) {
	response = api.CreateResponse[MemberOpenPlatformServiceAddOfflineMemberResponse](&MemberOpenPlatformServiceAddOfflineMemberResponse{})
	return
}
