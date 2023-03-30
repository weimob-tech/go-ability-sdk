package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponGetCouponById
// @id 1086
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询优惠券详情)
func (client *Client) CouponGetCouponById(request *CouponGetCouponByIdRequest) (response *CouponGetCouponByIdResponse, err error) {
	rpcResponse := CreateCouponGetCouponByIdResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponGetCouponByIdRequest struct {
	*api.BaseRequest

	CardTemplateId int64 `json:"cardTemplateId,omitempty"`
}

type CouponGetCouponByIdResponse struct {
}

func CreateCouponGetCouponByIdRequest() (request *CouponGetCouponByIdRequest) {
	request = &CouponGetCouponByIdRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "coupon/getCouponById", "api")
	request.Method = api.POST
	return
}

func CreateCouponGetCouponByIdResponse() (response *api.BaseResponse[CouponGetCouponByIdResponse]) {
	response = api.CreateResponse[CouponGetCouponByIdResponse](&CouponGetCouponByIdResponse{})
	return
}
