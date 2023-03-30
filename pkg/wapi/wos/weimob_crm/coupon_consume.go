package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponConsume
// @id 1614
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1614?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=核销优惠券)
func (client *Client) CouponConsume(request *CouponConsumeRequest) (response *CouponConsumeResponse, err error) {
	rpcResponse := CreateCouponConsumeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponConsumeRequest struct {
	*api.BaseRequest

	OperateCoupons  []CouponConsumeRequestOperateCoupons `json:"operateCoupons,omitempty"`
	OrderAmount     int64                                `json:"orderAmount,omitempty"`
	Operator        int64                                `json:"operator,omitempty"`
	UseResourceType int64                                `json:"useResourceType,omitempty"`
	UseScene        int64                                `json:"useScene,omitempty"`
	Wid             int64                                `json:"wid,omitempty"`
	VidType         int64                                `json:"vidType,omitempty"`
	Vid             int64                                `json:"vid,omitempty"`
}

type CouponConsumeResponse struct {
}

func CreateCouponConsumeRequest() (request *CouponConsumeRequest) {
	request = &CouponConsumeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/consume", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponConsumeResponse() (response *api.BaseResponse[CouponConsumeResponse]) {
	response = api.CreateResponse[CouponConsumeResponse](&CouponConsumeResponse{})
	return
}

type CouponConsumeRequestOperateCoupons struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	OrderNo          string `json:"orderNo,omitempty"`
	Code             string `json:"code,omitempty"`
	Amount           int64  `json:"amount,omitempty"`
}
