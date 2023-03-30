package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FulfillLogisticsInfoUpdate
// @id 1828
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1828?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=修改物流)
func (client *Client) FulfillLogisticsInfoUpdate(request *FulfillLogisticsInfoUpdateRequest) (response *FulfillLogisticsInfoUpdateResponse, err error) {
	rpcResponse := CreateFulfillLogisticsInfoUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FulfillLogisticsInfoUpdateRequest struct {
	*api.BaseRequest

	LogisticsVo   FulfillLogisticsInfoUpdateRequestLogisticsVo `json:"logisticsVo,omitempty"`
	BasicInfo     FulfillLogisticsInfoUpdateRequestBasicInfo   `json:"basicInfo,omitempty"`
	OrderNo       int64                                        `json:"orderNo,omitempty"`
	FulfillNo     int64                                        `json:"fulfillNo,omitempty"`
	FulfillMethod int64                                        `json:"fulfillMethod,omitempty"`
	Remark        string                                       `json:"remark,omitempty"`
}

type FulfillLogisticsInfoUpdateResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateFulfillLogisticsInfoUpdateRequest() (request *FulfillLogisticsInfoUpdateRequest) {
	request = &FulfillLogisticsInfoUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "fulfill/logistics/info/update", "apigw")
	request.Method = api.POST
	return
}

func CreateFulfillLogisticsInfoUpdateResponse() (response *api.BaseResponse[FulfillLogisticsInfoUpdateResponse]) {
	response = api.CreateResponse[FulfillLogisticsInfoUpdateResponse](&FulfillLogisticsInfoUpdateResponse{})
	return
}

type FulfillLogisticsInfoUpdateRequestLogisticsVo struct {
	DeliveryNo          string `json:"deliveryNo,omitempty"`
	DeliveryCompanyCode string `json:"deliveryCompanyCode,omitempty"`
	DeliveryCompanyName string `json:"deliveryCompanyName,omitempty"`
}

type FulfillLogisticsInfoUpdateRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}
