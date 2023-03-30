package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NewmembercardUpdateMemberGrowthValue
// @id 114
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=增加减少成长值接口)
func (client *Client) NewmembercardUpdateMemberGrowthValue(request *NewmembercardUpdateMemberGrowthValueRequest) (response *NewmembercardUpdateMemberGrowthValueResponse, err error) {
	rpcResponse := CreateNewmembercardUpdateMemberGrowthValueResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NewmembercardUpdateMemberGrowthValueRequest struct {
	*api.BaseRequest

	MemberCardNo string `json:"memberCardNo,omitempty"`
	ChangeValue  int    `json:"changeValue,omitempty"`
}

type NewmembercardUpdateMemberGrowthValueResponse struct {
}

func CreateNewmembercardUpdateMemberGrowthValueRequest() (request *NewmembercardUpdateMemberGrowthValueRequest) {
	request = &NewmembercardUpdateMemberGrowthValueRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "newmembercard/UpdateMemberGrowthValue", "api")
	request.Method = api.POST
	return
}

func CreateNewmembercardUpdateMemberGrowthValueResponse() (response *api.BaseResponse[NewmembercardUpdateMemberGrowthValueResponse]) {
	response = api.CreateResponse[NewmembercardUpdateMemberGrowthValueResponse](&NewmembercardUpdateMemberGrowthValueResponse{})
	return
}
