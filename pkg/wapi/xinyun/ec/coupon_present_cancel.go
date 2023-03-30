package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponPresentCancel
// @id 1908
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=取消优惠券转赠)
func (client *Client) CouponPresentCancel(request *CouponPresentCancelRequest) (response *CouponPresentCancelResponse, err error) {
	rpcResponse := CreateCouponPresentCancelResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponPresentCancelRequest struct {
	*api.BaseRequest

	StoreId    int64  `json:"storeId,omitempty"`
	Wid        int64  `json:"wid,omitempty"`
	CouponCode string `json:"couponCode,omitempty"`
}

type CouponPresentCancelResponse struct {
	Succeed bool `json:"succeed,omitempty"`
}

func CreateCouponPresentCancelRequest() (request *CouponPresentCancelRequest) {
	request = &CouponPresentCancelRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "coupon/presentCancel", "api")
	request.Method = api.POST
	return
}

func CreateCouponPresentCancelResponse() (response *api.BaseResponse[CouponPresentCancelResponse]) {
	response = api.CreateResponse[CouponPresentCancelResponse](&CouponPresentCancelResponse{})
	return
}
