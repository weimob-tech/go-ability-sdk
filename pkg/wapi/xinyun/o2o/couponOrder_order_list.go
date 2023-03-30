package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponOrderOrderList
// @id 1336
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询卡券购买订单接口)
func (client *Client) CouponOrderOrderList(request *CouponOrderOrderListRequest) (response *CouponOrderOrderListResponse, err error) {
	rpcResponse := CreateCouponOrderOrderListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponOrderOrderListRequest struct {
	*api.BaseRequest

	Index          string `json:"index,omitempty"`
	OrderTimeEnd   int64  `json:"orderTimeEnd,omitempty"`
	OrderTimeStart int64  `json:"orderTimeStart,omitempty"`
}

type CouponOrderOrderListResponse struct {
}

func CreateCouponOrderOrderListRequest() (request *CouponOrderOrderListRequest) {
	request = &CouponOrderOrderListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "couponOrder/orderList", "api")
	request.Method = api.POST
	return
}

func CreateCouponOrderOrderListResponse() (response *api.BaseResponse[CouponOrderOrderListResponse]) {
	response = api.CreateResponse[CouponOrderOrderListResponse](&CouponOrderOrderListResponse{})
	return
}
