package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberQueryDepositCardInfo
// @id 1338
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询储值卡信息)
func (client *Client) MemberQueryDepositCardInfo(request *MemberQueryDepositCardInfoRequest) (response *MemberQueryDepositCardInfoResponse, err error) {
	rpcResponse := CreateMemberQueryDepositCardInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberQueryDepositCardInfoRequest struct {
	*api.BaseRequest

	ConsumerWid int64 `json:"consumerWid,omitempty"`
	StoreId     int64 `json:"storeId,omitempty"`
}

type MemberQueryDepositCardInfoResponse struct {
}

func CreateMemberQueryDepositCardInfoRequest() (request *MemberQueryDepositCardInfoRequest) {
	request = &MemberQueryDepositCardInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "member/queryDepositCardInfo", "api")
	request.Method = api.POST
	return
}

func CreateMemberQueryDepositCardInfoResponse() (response *api.BaseResponse[MemberQueryDepositCardInfoResponse]) {
	response = api.CreateResponse[MemberQueryDepositCardInfoResponse](&MemberQueryDepositCardInfoResponse{})
	return
}
