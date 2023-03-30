package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DeliverySettingQueryCityDeliverySetting
// @id 2445
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询同城配送设置信息)
func (client *Client) DeliverySettingQueryCityDeliverySetting(request *DeliverySettingQueryCityDeliverySettingRequest) (response *DeliverySettingQueryCityDeliverySettingResponse, err error) {
	rpcResponse := CreateDeliverySettingQueryCityDeliverySettingResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DeliverySettingQueryCityDeliverySettingRequest struct {
	*api.BaseRequest

	Pid     int64 `json:"pid,omitempty"`
	StoreId int64 `json:"storeId,omitempty"`
}

type DeliverySettingQueryCityDeliverySettingResponse struct {
	Id                          int64  `json:"id,omitempty"`
	DeliveryTypeName            string `json:"deliveryTypeName,omitempty"`
	IsSupported                 bool   `json:"isSupported,omitempty"`
	IsDefault                   bool   `json:"isDefault,omitempty"`
	FreightConstituteRule       int    `json:"freightConstituteRule,omitempty"`
	Pid                         int64  `json:"pid,omitempty"`
	StoreId                     int64  `json:"storeId,omitempty"`
	IsSupportThirdPartLogistics bool   `json:"isSupportThirdPartLogistics,omitempty"`
	IsAutoSendOrder             bool   `json:"isAutoSendOrder,omitempty"`
	LeastAheadMinute            int    `json:"leastAheadMinute,omitempty"`
}

func CreateDeliverySettingQueryCityDeliverySettingRequest() (request *DeliverySettingQueryCityDeliverySettingRequest) {
	request = &DeliverySettingQueryCityDeliverySettingRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "deliverySetting/queryCityDeliverySetting", "api")
	request.Method = api.POST
	return
}

func CreateDeliverySettingQueryCityDeliverySettingResponse() (response *api.BaseResponse[DeliverySettingQueryCityDeliverySettingResponse]) {
	response = api.CreateResponse[DeliverySettingQueryCityDeliverySettingResponse](&DeliverySettingQueryCityDeliverySettingResponse{})
	return
}
