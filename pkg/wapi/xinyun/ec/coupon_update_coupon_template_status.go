package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponUpdateCouponTemplateStatus
// @id 1695
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新优惠券模板状态)
func (client *Client) CouponUpdateCouponTemplateStatus(request *CouponUpdateCouponTemplateStatusRequest) (response *CouponUpdateCouponTemplateStatusResponse, err error) {
	rpcResponse := CreateCouponUpdateCouponTemplateStatusResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponUpdateCouponTemplateStatusRequest struct {
	*api.BaseRequest

	StoreId        int64 `json:"storeId,omitempty"`
	CardTemplateId int64 `json:"cardTemplateId,omitempty"`
	ActionId       int   `json:"actionId,omitempty"`
}

type CouponUpdateCouponTemplateStatusResponse struct {
}

func CreateCouponUpdateCouponTemplateStatusRequest() (request *CouponUpdateCouponTemplateStatusRequest) {
	request = &CouponUpdateCouponTemplateStatusRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "coupon/updateCouponTemplateStatus", "api")
	request.Method = api.POST
	return
}

func CreateCouponUpdateCouponTemplateStatusResponse() (response *api.BaseResponse[CouponUpdateCouponTemplateStatusResponse]) {
	response = api.CreateResponse[CouponUpdateCouponTemplateStatusResponse](&CouponUpdateCouponTemplateStatusResponse{})
	return
}
