package kld

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// IMemberCouponServiceVerificationGood
// @id 229
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=核销产品)
func (client *Client) IMemberCouponServiceVerificationGood(request *IMemberCouponServiceVerificationGoodRequest) (response *IMemberCouponServiceVerificationGoodResponse, err error) {
	rpcResponse := CreateIMemberCouponServiceVerificationGoodResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type IMemberCouponServiceVerificationGoodRequest struct {
	*api.BaseRequest

	Data          IMemberCouponServiceVerificationGoodRequestData `json:"data,omitempty"`
	CurrentUserId string                                          `json:"currentUserId,omitempty"`
	SnNo          string                                          `json:"snNo,omitempty"`
	StoreId       int64                                           `json:"storeId,omitempty"`
	FromType      int                                             `json:"fromType,omitempty"`
}

type IMemberCouponServiceVerificationGoodResponse struct {
}

func CreateIMemberCouponServiceVerificationGoodRequest() (request *IMemberCouponServiceVerificationGoodRequest) {
	request = &IMemberCouponServiceVerificationGoodRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("kld", "1_0", "IMemberCouponService/verificationGood", "api")
	request.Method = api.POST
	return
}

func CreateIMemberCouponServiceVerificationGoodResponse() (response *api.BaseResponse[IMemberCouponServiceVerificationGoodResponse]) {
	response = api.CreateResponse[IMemberCouponServiceVerificationGoodResponse](&IMemberCouponServiceVerificationGoodResponse{})
	return
}

type IMemberCouponServiceVerificationGoodRequestData struct {
}
