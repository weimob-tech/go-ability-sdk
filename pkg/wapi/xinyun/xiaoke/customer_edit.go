package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerEdit
// @id 1714
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=编辑客户)
func (client *Client) CustomerEdit(request *CustomerEditRequest) (response *CustomerEditResponse, err error) {
	rpcResponse := CreateCustomerEditResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerEditRequest struct {
	*api.BaseRequest

	CustomerInfo CustomerEditRequestCustomerInfo `json:"customerInfo,omitempty"`
	Key          string                          `json:"key,omitempty"`
	Wid          int64                           `json:"wid,omitempty"`
}

type CustomerEditResponse struct {
}

func CreateCustomerEditRequest() (request *CustomerEditRequest) {
	request = &CustomerEditRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "customer/edit", "api")
	request.Method = api.POST
	return
}

func CreateCustomerEditResponse() (response *api.BaseResponse[CustomerEditResponse]) {
	response = api.CreateResponse[CustomerEditResponse](&CustomerEditResponse{})
	return
}

type CustomerEditRequestCustomerInfo struct {
	SCauseOfAbandonment string `json:"s_cause_of_abandonment,omitempty"`
	SClaimTime          string `json:"s_claim_time,omitempty"`
	SConversionTime     string `json:"s_conversion_time,omitempty"`
	SCreateTime         string `json:"s_create_time,omitempty"`
	SCreateUserWid      int64  `json:"s_create_user_wid,omitempty"`
	SCustomerLevel      string `json:"s_customer_level,omitempty"`
	SCustomerName       string `json:"s_customer_name,omitempty"`
	SCustomerStatus     string `json:"s_customer_status,omitempty"`
	SAddress            string `json:"s_address,omitempty"`
	SDeadline           string `json:"s_deadline,omitempty"`
	SDealStatus         string `json:"s_deal_status,omitempty"`
	SEnterpriseScale    string `json:"s_enterprise_scale,omitempty"`
	SFollowStatus       string `json:"s_follow_status,omitempty"`
	SFormerOwner        int64  `json:"s_former_owner,omitempty"`
	SIndustryType       string `json:"s_industry_type,omitempty"`
	SInformationSource  string `json:"s_information_source,omitempty"`
	SLastDealTime       string `json:"s_last_deal_time,omitempty"`
	SLastFollowTime     string `json:"s_last_follow_time,omitempty"`
	SLastFollowUserWid  int64  `json:"s_last_follow_user_wid,omitempty"`
	SLastUpdateTime     string `json:"s_last_update_time,omitempty"`
	SLastUpdateUserWid  int64  `json:"s_last_update_user_wid,omitempty"`
	SOwner              int64  `json:"s_owner,omitempty"`
	SPoolingTime        string `json:"s_pooling_time,omitempty"`
	SRemark             string `json:"s_remark,omitempty"`
	SWebsite            string `json:"s_website,omitempty"`
	SCustomerSource     string `json:"s_customer_source,omitempty"`
}
