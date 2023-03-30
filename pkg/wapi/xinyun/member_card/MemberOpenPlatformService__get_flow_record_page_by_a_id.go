package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberOpenPlatformServiceGetFlowRecordPageByAId
// @id 183
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取余额记录)
func (client *Client) MemberOpenPlatformServiceGetFlowRecordPageByAId(request *MemberOpenPlatformServiceGetFlowRecordPageByAIdRequest) (response *MemberOpenPlatformServiceGetFlowRecordPageByAIdResponse, err error) {
	rpcResponse := CreateMemberOpenPlatformServiceGetFlowRecordPageByAIdResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberOpenPlatformServiceGetFlowRecordPageByAIdRequest struct {
	*api.BaseRequest

	Customerno            string `json:"customerno,omitempty"`
	NicknameOrCardcode    string `json:"nickname_or_cardcode,omitempty"`
	TransactionDateFrom   int64  `json:"transaction_date_from,omitempty"`
	TransactionDateTo     int64  `json:"transaction_date_to,omitempty"`
	FlowMode              int    `json:"flow_mode,omitempty"`
	TransactionAmountFrom int64  `json:"transaction_amount_from,omitempty"`
	TransactionAmountTo   int64  `json:"transaction_amount_to,omitempty"`
	Pageindex             int    `json:"pageindex,omitempty"`
	Pagesize              int    `json:"pagesize,omitempty"`
}

type MemberOpenPlatformServiceGetFlowRecordPageByAIdResponse struct {
}

func CreateMemberOpenPlatformServiceGetFlowRecordPageByAIdRequest() (request *MemberOpenPlatformServiceGetFlowRecordPageByAIdRequest) {
	request = &MemberOpenPlatformServiceGetFlowRecordPageByAIdRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "MemberOpenPlatformService/GetFlowRecordPageByAId", "api")
	request.Method = api.POST
	return
}

func CreateMemberOpenPlatformServiceGetFlowRecordPageByAIdResponse() (response *api.BaseResponse[MemberOpenPlatformServiceGetFlowRecordPageByAIdResponse]) {
	response = api.CreateResponse[MemberOpenPlatformServiceGetFlowRecordPageByAIdResponse](&MemberOpenPlatformServiceGetFlowRecordPageByAIdResponse{})
	return
}
