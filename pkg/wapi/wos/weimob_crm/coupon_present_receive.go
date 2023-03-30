package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponPresentReceive
// @id 2104
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2104?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=券转赠领取)
func (client *Client) CouponPresentReceive(request *CouponPresentReceiveRequest) (response *CouponPresentReceiveResponse, err error) {
	rpcResponse := CreateCouponPresentReceiveResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponPresentReceiveRequest struct {
	*api.BaseRequest

	FromWid   int64  `json:"fromWid,omitempty"`
	TargetWid int64  `json:"targetWid,omitempty"`
	Code      string `json:"code,omitempty"`
}

type CouponPresentReceiveResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateCouponPresentReceiveRequest() (request *CouponPresentReceiveRequest) {
	request = &CouponPresentReceiveRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/present/receive", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponPresentReceiveResponse() (response *api.BaseResponse[CouponPresentReceiveResponse]) {
	response = api.CreateResponse[CouponPresentReceiveResponse](&CouponPresentReceiveResponse{})
	return
}
