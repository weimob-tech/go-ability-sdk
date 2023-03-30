package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberCheckTradePassword
// @id 2451
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=校验交易密码)
func (client *Client) MemberCheckTradePassword(request *MemberCheckTradePasswordRequest) (response *MemberCheckTradePasswordResponse, err error) {
	rpcResponse := CreateMemberCheckTradePasswordResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberCheckTradePasswordRequest struct {
	*api.BaseRequest

	Password    string `json:"password,omitempty"`
	ConsumerWid int    `json:"consumerWid,omitempty"`
}

type MemberCheckTradePasswordResponse struct {
}

func CreateMemberCheckTradePasswordRequest() (request *MemberCheckTradePasswordRequest) {
	request = &MemberCheckTradePasswordRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "member/checkTradePassword", "api")
	request.Method = api.POST
	return
}

func CreateMemberCheckTradePasswordResponse() (response *api.BaseResponse[MemberCheckTradePasswordResponse]) {
	response = api.CreateResponse[MemberCheckTradePasswordResponse](&MemberCheckTradePasswordResponse{})
	return
}
