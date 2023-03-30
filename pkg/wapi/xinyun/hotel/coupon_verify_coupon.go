package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponVerifyCoupon
// @id 1534
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=核销优惠券(线上/线下))
func (client *Client) CouponVerifyCoupon(request *CouponVerifyCouponRequest) (response *CouponVerifyCouponResponse, err error) {
	rpcResponse := CreateCouponVerifyCouponResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponVerifyCouponRequest struct {
	*api.BaseRequest

	CouponCode string  `json:"couponCode,omitempty"`
	StoreId    int64   `json:"storeId,omitempty"`
	VerifyType int     `json:"verifyType,omitempty"`
	OrderNo    string  `json:"orderNo,omitempty"`
	Amount     float64 `json:"amount,omitempty"`
}

type CouponVerifyCouponResponse struct {
}

func CreateCouponVerifyCouponRequest() (request *CouponVerifyCouponRequest) {
	request = &CouponVerifyCouponRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "coupon/verifyCoupon", "api")
	request.Method = api.POST
	return
}

func CreateCouponVerifyCouponResponse() (response *api.BaseResponse[CouponVerifyCouponResponse]) {
	response = api.CreateResponse[CouponVerifyCouponResponse](&CouponVerifyCouponResponse{})
	return
}
