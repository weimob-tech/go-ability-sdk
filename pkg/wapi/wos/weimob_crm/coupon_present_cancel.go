package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponPresentCancel
// @id 2269
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2269?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=优惠券取消转赠)
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

	BosId int64  `json:"bosId,omitempty"`
	Wid   int64  `json:"wid,omitempty"`
	Code  string `json:"code,omitempty"`
}

type CouponPresentCancelResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateCouponPresentCancelRequest() (request *CouponPresentCancelRequest) {
	request = &CouponPresentCancelRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/present/cancel", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponPresentCancelResponse() (response *api.BaseResponse[CouponPresentCancelResponse]) {
	response = api.CreateResponse[CouponPresentCancelResponse](&CouponPresentCancelResponse{})
	return
}
