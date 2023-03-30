package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponCouponPresentReceive
// @id 1490
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=转赠领取)
func (client *Client) CouponCouponPresentReceive(request *CouponCouponPresentReceiveRequest) (response *CouponCouponPresentReceiveResponse, err error) {
	rpcResponse := CreateCouponCouponPresentReceiveResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponCouponPresentReceiveRequest struct {
	*api.BaseRequest

	Code      string `json:"code,omitempty"`
	FromWid   int64  `json:"fromWid,omitempty"`
	TargetWid int64  `json:"targetWid,omitempty"`
}

type CouponCouponPresentReceiveResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateCouponCouponPresentReceiveRequest() (request *CouponCouponPresentReceiveRequest) {
	request = &CouponCouponPresentReceiveRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "coupon/couponPresentReceive", "api")
	request.Method = api.POST
	return
}

func CreateCouponCouponPresentReceiveResponse() (response *api.BaseResponse[CouponCouponPresentReceiveResponse]) {
	response = api.CreateResponse[CouponCouponPresentReceiveResponse](&CouponCouponPresentReceiveResponse{})
	return
}
