package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpGetMemberField
// @id 1681
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询客户资料设置)
func (client *Client) MbpGetMemberField(request *MbpGetMemberFieldRequest) (response *MbpGetMemberFieldResponse, err error) {
	rpcResponse := CreateMbpGetMemberFieldResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpGetMemberFieldRequest struct {
	*api.BaseRequest
}

type MbpGetMemberFieldResponse struct {
}

func CreateMbpGetMemberFieldRequest() (request *MbpGetMemberFieldRequest) {
	request = &MbpGetMemberFieldRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "mbp/getMemberField", "api")
	request.Method = api.POST
	return
}

func CreateMbpGetMemberFieldResponse() (response *api.BaseResponse[MbpGetMemberFieldResponse]) {
	response = api.CreateResponse[MbpGetMemberFieldResponse](&MbpGetMemberFieldResponse{})
	return
}
