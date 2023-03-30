package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponRecycleCoupon
// @id 1531
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=回收优惠券(线上))
func (client *Client) CouponRecycleCoupon(request *CouponRecycleCouponRequest) (response *CouponRecycleCouponResponse, err error) {
	rpcResponse := CreateCouponRecycleCouponResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponRecycleCouponRequest struct {
	*api.BaseRequest

	CouponCode string `json:"couponCode,omitempty"`
	OrderNo    string `json:"orderNo,omitempty"`
}

type CouponRecycleCouponResponse struct {
}

func CreateCouponRecycleCouponRequest() (request *CouponRecycleCouponRequest) {
	request = &CouponRecycleCouponRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "coupon/recycleCoupon", "api")
	request.Method = api.POST
	return
}

func CreateCouponRecycleCouponResponse() (response *api.BaseResponse[CouponRecycleCouponResponse]) {
	response = api.CreateResponse[CouponRecycleCouponResponse](&CouponRecycleCouponResponse{})
	return
}
