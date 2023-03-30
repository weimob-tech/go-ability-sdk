package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponCancel
// @id 1615
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1615?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=作废优惠券)
func (client *Client) CouponCancel(request *CouponCancelRequest) (response *CouponCancelResponse, err error) {
	rpcResponse := CreateCouponCancelResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponCancelRequest struct {
	*api.BaseRequest

	OperateCoupons []CouponCancelRequestOperateCoupons `json:"operateCoupons,omitempty"`
	Wid            int64                               `json:"wid,omitempty"`
	VidType        int64                               `json:"vidType,omitempty"`
	Vid            int64                               `json:"vid,omitempty"`
	Type           int64                               `json:"type,omitempty"`
}

type CouponCancelResponse struct {
}

func CreateCouponCancelRequest() (request *CouponCancelRequest) {
	request = &CouponCancelRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/cancel", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponCancelResponse() (response *api.BaseResponse[CouponCancelResponse]) {
	response = api.CreateResponse[CouponCancelResponse](&CouponCancelResponse{})
	return
}

type CouponCancelRequestOperateCoupons struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	OrderNo          string `json:"orderNo,omitempty"`
	Code             string `json:"code,omitempty"`
}
