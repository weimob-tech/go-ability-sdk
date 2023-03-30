package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FulfillPickupChargeOff
// @id 1735
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1735?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=自提订单核销)
func (client *Client) FulfillPickupChargeOff(request *FulfillPickupChargeOffRequest) (response *FulfillPickupChargeOffResponse, err error) {
	rpcResponse := CreateFulfillPickupChargeOffResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FulfillPickupChargeOffRequest struct {
	*api.BaseRequest

	ChargeOffUserInfo  FulfillPickupChargeOffRequestChargeOffUserInfo    `json:"chargeOffUserInfo,omitempty"`
	BasicInfo          FulfillPickupChargeOffRequestBasicInfo            `json:"basicInfo,omitempty"`
	SelfPickupCodeList []FulfillPickupChargeOffRequestSelfPickupCodeList `json:"selfPickupCodeList,omitempty"`
	OrderNo            int64                                             `json:"orderNo,omitempty"`
	FulfillNoList      []int64                                           `json:"fulfillNoList,omitempty"`
}

type FulfillPickupChargeOffResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateFulfillPickupChargeOffRequest() (request *FulfillPickupChargeOffRequest) {
	request = &FulfillPickupChargeOffRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "fulfill/pickup/chargeOff", "apigw")
	request.Method = api.POST
	return
}

func CreateFulfillPickupChargeOffResponse() (response *api.BaseResponse[FulfillPickupChargeOffResponse]) {
	response = api.CreateResponse[FulfillPickupChargeOffResponse](&FulfillPickupChargeOffResponse{})
	return
}

type FulfillPickupChargeOffRequestChargeOffUserInfo struct {
	OperatorId    string `json:"operatorId,omitempty"`
	OperatorName  string `json:"operatorName,omitempty"`
	OperatorPhone string `json:"operatorPhone,omitempty"`
}

type FulfillPickupChargeOffRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type FulfillPickupChargeOffRequestSelfPickupCodeList struct {
}
