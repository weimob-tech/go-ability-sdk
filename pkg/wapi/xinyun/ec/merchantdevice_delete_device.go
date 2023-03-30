package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantdeviceDeleteDevice
// @id 1477
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除设备)
func (client *Client) MerchantdeviceDeleteDevice(request *MerchantdeviceDeleteDeviceRequest) (response *MerchantdeviceDeleteDeviceResponse, err error) {
	rpcResponse := CreateMerchantdeviceDeleteDeviceResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantdeviceDeleteDeviceRequest struct {
	*api.BaseRequest

	DeviceNumber string `json:"deviceNumber,omitempty"`
}

type MerchantdeviceDeleteDeviceResponse struct {
}

func CreateMerchantdeviceDeleteDeviceRequest() (request *MerchantdeviceDeleteDeviceRequest) {
	request = &MerchantdeviceDeleteDeviceRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchantdevice/deleteDevice", "api")
	request.Method = api.POST
	return
}

func CreateMerchantdeviceDeleteDeviceResponse() (response *api.BaseResponse[MerchantdeviceDeleteDeviceResponse]) {
	response = api.CreateResponse[MerchantdeviceDeleteDeviceResponse](&MerchantdeviceDeleteDeviceResponse{})
	return
}
