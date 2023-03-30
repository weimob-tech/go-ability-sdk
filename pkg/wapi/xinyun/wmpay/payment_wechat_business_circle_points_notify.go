package wmpay

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PaymentWechatBusinessCirclePointsNotify
// @id 1554
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=微信商圈积分通知)
func (client *Client) PaymentWechatBusinessCirclePointsNotify(request *PaymentWechatBusinessCirclePointsNotifyRequest) (response *PaymentWechatBusinessCirclePointsNotifyResponse, err error) {
	rpcResponse := CreatePaymentWechatBusinessCirclePointsNotifyResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PaymentWechatBusinessCirclePointsNotifyRequest struct {
	*api.BaseRequest

	Mchid            string `json:"mchid,omitempty"`
	SubMchid         string `json:"sub_mchid,omitempty"`
	Appid            string `json:"appid,omitempty"`
	Openid           string `json:"openid,omitempty"`
	EarnPoints       bool   `json:"earn_points,omitempty"`
	IncreasedPoints  int    `json:"increased_points,omitempty"`
	PointsUpdateTime string `json:"points_update_time,omitempty"`
	NoPointsRemarks  string `json:"no_points_remarks,omitempty"`
	TotalPoints      int    `json:"total_points,omitempty"`
}

type PaymentWechatBusinessCirclePointsNotifyResponse struct {
}

func CreatePaymentWechatBusinessCirclePointsNotifyRequest() (request *PaymentWechatBusinessCirclePointsNotifyRequest) {
	request = &PaymentWechatBusinessCirclePointsNotifyRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wmpay", "1_0", "payment/wechatBusinessCirclePointsNotify", "api")
	request.Method = api.POST
	return
}

func CreatePaymentWechatBusinessCirclePointsNotifyResponse() (response *api.BaseResponse[PaymentWechatBusinessCirclePointsNotifyResponse]) {
	response = api.CreateResponse[PaymentWechatBusinessCirclePointsNotifyResponse](&PaymentWechatBusinessCirclePointsNotifyResponse{})
	return
}
