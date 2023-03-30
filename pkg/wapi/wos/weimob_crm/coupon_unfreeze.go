package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponUnfreeze
// @id 1616
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1616?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=解冻优惠券)
func (client *Client) CouponUnfreeze(request *CouponUnfreezeRequest) (response *CouponUnfreezeResponse, err error) {
	rpcResponse := CreateCouponUnfreezeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponUnfreezeRequest struct {
	*api.BaseRequest

	OperateCoupons []CouponUnfreezeRequestOperateCoupons `json:"operateCoupons,omitempty"`
	Wid            int64                                 `json:"wid,omitempty"`
	VidType        int64                                 `json:"vidType,omitempty"`
	Vid            int64                                 `json:"vid,omitempty"`
}

type CouponUnfreezeResponse struct {
}

func CreateCouponUnfreezeRequest() (request *CouponUnfreezeRequest) {
	request = &CouponUnfreezeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/unfreeze", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponUnfreezeResponse() (response *api.BaseResponse[CouponUnfreezeResponse]) {
	response = api.CreateResponse[CouponUnfreezeResponse](&CouponUnfreezeResponse{})
	return
}

type CouponUnfreezeRequestOperateCoupons struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	OrderNo          string `json:"orderNo,omitempty"`
	Code             string `json:"code,omitempty"`
}
