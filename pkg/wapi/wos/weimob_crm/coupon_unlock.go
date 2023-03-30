package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponUnlock
// @id 1143
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1143?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=优惠券解锁)
func (client *Client) CouponUnlock(request *CouponUnlockRequest) (response *CouponUnlockResponse, err error) {
	rpcResponse := CreateCouponUnlockResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponUnlockRequest struct {
	*api.BaseRequest

	CouponList []CouponUnlockRequestCouponList `json:"couponList,omitempty"`
	UseScene   int64                           `json:"useScene,omitempty"`
	Wid        int64                           `json:"wid,omitempty"`
	VidType    int64                           `json:"vidType,omitempty"`
	Vid        int64                           `json:"vid,omitempty"`
}

type CouponUnlockResponse struct {
}

func CreateCouponUnlockRequest() (request *CouponUnlockRequest) {
	request = &CouponUnlockRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/unlock", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponUnlockResponse() (response *api.BaseResponse[CouponUnlockResponse]) {
	response = api.CreateResponse[CouponUnlockResponse](&CouponUnlockResponse{})
	return
}

type CouponUnlockRequestCouponList struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	OrderNo          string `json:"orderNo,omitempty"`
	Code             string `json:"code,omitempty"`
	Amount           string `json:"amount,omitempty"`
}
