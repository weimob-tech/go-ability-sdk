package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberOpenPlatformServiceGetPointsRecordPageByAId
// @id 187
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取积分记录)
func (client *Client) MemberOpenPlatformServiceGetPointsRecordPageByAId(request *MemberOpenPlatformServiceGetPointsRecordPageByAIdRequest) (response *MemberOpenPlatformServiceGetPointsRecordPageByAIdResponse, err error) {
	rpcResponse := CreateMemberOpenPlatformServiceGetPointsRecordPageByAIdResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberOpenPlatformServiceGetPointsRecordPageByAIdRequest struct {
	*api.BaseRequest

	Nickname            string `json:"nickname,omitempty"`
	CustomerNo          string `json:"customer_no,omitempty"`
	CardCode            string `json:"card_code,omitempty"`
	TransactionDateFrom int64  `json:"transaction_date_from,omitempty"`
	TransactionDateTo   int64  `json:"transaction_date_to,omitempty"`
	TransactionType     int64  `json:"transaction_type,omitempty"`
	TransactionMode     int64  `json:"transaction_mode,omitempty"`
	Pageindex           int    `json:"pageindex,omitempty"`
	Pagesize            int    `json:"pagesize,omitempty"`
}

type MemberOpenPlatformServiceGetPointsRecordPageByAIdResponse struct {
}

func CreateMemberOpenPlatformServiceGetPointsRecordPageByAIdRequest() (request *MemberOpenPlatformServiceGetPointsRecordPageByAIdRequest) {
	request = &MemberOpenPlatformServiceGetPointsRecordPageByAIdRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "MemberOpenPlatformService/GetPointsRecordPageByAId", "api")
	request.Method = api.POST
	return
}

func CreateMemberOpenPlatformServiceGetPointsRecordPageByAIdResponse() (response *api.BaseResponse[MemberOpenPlatformServiceGetPointsRecordPageByAIdResponse]) {
	response = api.CreateResponse[MemberOpenPlatformServiceGetPointsRecordPageByAIdResponse](&MemberOpenPlatformServiceGetPointsRecordPageByAIdResponse{})
	return
}
