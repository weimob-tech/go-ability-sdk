package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberOpenPlatformServiceModifyMemberGrowthValue
// @id 166
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改会员成长值)
func (client *Client) MemberOpenPlatformServiceModifyMemberGrowthValue(request *MemberOpenPlatformServiceModifyMemberGrowthValueRequest) (response *MemberOpenPlatformServiceModifyMemberGrowthValueResponse, err error) {
	rpcResponse := CreateMemberOpenPlatformServiceModifyMemberGrowthValueResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberOpenPlatformServiceModifyMemberGrowthValueRequest struct {
	*api.BaseRequest

	CustomerNos     []string `json:"customer_nos,omitempty"`
	Growthvalue     int      `json:"growthvalue,omitempty"`
	TransactionMode int      `json:"transaction_mode,omitempty"`
	Remark          string   `json:"remark,omitempty"`
}

type MemberOpenPlatformServiceModifyMemberGrowthValueResponse struct {
}

func CreateMemberOpenPlatformServiceModifyMemberGrowthValueRequest() (request *MemberOpenPlatformServiceModifyMemberGrowthValueRequest) {
	request = &MemberOpenPlatformServiceModifyMemberGrowthValueRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "MemberOpenPlatformService/ModifyMemberGrowthValue", "api")
	request.Method = api.POST
	return
}

func CreateMemberOpenPlatformServiceModifyMemberGrowthValueResponse() (response *api.BaseResponse[MemberOpenPlatformServiceModifyMemberGrowthValueResponse]) {
	response = api.CreateResponse[MemberOpenPlatformServiceModifyMemberGrowthValueResponse](&MemberOpenPlatformServiceModifyMemberGrowthValueResponse{})
	return
}
