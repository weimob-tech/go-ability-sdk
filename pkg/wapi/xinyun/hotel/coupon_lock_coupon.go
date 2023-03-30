package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponLockCoupon
// @id 1533
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=锁定优惠券(线上核销))
func (client *Client) CouponLockCoupon(request *CouponLockCouponRequest) (response *CouponLockCouponResponse, err error) {
	rpcResponse := CreateCouponLockCouponResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponLockCouponRequest struct {
	*api.BaseRequest

	CouponCode string `json:"couponCode,omitempty"`
	OrderNo    string `json:"orderNo,omitempty"`
	Wid        int64  `json:"wid,omitempty"`
}

type CouponLockCouponResponse struct {
}

func CreateCouponLockCouponRequest() (request *CouponLockCouponRequest) {
	request = &CouponLockCouponRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "coupon/lockCoupon", "api")
	request.Method = api.POST
	return
}

func CreateCouponLockCouponResponse() (response *api.BaseResponse[CouponLockCouponResponse]) {
	response = api.CreateResponse[CouponLockCouponResponse](&CouponLockCouponResponse{})
	return
}
