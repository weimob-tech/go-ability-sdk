package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponCustomerSync
// @id 2166
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2166?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=用户券信息同步)
func (client *Client) CouponCustomerSync(request *CouponCustomerSyncRequest) (response *CouponCustomerSyncResponse, err error) {
	rpcResponse := CreateCouponCustomerSyncResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponCustomerSyncRequest struct {
	*api.BaseRequest

	AcquireTime      int64  `json:"acquireTime,omitempty"`
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	Code             string `json:"code,omitempty"`
	Wid              int64  `json:"wid,omitempty"`
	ExpireEndTime    int64  `json:"expireEndTime,omitempty"`
	ExpireStartTime  int64  `json:"expireStartTime,omitempty"`
	Vid              int64  `json:"vid,omitempty"`
}

type CouponCustomerSyncResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateCouponCustomerSyncRequest() (request *CouponCustomerSyncRequest) {
	request = &CouponCustomerSyncRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/customer/sync", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponCustomerSyncResponse() (response *api.BaseResponse[CouponCustomerSyncResponse]) {
	response = api.CreateResponse[CouponCustomerSyncResponse](&CouponCustomerSyncResponse{})
	return
}
