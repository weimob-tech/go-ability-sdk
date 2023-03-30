package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberGetTradePasswordStatus
// @id 2452
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取交易密码设置状态)
func (client *Client) MemberGetTradePasswordStatus(request *MemberGetTradePasswordStatusRequest) (response *MemberGetTradePasswordStatusResponse, err error) {
	rpcResponse := CreateMemberGetTradePasswordStatusResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberGetTradePasswordStatusRequest struct {
	*api.BaseRequest

	ConsumerWid int `json:"consumerWid,omitempty"`
}

type MemberGetTradePasswordStatusResponse struct {
}

func CreateMemberGetTradePasswordStatusRequest() (request *MemberGetTradePasswordStatusRequest) {
	request = &MemberGetTradePasswordStatusRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "member/getTradePasswordStatus", "api")
	request.Method = api.POST
	return
}

func CreateMemberGetTradePasswordStatusResponse() (response *api.BaseResponse[MemberGetTradePasswordStatusResponse]) {
	response = api.CreateResponse[MemberGetTradePasswordStatusResponse](&MemberGetTradePasswordStatusResponse{})
	return
}
