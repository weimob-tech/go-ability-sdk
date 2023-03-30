package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponConsumeCouponForOffline
// @id 583
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=核销优惠券-线下)
func (client *Client) CouponConsumeCouponForOffline(request *CouponConsumeCouponForOfflineRequest) (response *CouponConsumeCouponForOfflineResponse, err error) {
	rpcResponse := CreateCouponConsumeCouponForOfflineResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponConsumeCouponForOfflineRequest struct {
	*api.BaseRequest

	Code          string `json:"code,omitempty"`
	OperatorPhone string `json:"operatorPhone,omitempty"`
	Wid           int64  `json:"wid,omitempty"`
	Operator      string `json:"operator,omitempty"`
	UseScene      int64  `json:"useScene,omitempty"`
	StoreId       int64  `json:"storeId,omitempty"`
}

type CouponConsumeCouponForOfflineResponse struct {
}

func CreateCouponConsumeCouponForOfflineRequest() (request *CouponConsumeCouponForOfflineRequest) {
	request = &CouponConsumeCouponForOfflineRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "coupon/consumeCouponForOffline", "api")
	request.Method = api.POST
	return
}

func CreateCouponConsumeCouponForOfflineResponse() (response *api.BaseResponse[CouponConsumeCouponForOfflineResponse]) {
	response = api.CreateResponse[CouponConsumeCouponForOfflineResponse](&CouponConsumeCouponForOfflineResponse{})
	return
}
