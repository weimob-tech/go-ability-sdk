package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponGetCouponHouseList
// @id 365
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取优惠券码库)
func (client *Client) CouponGetCouponHouseList(request *CouponGetCouponHouseListRequest) (response *CouponGetCouponHouseListResponse, err error) {
	rpcResponse := CreateCouponGetCouponHouseListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponGetCouponHouseListRequest struct {
	*api.BaseRequest

	CouponId int64 `json:"coupon_id,omitempty"`
	PageSize int64 `json:"page_size,omitempty"`
	PageNo   int   `json:"page_no,omitempty"`
}

type CouponGetCouponHouseListResponse struct {
}

func CreateCouponGetCouponHouseListRequest() (request *CouponGetCouponHouseListRequest) {
	request = &CouponGetCouponHouseListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "coupon/GetCouponHouseList", "api")
	request.Method = api.POST
	return
}

func CreateCouponGetCouponHouseListResponse() (response *api.BaseResponse[CouponGetCouponHouseListResponse]) {
	response = api.CreateResponse[CouponGetCouponHouseListResponse](&CouponGetCouponHouseListResponse{})
	return
}
