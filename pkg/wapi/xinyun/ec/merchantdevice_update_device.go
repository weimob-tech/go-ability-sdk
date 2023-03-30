package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantdeviceUpdateDevice
// @id 1478
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改设备)
func (client *Client) MerchantdeviceUpdateDevice(request *MerchantdeviceUpdateDeviceRequest) (response *MerchantdeviceUpdateDeviceResponse, err error) {
	rpcResponse := CreateMerchantdeviceUpdateDeviceResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantdeviceUpdateDeviceRequest struct {
	*api.BaseRequest

	DeviceType     int    `json:"deviceType,omitempty"`
	DeviceNumber   string `json:"deviceNumber,omitempty"`
	SupplierNumber string `json:"supplierNumber,omitempty"`
	DeviceName     string `json:"deviceName,omitempty"`
	DeviceRemark   string `json:"deviceRemark,omitempty"`
	ProvinceName   string `json:"provinceName,omitempty"`
	CityName       string `json:"cityName,omitempty"`
	CountyName     string `json:"countyName,omitempty"`
	Address        string `json:"address,omitempty"`
}

type MerchantdeviceUpdateDeviceResponse struct {
}

func CreateMerchantdeviceUpdateDeviceRequest() (request *MerchantdeviceUpdateDeviceRequest) {
	request = &MerchantdeviceUpdateDeviceRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchantdevice/updateDevice", "api")
	request.Method = api.POST
	return
}

func CreateMerchantdeviceUpdateDeviceResponse() (response *api.BaseResponse[MerchantdeviceUpdateDeviceResponse]) {
	response = api.CreateResponse[MerchantdeviceUpdateDeviceResponse](&MerchantdeviceUpdateDeviceResponse{})
	return
}
