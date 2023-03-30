package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NewmembercardOpenOrCloseMember
// @id 110
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=冻结/解冻会员)
func (client *Client) NewmembercardOpenOrCloseMember(request *NewmembercardOpenOrCloseMemberRequest) (response *NewmembercardOpenOrCloseMemberResponse, err error) {
	rpcResponse := CreateNewmembercardOpenOrCloseMemberResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NewmembercardOpenOrCloseMemberRequest struct {
	*api.BaseRequest

	IsClose          int      `json:"isClose,omitempty"`
	ListMemberCardNo []string `json:"listMemberCardNo,omitempty"`
	Operator         string   `json:"operator,omitempty"`
}

type NewmembercardOpenOrCloseMemberResponse struct {
}

func CreateNewmembercardOpenOrCloseMemberRequest() (request *NewmembercardOpenOrCloseMemberRequest) {
	request = &NewmembercardOpenOrCloseMemberRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "newmembercard/OpenOrCloseMember", "api")
	request.Method = api.POST
	return
}

func CreateNewmembercardOpenOrCloseMemberResponse() (response *api.BaseResponse[NewmembercardOpenOrCloseMemberResponse]) {
	response = api.CreateResponse[NewmembercardOpenOrCloseMemberResponse](&NewmembercardOpenOrCloseMemberResponse{})
	return
}
