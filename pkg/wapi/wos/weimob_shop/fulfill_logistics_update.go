package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FulfillLogisticsUpdate
// @id 1732
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1732?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商家配送发货)
func (client *Client) FulfillLogisticsUpdate(request *FulfillLogisticsUpdateRequest) (response *FulfillLogisticsUpdateResponse, err error) {
	rpcResponse := CreateFulfillLogisticsUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FulfillLogisticsUpdateRequest struct {
	*api.BaseRequest

	Logistics      FulfillLogisticsUpdateRequestLogistics      `json:"logistics,omitempty"`
	OperatorVo     FulfillLogisticsUpdateRequestOperatorVo     `json:"operatorVo,omitempty"`
	FulfillItems   []FulfillLogisticsUpdateRequestFulfillItems `json:"fulfillItems,omitempty"`
	BasicInfo      FulfillLogisticsUpdateRequestBasicInfo      `json:"basicInfo,omitempty"`
	OrderNo        int64                                       `json:"orderNo,omitempty"`
	IsSplitPackage bool                                        `json:"isSplitPackage,omitempty"`
	FulfillMethod  int64                                       `json:"fulfillMethod,omitempty"`
}

type FulfillLogisticsUpdateResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateFulfillLogisticsUpdateRequest() (request *FulfillLogisticsUpdateRequest) {
	request = &FulfillLogisticsUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "fulfill/logistics/update", "apigw")
	request.Method = api.POST
	return
}

func CreateFulfillLogisticsUpdateResponse() (response *api.BaseResponse[FulfillLogisticsUpdateResponse]) {
	response = api.CreateResponse[FulfillLogisticsUpdateResponse](&FulfillLogisticsUpdateResponse{})
	return
}

type FulfillLogisticsUpdateRequestLogistics struct {
	DeliveryNo          string `json:"deliveryNo,omitempty"`
	DeliveryCompanyCode string `json:"deliveryCompanyCode,omitempty"`
	DeliveryCompanyName string `json:"deliveryCompanyName,omitempty"`
}

type FulfillLogisticsUpdateRequestOperatorVo struct {
	OperatorId    string `json:"operatorId,omitempty"`
	OperatorName  string `json:"operatorName,omitempty"`
	OperatorPhone string `json:"operatorPhone,omitempty"`
}

type FulfillLogisticsUpdateRequestFulfillItems struct {
	OrderItemId int64 `json:"orderItemId,omitempty"`
	SkuNum      int64 `json:"skuNum,omitempty"`
}

type FulfillLogisticsUpdateRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}
