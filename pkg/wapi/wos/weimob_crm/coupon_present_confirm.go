package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponPresentConfirm
// @id 2106
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2106?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=优惠券转赠确认)
func (client *Client) CouponPresentConfirm(request *CouponPresentConfirmRequest) (response *CouponPresentConfirmResponse, err error) {
	rpcResponse := CreateCouponPresentConfirmResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponPresentConfirmRequest struct {
	*api.BaseRequest

	FromWid   int64  `json:"fromWid,omitempty"`
	TargetWid int64  `json:"targetWid,omitempty"`
	Code      string `json:"code,omitempty"`
}

type CouponPresentConfirmResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateCouponPresentConfirmRequest() (request *CouponPresentConfirmRequest) {
	request = &CouponPresentConfirmRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/present/confirm", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponPresentConfirmResponse() (response *api.BaseResponse[CouponPresentConfirmResponse]) {
	response = api.CreateResponse[CouponPresentConfirmResponse](&CouponPresentConfirmResponse{})
	return
}
