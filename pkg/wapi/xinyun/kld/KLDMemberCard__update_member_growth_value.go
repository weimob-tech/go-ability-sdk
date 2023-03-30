package kld

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// KLDMemberCardUpdateMemberGrowthValue
// @id 235
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=增加/减少会员成长值)
func (client *Client) KLDMemberCardUpdateMemberGrowthValue(request *KLDMemberCardUpdateMemberGrowthValueRequest) (response *KLDMemberCardUpdateMemberGrowthValueResponse, err error) {
	rpcResponse := CreateKLDMemberCardUpdateMemberGrowthValueResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type KLDMemberCardUpdateMemberGrowthValueRequest struct {
	*api.BaseRequest

	MemberCardNo string `json:"memberCardNo,omitempty"`
	ChangeValue  int    `json:"changeValue,omitempty"`
}

type KLDMemberCardUpdateMemberGrowthValueResponse struct {
}

func CreateKLDMemberCardUpdateMemberGrowthValueRequest() (request *KLDMemberCardUpdateMemberGrowthValueRequest) {
	request = &KLDMemberCardUpdateMemberGrowthValueRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("kld", "1_0", "KLDMemberCard/UpdateMemberGrowthValue", "api")
	request.Method = api.POST
	return
}

func CreateKLDMemberCardUpdateMemberGrowthValueResponse() (response *api.BaseResponse[KLDMemberCardUpdateMemberGrowthValueResponse]) {
	response = api.CreateResponse[KLDMemberCardUpdateMemberGrowthValueResponse](&KLDMemberCardUpdateMemberGrowthValueResponse{})
	return
}
