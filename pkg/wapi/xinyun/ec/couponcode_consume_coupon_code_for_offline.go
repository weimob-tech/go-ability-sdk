package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponcodeConsumeCouponCodeForOffline
// @id 1492
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=线下核销优惠码)
func (client *Client) CouponcodeConsumeCouponCodeForOffline(request *CouponcodeConsumeCouponCodeForOfflineRequest) (response *CouponcodeConsumeCouponCodeForOfflineResponse, err error) {
	rpcResponse := CreateCouponcodeConsumeCouponCodeForOfflineResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponcodeConsumeCouponCodeForOfflineRequest struct {
	*api.BaseRequest

	Code     string `json:"code,omitempty"`
	UseWid   int64  `json:"useWid,omitempty"`
	Operator string `json:"operator,omitempty"`
	StoreId  int64  `json:"storeId,omitempty"`
}

type CouponcodeConsumeCouponCodeForOfflineResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateCouponcodeConsumeCouponCodeForOfflineRequest() (request *CouponcodeConsumeCouponCodeForOfflineRequest) {
	request = &CouponcodeConsumeCouponCodeForOfflineRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "couponcode/consumeCouponCodeForOffline", "api")
	request.Method = api.POST
	return
}

func CreateCouponcodeConsumeCouponCodeForOfflineResponse() (response *api.BaseResponse[CouponcodeConsumeCouponCodeForOfflineResponse]) {
	response = api.CreateResponse[CouponcodeConsumeCouponCodeForOfflineResponse](&CouponcodeConsumeCouponCodeForOfflineResponse{})
	return
}
