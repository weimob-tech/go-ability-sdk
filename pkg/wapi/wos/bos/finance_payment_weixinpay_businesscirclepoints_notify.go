package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FinancePaymentWeixinpayBusinesscirclepointsNotify
// @id 1793
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1793?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=微信商圈积分)
func (client *Client) FinancePaymentWeixinpayBusinesscirclepointsNotify(request *FinancePaymentWeixinpayBusinesscirclepointsNotifyRequest) (response *FinancePaymentWeixinpayBusinesscirclepointsNotifyResponse, err error) {
	rpcResponse := CreateFinancePaymentWeixinpayBusinesscirclepointsNotifyResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FinancePaymentWeixinpayBusinesscirclepointsNotifyRequest struct {
	*api.BaseRequest

	Mchid            string `json:"mchid,omitempty"`
	SubMchid         string `json:"sub_mchid,omitempty"`
	Appid            string `json:"appid,omitempty"`
	Openid           string `json:"openid,omitempty"`
	EarnPoints       bool   `json:"earn_points,omitempty"`
	IncreasedPoints  int64  `json:"increased_points,omitempty"`
	PointsUpdateTime string `json:"points_update_time,omitempty"`
	NoPointsRemarks  string `json:"no_points_remarks,omitempty"`
	TotalPoints      int64  `json:"total_points,omitempty"`
}

type FinancePaymentWeixinpayBusinesscirclepointsNotifyResponse struct {
	InteractResult string `json:"interact_result,omitempty"`
	InteractCode   string `json:"interact_code,omitempty"`
	InteractMsg    string `json:"interact_msg,omitempty"`
}

func CreateFinancePaymentWeixinpayBusinesscirclepointsNotifyRequest() (request *FinancePaymentWeixinpayBusinesscirclepointsNotifyRequest) {
	request = &FinancePaymentWeixinpayBusinesscirclepointsNotifyRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "finance/payment/weixinpay/businesscirclepoints/notify", "apigw")
	request.Method = api.POST
	return
}

func CreateFinancePaymentWeixinpayBusinesscirclepointsNotifyResponse() (response *api.BaseResponse[FinancePaymentWeixinpayBusinesscirclepointsNotifyResponse]) {
	response = api.CreateResponse[FinancePaymentWeixinpayBusinesscirclepointsNotifyResponse](&FinancePaymentWeixinpayBusinesscirclepointsNotifyResponse{})
	return
}
