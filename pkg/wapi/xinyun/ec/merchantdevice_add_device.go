package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantdeviceAddDevice
// @id 1476
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=添加设备)
func (client *Client) MerchantdeviceAddDevice(request *MerchantdeviceAddDeviceRequest) (response *MerchantdeviceAddDeviceResponse, err error) {
	rpcResponse := CreateMerchantdeviceAddDeviceResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantdeviceAddDeviceRequest struct {
	*api.BaseRequest

	EcBizStoreId   int64  `json:"ecBizStoreId,omitempty"`
	DeviceType     int    `json:"deviceType,omitempty"`
	DeviceNumber   string `json:"deviceNumber,omitempty"`
	SupplierNumber string `json:"supplierNumber,omitempty"`
	DeviceName     string `json:"deviceName,omitempty"`
	DeviceRemark   string `json:"deviceRemark,omitempty"`
	ProvinceCode   string `json:"provinceCode,omitempty"`
	CityCode       string `json:"cityCode,omitempty"`
	CountyCode     string `json:"countyCode,omitempty"`
	ProvinceName   string `json:"provinceName,omitempty"`
	CityName       string `json:"cityName,omitempty"`
	CountyName     string `json:"countyName,omitempty"`
	Address        string `json:"address,omitempty"`
}

type MerchantdeviceAddDeviceResponse struct {
}

func CreateMerchantdeviceAddDeviceRequest() (request *MerchantdeviceAddDeviceRequest) {
	request = &MerchantdeviceAddDeviceRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchantdevice/addDevice", "api")
	request.Method = api.POST
	return
}

func CreateMerchantdeviceAddDeviceResponse() (response *api.BaseResponse[MerchantdeviceAddDeviceResponse]) {
	response = api.CreateResponse[MerchantdeviceAddDeviceResponse](&MerchantdeviceAddDeviceResponse{})
	return
}
