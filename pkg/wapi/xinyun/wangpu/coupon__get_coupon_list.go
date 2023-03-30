package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponGetCouponList
// @id 363
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=分页获取优惠券列表)
func (client *Client) CouponGetCouponList(request *CouponGetCouponListRequest) (response *CouponGetCouponListResponse, err error) {
	rpcResponse := CreateCouponGetCouponListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponGetCouponListRequest struct {
	*api.BaseRequest

	CouponType     string `json:"coupon_type,omitempty"`
	CouponStatus   string `json:"coupon_status,omitempty"`
	CouponSendType string `json:"coupon_send_type,omitempty"`
	PageSize       int    `json:"page_size,omitempty"`
	PageNo         int    `json:"page_no,omitempty"`
}

type CouponGetCouponListResponse struct {
}

func CreateCouponGetCouponListRequest() (request *CouponGetCouponListRequest) {
	request = &CouponGetCouponListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "coupon/GetCouponList", "api")
	request.Method = api.POST
	return
}

func CreateCouponGetCouponListResponse() (response *api.BaseResponse[CouponGetCouponListResponse]) {
	response = api.CreateResponse[CouponGetCouponListResponse](&CouponGetCouponListResponse{})
	return
}
