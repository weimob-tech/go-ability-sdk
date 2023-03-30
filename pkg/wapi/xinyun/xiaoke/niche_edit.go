package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NicheEdit
// @id 1735
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=编辑商机)
func (client *Client) NicheEdit(request *NicheEditRequest) (response *NicheEditResponse, err error) {
	rpcResponse := CreateNicheEditResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NicheEditRequest struct {
	*api.BaseRequest

	NicheInfo NicheEditRequestNicheInfo `json:"nicheInfo,omitempty"`
	Key       string                    `json:"key,omitempty"`
	Wid       int64                     `json:"wid,omitempty"`
}

type NicheEditResponse struct {
}

func CreateNicheEditRequest() (request *NicheEditRequest) {
	request = &NicheEditRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "niche/edit", "api")
	request.Method = api.POST
	return
}

func CreateNicheEditResponse() (response *api.BaseResponse[NicheEditResponse]) {
	response = api.CreateResponse[NicheEditResponse](&NicheEditResponse{})
	return
}

type NicheEditRequestNicheInfo struct {
	SIntentionalProduct []map[string]any `json:"s_intentional_product,omitempty"`
	SCauseOfLoseOrder   string           `json:"s_cause_of_lose_order,omitempty"`
	SCollaborators      int64            `json:"s_collaborators,omitempty"`
	SCreateTime         string           `json:"s_create_time,omitempty"`
	SCreateUserWid      int64            `json:"s_create_user_wid,omitempty"`
	SCustomerKey        int64            `json:"s_customer_key,omitempty"`
	SExpectedDealDate   string           `json:"s_expected_deal_date,omitempty"`
	SFollowStatus       string           `json:"s_follow_status,omitempty"`
	SFormerOwner        int64            `json:"s_former_owner,omitempty"`
	SLastDealTime       string           `json:"s_last_deal_time,omitempty"`
	SLastFollowTime     string           `json:"s_last_follow_time,omitempty"`
	SLastFollowUserWid  int64            `json:"s_last_follow_user_wid,omitempty"`
	SLastUpdateTime     string           `json:"s_last_update_time,omitempty"`
	SLastUpdateUserWid  int64            `json:"s_last_update_user_wid,omitempty"`
	SNicheName          string           `json:"s_niche_name,omitempty"`
	SNicheStatus        string           `json:"s_niche_status,omitempty"`
	SOwner              int64            `json:"s_owner,omitempty"`
	SSalesStage         int64            `json:"s_sales_stage,omitempty"`
	STransactionAmount  float64          `json:"s_transaction_amount,omitempty"`
	SWinRate            int64            `json:"s_win_rate,omitempty"`
}
