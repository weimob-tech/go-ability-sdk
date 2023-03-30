package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberOpenPlatformServiceGetPaymentRecordPage
// @id 186
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取收款记录)
func (client *Client) MemberOpenPlatformServiceGetPaymentRecordPage(request *MemberOpenPlatformServiceGetPaymentRecordPageRequest) (response *MemberOpenPlatformServiceGetPaymentRecordPageResponse, err error) {
	rpcResponse := CreateMemberOpenPlatformServiceGetPaymentRecordPageResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberOpenPlatformServiceGetPaymentRecordPageRequest struct {
	*api.BaseRequest

	CustomerNo          string `json:"customer_no,omitempty"`
	OrderNo             string `json:"order_no,omitempty"`
	MemberCardCode      string `json:"member_card_code,omitempty"`
	Nickname            string `json:"nickname,omitempty"`
	TransactionDateFrom int64  `json:"transaction_date_from,omitempty"`
	TransactionDateTo   int64  `json:"transaction_date_to,omitempty"`
	TransactionAmtFrom  int64  `json:"transaction_amt_from,omitempty"`
	TransactionAmtTo    int64  `json:"transaction_amt_to,omitempty"`
	Pageindex           int    `json:"pageindex,omitempty"`
	Pagesize            int    `json:"pagesize,omitempty"`
}

type MemberOpenPlatformServiceGetPaymentRecordPageResponse struct {
}

func CreateMemberOpenPlatformServiceGetPaymentRecordPageRequest() (request *MemberOpenPlatformServiceGetPaymentRecordPageRequest) {
	request = &MemberOpenPlatformServiceGetPaymentRecordPageRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "MemberOpenPlatformService/GetPaymentRecordPage", "api")
	request.Method = api.POST
	return
}

func CreateMemberOpenPlatformServiceGetPaymentRecordPageResponse() (response *api.BaseResponse[MemberOpenPlatformServiceGetPaymentRecordPageResponse]) {
	response = api.CreateResponse[MemberOpenPlatformServiceGetPaymentRecordPageResponse](&MemberOpenPlatformServiceGetPaymentRecordPageResponse{})
	return
}
