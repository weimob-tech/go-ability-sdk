package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponCouponStatusChange
// @id 1453
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=优惠券状态变更)
func (client *Client) CouponCouponStatusChange(request *CouponCouponStatusChangeRequest) (response *CouponCouponStatusChangeResponse, err error) {
	rpcResponse := CreateCouponCouponStatusChangeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponCouponStatusChangeRequest struct {
	*api.BaseRequest

	Code       string `json:"code,omitempty"`
	ChangeType int    `json:"changeType,omitempty"`
}

type CouponCouponStatusChangeResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateCouponCouponStatusChangeRequest() (request *CouponCouponStatusChangeRequest) {
	request = &CouponCouponStatusChangeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "coupon/couponStatusChange", "api")
	request.Method = api.POST
	return
}

func CreateCouponCouponStatusChangeResponse() (response *api.BaseResponse[CouponCouponStatusChangeResponse]) {
	response = api.CreateResponse[CouponCouponStatusChangeResponse](&CouponCouponStatusChangeResponse{})
	return
}
