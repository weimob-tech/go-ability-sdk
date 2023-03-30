package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DeliverySettingUpdateCityDeliverySetting
// @id 2444
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新同城配送设置)
func (client *Client) DeliverySettingUpdateCityDeliverySetting(request *DeliverySettingUpdateCityDeliverySettingRequest) (response *DeliverySettingUpdateCityDeliverySettingResponse, err error) {
	rpcResponse := CreateDeliverySettingUpdateCityDeliverySettingResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DeliverySettingUpdateCityDeliverySettingRequest struct {
	*api.BaseRequest

	Id                          string `json:"id,omitempty"`
	DeliveryTypeName            string `json:"deliveryTypeName,omitempty"`
	IsSupported                 bool   `json:"isSupported,omitempty"`
	IsDefault                   bool   `json:"isDefault,omitempty"`
	FreightConstituteRule       int    `json:"freightConstituteRule,omitempty"`
	Pid                         string `json:"pid,omitempty"`
	StoreId                     string `json:"storeId,omitempty"`
	IsSupportThirdPartLogistics bool   `json:"isSupportThirdPartLogistics,omitempty"`
	IsAutoSendOrder             bool   `json:"isAutoSendOrder,omitempty"`
	LeastAheadMinute            int    `json:"leastAheadMinute,omitempty"`
	DefaultDeliveryType         int64  `json:"defaultDeliveryType,omitempty"`
}

type DeliverySettingUpdateCityDeliverySettingResponse struct {
}

func CreateDeliverySettingUpdateCityDeliverySettingRequest() (request *DeliverySettingUpdateCityDeliverySettingRequest) {
	request = &DeliverySettingUpdateCityDeliverySettingRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "deliverySetting/updateCityDeliverySetting", "api")
	request.Method = api.POST
	return
}

func CreateDeliverySettingUpdateCityDeliverySettingResponse() (response *api.BaseResponse[DeliverySettingUpdateCityDeliverySettingResponse]) {
	response = api.CreateResponse[DeliverySettingUpdateCityDeliverySettingResponse](&DeliverySettingUpdateCityDeliverySettingResponse{})
	return
}
