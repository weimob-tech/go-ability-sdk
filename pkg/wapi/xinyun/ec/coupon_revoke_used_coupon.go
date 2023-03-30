package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponRevokeUsedCoupon
// @id 2191
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=撤销核销)
func (client *Client) CouponRevokeUsedCoupon(request *CouponRevokeUsedCouponRequest) (response *CouponRevokeUsedCouponResponse, err error) {
	rpcResponse := CreateCouponRevokeUsedCouponResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponRevokeUsedCouponRequest struct {
	*api.BaseRequest

	Code     CouponRevokeUsedCouponRequestCode `json:"code,omitempty"`
	CodeList []string                          `json:"codeList,omitempty"`
}

type CouponRevokeUsedCouponResponse struct {
	Succeed bool `json:"succeed,omitempty"`
}

func CreateCouponRevokeUsedCouponRequest() (request *CouponRevokeUsedCouponRequest) {
	request = &CouponRevokeUsedCouponRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "coupon/revokeUsedCoupon", "api")
	request.Method = api.POST
	return
}

func CreateCouponRevokeUsedCouponResponse() (response *api.BaseResponse[CouponRevokeUsedCouponResponse]) {
	response = api.CreateResponse[CouponRevokeUsedCouponResponse](&CouponRevokeUsedCouponResponse{})
	return
}

type CouponRevokeUsedCouponRequestCode struct {
	Errcode string `json:"errcode,omitempty"`
	Errmag  string `json:"errmag,omitempty"`
}
