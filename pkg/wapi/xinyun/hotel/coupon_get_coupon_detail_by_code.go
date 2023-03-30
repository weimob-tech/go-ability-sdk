package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponGetCouponDetailByCode
// @id 1536
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=通过优惠券码获取优惠券信息)
func (client *Client) CouponGetCouponDetailByCode(request *CouponGetCouponDetailByCodeRequest) (response *CouponGetCouponDetailByCodeResponse, err error) {
	rpcResponse := CreateCouponGetCouponDetailByCodeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponGetCouponDetailByCodeRequest struct {
	*api.BaseRequest

	CouponCode string `json:"couponCode,omitempty"`
}

type CouponGetCouponDetailByCodeResponse struct {
}

func CreateCouponGetCouponDetailByCodeRequest() (request *CouponGetCouponDetailByCodeRequest) {
	request = &CouponGetCouponDetailByCodeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "coupon/getCouponDetailByCode", "api")
	request.Method = api.POST
	return
}

func CreateCouponGetCouponDetailByCodeResponse() (response *api.BaseResponse[CouponGetCouponDetailByCodeResponse]) {
	response = api.CreateResponse[CouponGetCouponDetailByCodeResponse](&CouponGetCouponDetailByCodeResponse{})
	return
}
