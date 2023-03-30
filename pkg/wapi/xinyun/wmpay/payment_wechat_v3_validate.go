package wmpay

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PaymentWechatV3Validate
// @id 1553
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=微信v3参数验签解密)
func (client *Client) PaymentWechatV3Validate(request *PaymentWechatV3ValidateRequest) (response *PaymentWechatV3ValidateResponse, err error) {
	rpcResponse := CreatePaymentWechatV3ValidateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PaymentWechatV3ValidateRequest struct {
	*api.BaseRequest

	Mchid           string `json:"mchid,omitempty"`
	WechatV3Content string `json:"wechat_v3_content,omitempty"`
	WechatV3Header  string `json:"wechat_v3_header,omitempty"`
}

type PaymentWechatV3ValidateResponse struct {
}

func CreatePaymentWechatV3ValidateRequest() (request *PaymentWechatV3ValidateRequest) {
	request = &PaymentWechatV3ValidateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wmpay", "1_0", "payment/wechatV3Validate", "api")
	request.Method = api.POST
	return
}

func CreatePaymentWechatV3ValidateResponse() (response *api.BaseResponse[PaymentWechatV3ValidateResponse]) {
	response = api.CreateResponse[PaymentWechatV3ValidateResponse](&PaymentWechatV3ValidateResponse{})
	return
}
