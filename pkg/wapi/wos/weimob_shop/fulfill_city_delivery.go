package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FulfillCityDelivery
// @id 1736
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1736?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=同城配送发货)
func (client *Client) FulfillCityDelivery(request *FulfillCityDeliveryRequest) (response *FulfillCityDeliveryResponse, err error) {
	rpcResponse := CreateFulfillCityDeliveryResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FulfillCityDeliveryRequest struct {
	*api.BaseRequest

	OperatorVo     FulfillCityDeliveryRequestOperatorVo     `json:"operatorVo,omitempty"`
	FulfillItems   []FulfillCityDeliveryRequestFulfillItems `json:"fulfillItems,omitempty"`
	BasicInfo      FulfillCityDeliveryRequestBasicInfo      `json:"basicInfo,omitempty"`
	OrderNo        int64                                    `json:"orderNo,omitempty"`
	IsSplitPackage bool                                     `json:"isSplitPackage,omitempty"`
	FulfillMethod  int64                                    `json:"fulfillMethod,omitempty"`
	Remark         string                                   `json:"remark,omitempty"`
}

type FulfillCityDeliveryResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateFulfillCityDeliveryRequest() (request *FulfillCityDeliveryRequest) {
	request = &FulfillCityDeliveryRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "fulfill/city/delivery", "apigw")
	request.Method = api.POST
	return
}

func CreateFulfillCityDeliveryResponse() (response *api.BaseResponse[FulfillCityDeliveryResponse]) {
	response = api.CreateResponse[FulfillCityDeliveryResponse](&FulfillCityDeliveryResponse{})
	return
}

type FulfillCityDeliveryRequestOperatorVo struct {
	OperatorId    string `json:"operatorId,omitempty"`
	OperatorName  string `json:"operatorName,omitempty"`
	OperatorPhone string `json:"operatorPhone,omitempty"`
}

type FulfillCityDeliveryRequestFulfillItems struct {
	OrderItemId int64 `json:"orderItemId,omitempty"`
	SkuNum      int64 `json:"skuNum,omitempty"`
}

type FulfillCityDeliveryRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}
