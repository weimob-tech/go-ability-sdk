package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponUnlockCoupon
// @id 1532
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=解锁优惠券(线上核销))
func (client *Client) CouponUnlockCoupon(request *CouponUnlockCouponRequest) (response *CouponUnlockCouponResponse, err error) {
	rpcResponse := CreateCouponUnlockCouponResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponUnlockCouponRequest struct {
	*api.BaseRequest

	CouponCode string `json:"couponCode,omitempty"`
	OrderNo    string `json:"orderNo,omitempty"`
	Wid        int64  `json:"wid,omitempty"`
}

type CouponUnlockCouponResponse struct {
}

func CreateCouponUnlockCouponRequest() (request *CouponUnlockCouponRequest) {
	request = &CouponUnlockCouponRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "coupon/unlockCoupon", "api")
	request.Method = api.POST
	return
}

func CreateCouponUnlockCouponResponse() (response *api.BaseResponse[CouponUnlockCouponResponse]) {
	response = api.CreateResponse[CouponUnlockCouponResponse](&CouponUnlockCouponResponse{})
	return
}
