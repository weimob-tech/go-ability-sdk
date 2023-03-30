package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponPrecouponSend
// @id 1610
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1610?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=预发优惠券)
func (client *Client) CouponPrecouponSend(request *CouponPrecouponSendRequest) (response *CouponPrecouponSendResponse, err error) {
	rpcResponse := CreateCouponPrecouponSendResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponPrecouponSendRequest struct {
	*api.BaseRequest

	CouponNums []CouponPrecouponSendRequestCouponNums `json:"couponNums,omitempty"`
	Wid        int64                                  `json:"wid,omitempty"`
	Channel    int64                                  `json:"channel,omitempty"`
	Scene      int64                                  `json:"scene,omitempty"`
	SceneId    string                                 `json:"sceneId,omitempty"`
	VidType    int64                                  `json:"vidType,omitempty"`
	Vid        int64                                  `json:"vid,omitempty"`
}

type CouponPrecouponSendResponse struct {
	Result  bool   `json:"result,omitempty"`
	ApplyNo string `json:"applyNo,omitempty"`
}

func CreateCouponPrecouponSendRequest() (request *CouponPrecouponSendRequest) {
	request = &CouponPrecouponSendRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/precoupon/send", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponPrecouponSendResponse() (response *api.BaseResponse[CouponPrecouponSendResponse]) {
	response = api.CreateResponse[CouponPrecouponSendResponse](&CouponPrecouponSendResponse{})
	return
}

type CouponPrecouponSendRequestCouponNums struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	Num              int64  `json:"num,omitempty"`
	RequestId        string `json:"requestId,omitempty"`
}
