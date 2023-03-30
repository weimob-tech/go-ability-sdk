package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponDeleteCoupon
// @id 1085
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除优惠券)
func (client *Client) CouponDeleteCoupon(request *CouponDeleteCouponRequest) (response *CouponDeleteCouponResponse, err error) {
	rpcResponse := CreateCouponDeleteCouponResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponDeleteCouponRequest struct {
	*api.BaseRequest

	CardTemplateIds []int64 `json:"cardTemplateIds,omitempty"`
}

type CouponDeleteCouponResponse struct {
}

func CreateCouponDeleteCouponRequest() (request *CouponDeleteCouponRequest) {
	request = &CouponDeleteCouponRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "coupon/deleteCoupon", "api")
	request.Method = api.POST
	return
}

func CreateCouponDeleteCouponResponse() (response *api.BaseResponse[CouponDeleteCouponResponse]) {
	response = api.CreateResponse[CouponDeleteCouponResponse](&CouponDeleteCouponResponse{})
	return
}
