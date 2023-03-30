package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberSetTradePassword
// @id 2450
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=设置交易密码)
func (client *Client) MemberSetTradePassword(request *MemberSetTradePasswordRequest) (response *MemberSetTradePasswordResponse, err error) {
	rpcResponse := CreateMemberSetTradePasswordResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberSetTradePasswordRequest struct {
	*api.BaseRequest

	Password        string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirmPassword,omitempty"`
	ConsumerWid     int    `json:"consumerWid,omitempty"`
}

type MemberSetTradePasswordResponse struct {
}

func CreateMemberSetTradePasswordRequest() (request *MemberSetTradePasswordRequest) {
	request = &MemberSetTradePasswordRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "member/setTradePassword", "api")
	request.Method = api.POST
	return
}

func CreateMemberSetTradePasswordResponse() (response *api.BaseResponse[MemberSetTradePasswordResponse]) {
	response = api.CreateResponse[MemberSetTradePasswordResponse](&MemberSetTradePasswordResponse{})
	return
}
