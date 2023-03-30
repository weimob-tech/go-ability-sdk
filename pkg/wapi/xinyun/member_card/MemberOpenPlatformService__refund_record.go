package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberOpenPlatformServiceRefundRecord
// @id 406
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=退款退积分)
func (client *Client) MemberOpenPlatformServiceRefundRecord(request *MemberOpenPlatformServiceRefundRecordRequest) (response *MemberOpenPlatformServiceRefundRecordResponse, err error) {
	rpcResponse := CreateMemberOpenPlatformServiceRefundRecordResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberOpenPlatformServiceRefundRecordRequest struct {
	*api.BaseRequest

	FlowNo string `json:"flow_no,omitempty"`
}

type MemberOpenPlatformServiceRefundRecordResponse struct {
}

func CreateMemberOpenPlatformServiceRefundRecordRequest() (request *MemberOpenPlatformServiceRefundRecordRequest) {
	request = &MemberOpenPlatformServiceRefundRecordRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "MemberOpenPlatformService/RefundRecord", "api")
	request.Method = api.POST
	return
}

func CreateMemberOpenPlatformServiceRefundRecordResponse() (response *api.BaseResponse[MemberOpenPlatformServiceRefundRecordResponse]) {
	response = api.CreateResponse[MemberOpenPlatformServiceRefundRecordResponse](&MemberOpenPlatformServiceRefundRecordResponse{})
	return
}
