package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponCouponPresent
// @id 1489
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=发起转赠)
func (client *Client) CouponCouponPresent(request *CouponCouponPresentRequest) (response *CouponCouponPresentResponse, err error) {
	rpcResponse := CreateCouponCouponPresentResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponCouponPresentRequest struct {
	*api.BaseRequest

	FromWid int64  `json:"fromWid,omitempty"`
	Code    string `json:"code,omitempty"`
}

type CouponCouponPresentResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateCouponCouponPresentRequest() (request *CouponCouponPresentRequest) {
	request = &CouponCouponPresentRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "coupon/couponPresent", "api")
	request.Method = api.POST
	return
}

func CreateCouponCouponPresentResponse() (response *api.BaseResponse[CouponCouponPresentResponse]) {
	response = api.CreateResponse[CouponCouponPresentResponse](&CouponCouponPresentResponse{})
	return
}
