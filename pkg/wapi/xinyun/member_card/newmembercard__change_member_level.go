package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NewmembercardChangeMemberLevel
// @id 111
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改会员等级)
func (client *Client) NewmembercardChangeMemberLevel(request *NewmembercardChangeMemberLevelRequest) (response *NewmembercardChangeMemberLevelResponse, err error) {
	rpcResponse := CreateNewmembercardChangeMemberLevelResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NewmembercardChangeMemberLevelRequest struct {
	*api.BaseRequest

	NewLevelNo       int      `json:"newLevelNo,omitempty"`
	Operator         string   `json:"operator,omitempty"`
	ListMemberCardNo []string `json:"listMemberCardNo,omitempty"`
}

type NewmembercardChangeMemberLevelResponse struct {
}

func CreateNewmembercardChangeMemberLevelRequest() (request *NewmembercardChangeMemberLevelRequest) {
	request = &NewmembercardChangeMemberLevelRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "newmembercard/ChangeMemberLevel", "api")
	request.Method = api.POST
	return
}

func CreateNewmembercardChangeMemberLevelResponse() (response *api.BaseResponse[NewmembercardChangeMemberLevelResponse]) {
	response = api.CreateResponse[NewmembercardChangeMemberLevelResponse](&NewmembercardChangeMemberLevelResponse{})
	return
}
