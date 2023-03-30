package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponFreeze
// @id 1617
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1617?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=冻结优惠券)
func (client *Client) CouponFreeze(request *CouponFreezeRequest) (response *CouponFreezeResponse, err error) {
	rpcResponse := CreateCouponFreezeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponFreezeRequest struct {
	*api.BaseRequest

	OperateCoupons []CouponFreezeRequestOperateCoupons `json:"operateCoupons,omitempty"`
	UseScene       int64                               `json:"useScene,omitempty"`
	Wid            int64                               `json:"wid,omitempty"`
	VidType        int64                               `json:"vidType,omitempty"`
	Vid            int64                               `json:"vid,omitempty"`
}

type CouponFreezeResponse struct {
}

func CreateCouponFreezeRequest() (request *CouponFreezeRequest) {
	request = &CouponFreezeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/freeze", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponFreezeResponse() (response *api.BaseResponse[CouponFreezeResponse]) {
	response = api.CreateResponse[CouponFreezeResponse](&CouponFreezeResponse{})
	return
}

type CouponFreezeRequestOperateCoupons struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	OrderNo          string `json:"orderNo,omitempty"`
	Code             string `json:"code,omitempty"`
	Amount           string `json:"amount,omitempty"`
}
