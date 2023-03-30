package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponOpenCouponRecycle
// @id 1311
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=核销券后撤回)
func (client *Client) CouponOpenCouponRecycle(request *CouponOpenCouponRecycleRequest) (response *CouponOpenCouponRecycleResponse, err error) {
	rpcResponse := CreateCouponOpenCouponRecycleResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponOpenCouponRecycleRequest struct {
	*api.BaseRequest

	CardCode string `json:"cardCode,omitempty"`
}

type CouponOpenCouponRecycleResponse struct {
}

func CreateCouponOpenCouponRecycleRequest() (request *CouponOpenCouponRecycleRequest) {
	request = &CouponOpenCouponRecycleRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "couponOpen/couponRecycle", "api")
	request.Method = api.POST
	return
}

func CreateCouponOpenCouponRecycleResponse() (response *api.BaseResponse[CouponOpenCouponRecycleResponse]) {
	response = api.CreateResponse[CouponOpenCouponRecycleResponse](&CouponOpenCouponRecycleResponse{})
	return
}
