package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FulfillVirtualCouponStatusNotify
// @id 1625
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1625?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=产品融合-履约-券码状态通知回调)
func (client *Client) FulfillVirtualCouponStatusNotify(request *FulfillVirtualCouponStatusNotifyRequest) (response *FulfillVirtualCouponStatusNotifyResponse, err error) {
	rpcResponse := CreateFulfillVirtualCouponStatusNotifyResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FulfillVirtualCouponStatusNotifyRequest struct {
	*api.BaseRequest

	BizId           int64 `json:"bizId,omitempty"`
	CouponFinalTime int64 `json:"couponFinalTime,omitempty"`
	CouponStatus    int64 `json:"couponStatus,omitempty"`
	CouponUseTime   int64 `json:"couponUseTime,omitempty"`
}

type FulfillVirtualCouponStatusNotifyResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateFulfillVirtualCouponStatusNotifyRequest() (request *FulfillVirtualCouponStatusNotifyRequest) {
	request = &FulfillVirtualCouponStatusNotifyRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "fulfill/virtualCouponStatusNotify", "apigw")
	request.Method = api.POST
	return
}

func CreateFulfillVirtualCouponStatusNotifyResponse() (response *api.BaseResponse[FulfillVirtualCouponStatusNotifyResponse]) {
	response = api.CreateResponse[FulfillVirtualCouponStatusNotifyResponse](&FulfillVirtualCouponStatusNotifyResponse{})
	return
}
