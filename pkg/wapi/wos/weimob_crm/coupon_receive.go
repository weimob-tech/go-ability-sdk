package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponReceive
// @id 1136
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1136?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=发送优惠券给指定客户)
func (client *Client) CouponReceive(request *CouponReceiveRequest) (response *CouponReceiveResponse, err error) {
	rpcResponse := CreateCouponReceiveResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponReceiveRequest struct {
	*api.BaseRequest

	CouponNums []CouponReceiveRequestCouponNums `json:"couponNums,omitempty"`
	Wid        int64                            `json:"wid,omitempty"`
	Scene      int64                            `json:"scene,omitempty"`
	VidType    int64                            `json:"vidType,omitempty"`
	Vid        int64                            `json:"vid,omitempty"`
	SceneId    int64                            `json:"sceneId,omitempty"`
}

type CouponReceiveResponse struct {
	CouponResultList []CouponReceiveResponseCouponResultList `json:"couponResultList,omitempty"`
	SuccessCount     int64                                   `json:"successCount,omitempty"`
	FailedCount      int64                                   `json:"failedCount,omitempty"`
	Status           int64                                   `json:"status,omitempty"`
}

func CreateCouponReceiveRequest() (request *CouponReceiveRequest) {
	request = &CouponReceiveRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/receive", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponReceiveResponse() (response *api.BaseResponse[CouponReceiveResponse]) {
	response = api.CreateResponse[CouponReceiveResponse](&CouponReceiveResponse{})
	return
}

type CouponReceiveRequestCouponNums struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	Num              int64  `json:"num,omitempty"`
	RequestId        string `json:"requestId,omitempty"`
	Code             string `json:"code,omitempty"`
}

type CouponReceiveResponseCouponResultList struct {
	CouponTemplateId int64   `json:"couponTemplateId,omitempty"`
	Name             string  `json:"name,omitempty"`
	Wid              int64   `json:"wid,omitempty"`
	IsSuccess        bool    `json:"isSuccess,omitempty"`
	FailureNum       int64   `json:"failureNum,omitempty"`
	ErrCode          string  `json:"errCode,omitempty"`
	ErrMsg           string  `json:"errMsg,omitempty"`
	Codes            []int64 `json:"codes,omitempty"`
	CanReceive       int64   `json:"canReceive,omitempty"`
}
