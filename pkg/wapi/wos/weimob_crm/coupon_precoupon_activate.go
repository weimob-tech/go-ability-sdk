package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponPrecouponActivate
// @id 1609
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1609?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=激活预发优惠券)
func (client *Client) CouponPrecouponActivate(request *CouponPrecouponActivateRequest) (response *CouponPrecouponActivateResponse, err error) {
	rpcResponse := CreateCouponPrecouponActivateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponPrecouponActivateRequest struct {
	*api.BaseRequest

	Wid     int64  `json:"wid,omitempty"`
	ApplyNo string `json:"applyNo,omitempty"`
}

type CouponPrecouponActivateResponse struct {
	CouponFailList    []CouponPrecouponActivateResponseCouponFailList    `json:"couponFailList,omitempty"`
	CouponSuccessList []CouponPrecouponActivateResponseCouponSuccessList `json:"couponSuccessList,omitempty"`
	FailCodes         []int64                                            `json:"failCodes,omitempty"`
	SuccessCount      int64                                              `json:"successCount,omitempty"`
	FailedCount       int64                                              `json:"failedCount,omitempty"`
}

func CreateCouponPrecouponActivateRequest() (request *CouponPrecouponActivateRequest) {
	request = &CouponPrecouponActivateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/precoupon/activate", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponPrecouponActivateResponse() (response *api.BaseResponse[CouponPrecouponActivateResponse]) {
	response = api.CreateResponse[CouponPrecouponActivateResponse](&CouponPrecouponActivateResponse{})
	return
}

type CouponPrecouponActivateResponseCouponFailList struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	Code             string `json:"code,omitempty"`
	ErrCode          string `json:"errCode,omitempty"`
	ErrMsg           string `json:"errMsg,omitempty"`
}

type CouponPrecouponActivateResponseCouponSuccessList struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	Code             string `json:"code,omitempty"`
}
