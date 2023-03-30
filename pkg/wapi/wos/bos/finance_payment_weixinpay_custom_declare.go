package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FinancePaymentWeixinpayCustomDeclare
// @id 1794
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1794?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=微信支付报关)
func (client *Client) FinancePaymentWeixinpayCustomDeclare(request *FinancePaymentWeixinpayCustomDeclareRequest) (response *FinancePaymentWeixinpayCustomDeclareResponse, err error) {
	rpcResponse := CreateFinancePaymentWeixinpayCustomDeclareResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FinancePaymentWeixinpayCustomDeclareRequest struct {
	*api.BaseRequest

	Appid         string `json:"appid,omitempty"`
	MchId         string `json:"mch_id,omitempty"`
	OutTradeNo    string `json:"out_trade_no,omitempty"`
	TransactionId string `json:"transaction_id,omitempty"`
	Customs       string `json:"customs,omitempty"`
	MchCustomsNo  string `json:"mch_customs_no,omitempty"`
	Duty          string `json:"duty,omitempty"`
	ActionType    string `json:"action_type,omitempty"`
}

type FinancePaymentWeixinpayCustomDeclareResponse struct {
	ReturnCode              string `json:"return_code,omitempty"`
	ReturnMsg               string `json:"return_msg,omitempty"`
	Appid                   string `json:"appid,omitempty"`
	MchId                   string `json:"mch_id,omitempty"`
	ResultCode              string `json:"result_code,omitempty"`
	ErrCode                 string `json:"err_code,omitempty"`
	ErrCodeDes              string `json:"err_code_des,omitempty"`
	State                   string `json:"state,omitempty"`
	TransactionId           string `json:"transaction_id,omitempty"`
	OutTradeNo              string `json:"out_trade_no,omitempty"`
	SubOrderNo              string `json:"sub_order_no,omitempty"`
	SubOrderId              string `json:"sub_order_id,omitempty"`
	ModifyTime              string `json:"modify_time,omitempty"`
	CertCheckResult         string `json:"cert_check_result,omitempty"`
	VerifyDepartment        string `json:"verify_department,omitempty"`
	VerifyDepartmentTradeId string `json:"verify_department_trade_id,omitempty"`
}

func CreateFinancePaymentWeixinpayCustomDeclareRequest() (request *FinancePaymentWeixinpayCustomDeclareRequest) {
	request = &FinancePaymentWeixinpayCustomDeclareRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "finance/payment/weixinpay/custom/declare", "apigw")
	request.Method = api.POST
	return
}

func CreateFinancePaymentWeixinpayCustomDeclareResponse() (response *api.BaseResponse[FinancePaymentWeixinpayCustomDeclareResponse]) {
	response = api.CreateResponse[FinancePaymentWeixinpayCustomDeclareResponse](&FinancePaymentWeixinpayCustomDeclareResponse{})
	return
}
