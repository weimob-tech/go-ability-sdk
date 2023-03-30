package wmpay

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PaymentCustomDeclare
// @id 1485
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=微信报关接口)
func (client *Client) PaymentCustomDeclare(request *PaymentCustomDeclareRequest) (response *PaymentCustomDeclareResponse, err error) {
	rpcResponse := CreatePaymentCustomDeclareResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PaymentCustomDeclareRequest struct {
	*api.BaseRequest

	Appid         string `json:"appid,omitempty"`
	MchId         string `json:"mch_id,omitempty"`
	TransactionId string `json:"transaction_id,omitempty"`
	OutTradeNo    string `json:"out_trade_no,omitempty"`
	Customs       string `json:"customs,omitempty"`
	MchCustomsNo  string `json:"mch_customs_no,omitempty"`
	Duty          string `json:"duty,omitempty"`
	ActionType    string `json:"action_type,omitempty"`
	SubOrderNo    string `json:"sub_order_no,omitempty"`
	FeeType       string `json:"fee_type,omitempty"`
	OrderFee      string `json:"order_fee,omitempty"`
	TransportFee  string `json:"transport_fee,omitempty"`
	ProductFee    string `json:"product_fee,omitempty"`
	CertType      string `json:"cert_type,omitempty"`
	CertId        string `json:"cert_id,omitempty"`
	Name          string `json:"name,omitempty"`
}

type PaymentCustomDeclareResponse struct {
}

func CreatePaymentCustomDeclareRequest() (request *PaymentCustomDeclareRequest) {
	request = &PaymentCustomDeclareRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wmpay", "1_0", "payment/customDeclare", "api")
	request.Method = api.POST
	return
}

func CreatePaymentCustomDeclareResponse() (response *api.BaseResponse[PaymentCustomDeclareResponse]) {
	response = api.CreateResponse[PaymentCustomDeclareResponse](&PaymentCustomDeclareResponse{})
	return
}
