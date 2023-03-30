package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponTemplateStatusUpdate
// @id 2056
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2056?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新优惠券状态)
func (client *Client) CouponTemplateStatusUpdate(request *CouponTemplateStatusUpdateRequest) (response *CouponTemplateStatusUpdateResponse, err error) {
	rpcResponse := CreateCouponTemplateStatusUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponTemplateStatusUpdateRequest struct {
	*api.BaseRequest

	CouponTemplateId int64 `json:"couponTemplateId,omitempty"`
	Status           int64 `json:"status,omitempty"`
}

type CouponTemplateStatusUpdateResponse struct {
}

func CreateCouponTemplateStatusUpdateRequest() (request *CouponTemplateStatusUpdateRequest) {
	request = &CouponTemplateStatusUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/template/status/update", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponTemplateStatusUpdateResponse() (response *api.BaseResponse[CouponTemplateStatusUpdateResponse]) {
	response = api.CreateResponse[CouponTemplateStatusUpdateResponse](&CouponTemplateStatusUpdateResponse{})
	return
}
